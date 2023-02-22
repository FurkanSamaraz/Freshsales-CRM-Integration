package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Phone"

	"gorm.io/gorm"
)

type PhoneDB struct {
	CrmCollection *gorm.DB
}
type PhoneRepositoryInterface interface {
	CreatePhone(model crm_structures.CreatePhone) (bool, error)
}

func (t PhoneDB) CreatePhone(model crm_structures.CreatePhone) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Phone |")
	return true, nil
}
