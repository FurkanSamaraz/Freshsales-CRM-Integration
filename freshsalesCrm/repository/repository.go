package repository

import "gorm.io/gorm"

type CrmRepositoryDB struct {
	CrmCollection *gorm.DB
}

func NewAccountsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}

func NewActivitiesRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewDealsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewContactsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewAppointmentsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewLeadsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewSalesActivitiesRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewNotesRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewSearchRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewProductsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewTasksRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewUsersRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewPhoneRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewDocumentsRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewFilesRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
func NewJobStatusRepositoryDB(dbClient *gorm.DB) CrmRepositoryDB {
	return CrmRepositoryDB{CrmCollection: dbClient}
}
