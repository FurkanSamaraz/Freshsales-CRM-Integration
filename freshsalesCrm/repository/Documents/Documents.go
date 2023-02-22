package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Documents"

	"gorm.io/gorm"
)

type DocumentsDB struct {
	CrmCollection *gorm.DB
}
type DocumentsRepositoryInterface interface {
	GetDocuments(model crm_structures.GetDocuments) (bool, error)
	CreateDocuments(model crm_structures.CreateDocuments) (bool, error)
	UpdateDocuments(model crm_structures.UpdateDocuments) (bool, error)
}

func (t DocumentsDB) GetDocuments(model crm_structures.GetDocuments) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Documents |")
	return true, nil
}
func (t DocumentsDB) CreateDocuments(model crm_structures.CreateDocuments) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Documents |")
	return true, nil
}
func (t DocumentsDB) UpdateDocuments(model crm_structures.UpdateDocuments) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Documents |")
	return true, nil
}
