package crm_repository

import (
	"fmt"
	crm_structures "main/structure/SalesActivities"

	"gorm.io/gorm"
)

type SalesActivitiesDB struct {
	CrmCollection *gorm.DB
}
type SalesActivitiesRepositoryInterface interface {
	GetSalesActivities(model crm_structures.GetSalesActivities) (bool, error)
	CreateSalesActivities(model crm_structures.CreateSalesActivities) (bool, error)
	UpdateSalesActivities(model crm_structures.UpdateSalesActivities) (bool, error)
}

func (t SalesActivitiesDB) GetSalesActivities(model crm_structures.GetSalesActivities) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - SalesActivities |")
	return true, nil
}
func (t SalesActivitiesDB) CreateSalesActivities(model crm_structures.CreateSalesActivities) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - SalesActivities |")
	return true, nil
}
func (t SalesActivitiesDB) UpdateSalesActivities(model crm_structures.UpdateSalesActivities) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - SalesActivities |")
	return true, nil
}
