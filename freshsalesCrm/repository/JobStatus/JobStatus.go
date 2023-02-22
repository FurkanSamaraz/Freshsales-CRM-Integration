package crm_repository

import (
	"fmt"
	crm_structures "main/structure/JobStatus"

	"gorm.io/gorm"
)

type JobStatusDB struct {
	CrmCollection *gorm.DB
}
type JobStatusRepositoryInterface interface {
	GetJobStatus(model crm_structures.GetJobStatus) (bool, error)
}

func (t JobStatusDB) GetJobStatus(model crm_structures.GetJobStatus) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - JobStatus |")
	return true, nil
}
