package service

import (
	"api/models"
	"api/repository"
	"fmt"
)

func GetServicesById(id int64) (services models.Services, err error) {
	services, err = repository.GetServicesById(int64(id))
	return services, err
}

func PostServices(services models.Services) (err error) {
	found := repository.VerifyServicesByDescription(services.Desc)

	if found {
		return fmt.Errorf("Serviço com a DESCRIÇÃO %s já cadastrado!", services.Desc)
	} else {
		err = repository.PostServices(services)
	}

	return err
}

func GetAllServices() (servicess []models.Services, err error) {
	servicess, err = repository.GetAllServices()
	return servicess, err
}

func PutServices(servicesUpdate models.Services) (err error) {
	var services models.Services

	services, err = repository.GetServicesById(servicesUpdate.IdServices)
	if err != nil {
		return err
	}

	if services.IdServices == 0 {
		return fmt.Errorf("Serviço não está cadastrado no sistema!")
	}

	err = repository.PutServices(servicesUpdate)

	return
}
