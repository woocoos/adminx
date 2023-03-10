package resource

import (
	"ariga.io/entcache"
	"context"
	"fmt"
	"github.com/woocoos/adminx/ent"
	"github.com/woocoos/adminx/ent/organization"
	"github.com/woocoos/adminx/ent/user"
	"github.com/woocoos/adminx/ent/useridentity"
	"github.com/woocoos/adminx/graph/entgen/types"
	"github.com/woocoos/adminx/graph/model"
	"github.com/woocoos/adminx/security"
	"strings"
)

// Service 企业目录服务管理
type Service struct {
	Client *ent.Client
}

// EnableOrganization 开启组织目录
func (s *Service) EnableOrganization(ctx context.Context, input model.EnableDirectoryInput) (*ent.Organization, error) {
	uid := security.UserIDFromContext(ctx)

	exist, err := s.Client.Organization.Query().Where(organization.OwnerID(uid)).Exist(entcache.Evict(ctx))
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("directory service has enable")
	}
	bulk := make([]*ent.OrganizationCreate, 1)
	bulk[0] = s.Client.Organization.Create().SetOwnerID(uid).SetName(input.Name).SetDomain(input.Domain).
		SetKind(organization.KindRoot).SetStatus(types.SimpleStatusActive)
	os, err := s.Client.Organization.CreateBulk(bulk...).Save(ctx)
	return os[0], err
}

// CreateOrganization 创建组织目录
func (s *Service) CreateOrganization(ctx context.Context, input ent.CreateOrganizationInput) (*ent.Organization, error) {
	return s.Client.Organization.Create().SetInput(input).Save(ctx)
}

// UpdateOrganization 更新组织目录
//
// - 如果更新的组织目录的管理账号，那么指向的用户必须是账户类型用户
func (s *Service) UpdateOrganization(ctx context.Context, id int, input ent.UpdateOrganizationInput) (*ent.Organization, error) {
	if input.OwnerID != nil {
		u := s.Client.User.Query().Where(user.ID(*input.OwnerID)).Select(user.FieldUserType).FirstX(ctx)
		if u.UserType != user.UserTypeAccount {
			return nil, fmt.Errorf("owner must be account")
		}
	}
	return s.Client.Organization.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteOrganization 删除组织目录
func (s *Service) DeleteOrganization(ctx context.Context, id int) error {
	count, err := s.Client.Organization.Query().Where(
		organization.ParentID(id),
	).Count(ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("organization has children")
	}
	count, err = s.Client.Organization.Query().Where(organization.ID(id), organization.HasUsers()).Count(ctx)
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("organization has users")
	}
	return s.Client.Organization.DeleteOneID(id).Exec(ctx)
}

// CreateOrganizationAccount 创建组织目录账户,进入账户激活流程
//
// - 管理员账户才能创建下级组织目录的账户
func (s *Service) CreateOrganizationAccount(ctx context.Context, input model.CreateOrganizationAccountInput) (*ent.User, error) {
	uid := security.UserIDFromContext(ctx)
	morg := s.Client.Organization.Query().Where(organization.OwnerID(uid)).FirstX(ctx)
	porg := s.Client.Organization.Query().Where(organization.ID(input.OrgID)).FirstX(ctx)
	if strings.HasPrefix(porg.Path, morg.Path) {
		return nil, fmt.Errorf("organization not found")
	}

	err := s.Client.OrganizationUser.Create().SetOrganizationID(input.OrgID).SetUserID(uid).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return s.Client.User.Create().SetInput(ent.CreateUserInput{
		PrincipalName: input.PrincipalName,
		DisplayName:   input.DisplayName,
	}).SetUserType(user.UserTypeAccount).SetCreationType(user.CreationTypeManual).SetStatus(types.SimpleStatusInactive).Save(ctx)
}

// CreateOrganizationUser 创建组织目录用户
//
// TODO 新用户需要激活,如在国内,用户往往需要绑定手机或邮箱,然后通过邮件或短信激活.
func (s *Service) CreateOrganizationUser(ctx context.Context, orgId int, input ent.CreateUserInput) (*ent.User, error) {
	_, err := s.Client.Organization.Query().Where(organization.ID(orgId), organization.StatusEQ(types.SimpleStatusActive)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("organization not exists or inactive")
	}
	client := ent.FromContext(ctx)
	us, err := client.User.Create().SetID(*input.ID).SetInput(input).
		SetUserType(user.UserTypeMember).
		SetCreationType(user.CreationTypeManual).
		SetRegisterIP("").
		Save(ctx)
	if err != nil {
		return nil, err
	}
	_, err = client.UserIdentity.Create().SetUserID(us.ID).SetCode(input.PrincipalName).SetKind(useridentity.KindName).
		SetStatus(types.SimpleStatusActive).Save(ctx)
	if err != nil {
		return nil, err
	}

	_, err = client.OrganizationUser.Create().SetOrganizationID(orgId).SetUserID(us.ID).SetDisplayName(us.DisplayName).Save(ctx)
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (s *Service) CreateUserPassword(ctx context.Context, input *ent.CreateUserPasswordInput) (pw *ent.UserPassword, err error) {
	if input == nil {
		return
	}
	salt := RandomStrByLetterAndDigit(5)
	var hashPwd string
	if input.Password != nil || *input.Password != "" {
		hashPwd = SHA256(*input.Password + salt)
	} else {
		hashPwd = RandomStrByLetterAndDigit(6)
		hashPwd = SHA256(hashPwd + salt)
	}
	pw, err = ent.FromContext(ctx).UserPassword.Create().
		SetInput(*input).
		SetPassword(hashPwd).
		SetSalt(salt).
		Save(ctx)

	return
}

// DeleteApp 删除应用
//
// 该应用没有被关联才可以删除: 1. 没有组织目录关联 2. 没有用户关联
func (s *Service) DeleteApp() (bool, error) {
	panic("implement me")
}
