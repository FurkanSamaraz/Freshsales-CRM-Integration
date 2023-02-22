package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Activities"
	crm_structures "main/structure/Activities"
)

type Service_Activities interface {
	ActivitiesGet(model crm_structures.GetActivities) (*dto.Crm_DTO, error)
}
type ActivitiesCRMService struct {
	Repo crm_repository.ActivitiesRepositoryInterface
}

func (t ActivitiesCRMService) ActivitiesGet(model crm_structures.GetActivities) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Activities |")
	result, err := t.Repo.GetActivities(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
