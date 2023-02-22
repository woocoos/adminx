package resource

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/woocoos/adminx/ent"
	"github.com/woocoos/adminx/ent/app"
	"github.com/woocoos/adminx/ent/apppolicy"
	"github.com/woocoos/adminx/ent/approle"
	"github.com/woocoos/adminx/ent/organization"
	"github.com/woocoos/adminx/ent/organizationrole"
	"strings"
)

// CreateAppPolicies 创建应用策略.
//
// 该方法会检查应用策略的规则中的action是否以应用代码开头.
func (s *Service) CreateAppPolicies(ctx context.Context, input []*ent.CreateAppPolicyInput) error {
	c := ent.FromContext(ctx)
	builders := make([]*ent.AppPolicyCreate, len(input))
	for i, data := range input {
		apl, err := s.Client.App.Query().Where(app.ID(data.AppID)).Only(ctx)
		if err != nil {
			return err
		}
		for _, rule := range data.Rules {
			for _, action := range rule.Actions {
				if action == "*" {
					continue
				}
				if !strings.HasPrefix(action, apl.Code+":") {
					return fmt.Errorf("action %s must start with app code %s", action, apl.Code)
				}
			}
		}
		builders[i] = c.AppPolicy.Create().SetInput(*input[i])
	}
	err := c.AppPolicy.CreateBulk(builders...).Exec(ctx)
	return err
}

// AssignOrganizationApp 分配应用到根组织下. 如: 新账户创建时, 根账户分配已有应用给子账户(需要验证根用户是否该应用权限,可在外层验证).
func (s *Service) AssignOrganizationApp(ctx context.Context, orgID int, appID int) (bool, error) {
	org := s.Client.Organization.Query().Where(organization.ID(orgID), organization.KindEQ(organization.KindRoot),
		organization.Not(organization.HasAppsWith(func(selector *sql.Selector) {
			selector.Where(sql.EQ(app.FieldID, appID))
		}))).OnlyX(ctx)
	ap := s.Client.App.Query().Where(app.ID(appID)).
		WithPolicies(func(query *ent.AppPolicyQuery) {
			query.Where(apppolicy.AutoGrant(true)).Unique(true).
				WithRoles(func(query *ent.AppRoleQuery) {
					query.Where(approle.AutoGrant(true))
				})
		}).OnlyX(ctx)
	// 创建应用策略
	ps, err := ap.Policies(ctx)
	if err != nil {
		return false, err
	}
	// 角色
	rs, err := ap.Roles(ctx)
	if err != nil {
		return false, err
	}
	c := ent.FromContext(ctx)
	pbk := make([]*ent.PermissionPolicyCreate, len(ps))
	for i, p := range ps {
		pbk[i] = c.PermissionPolicy.Create().SetOrgID(org.ID).SetAppID(ap.ID).SetAppPolicyID(p.ID).
			SetComments(p.Comments).SetRules(p.Rules).SetComments(p.Comments).SetName(p.Name)
	}
	err = c.PermissionPolicy.CreateBulk(pbk...).Exec(ctx)
	if err != nil {
		return false, err
	}
	rbk := make([]*ent.OrganizationRoleCreate, len(rs))
	for i, r := range rs {
		rbk[i] = c.OrganizationRole.Create().SetOrgID(org.ID).SetKind(organizationrole.KindRole).SetAppRoleID(r.ID).
			SetComments(r.Comments).SetName(r.Name)
	}
	err = c.OrganizationRole.CreateBulk(rbk...).Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

// assignOrganizationPolicy 分配应用权限策略到组织.
func (s *Service) assignAppPolicyToOrganization(ctx context.Context, orgID int, policyID int) (*ent.PermissionPolicy, error) {
	c := ent.FromContext(ctx)
	pm := c.AppPolicy.Query().Where(apppolicy.ID(policyID)).OnlyX(ctx)
	// 如应用存在,确定应用是否分配入组织
	// 对组织授权
	return c.PermissionPolicy.Create().SetOrgID(orgID).SetAppID(pm.AppID).SetAppPolicyID(pm.ID).
		SetComments(pm.Comments).
		Save(ctx)
}
