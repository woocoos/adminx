// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/woocoos/adminx/ent/app"
	"github.com/woocoos/adminx/ent/appaction"
	"github.com/woocoos/adminx/ent/appmenu"
	"github.com/woocoos/adminx/ent/apppolicy"
	"github.com/woocoos/adminx/ent/appres"
	"github.com/woocoos/adminx/ent/approle"
	"github.com/woocoos/adminx/ent/approlepolicy"
	"github.com/woocoos/adminx/ent/organization"
	"github.com/woocoos/adminx/ent/organizationapp"
	"github.com/woocoos/adminx/ent/organizationrole"
	"github.com/woocoos/adminx/ent/organizationuser"
	"github.com/woocoos/adminx/ent/permission"
	"github.com/woocoos/adminx/ent/permissionpolicy"
	"github.com/woocoos/adminx/ent/user"
	"github.com/woocoos/adminx/ent/userdevice"
	"github.com/woocoos/adminx/ent/useridentity"
	"github.com/woocoos/adminx/ent/userloginprofile"
	"github.com/woocoos/adminx/ent/userpassword"
	"github.com/woocoos/adminx/graph/entgen/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	appMixin := schema.App{}.Mixin()
	appMixinHooks1 := appMixin[1].Hooks()
	app.Hooks[0] = appMixinHooks1[0]
	appMixinFields0 := appMixin[0].Fields()
	_ = appMixinFields0
	appMixinFields1 := appMixin[1].Fields()
	_ = appMixinFields1
	appFields := schema.App{}.Fields()
	_ = appFields
	// appDescCreatedAt is the schema descriptor for created_at field.
	appDescCreatedAt := appMixinFields1[1].Descriptor()
	// app.DefaultCreatedAt holds the default value on creation for the created_at field.
	app.DefaultCreatedAt = appDescCreatedAt.Default.(func() time.Time)
	// appDescName is the schema descriptor for name field.
	appDescName := appFields[0].Descriptor()
	// app.NameValidator is a validator for the "name" field. It is called by the builders before save.
	app.NameValidator = appDescName.Validators[0].(func(string) error)
	// appDescCode is the schema descriptor for code field.
	appDescCode := appFields[1].Descriptor()
	// app.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	app.CodeValidator = appDescCode.Validators[0].(func(string) error)
	// appDescRedirectURI is the schema descriptor for redirect_uri field.
	appDescRedirectURI := appFields[3].Descriptor()
	// app.RedirectURIValidator is a validator for the "redirect_uri" field. It is called by the builders before save.
	app.RedirectURIValidator = appDescRedirectURI.Validators[0].(func(string) error)
	// appDescAppSecret is the schema descriptor for app_secret field.
	appDescAppSecret := appFields[5].Descriptor()
	// app.AppSecretValidator is a validator for the "app_secret" field. It is called by the builders before save.
	app.AppSecretValidator = appDescAppSecret.Validators[0].(func(string) error)
	// appDescScopes is the schema descriptor for scopes field.
	appDescScopes := appFields[6].Descriptor()
	// app.ScopesValidator is a validator for the "scopes" field. It is called by the builders before save.
	app.ScopesValidator = appDescScopes.Validators[0].(func(string) error)
	// appDescID is the schema descriptor for id field.
	appDescID := appMixinFields0[0].Descriptor()
	// app.DefaultID holds the default value on creation for the id field.
	app.DefaultID = appDescID.Default.(func() int)
	appactionMixin := schema.AppAction{}.Mixin()
	appactionMixinHooks1 := appactionMixin[1].Hooks()
	appactionHooks := schema.AppAction{}.Hooks()
	appaction.Hooks[0] = appactionMixinHooks1[0]
	appaction.Hooks[1] = appactionHooks[0]
	appactionMixinFields0 := appactionMixin[0].Fields()
	_ = appactionMixinFields0
	appactionMixinFields1 := appactionMixin[1].Fields()
	_ = appactionMixinFields1
	appactionFields := schema.AppAction{}.Fields()
	_ = appactionFields
	// appactionDescCreatedAt is the schema descriptor for created_at field.
	appactionDescCreatedAt := appactionMixinFields1[1].Descriptor()
	// appaction.DefaultCreatedAt holds the default value on creation for the created_at field.
	appaction.DefaultCreatedAt = appactionDescCreatedAt.Default.(func() time.Time)
	// appactionDescName is the schema descriptor for name field.
	appactionDescName := appactionFields[1].Descriptor()
	// appaction.NameValidator is a validator for the "name" field. It is called by the builders before save.
	appaction.NameValidator = appactionDescName.Validators[0].(func(string) error)
	// appactionDescID is the schema descriptor for id field.
	appactionDescID := appactionMixinFields0[0].Descriptor()
	// appaction.DefaultID holds the default value on creation for the id field.
	appaction.DefaultID = appactionDescID.Default.(func() int)
	appmenuMixin := schema.AppMenu{}.Mixin()
	appmenuMixinHooks1 := appmenuMixin[1].Hooks()
	appmenuHooks := schema.AppMenu{}.Hooks()
	appmenu.Hooks[0] = appmenuMixinHooks1[0]
	appmenu.Hooks[1] = appmenuHooks[0]
	appmenuMixinFields0 := appmenuMixin[0].Fields()
	_ = appmenuMixinFields0
	appmenuMixinFields1 := appmenuMixin[1].Fields()
	_ = appmenuMixinFields1
	appmenuFields := schema.AppMenu{}.Fields()
	_ = appmenuFields
	// appmenuDescCreatedAt is the schema descriptor for created_at field.
	appmenuDescCreatedAt := appmenuMixinFields1[1].Descriptor()
	// appmenu.DefaultCreatedAt holds the default value on creation for the created_at field.
	appmenu.DefaultCreatedAt = appmenuDescCreatedAt.Default.(func() time.Time)
	// appmenuDescName is the schema descriptor for name field.
	appmenuDescName := appmenuFields[3].Descriptor()
	// appmenu.NameValidator is a validator for the "name" field. It is called by the builders before save.
	appmenu.NameValidator = appmenuDescName.Validators[0].(func(string) error)
	// appmenuDescID is the schema descriptor for id field.
	appmenuDescID := appmenuMixinFields0[0].Descriptor()
	// appmenu.DefaultID holds the default value on creation for the id field.
	appmenu.DefaultID = appmenuDescID.Default.(func() int)
	apppolicyMixin := schema.AppPolicy{}.Mixin()
	apppolicyMixinHooks1 := apppolicyMixin[1].Hooks()
	apppolicy.Hooks[0] = apppolicyMixinHooks1[0]
	apppolicyMixinFields0 := apppolicyMixin[0].Fields()
	_ = apppolicyMixinFields0
	apppolicyMixinFields1 := apppolicyMixin[1].Fields()
	_ = apppolicyMixinFields1
	apppolicyFields := schema.AppPolicy{}.Fields()
	_ = apppolicyFields
	// apppolicyDescCreatedAt is the schema descriptor for created_at field.
	apppolicyDescCreatedAt := apppolicyMixinFields1[1].Descriptor()
	// apppolicy.DefaultCreatedAt holds the default value on creation for the created_at field.
	apppolicy.DefaultCreatedAt = apppolicyDescCreatedAt.Default.(func() time.Time)
	// apppolicyDescAutoGrant is the schema descriptor for auto_grant field.
	apppolicyDescAutoGrant := apppolicyFields[5].Descriptor()
	// apppolicy.DefaultAutoGrant holds the default value on creation for the auto_grant field.
	apppolicy.DefaultAutoGrant = apppolicyDescAutoGrant.Default.(bool)
	// apppolicyDescID is the schema descriptor for id field.
	apppolicyDescID := apppolicyMixinFields0[0].Descriptor()
	// apppolicy.DefaultID holds the default value on creation for the id field.
	apppolicy.DefaultID = apppolicyDescID.Default.(func() int)
	appresMixin := schema.AppRes{}.Mixin()
	appresMixinHooks1 := appresMixin[1].Hooks()
	appres.Hooks[0] = appresMixinHooks1[0]
	appresMixinFields0 := appresMixin[0].Fields()
	_ = appresMixinFields0
	appresMixinFields1 := appresMixin[1].Fields()
	_ = appresMixinFields1
	appresFields := schema.AppRes{}.Fields()
	_ = appresFields
	// appresDescCreatedAt is the schema descriptor for created_at field.
	appresDescCreatedAt := appresMixinFields1[1].Descriptor()
	// appres.DefaultCreatedAt holds the default value on creation for the created_at field.
	appres.DefaultCreatedAt = appresDescCreatedAt.Default.(func() time.Time)
	// appresDescID is the schema descriptor for id field.
	appresDescID := appresMixinFields0[0].Descriptor()
	// appres.DefaultID holds the default value on creation for the id field.
	appres.DefaultID = appresDescID.Default.(func() int)
	approleMixin := schema.AppRole{}.Mixin()
	approleMixinHooks1 := approleMixin[1].Hooks()
	approle.Hooks[0] = approleMixinHooks1[0]
	approleMixinFields0 := approleMixin[0].Fields()
	_ = approleMixinFields0
	approleMixinFields1 := approleMixin[1].Fields()
	_ = approleMixinFields1
	approleFields := schema.AppRole{}.Fields()
	_ = approleFields
	// approleDescCreatedAt is the schema descriptor for created_at field.
	approleDescCreatedAt := approleMixinFields1[1].Descriptor()
	// approle.DefaultCreatedAt holds the default value on creation for the created_at field.
	approle.DefaultCreatedAt = approleDescCreatedAt.Default.(func() time.Time)
	// approleDescAutoGrant is the schema descriptor for auto_grant field.
	approleDescAutoGrant := approleFields[3].Descriptor()
	// approle.DefaultAutoGrant holds the default value on creation for the auto_grant field.
	approle.DefaultAutoGrant = approleDescAutoGrant.Default.(bool)
	// approleDescEditable is the schema descriptor for editable field.
	approleDescEditable := approleFields[4].Descriptor()
	// approle.DefaultEditable holds the default value on creation for the editable field.
	approle.DefaultEditable = approleDescEditable.Default.(bool)
	// approleDescID is the schema descriptor for id field.
	approleDescID := approleMixinFields0[0].Descriptor()
	// approle.DefaultID holds the default value on creation for the id field.
	approle.DefaultID = approleDescID.Default.(func() int)
	approlepolicyMixin := schema.AppRolePolicy{}.Mixin()
	approlepolicyMixinHooks0 := approlepolicyMixin[0].Hooks()
	approlepolicy.Hooks[0] = approlepolicyMixinHooks0[0]
	approlepolicyMixinFields0 := approlepolicyMixin[0].Fields()
	_ = approlepolicyMixinFields0
	approlepolicyFields := schema.AppRolePolicy{}.Fields()
	_ = approlepolicyFields
	// approlepolicyDescCreatedAt is the schema descriptor for created_at field.
	approlepolicyDescCreatedAt := approlepolicyMixinFields0[1].Descriptor()
	// approlepolicy.DefaultCreatedAt holds the default value on creation for the created_at field.
	approlepolicy.DefaultCreatedAt = approlepolicyDescCreatedAt.Default.(func() time.Time)
	organizationMixin := schema.Organization{}.Mixin()
	organizationMixinHooks1 := organizationMixin[1].Hooks()
	organizationMixinHooks2 := organizationMixin[2].Hooks()
	organizationHooks := schema.Organization{}.Hooks()
	organization.Hooks[0] = organizationMixinHooks1[0]
	organization.Hooks[1] = organizationMixinHooks2[0]
	organization.Hooks[2] = organizationHooks[0]
	organization.Hooks[3] = organizationHooks[1]
	organization.Hooks[4] = organizationHooks[2]
	organizationMixinInters2 := organizationMixin[2].Interceptors()
	organization.Interceptors[0] = organizationMixinInters2[0]
	organizationMixinFields0 := organizationMixin[0].Fields()
	_ = organizationMixinFields0
	organizationMixinFields1 := organizationMixin[1].Fields()
	_ = organizationMixinFields1
	organizationFields := schema.Organization{}.Fields()
	_ = organizationFields
	// organizationDescCreatedAt is the schema descriptor for created_at field.
	organizationDescCreatedAt := organizationMixinFields1[1].Descriptor()
	// organization.DefaultCreatedAt holds the default value on creation for the created_at field.
	organization.DefaultCreatedAt = organizationDescCreatedAt.Default.(func() time.Time)
	// organizationDescParentID is the schema descriptor for parent_id field.
	organizationDescParentID := organizationFields[2].Descriptor()
	// organization.DefaultParentID holds the default value on creation for the parent_id field.
	organization.DefaultParentID = organizationDescParentID.Default.(int)
	// organizationDescDomain is the schema descriptor for domain field.
	organizationDescDomain := organizationFields[3].Descriptor()
	// organization.DomainValidator is a validator for the "domain" field. It is called by the builders before save.
	organization.DomainValidator = organizationDescDomain.Validators[0].(func(string) error)
	// organizationDescCode is the schema descriptor for code field.
	organizationDescCode := organizationFields[4].Descriptor()
	// organization.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	organization.CodeValidator = organizationDescCode.Validators[0].(func(string) error)
	// organizationDescName is the schema descriptor for name field.
	organizationDescName := organizationFields[5].Descriptor()
	// organization.NameValidator is a validator for the "name" field. It is called by the builders before save.
	organization.NameValidator = organizationDescName.Validators[0].(func(string) error)
	// organizationDescCountryCode is the schema descriptor for country_code field.
	organizationDescCountryCode := organizationFields[10].Descriptor()
	// organization.CountryCodeValidator is a validator for the "country_code" field. It is called by the builders before save.
	organization.CountryCodeValidator = organizationDescCountryCode.Validators[0].(func(string) error)
	// organizationDescTimezone is the schema descriptor for timezone field.
	organizationDescTimezone := organizationFields[11].Descriptor()
	// organization.TimezoneValidator is a validator for the "timezone" field. It is called by the builders before save.
	organization.TimezoneValidator = organizationDescTimezone.Validators[0].(func(string) error)
	// organizationDescID is the schema descriptor for id field.
	organizationDescID := organizationMixinFields0[0].Descriptor()
	// organization.DefaultID holds the default value on creation for the id field.
	organization.DefaultID = organizationDescID.Default.(func() int)
	organizationappMixin := schema.OrganizationApp{}.Mixin()
	organizationappMixinHooks0 := organizationappMixin[0].Hooks()
	organizationapp.Hooks[0] = organizationappMixinHooks0[0]
	organizationappMixinFields0 := organizationappMixin[0].Fields()
	_ = organizationappMixinFields0
	organizationappFields := schema.OrganizationApp{}.Fields()
	_ = organizationappFields
	// organizationappDescCreatedAt is the schema descriptor for created_at field.
	organizationappDescCreatedAt := organizationappMixinFields0[1].Descriptor()
	// organizationapp.DefaultCreatedAt holds the default value on creation for the created_at field.
	organizationapp.DefaultCreatedAt = organizationappDescCreatedAt.Default.(func() time.Time)
	organizationroleMixin := schema.OrganizationRole{}.Mixin()
	organizationroleMixinHooks0 := organizationroleMixin[0].Hooks()
	organizationrole.Hooks[0] = organizationroleMixinHooks0[0]
	organizationroleMixinFields0 := organizationroleMixin[0].Fields()
	_ = organizationroleMixinFields0
	organizationroleFields := schema.OrganizationRole{}.Fields()
	_ = organizationroleFields
	// organizationroleDescCreatedAt is the schema descriptor for created_at field.
	organizationroleDescCreatedAt := organizationroleMixinFields0[1].Descriptor()
	// organizationrole.DefaultCreatedAt holds the default value on creation for the created_at field.
	organizationrole.DefaultCreatedAt = organizationroleDescCreatedAt.Default.(func() time.Time)
	organizationuserMixin := schema.OrganizationUser{}.Mixin()
	organizationuserMixinHooks1 := organizationuserMixin[1].Hooks()
	organizationuser.Hooks[0] = organizationuserMixinHooks1[0]
	organizationuserMixinFields1 := organizationuserMixin[1].Fields()
	_ = organizationuserMixinFields1
	organizationuserFields := schema.OrganizationUser{}.Fields()
	_ = organizationuserFields
	// organizationuserDescCreatedAt is the schema descriptor for created_at field.
	organizationuserDescCreatedAt := organizationuserMixinFields1[1].Descriptor()
	// organizationuser.DefaultCreatedAt holds the default value on creation for the created_at field.
	organizationuser.DefaultCreatedAt = organizationuserDescCreatedAt.Default.(func() time.Time)
	permissionMixin := schema.Permission{}.Mixin()
	permissionMixinHooks1 := permissionMixin[1].Hooks()
	permission.Hooks[0] = permissionMixinHooks1[0]
	permissionMixinFields0 := permissionMixin[0].Fields()
	_ = permissionMixinFields0
	permissionMixinFields1 := permissionMixin[1].Fields()
	_ = permissionMixinFields1
	permissionFields := schema.Permission{}.Fields()
	_ = permissionFields
	// permissionDescCreatedAt is the schema descriptor for created_at field.
	permissionDescCreatedAt := permissionMixinFields1[1].Descriptor()
	// permission.DefaultCreatedAt holds the default value on creation for the created_at field.
	permission.DefaultCreatedAt = permissionDescCreatedAt.Default.(func() time.Time)
	// permissionDescID is the schema descriptor for id field.
	permissionDescID := permissionMixinFields0[0].Descriptor()
	// permission.DefaultID holds the default value on creation for the id field.
	permission.DefaultID = permissionDescID.Default.(func() int)
	permissionpolicyMixin := schema.PermissionPolicy{}.Mixin()
	permissionpolicyMixinHooks1 := permissionpolicyMixin[1].Hooks()
	permissionpolicy.Hooks[0] = permissionpolicyMixinHooks1[0]
	permissionpolicyMixinFields0 := permissionpolicyMixin[0].Fields()
	_ = permissionpolicyMixinFields0
	permissionpolicyMixinFields1 := permissionpolicyMixin[1].Fields()
	_ = permissionpolicyMixinFields1
	permissionpolicyFields := schema.PermissionPolicy{}.Fields()
	_ = permissionpolicyFields
	// permissionpolicyDescCreatedAt is the schema descriptor for created_at field.
	permissionpolicyDescCreatedAt := permissionpolicyMixinFields1[1].Descriptor()
	// permissionpolicy.DefaultCreatedAt holds the default value on creation for the created_at field.
	permissionpolicy.DefaultCreatedAt = permissionpolicyDescCreatedAt.Default.(func() time.Time)
	// permissionpolicyDescID is the schema descriptor for id field.
	permissionpolicyDescID := permissionpolicyMixinFields0[0].Descriptor()
	// permissionpolicy.DefaultID holds the default value on creation for the id field.
	permissionpolicy.DefaultID = permissionpolicyDescID.Default.(func() int)
	userMixin := schema.User{}.Mixin()
	userMixinHooks1 := userMixin[1].Hooks()
	userMixinHooks2 := userMixin[2].Hooks()
	user.Hooks[0] = userMixinHooks1[0]
	user.Hooks[1] = userMixinHooks2[0]
	userMixinInters2 := userMixin[2].Interceptors()
	user.Interceptors[0] = userMixinInters2[0]
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userMixinFields1 := userMixin[1].Fields()
	_ = userMixinFields1
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields1[1].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = func() func(string) error {
		validators := userDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescMobile is the schema descriptor for mobile field.
	userDescMobile := userFields[3].Descriptor()
	// user.MobileValidator is a validator for the "mobile" field. It is called by the builders before save.
	user.MobileValidator = userDescMobile.Validators[0].(func(string) error)
	// userDescRegisterIP is the schema descriptor for register_ip field.
	userDescRegisterIP := userFields[6].Descriptor()
	// user.RegisterIPValidator is a validator for the "register_ip" field. It is called by the builders before save.
	user.RegisterIPValidator = userDescRegisterIP.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields0[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() int)
	userdeviceMixin := schema.UserDevice{}.Mixin()
	userdeviceMixinHooks1 := userdeviceMixin[1].Hooks()
	userdevice.Hooks[0] = userdeviceMixinHooks1[0]
	userdeviceMixinFields1 := userdeviceMixin[1].Fields()
	_ = userdeviceMixinFields1
	userdeviceFields := schema.UserDevice{}.Fields()
	_ = userdeviceFields
	// userdeviceDescCreatedAt is the schema descriptor for created_at field.
	userdeviceDescCreatedAt := userdeviceMixinFields1[1].Descriptor()
	// userdevice.DefaultCreatedAt holds the default value on creation for the created_at field.
	userdevice.DefaultCreatedAt = userdeviceDescCreatedAt.Default.(func() time.Time)
	// userdeviceDescDeviceUID is the schema descriptor for device_uid field.
	userdeviceDescDeviceUID := userdeviceFields[1].Descriptor()
	// userdevice.DeviceUIDValidator is a validator for the "device_uid" field. It is called by the builders before save.
	userdevice.DeviceUIDValidator = userdeviceDescDeviceUID.Validators[0].(func(string) error)
	// userdeviceDescDeviceName is the schema descriptor for device_name field.
	userdeviceDescDeviceName := userdeviceFields[2].Descriptor()
	// userdevice.DeviceNameValidator is a validator for the "device_name" field. It is called by the builders before save.
	userdevice.DeviceNameValidator = userdeviceDescDeviceName.Validators[0].(func(string) error)
	// userdeviceDescSystemName is the schema descriptor for system_name field.
	userdeviceDescSystemName := userdeviceFields[3].Descriptor()
	// userdevice.SystemNameValidator is a validator for the "system_name" field. It is called by the builders before save.
	userdevice.SystemNameValidator = userdeviceDescSystemName.Validators[0].(func(string) error)
	// userdeviceDescSystemVersion is the schema descriptor for system_version field.
	userdeviceDescSystemVersion := userdeviceFields[4].Descriptor()
	// userdevice.SystemVersionValidator is a validator for the "system_version" field. It is called by the builders before save.
	userdevice.SystemVersionValidator = userdeviceDescSystemVersion.Validators[0].(func(string) error)
	// userdeviceDescAppVersion is the schema descriptor for app_version field.
	userdeviceDescAppVersion := userdeviceFields[5].Descriptor()
	// userdevice.AppVersionValidator is a validator for the "app_version" field. It is called by the builders before save.
	userdevice.AppVersionValidator = userdeviceDescAppVersion.Validators[0].(func(string) error)
	// userdeviceDescDeviceModel is the schema descriptor for device_model field.
	userdeviceDescDeviceModel := userdeviceFields[6].Descriptor()
	// userdevice.DeviceModelValidator is a validator for the "device_model" field. It is called by the builders before save.
	userdevice.DeviceModelValidator = userdeviceDescDeviceModel.Validators[0].(func(string) error)
	useridentityMixin := schema.UserIdentity{}.Mixin()
	useridentityMixinHooks1 := useridentityMixin[1].Hooks()
	useridentity.Hooks[0] = useridentityMixinHooks1[0]
	useridentityMixinFields1 := useridentityMixin[1].Fields()
	_ = useridentityMixinFields1
	useridentityFields := schema.UserIdentity{}.Fields()
	_ = useridentityFields
	// useridentityDescCreatedAt is the schema descriptor for created_at field.
	useridentityDescCreatedAt := useridentityMixinFields1[1].Descriptor()
	// useridentity.DefaultCreatedAt holds the default value on creation for the created_at field.
	useridentity.DefaultCreatedAt = useridentityDescCreatedAt.Default.(func() time.Time)
	userloginprofileMixin := schema.UserLoginProfile{}.Mixin()
	userloginprofileMixinHooks1 := userloginprofileMixin[1].Hooks()
	userloginprofile.Hooks[0] = userloginprofileMixinHooks1[0]
	userloginprofileMixinFields1 := userloginprofileMixin[1].Fields()
	_ = userloginprofileMixinFields1
	userloginprofileFields := schema.UserLoginProfile{}.Fields()
	_ = userloginprofileFields
	// userloginprofileDescCreatedAt is the schema descriptor for created_at field.
	userloginprofileDescCreatedAt := userloginprofileMixinFields1[1].Descriptor()
	// userloginprofile.DefaultCreatedAt holds the default value on creation for the created_at field.
	userloginprofile.DefaultCreatedAt = userloginprofileDescCreatedAt.Default.(func() time.Time)
	// userloginprofileDescMfaSecret is the schema descriptor for mfa_secret field.
	userloginprofileDescMfaSecret := userloginprofileFields[8].Descriptor()
	// userloginprofile.MfaSecretValidator is a validator for the "mfa_secret" field. It is called by the builders before save.
	userloginprofile.MfaSecretValidator = userloginprofileDescMfaSecret.Validators[0].(func(string) error)
	userpasswordMixin := schema.UserPassword{}.Mixin()
	userpasswordMixinHooks1 := userpasswordMixin[1].Hooks()
	userpassword.Hooks[0] = userpasswordMixinHooks1[0]
	userpasswordMixinFields1 := userpasswordMixin[1].Fields()
	_ = userpasswordMixinFields1
	userpasswordFields := schema.UserPassword{}.Fields()
	_ = userpasswordFields
	// userpasswordDescCreatedAt is the schema descriptor for created_at field.
	userpasswordDescCreatedAt := userpasswordMixinFields1[1].Descriptor()
	// userpassword.DefaultCreatedAt holds the default value on creation for the created_at field.
	userpassword.DefaultCreatedAt = userpasswordDescCreatedAt.Default.(func() time.Time)
	// userpasswordDescPassword is the schema descriptor for password field.
	userpasswordDescPassword := userpasswordFields[2].Descriptor()
	// userpassword.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	userpassword.PasswordValidator = userpasswordDescPassword.Validators[0].(func(string) error)
	// userpasswordDescSalt is the schema descriptor for salt field.
	userpasswordDescSalt := userpasswordFields[3].Descriptor()
	// userpassword.SaltValidator is a validator for the "salt" field. It is called by the builders before save.
	userpassword.SaltValidator = userpasswordDescSalt.Validators[0].(func(string) error)
}

const (
	Version = "v0.11.8"                                         // Version of ent codegen.
	Sum     = "h1:M/M0QL1CYCUSdqGRXUrXhFYSDRJPsOOrr+RLEej/gyQ=" // Sum of ent codegen.
)
