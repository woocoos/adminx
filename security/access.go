package security

import (
	"context"
	"entgo.io/ent/dialect"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/tsingsun/woocoo/pkg/authz"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/security"
	"github.com/woocoos/adminx/graph/entgen/types"
	entadapter "github.com/woocoos/casbin-ent-adapter"
	entadapterent "github.com/woocoos/casbin-ent-adapter/ent"
)

var (
	authorizer      *authz.Authorization
	DomainHeaderKey = "X-Domain-ID"
)

// SetAuthorization 设置授权器
func SetAuthorization(cnf *conf.Configuration, driver dialect.Driver) error {
	casbinClient := entadapterent.NewClient(entadapterent.Driver(driver))
	adp, err := entadapter.NewAdapterWithClient(casbinClient)
	if err != nil {
		return err
	}
	err = casbinClient.Schema.Create(context.Background())
	if err != nil {
		return err
	}
	authz.SetAdapter(adp)
	authorizer, err = authz.NewAuthorization(cnf, authz.WithRequestParseFunc(RBACWithDomainRequestParserFunc))
	if err != nil {
		return err
	}
	authz.SetDefaultAuthorization(authorizer)
	return nil
}

// RBACWithDomainRequestParserFunc 以RBAC with domain模型生成casbin请求
//
// ctx: 一般就是gin.Context
func RBACWithDomainRequestParserFunc(ctx context.Context, identity security.Identity, item *security.PermissionItem) []any {
	gctx := ctx.Value(gin.ContextKey).(*gin.Context)
	domain := gctx.GetHeader(DomainHeaderKey)
	return []any{identity.Name(), domain, item.AppCode + ":" + item.Action, item.Operator}
}

// GrantPolicy 同步授权信息到cashbin.
//
//  1. 将授权信息同步到cashbin中.
//     如果授权主体为用户,则将用户作为特定角色.
//  2. 将授权信息同步到redis中
func GrantPolicy(rules []types.PolicyRule, principal, principalKind, domain string) error {
	role := principal
	switch principalKind {
	case "user":
		role = "r_" + principal
		_, err := authorizer.Enforcer.AddRoleForUserInDomain(principal, role, domain)
		if err != nil {
			return err
		}
	case "role":
	default:
		return errors.New("invalid principal kind")
	}
	pls := make([][]string, 0, len(rules))
	for _, rule := range rules {
		for _, action := range rule.Actions {
			pls = append(pls, []string{role, domain, action, "read", "allow"})
		}
	}
	_, err := authorizer.Enforcer.AddPolicies(pls)
	if err != nil {
		return err
	}
	return authorizer.Enforcer.SavePolicy()
}
