package crm_repository

import (
	"fmt"
	crm_structures "main/structure/Search"

	"gorm.io/gorm"
)

type SearchDB struct {
	CrmCollection *gorm.DB
}
type SearchRepositoryInterface interface {
	GetSearch(model crm_structures.GetSearch) (bool, error)
}

func (t SearchDB) GetSearch(model crm_structures.GetSearch) (bool, error) {
	fmt.Println("| REPOSITORY SUCCESS - Search |")
	return true, nil
}
