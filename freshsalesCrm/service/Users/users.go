package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Users"
	crm_structures "main/structure/Users"
)

type Service_Users interface {
	UsersGet(model crm_structures.GetUsers) (*dto.Crm_DTO, error)
}
type UsersCRMService struct {
	Repo crm_repository.UsersRepositoryInterface
}

func (t UsersCRMService) UsersGet(model crm_structures.GetUsers) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Users |")
	result, err := t.Repo.GetUsers(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
