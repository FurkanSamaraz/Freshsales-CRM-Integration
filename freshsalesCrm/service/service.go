package crm_services

import (
	crm_repository_Accounts "main/repository/Accounts"
	crm_service_Accounts "main/service/Accounts"

	crm_repository_Activities "main/repository/Activities"
	crm_service_Activities "main/service/Activities"

	crm_repository_Deals "main/repository/Deals"
	crm_service_Deals "main/service/Deals"

	crm_repository_Contacts "main/repository/Contacts"
	crm_service_Contacts "main/service/Contacts"

	crm_repository_Appointments "main/repository/Appointments"
	crm_service_Appointments "main/service/Appointments"

	crm_repository_Leads "main/repository/Leads"
	crm_service_Leads "main/service/Leads"

	crm_repository_SalesActivities "main/repository/SalesActivities"
	crm_service_SalesActivities "main/service/SalesActivities"

	crm_repository_Notes "main/repository/Notes"
	crm_service_Notes "main/service/Notes"

	crm_repository_Search "main/repository/Search"
	crm_service_Search "main/service/Search"

	crm_repository_Products "main/repository/Products"
	crm_service_Products "main/service/Products"

	crm_repository_Tasks "main/repository/Tasks"
	crm_service_Tasks "main/service/Tasks"

	crm_repository_Users "main/repository/Users"
	crm_service_Users "main/service/Users"

	crm_repository_Phone "main/repository/Phone"
	crm_service_Phone "main/service/Phone"

	crm_repository_Documents "main/repository/Documents"
	crm_service_Documents "main/service/Documents"

	crm_repository_Files "main/repository/Files"
	crm_service_Files "main/service/Files"

	crm_repository_JobStatus "main/repository/JobStatus"
	crm_service_JobStatus "main/service/JobStatus"
)

func NewDocumentsService(Repo crm_repository_Documents.DocumentsRepositoryInterface) crm_service_Documents.DocumentsCRMService {
	return crm_service_Documents.DocumentsCRMService{Repo: Repo}
}
func NewFilesService(Repo crm_repository_Files.FilesRepositoryInterface) crm_service_Files.FilesCRMService {
	return crm_service_Files.FilesCRMService{Repo: Repo}
}
func NewJobStatusService(Repo crm_repository_JobStatus.JobStatusRepositoryInterface) crm_service_JobStatus.JobStatusCRMService {
	return crm_service_JobStatus.JobStatusCRMService{Repo: Repo}
}

func NewPhoneService(Repo crm_repository_Phone.PhoneRepositoryInterface) crm_service_Phone.PhoneCRMService {
	return crm_service_Phone.PhoneCRMService{Repo: Repo}
}
func NewAccountsService(Repo crm_repository_Accounts.AccountsRepositoryInterface) crm_service_Accounts.AccountsCRMService {
	return crm_service_Accounts.AccountsCRMService{Repo: Repo}
}

func NewActivitiesService(Repo crm_repository_Activities.ActivitiesRepositoryInterface) crm_service_Activities.ActivitiesCRMService {
	return crm_service_Activities.ActivitiesCRMService{Repo: Repo}
}

func NewDealsService(Repo crm_repository_Deals.DealsRepositoryInterface) crm_service_Deals.DealsCRMService {
	return crm_service_Deals.DealsCRMService{Repo: Repo}
}

func NewContactsService(Repo crm_repository_Contacts.ContactsRepositoryInterface) crm_service_Contacts.ContactsCRMService {
	return crm_service_Contacts.ContactsCRMService{Repo: Repo}
}
func NewAppointmentsService(Repo crm_repository_Appointments.KnowledgerArticlesRepositoryInterface) crm_service_Appointments.AppointmentsCRMService {
	return crm_service_Appointments.AppointmentsCRMService{Repo: Repo}
}
func NewLeadsService(Repo crm_repository_Leads.LeadsRepositoryInterface) crm_service_Leads.LeadsCRMService {
	return crm_service_Leads.LeadsCRMService{Repo: Repo}
}
func NewSalesActivitiesService(Repo crm_repository_SalesActivities.SalesActivitiesRepositoryInterface) crm_service_SalesActivities.SalesActivitiesCRMService {
	return crm_service_SalesActivities.SalesActivitiesCRMService{Repo: Repo}
}
func NewNotesService(Repo crm_repository_Notes.NotesRepositoryInterface) crm_service_Notes.NotesCRMService {
	return crm_service_Notes.NotesCRMService{Repo: Repo}
}
func NewProductsService(Repo crm_repository_Products.ProductsRepositoryInterface) crm_service_Products.ProductsCRMService {
	return crm_service_Products.ProductsCRMService{Repo: Repo}
}

func NewTasksService(Repo crm_repository_Tasks.TasksRepositoryInterface) crm_service_Tasks.TasksCRMService {
	return crm_service_Tasks.TasksCRMService{Repo: Repo}
}
func NewSearchService(Repo crm_repository_Search.SearchRepositoryInterface) crm_service_Search.SearchCRMService {
	return crm_service_Search.SearchCRMService{Repo: Repo}
}
func NewUsersService(Repo crm_repository_Users.UsersRepositoryInterface) crm_service_Users.UsersCRMService {
	return crm_service_Users.UsersCRMService{Repo: Repo}
}
