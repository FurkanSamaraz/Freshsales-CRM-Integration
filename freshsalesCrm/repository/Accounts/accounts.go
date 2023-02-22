package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Accounts"

	"gorm.io/gorm"
)

type AccountsDB struct {
	CrmCollection *gorm.DB
}
type AccountsRepositoryInterface interface {
	GetAccounts(model crm_structures.GetAccounts) (bool, error)
	CreateAccounts(model crm_structures.CreateAccounts) (bool, error)
	UpdateAccounts(model crm_structures.UpdateAccounts) (bool, error)
}

func (t AccountsDB) GetAccounts(model crm_structures.GetAccounts) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Accounts |")
	return true, nil
}
func (t AccountsDB) CreateAccounts(model crm_structures.CreateAccounts) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Accounts |")
	return true, nil
}
func (t AccountsDB) UpdateAccounts(model crm_structures.UpdateAccounts) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Accounts |")
	return true, nil
}
