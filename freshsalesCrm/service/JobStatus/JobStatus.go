package crm_serices

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/JobStatus"
	crm_structures "main/structure/JobStatus"
)

type Service_JobStatus interface {
	JobStatusGet(model crm_structures.GetJobStatus) (*dto.Crm_DTO, error)
}
type JobStatusCRMService struct {
	Repo crm_repository.JobStatusRepositoryInterface
}

func (t JobStatusCRMService) JobStatusGet(model crm_structures.GetJobStatus) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - JobStatus |")
	result, err := t.Repo.GetJobStatus(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
