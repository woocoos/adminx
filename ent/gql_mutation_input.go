// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/woocoos/adminx/ent/app"
	"github.com/woocoos/adminx/ent/organization"
	"github.com/woocoos/adminx/ent/user"
	"github.com/woocoos/adminx/ent/useridentity"
	"github.com/woocoos/adminx/ent/userloginprofile"
	"github.com/woocoos/adminx/ent/userpassword"
)

// CreateAppInput represents a mutation input for creating apps.
type CreateAppInput struct {
	ID                   *int
	Name                 *string
	Kind                 app.Kind
	RedirectURI          *string
	AppKey               *string
	AppSecret            *string
	Scopes               *string
	TokenValidity        *int32
	RefreshTokenValidity *int32
	Logo                 *string
	Comments             *string
	Status               *app.Status
	MenuIDs              []int
	PermissionIDs        []int
}

// Mutate applies the CreateAppInput on the AppMutation builder.
func (i *CreateAppInput) Mutate(m *AppMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	m.SetKind(i.Kind)
	if v := i.RedirectURI; v != nil {
		m.SetRedirectURI(*v)
	}
	if v := i.AppKey; v != nil {
		m.SetAppKey(*v)
	}
	if v := i.AppSecret; v != nil {
		m.SetAppSecret(*v)
	}
	if v := i.Scopes; v != nil {
		m.SetScopes(*v)
	}
	if v := i.TokenValidity; v != nil {
		m.SetTokenValidity(*v)
	}
	if v := i.RefreshTokenValidity; v != nil {
		m.SetRefreshTokenValidity(*v)
	}
	if v := i.Logo; v != nil {
		m.SetLogo(*v)
	}
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.MenuIDs; len(v) > 0 {
		m.AddMenuIDs(v...)
	}
	if v := i.PermissionIDs; len(v) > 0 {
		m.AddPermissionIDs(v...)
	}
}

// SetInput applies the change-set in the CreateAppInput on the AppCreate builder.
func (c *AppCreate) SetInput(i CreateAppInput) *AppCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateAppInput represents a mutation input for updating apps.
type UpdateAppInput struct {
	ID                        *int
	ClearName                 bool
	Name                      *string
	Kind                      *app.Kind
	ClearRedirectURI          bool
	RedirectURI               *string
	ClearAppKey               bool
	AppKey                    *string
	ClearAppSecret            bool
	AppSecret                 *string
	ClearScopes               bool
	Scopes                    *string
	ClearTokenValidity        bool
	TokenValidity             *int32
	ClearRefreshTokenValidity bool
	RefreshTokenValidity      *int32
	ClearLogo                 bool
	Logo                      *string
	ClearComments             bool
	Comments                  *string
	ClearStatus               bool
	Status                    *app.Status
	ClearMenus                bool
	AddMenuIDs                []int
	RemoveMenuIDs             []int
	ClearPermissions          bool
	AddPermissionIDs          []int
	RemovePermissionIDs       []int
}

