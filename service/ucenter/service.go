package ucenter

import (
	"context"
	"fmt"
	"github.com/woocoos/adminx/ent"
	"github.com/woocoos/adminx/ent/organization"
	"strconv"
	"strings"
)

type Service struct {
	Client *ent.Client
}

func (uc *Service) CanEditOrgByLoginOrg(ctx context.Context, orgID int) (bool, error) {
	org, err := uc.Client.Organization.Query().Where(
		organization.IDEQ(orgID),
	).Only(context.Background())
	if err != nil {
		return false, err
	}
	// 查组织以及子组织
	tps := strings.Split(org.Path, "/")
	for _, tp := range tps {
		// gfc.LoginOrgID
		if tp == strconv.Itoa(0) {
			return true, nil
		}
	}
	return false, fmt.Errorf("没有操作组织权限")
}
