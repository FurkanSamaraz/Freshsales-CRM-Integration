package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/SalesActivities"
	crm_structures "main/structure/SalesActivities"
)

type Service_SalesActivities interface {
	SalesActivitiesCreate(model crm_structures.CreateSalesActivities) (*dto.Crm_DTO, error)
	SalesActivitiesGet(model crm_structures.GetSalesActivities) (*dto.Crm_DTO, error)
	SalesActivitiesUpdate(model crm_structures.UpdateSalesActivities) (*dto.Crm_DTO, error)
}
type SalesActivitiesCRMService struct {
	Repo crm_repository.SalesActivitiesRepositoryInterface
}

func (t SalesActivitiesCRMService) SalesActivitiesGet(model crm_structures.GetSalesActivities) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - SalesActivities |")
	result, err := t.Repo.GetSalesActivities(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t SalesActivitiesCRMService) SalesActivitiesCreate(model crm_structures.CreateSalesActivities) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - SalesActivities |")
	result, err := t.Repo.CreateSalesActivities(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t SalesActivitiesCRMService) SalesActivitiesUpdate(model crm_structures.UpdateSalesActivities) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - SalesActivities |")
	result, err := t.Repo.UpdateSalesActivities(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
