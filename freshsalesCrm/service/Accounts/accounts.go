package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Accounts"
	crm_structures "main/structure/Accounts"
)

type AccountsCRMService struct {
	Repo crm_repository.AccountsRepositoryInterface
}

type Service_Accounts interface {
	AccountsCreate(model crm_structures.CreateAccounts) (*dto.Crm_DTO, error)
	AccountsGet(model crm_structures.GetAccounts) (*dto.Crm_DTO, error)
	AccountsUpdate(model crm_structures.UpdateAccounts) (*dto.Crm_DTO, error)
}

func (t AccountsCRMService) AccountsGet(model crm_structures.GetAccounts) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Accounts |")
	result, err := t.Repo.GetAccounts(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t AccountsCRMService) AccountsCreate(model crm_structures.CreateAccounts) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Accounts |")
	result, err := t.Repo.CreateAccounts(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t AccountsCRMService) AccountsUpdate(model crm_structures.UpdateAccounts) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Accounts |")
	result, err := t.Repo.UpdateAccounts(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
