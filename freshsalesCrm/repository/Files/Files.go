package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Files"

	"gorm.io/gorm"
)

type FilesDB struct {
	CrmCollection *gorm.DB
}
type FilesRepositoryInterface interface {
	GetFiles(model crm_structures.GetFiles) (bool, error)
	CreateFiles(model crm_structures.CreateFiles) (bool, error)
}

func (t FilesDB) GetFiles(model crm_structures.GetFiles) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Files |")
	return true, nil
}
func (t FilesDB) CreateFiles(model crm_structures.CreateFiles) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Files |")
	return true, nil
}
