package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Files"
	crm_structures "main/structure/Files"
)

type Service_Files interface {
	FilesCreate(model crm_structures.CreateFiles) (*dto.Crm_DTO, error)
	FilesGet(model crm_structures.GetFiles) (*dto.Crm_DTO, error)
}
type FilesCRMService struct {
	Repo crm_repository.FilesRepositoryInterface
}

func (t FilesCRMService) FilesGet(model crm_structures.GetFiles) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Files |")
	result, err := t.Repo.GetFiles(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t FilesCRMService) FilesCreate(model crm_structures.CreateFiles) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Files |")
	result, err := t.Repo.CreateFiles(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
