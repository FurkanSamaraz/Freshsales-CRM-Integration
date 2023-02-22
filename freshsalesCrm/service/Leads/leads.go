package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Leads"
	crm_structures "main/structure/Leads"
)

type Service_Leads interface {
	LeadsCreate(model crm_structures.CreateLeads) (*dto.Crm_DTO, error)
	LeadsGet(model crm_structures.GetLeads) (*dto.Crm_DTO, error)
	LeadsUpdate(model crm_structures.UpdateLeads) (*dto.Crm_DTO, error)
}
type LeadsCRMService struct {
	Repo crm_repository.LeadsRepositoryInterface
}

func (t LeadsCRMService) LeadsGet(model crm_structures.GetLeads) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Leads |")
	result, err := t.Repo.GetLeads(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t LeadsCRMService) LeadsCreate(model crm_structures.CreateLeads) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Leads |")
	result, err := t.Repo.CreateLeads(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t LeadsCRMService) LeadsUpdate(model crm_structures.UpdateLeads) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Leads |")
	result, err := t.Repo.UpdateLeads(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
