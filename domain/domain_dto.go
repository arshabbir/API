package domain

//DB connectivity and quries functionality will be executed in this area

type UserDTO interface {
	GetUser(id int) User
}

type user struct{}

var userStruct user
var UserInter UserDTO

func init() {

	UserInter = userStruct
}

func (u user) GetUser(id int) User {

	return User{Age: 30, HomePage: "www.google.com", Name: "arshabbir"}

}
