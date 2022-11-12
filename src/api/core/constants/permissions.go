package constants

import (
	"fmt"
)

type UserType string

const (
	UserAdmin  UserType = "admin"
	UserRenter UserType = "renter"
	UserOwner  UserType = "owner"
)

func AnyUserType() []UserType {
	return []UserType{UserRenter, UserAdmin, UserOwner}
}

func (userType UserType) Equals(t interface{}) bool {
	stringType := fmt.Sprintf("%v", t)
	return stringType == string(userType)
}

func ToUserType(userType string) (*UserType, error) {
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
