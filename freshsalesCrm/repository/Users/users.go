package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Users"

	"gorm.io/gorm"
)

type UsersDB struct {
	CrmCollection *gorm.DB
}
type UsersRepositoryInterface interface {
	GetUsers(model crm_structures.GetUsers) (bool, error)
}

func (t UsersDB) GetUsers(model crm_structures.GetUsers) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Users |")
	return true, nil
}
