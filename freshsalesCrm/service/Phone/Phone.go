package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Phone"
	crm_structures "main/structure/Phone"
)

type Service_Phone interface {
	PhoneCreate(model crm_structures.CreatePhone) (*dto.Crm_DTO, error)
}
type PhoneCRMService struct {
	Repo crm_repository.PhoneRepositoryInterface
}

func (t PhoneCRMService) PhoneCreate(model crm_structures.CreatePhone) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Phone |")
	result, err := t.Repo.CreatePhone(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
