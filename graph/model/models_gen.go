// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/woocoos/adminx/ent"
)

type CreateOrganizationAccountInput struct {
	// 账号名称,组织内不可重复
	DisplayName string `json:"displayName"`
	// 账号登录名(含默认域名)
	PrincipalName string `json:"principalName"`
	// 邮箱
	Email string `json:"email"`
	// 所属组织ID
	OrgID int `json:"orgId"`
}

type EnableDirectoryInput struct {
	// 域名
	Domain string `json:"domain"`
	Name   string `json:"name"`
}

// Ordering options for OrganizationUser connections
type OrganizationUserOrder struct {
	// The ordering direction.
	Direction ent.OrderDirection `json:"direction"`
	// The field by which to order OrganizationUsers.
	Field OrganizationUserOrderField `json:"field"`
}

// Properties by which OrganizationUser connections can be ordered.
type OrganizationUserOrderField string

const (
	OrganizationUserOrderFieldCreatedAt OrganizationUserOrderField = "createdAt"
)

var AllOrganizationUserOrderField = []OrganizationUserOrderField{
	OrganizationUserOrderFieldCreatedAt,
}

func (e OrganizationUserOrderField) IsValid() bool {
	switch e {
	case OrganizationUserOrderFieldCreatedAt:
		return true
	}
	return false
}

func (e OrganizationUserOrderField) String() string {
	return string(e)
}

func (e *OrganizationUserOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrganizationUserOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrganizationUserOrderField", str)
	}
	return nil
}

func (e OrganizationUserOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}