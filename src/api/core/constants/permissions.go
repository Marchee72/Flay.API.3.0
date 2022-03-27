package constants

import "fmt"

type UserType string

const (
	UserAdmin  UserType = "admin"
	UserRenter UserType = "renter"
	UserOwner  UserType = "owner"
)

func (userType UserType) Equals(t interface{}) bool {
	stringType := fmt.Sprintf("%v", t)
	return stringType == string(userType)
}