// Mutate applies the UpdateAppInput on the AppMutation builder.
func (i *UpdateAppInput) Mutate(m *AppMutation) {
	if i.ClearName {
		m.ClearName()
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Kind; v != nil {
		m.SetKind(*v)
	}
	if i.ClearRedirectURI {
		m.ClearRedirectURI()
	}
	if v := i.RedirectURI; v != nil {
		m.SetRedirectURI(*v)
	}
	if i.ClearAppKey {
		m.ClearAppKey()
	}
	if v := i.AppKey; v != nil {
		m.SetAppKey(*v)
	}
	if i.ClearAppSecret {
		m.ClearAppSecret()
	}
	if v := i.AppSecret; v != nil {
		m.SetAppSecret(*v)
	}
	if i.ClearScopes {
		m.ClearScopes()
	}
	if v := i.Scopes; v != nil {
		m.SetScopes(*v)
	}
	if i.ClearTokenValidity {
		m.ClearTokenValidity()
	}
	if v := i.TokenValidity; v != nil {
		m.SetTokenValidity(*v)
	}
	if i.ClearRefreshTokenValidity {
		m.ClearRefreshTokenValidity()
	}
	if v := i.RefreshTokenValidity; v != nil {
		m.SetRefreshTokenValidity(*v)
	}
	if i.ClearLogo {
		m.ClearLogo()
	}
	if v := i.Logo; v != nil {
		m.SetLogo(*v)
	}
	if i.ClearComments {
		m.ClearComments()
	}
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	if i.ClearStatus {
		m.ClearStatus()
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if i.ClearMenus {
		m.ClearMenus()
	}
	if v := i.AddMenuIDs; len(v) > 0 {
		m.AddMenuIDs(v...)
	}
	if v := i.RemoveMenuIDs; len(v) > 0 {
		m.RemoveMenuIDs(v...)
	}
	if i.ClearPermissions {
		m.ClearPermissions()
	}
	if v := i.AddPermissionIDs; len(v) > 0 {
		m.AddPermissionIDs(v...)
	}
	if v := i.RemovePermissionIDs; len(v) > 0 {
		m.RemovePermissionIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateAppInput on the AppUpdate builder.
func (c *AppUpdate) SetInput(i UpdateAppInput) *AppUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateAppInput on the AppUpdateOne builder.
func (c *AppUpdateOne) SetInput(i UpdateAppInput) *AppUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateOrganizationInput represents a mutation input for creating organizations.
type CreateOrganizationInput struct {
	ID          *int
	Domain      *string
	Name        string
	Profile     *string
	Status      *organization.Status
	CountryCode *string
	Timezone    *string
	ParentID    int
	OwnerID     *int
}

// Mutate applies the CreateOrganizationInput on the OrganizationMutation builder.
func (i *CreateOrganizationInput) Mutate(m *OrganizationMutation) {
	if v := i.Domain; v != nil {
		m.SetDomain(*v)
	}
	m.SetName(i.Name)
	if v := i.Profile; v != nil {
		m.SetProfile(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.CountryCode; v != nil {
		m.SetCountryCode(*v)
	}
	if v := i.Timezone; v != nil {
		m.SetTimezone(*v)
	}
	m.SetParentID(i.ParentID)
	if v := i.OwnerID; v != nil {
		m.SetOwnerID(*v)
	}
}

// SetInput applies the change-set in the CreateOrganizationInput on the OrganizationCreate builder.
func (c *OrganizationCreate) SetInput(i CreateOrganizationInput) *OrganizationCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateOrganizationInput represents a mutation input for updating organizations.
type UpdateOrganizationInput struct {
	ID               *int
	ClearDomain      bool
	Domain           *string
	Name             *string
	ClearProfile     bool
	Profile          *string
	ClearStatus      bool
	Status           *organization.Status
	ClearCountryCode bool
	CountryCode      *string
	ClearTimezone    bool
	Timezone         *string
	ClearParent      bool
	ParentID         *int
	ClearOwner       bool
	OwnerID          *int
}

// Mutate applies the UpdateOrganizationInput on the OrganizationMutation builder.
func (i *UpdateOrganizationInput) Mutate(m *OrganizationMutation) {
	if i.ClearDomain {
		m.ClearDomain()
	}
	if v := i.Domain; v != nil {
		m.SetDomain(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearProfile {
		m.ClearProfile()
	}
	if v := i.Profile; v != nil {
		m.SetProfile(*v)
	}
	if i.ClearStatus {
		m.ClearStatus()
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if i.ClearCountryCode {
		m.ClearCountryCode()
	}
	if v := i.CountryCode; v != nil {
		m.SetCountryCode(*v)
	}
	if i.ClearTimezone {
		m.ClearTimezone()
	}
	if v := i.Timezone; v != nil {
		m.SetTimezone(*v)
	}
	if i.ClearParent {
		m.ClearParent()
	}
	if v := i.ParentID; v != nil {
		m.SetParentID(*v)
	}
	if i.ClearOwner {
		m.ClearOwner()
	}
	if v := i.OwnerID; v != nil {
		m.SetOwnerID(*v)
	}
}

// SetInput applies the change-set in the UpdateOrganizationInput on the OrganizationUpdate builder.
func (c *OrganizationUpdate) SetInput(i UpdateOrganizationInput) *OrganizationUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateOrganizationInput on the OrganizationUpdateOne builder.
func (c *OrganizationUpdateOne) SetInput(i UpdateOrganizationInput) *OrganizationUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	ID             *int
	PrincipalName  string
	DisplayName    string
	Status         *user.Status
	Comments       *string
	IdentityIDs    []int
	LoginProfileID *int
	PasswordIDs    []int
	DeviceIDs      []int
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	m.SetPrincipalName(i.PrincipalName)
	m.SetDisplayName(i.DisplayName)
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	if v := i.IdentityIDs; len(v) > 0 {
		m.AddIdentityIDs(v...)
	}
	if v := i.LoginProfileID; v != nil {
		m.SetLoginProfileID(*v)
	}
	if v := i.PasswordIDs; len(v) > 0 {
		m.AddPasswordIDs(v...)
	}
	if v := i.DeviceIDs; len(v) > 0 {
		m.AddDeviceIDs(v...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	ID                *int
	PrincipalName     *string
	DisplayName       *string
	ClearComments     bool
	Comments          *string
	ClearIdentities   bool
	AddIdentityIDs    []int
	RemoveIdentityIDs []int
	ClearLoginProfile bool
	LoginProfileID    *int
	ClearPasswords    bool
	AddPasswordIDs    []int
	RemovePasswordIDs []int
	ClearDevices      bool
	AddDeviceIDs      []int
	RemoveDeviceIDs   []int
}

// Mutate applies the UpdateUserInput on the UserMutation builder.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.PrincipalName; v != nil {
		m.SetPrincipalName(*v)
	}
	if v := i.DisplayName; v != nil {
		m.SetDisplayName(*v)
	}
	if i.ClearComments {
		m.ClearComments()
	}
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	if i.ClearIdentities {
		m.ClearIdentities()
	}
	if v := i.AddIdentityIDs; len(v) > 0 {
		m.AddIdentityIDs(v...)
	}
	if v := i.RemoveIdentityIDs; len(v) > 0 {
		m.RemoveIdentityIDs(v...)
	}
	if i.ClearLoginProfile {
		m.ClearLoginProfile()
	}
	if v := i.LoginProfileID; v != nil {
		m.SetLoginProfileID(*v)
	}
	if i.ClearPasswords {
		m.ClearPasswords()
	}
	if v := i.AddPasswordIDs; len(v) > 0 {
		m.AddPasswordIDs(v...)
	}
	if v := i.RemovePasswordIDs; len(v) > 0 {
		m.RemovePasswordIDs(v...)
	}
	if i.ClearDevices {
		m.ClearDevices()
	}
	if v := i.AddDeviceIDs; len(v) > 0 {
		m.AddDeviceIDs(v...)
	}
	if v := i.RemoveDeviceIDs; len(v) > 0 {
		m.RemoveDeviceIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdate builder.
func (c *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdateOne builder.
func (c *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserIdentityInput represents a mutation input for creating useridentities.
type CreateUserIdentityInput struct {
	ID         *int
	Kind       useridentity.Kind
	Code       *string
	CodeExtend *string
	Status     *useridentity.Status
	UserID     *int
}

// Mutate applies the CreateUserIdentityInput on the UserIdentityMutation builder.
func (i *CreateUserIdentityInput) Mutate(m *UserIdentityMutation) {
	m.SetKind(i.Kind)
	if v := i.Code; v != nil {
		m.SetCode(*v)
	}
	if v := i.CodeExtend; v != nil {
		m.SetCodeExtend(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
}

// SetInput applies the change-set in the CreateUserIdentityInput on the UserIdentityCreate builder.
func (c *UserIdentityCreate) SetInput(i CreateUserIdentityInput) *UserIdentityCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserIdentityInput represents a mutation input for updating useridentities.
type UpdateUserIdentityInput struct {
	ID              *int
	Kind            *useridentity.Kind
	ClearCode       bool
	Code            *string
	ClearCodeExtend bool
	CodeExtend      *string
	ClearStatus     bool
	Status          *useridentity.Status
}

// Mutate applies the UpdateUserIdentityInput on the UserIdentityMutation builder.
func (i *UpdateUserIdentityInput) Mutate(m *UserIdentityMutation) {
	if v := i.Kind; v != nil {
		m.SetKind(*v)
	}
	if i.ClearCode {
		m.ClearCode()
	}
	if v := i.Code; v != nil {
		m.SetCode(*v)
	}
	if i.ClearCodeExtend {
		m.ClearCodeExtend()
	}
	if v := i.CodeExtend; v != nil {
		m.SetCodeExtend(*v)
	}
	if i.ClearStatus {
		m.ClearStatus()
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
}

// SetInput applies the change-set in the UpdateUserIdentityInput on the UserIdentityUpdate builder.
func (c *UserIdentityUpdate) SetInput(i UpdateUserIdentityInput) *UserIdentityUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserIdentityInput on the UserIdentityUpdateOne builder.
func (c *UserIdentityUpdateOne) SetInput(i UpdateUserIdentityInput) *UserIdentityUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserLoginProfileInput represents a mutation input for creating userloginprofiles.
type CreateUserLoginProfileInput struct {
	ID            *int
	CanLogin      *bool
	SetKind       userloginprofile.SetKind
	PasswordReset *bool
	VerifyDevice  bool
	MfaEnabled    *bool
	MfaSecret     *string
	MfaStatus     *userloginprofile.MfaStatus
	UserID        *int
}

// Mutate applies the CreateUserLoginProfileInput on the UserLoginProfileMutation builder.
func (i *CreateUserLoginProfileInput) Mutate(m *UserLoginProfileMutation) {
	if v := i.CanLogin; v != nil {
		m.SetCanLogin(*v)
	}
	m.SetSetKind(i.SetKind)
	if v := i.PasswordReset; v != nil {
		m.SetPasswordReset(*v)
	}
	m.SetVerifyDevice(i.VerifyDevice)
	if v := i.MfaEnabled; v != nil {
		m.SetMfaEnabled(*v)
	}
	if v := i.MfaSecret; v != nil {
		m.SetMfaSecret(*v)
	}
	if v := i.MfaStatus; v != nil {
		m.SetMfaStatus(*v)
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
}

// SetInput applies the change-set in the CreateUserLoginProfileInput on the UserLoginProfileCreate builder.
func (c *UserLoginProfileCreate) SetInput(i CreateUserLoginProfileInput) *UserLoginProfileCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserLoginProfileInput represents a mutation input for updating userloginprofiles.
type UpdateUserLoginProfileInput struct {
	ID                 *int
	ClearCanLogin      bool
	CanLogin           *bool
	SetKind            *userloginprofile.SetKind
	ClearPasswordReset bool
	PasswordReset      *bool
	VerifyDevice       *bool
	ClearMfaEnabled    bool
	MfaEnabled         *bool
	ClearMfaSecret     bool
	MfaSecret          *string
	ClearMfaStatus     bool
	MfaStatus          *userloginprofile.MfaStatus
}

// Mutate applies the UpdateUserLoginProfileInput on the UserLoginProfileMutation builder.
func (i *UpdateUserLoginProfileInput) Mutate(m *UserLoginProfileMutation) {
	if i.ClearCanLogin {
		m.ClearCanLogin()
	}
	if v := i.CanLogin; v != nil {
		m.SetCanLogin(*v)
	}
	if v := i.SetKind; v != nil {
		m.SetSetKind(*v)
	}
	if i.ClearPasswordReset {
		m.ClearPasswordReset()
	}
	if v := i.PasswordReset; v != nil {
		m.SetPasswordReset(*v)
	}
	if v := i.VerifyDevice; v != nil {
		m.SetVerifyDevice(*v)
	}
	if i.ClearMfaEnabled {
		m.ClearMfaEnabled()
	}
	if v := i.MfaEnabled; v != nil {
		m.SetMfaEnabled(*v)
	}
	if i.ClearMfaSecret {
		m.ClearMfaSecret()
	}
	if v := i.MfaSecret; v != nil {
		m.SetMfaSecret(*v)
	}
	if i.ClearMfaStatus {
		m.ClearMfaStatus()
	}
	if v := i.MfaStatus; v != nil {
		m.SetMfaStatus(*v)
	}
}

// SetInput applies the change-set in the UpdateUserLoginProfileInput on the UserLoginProfileUpdate builder.
func (c *UserLoginProfileUpdate) SetInput(i UpdateUserLoginProfileInput) *UserLoginProfileUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserLoginProfileInput on the UserLoginProfileUpdateOne builder.
func (c *UserLoginProfileUpdateOne) SetInput(i UpdateUserLoginProfileInput) *UserLoginProfileUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserPasswordInput represents a mutation input for creating userpasswords.
type CreateUserPasswordInput struct {
	ID       *int
	Scene    *userpassword.Scene
	Password *string
	Status   *userpassword.Status
	Memo     *string
	UserID   *int
}

// Mutate applies the CreateUserPasswordInput on the UserPasswordMutation builder.
func (i *CreateUserPasswordInput) Mutate(m *UserPasswordMutation) {
	if v := i.Scene; v != nil {
		m.SetScene(*v)
	}
	if v := i.Password; v != nil {
		m.SetPassword(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.Memo; v != nil {
		m.SetMemo(*v)
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
}

// SetInput applies the change-set in the CreateUserPasswordInput on the UserPasswordCreate builder.
func (c *UserPasswordCreate) SetInput(i CreateUserPasswordInput) *UserPasswordCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserPasswordInput represents a mutation input for updating userpasswords.
type UpdateUserPasswordInput struct {
	ID            *int
	ClearScene    bool
	Scene         *userpassword.Scene
	ClearPassword bool
	Password      *string
	ClearStatus   bool
	Status        *userpassword.Status
	ClearMemo     bool
	Memo          *string
}

// Mutate applies the UpdateUserPasswordInput on the UserPasswordMutation builder.
func (i *UpdateUserPasswordInput) Mutate(m *UserPasswordMutation) {
	if i.ClearScene {
		m.ClearScene()
	}
	if v := i.Scene; v != nil {
		m.SetScene(*v)
	}
	if i.ClearPassword {
		m.ClearPassword()
	}
	if v := i.Password; v != nil {
		m.SetPassword(*v)
	}
	if i.ClearStatus {
		m.ClearStatus()
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if i.ClearMemo {
		m.ClearMemo()
	}
	if v := i.Memo; v != nil {
		m.SetMemo(*v)
	}
}

// SetInput applies the change-set in the UpdateUserPasswordInput on the UserPasswordUpdate builder.
func (c *UserPasswordUpdate) SetInput(i UpdateUserPasswordInput) *UserPasswordUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserPasswordInput on the UserPasswordUpdateOne builder.
func (c *UserPasswordUpdateOne) SetInput(i UpdateUserPasswordInput) *UserPasswordUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
