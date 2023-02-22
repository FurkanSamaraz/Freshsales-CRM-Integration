package crm_services

import (
	"fmt"
	dto "main/dto"
	crm_repository "main/repository/Appointments"
	crm_structures "main/structure/Appointments"
)

type Service_Appointments interface {
	AppointmentsCreate(model crm_structures.Create_Appointments) (*dto.Crm_DTO, error)
	AppointmentsGet(model crm_structures.Get_Appointments) (*dto.Crm_DTO, error)
	AppointmentsUpdate(model crm_structures.Update_Appointments) (*dto.Crm_DTO, error)
}
type AppointmentsCRMService struct {
	Repo crm_repository.KnowledgerArticlesRepositoryInterface
}

func (t AppointmentsCRMService) AppointmentsGet(model crm_structures.Get_Appointments) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Appointments |")
	result, err := t.Repo.GetKnowledgerArticles(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t AppointmentsCRMService) AppointmentsCreate(model crm_structures.Create_Appointments) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Appointments |")
	result, err := t.Repo.CreateKnowledgerArticles(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
func (t AppointmentsCRMService) AppointmentsUpdate(model crm_structures.Update_Appointments) (*dto.Crm_DTO, error) {
	var res dto.Crm_DTO

	fmt.Println("| SERVICE    SUCCESS - Appointments |")
	result, err := t.Repo.UpdateKnowledgerArticles(model)

	if err != nil || !result {
		res.Status = false
		return &res, err
	}

	res = dto.Crm_DTO{Status: result}
	return &res, nil
}
