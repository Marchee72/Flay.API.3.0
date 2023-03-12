package constants

import (
	"fmt"
	"strings"
)

type UserType string

const (
	UserAdmin  UserType = "ADMIN"
	UserRenter UserType = "RENTER"
	UserOwner  UserType = "OWNER"
)

func AnyUserType() []UserType {
	return []UserType{UserRenter, UserAdmin, UserOwner}
}

func (userType UserType) Equals(t interface{}) bool {
	stringType := fmt.Sprintf("%v", t)
	return strings.ToUpper(stringType) == string(userType)
}

func ToUserType(userType string) (*UserType, error) {
	userType = strings.ToUpper(userType)
	var parsedType UserType
	var err error
	switch userType {
	case string(UserAdmin):
		parsedType = UserAdmin
	case string(UserRenter):
		parsedType = UserRenter
	case string(UserOwner):
		parsedType = UserOwner
	default:
		err = fmt.Errorf("can not parse user type '%s'", userType)
	}
	if err != nil {
		return nil, err
	}
	return &parsedType, nil
}
