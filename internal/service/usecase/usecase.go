package usecase

import (
	"lonkidely/technopark-dbms-forum/internal/models"
	"lonkidely/technopark-dbms-forum/internal/service/repository"
)

type ServiceUsecase interface {
	Clear() error
	Status() (models.Status, error)
}

type serviceUsecase struct {
	serviceRepo repository.ServiceRepository
}

func NewServiceUsecase(serviceRepo repository.ServiceRepository) ServiceUsecase {
	return &serviceUsecase{
		serviceRepo: serviceRepo,
	}
}

func (su *serviceUsecase) Clear() error {
	return su.serviceRepo.Clear()
}

func (su *serviceUsecase) Status() (models.Status, error) {
	return su.serviceRepo.Status()
}
