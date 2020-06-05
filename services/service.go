package services

import "master_go_programming/api/domain"

func GetUser(id int) domain.User {

	return domain.UserInter.GetUser(id)

}
