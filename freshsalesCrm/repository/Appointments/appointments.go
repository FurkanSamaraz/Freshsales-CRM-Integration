package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Appointments"

	"gorm.io/gorm"
)

type KnowledgerArticlesDB struct {
	CrmCollection *gorm.DB
}
type KnowledgerArticlesRepositoryInterface interface {
	GetKnowledgerArticles(model crm_structures.Get_Appointments) (bool, error)
	CreateKnowledgerArticles(model crm_structures.Create_Appointments) (bool, error)
	UpdateKnowledgerArticles(model crm_structures.Update_Appointments) (bool, error)
}

func (t KnowledgerArticlesDB) GetKnowledgerArticles(model crm_structures.Get_Appointments) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - KnowledgerArticles |")
	return true, nil
}
func (t KnowledgerArticlesDB) CreateKnowledgerArticles(model crm_structures.Create_Appointments) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - KnowledgerArticles |")
	return true, nil
}
func (t KnowledgerArticlesDB) UpdateKnowledgerArticles(model crm_structures.Update_Appointments) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - KnowledgerArticles |")
	return true, nil
}
