package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Activities"

	"gorm.io/gorm"
)

type ActivitiesDB struct {
	CrmCollection *gorm.DB
}
type ActivitiesRepositoryInterface interface {
	GetActivities(model crm_structures.GetActivities) (bool, error)
}

func (t ActivitiesDB) GetActivities(model crm_structures.GetActivities) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Activities |")
	return true, nil
}
