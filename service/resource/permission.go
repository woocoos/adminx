package resource

import (
	"context"
	"fmt"
	"github.com/woocoos/adminx/ent"
	"github.com/woocoos/adminx/ent/app"
	"github.com/woocoos/adminx/ent/apppolicy"
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

func (s *Service) assignOrganizationApp(ctx context.Context, orgID int, appID int) (bool, error) {
	return false, nil
}

// assignOrganizationPolicy 分配应用权限策略到组织.
func (s *Service) assignAppPolicyToOrganization(ctx context.Context, orgID int, policyID int) (*ent.OrganizationPolicy, error) {
	c := ent.FromContext(ctx)
	pm := c.AppPolicy.Query().Where(apppolicy.ID(policyID)).OnlyX(ctx)
	// 如应用存在,确定应用是否分配入组织
	// 对组织授权
	return c.OrganizationPolicy.Create().SetOrgID(orgID).SetAppID(pm.AppID).SetAppPolicyID(pm.ID).
		SetComments(pm.Comments).
		Save(ctx)
}
