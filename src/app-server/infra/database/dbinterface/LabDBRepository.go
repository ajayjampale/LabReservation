package dbinterface

import "app-server/infra/models"

type LabDBRepository interface {
	GetResources() models.Resources
	GetResourceDetail(string) models.Resource
	AddResource(models.Resource) error

	GetResourceType(resourceType string) (models.ResourceType, error)
	GetResourceTypes() ([]models.ResourceType, error)
	GetChildResourceTypes (resourceType string) ([]models.ResourceType, error)
	AddResourceType(resourceType models.ResourceType) error
	DeleteAllResourceTypes() error

	ReserveResource(string, matrix models.ReservationMatrix) error
}
