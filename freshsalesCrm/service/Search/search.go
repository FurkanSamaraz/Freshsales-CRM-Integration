package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Search"
	crm_structures "main/structure/Search"
)

type Service_Search interface {
	SearchGet(model crm_structures.GetSearch) (*dto.Crm_DTO, error)
}
type SearchCRMService struct {
	Repo crm_repository.SearchRepositoryInterface
}

func (t SearchCRMService) SearchGet(model crm_structures.GetSearch) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Search |")
	result, err := t.Repo.GetSearch(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
