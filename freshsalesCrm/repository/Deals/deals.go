package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Deals"

	"gorm.io/gorm"
)

type DealsDB struct {
	CrmCollection *gorm.DB
}
type DealsRepositoryInterface interface {
	GetDeals(model crm_structures.GetDeals) (bool, error)
	CreateDeals(model crm_structures.CreateDeals) (bool, error)
}

func (t DealsDB) GetDeals(model crm_structures.GetDeals) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Deals |")
	return true, nil
}
func (t DealsDB) CreateDeals(model crm_structures.CreateDeals) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Deals |")
	return true, nil
}
