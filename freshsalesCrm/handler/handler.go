package crm_handlers

import (
	"fmt"
	configs "main/config"
	handler_accounts "main/handler/accounts"
	handler_activities "main/handler/activities"
	crm_integration "main/package"
	repository "main/repository"
	crm_repository_accounts "main/repository/Accounts"
	crm_repository_activities "main/repository/Activities"

	handler_Deals "main/handler/Deals"
	crm_repository_Deals "main/repository/Deals"

	handler_Contacts "main/handler/Contacts"
	crm_repository_Contacts "main/repository/Contacts"

	handler_Appointments "main/handler/Appointments"
	crm_repository_Appointments "main/repository/Appointments"

	handler_Leads "main/handler/Leads"
	crm_repository_Leads "main/repository/Leads"

	handler_SalesActivities "main/handler/SalesActivities"
	crm_repository_SalesActivities "main/repository/SalesActivities"

	handler_Search "main/handler/Search"
	crm_repository_Search "main/repository/Search"

	handler_Products "main/handler/Products"
	crm_repository_Products "main/repository/Products"

	handler_Tasks "main/handler/Tasks"
	crm_repository_Tasks "main/repository/Tasks"

	handler_Users "main/handler/Users"
	crm_repository_Users "main/repository/Users"

	handler_Notes "main/handler/Notes"
	crm_repository_Notes "main/repository/Notes"

	handler_Phone "main/handler/Phone"
	crm_repository_Phone "main/repository/Phone"

	handler_Documents "main/handler/Documents"
	crm_repository_Documents "main/repository/Documents"

	handler_Files "main/handler/Files"
	crm_repository_Files "main/repository/Files"

	handler_JobStatus "main/handler/JobStatus"
	crm_repository_JobStatus "main/repository/JobStatus"

	service "main/service"
)

type CRM_CONTROLLER struct {
	Base *crm_integration.CRM_BASE
}

func App() {
	dbClient := configs.GetCollection()

	AccountsRepositoryDb := repository.NewAccountsRepositoryDB(dbClient)
	accounts := handler_accounts.AccountsHandler{Service: service.NewAccountsService(crm_repository_accounts.AccountsDB{CrmCollection: AccountsRepositoryDb.CrmCollection})}

	ActivitiesRepositoryDb := repository.NewActivitiesRepositoryDB(dbClient)
	activities := handler_activities.ActivitiesHandler{Service: service.NewActivitiesService(crm_repository_activities.ActivitiesDB{CrmCollection: ActivitiesRepositoryDb.CrmCollection})}

	DealsRepositoryDb := repository.NewDealsRepositoryDB(dbClient)
	Deals := handler_Deals.DealsHandler{Service: service.NewDealsService(crm_repository_Deals.DealsDB{CrmCollection: DealsRepositoryDb.CrmCollection})}

	ContactsRepositoryDb := repository.NewContactsRepositoryDB(dbClient)
	Contacts := handler_Contacts.ContactsHandler{Service: service.NewContactsService(crm_repository_Contacts.ContactsDB{CrmCollection: ContactsRepositoryDb.CrmCollection})}

	AppointmentsRepositoryDb := repository.NewAppointmentsRepositoryDB(dbClient)
	Appointments := handler_Appointments.AppointmentsHandler{Service: service.NewAppointmentsService(crm_repository_Appointments.KnowledgerArticlesDB{CrmCollection: AppointmentsRepositoryDb.CrmCollection})}

	LeadsRepositoryDb := repository.NewLeadsRepositoryDB(dbClient)
	Leads := handler_Leads.LeadsHandler{Service: service.NewLeadsService(crm_repository_Leads.LeadsDB{CrmCollection: LeadsRepositoryDb.CrmCollection})}

	SalesActivitiesRepositorysDb := repository.NewSalesActivitiesRepositoryDB(dbClient)
	SalesActivities := handler_SalesActivities.SalesActivitiesHandler{Service: service.NewSalesActivitiesService(crm_repository_SalesActivities.SalesActivitiesDB{CrmCollection: SalesActivitiesRepositorysDb.CrmCollection})}

	SearchRepositoryDb := repository.NewSearchRepositoryDB(dbClient)
	Search := handler_Search.SearchHandler{Service: service.NewSearchService(crm_repository_Search.SearchDB{CrmCollection: SearchRepositoryDb.CrmCollection})}

	ProductsRepositoryDb := repository.NewProductsRepositoryDB(dbClient)
	Products := handler_Products.ProductsHandler{Service: service.NewProductsService(crm_repository_Products.ProductsDB{CrmCollection: ProductsRepositoryDb.CrmCollection})}

	UsersRepositoryDb := repository.NewUsersRepositoryDB(dbClient)
	Users := handler_Users.UsersHandler{Service: service.NewUsersService(crm_repository_Users.UsersDB{CrmCollection: UsersRepositoryDb.CrmCollection})}

	TasksRepositoryDb := repository.NewTasksRepositoryDB(dbClient)
	Tasks := handler_Tasks.TasksHandler{Service: service.NewTasksService(crm_repository_Tasks.TasksDB{CrmCollection: TasksRepositoryDb.CrmCollection})}

	NotesRepositoryDb := repository.NewNotesRepositoryDB(dbClient)
	Notes := handler_Notes.NotesHandler{Service: service.NewNotesService(crm_repository_Notes.NotesDB{CrmCollection: NotesRepositoryDb.CrmCollection})}

	PhoneRepositoryDb := repository.NewPhoneRepositoryDB(dbClient)
	Phone := handler_Phone.PhoneHandler{Service: service.NewPhoneService(crm_repository_Phone.PhoneDB{CrmCollection: PhoneRepositoryDb.CrmCollection})}

	DocumentsRepositoryDb := repository.NewDocumentsRepositoryDB(dbClient)
	Documents := handler_Documents.DocumentsHandler{Service: service.NewDocumentsService(crm_repository_Documents.DocumentsDB{CrmCollection: DocumentsRepositoryDb.CrmCollection})}

	FilesRepositoryDb := repository.NewFilesRepositoryDB(dbClient)
	Files := handler_Files.FilesHandler{Service: service.NewFilesService(crm_repository_Files.FilesDB{CrmCollection: FilesRepositoryDb.CrmCollection})}

	JobStatusRepositoryDb := repository.NewJobStatusRepositoryDB(dbClient)
	JobStatus := handler_JobStatus.JobStatusHandler{Service: service.NewJobStatusService(crm_repository_JobStatus.JobStatusDB{CrmCollection: JobStatusRepositoryDb.CrmCollection})}

	fmt.Println(
		accounts,
		activities,
		Deals,
		Contacts,
		Appointments,
		Leads,
		SalesActivities,
		Search,
		Products,
		Users,
		Tasks,
		Notes,
		Phone,
		Documents,
		Files,
		JobStatus,
	)

}
