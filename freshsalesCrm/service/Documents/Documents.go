package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Documents"
	crm_structures "main/structure/Documents"
)

type Service_Documents interface {
	DocumentsCreate(model crm_structures.CreateDocuments) (*dto.Crm_DTO, error)
	DocumentsGet(model crm_structures.GetDocuments) (*dto.Crm_DTO, error)
	DocumentsUpdate(model crm_structures.UpdateDocuments) (*dto.Crm_DTO, error)
}
type DocumentsCRMService struct {
	Repo crm_repository.DocumentsRepositoryInterface
}

func (t DocumentsCRMService) DocumentsGet(model crm_structures.GetDocuments) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Documents |")
	result, err := t.Repo.GetDocuments(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t DocumentsCRMService) DocumentsCreate(model crm_structures.CreateDocuments) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Documents |")
	result, err := t.Repo.CreateDocuments(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t DocumentsCRMService) DocumentsUpdate(model crm_structures.UpdateDocuments) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Documents |")
	result, err := t.Repo.UpdateDocuments(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
