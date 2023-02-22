package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Leads"

	"gorm.io/gorm"
)

type LeadsDB struct {
	CrmCollection *gorm.DB
}
type LeadsRepositoryInterface interface {
	GetLeads(model crm_structures.GetLeads) (bool, error)
	CreateLeads(model crm_structures.CreateLeads) (bool, error)
	UpdateLeads(model crm_structures.UpdateLeads) (bool, error)
}

func (t LeadsDB) GetLeads(model crm_structures.GetLeads) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Leads |")
	return true, nil
}
func (t LeadsDB) CreateLeads(model crm_structures.CreateLeads) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Leads |")
	return true, nil
}
func (t LeadsDB) UpdateLeads(model crm_structures.UpdateLeads) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Leads |")
	return true, nil
}
