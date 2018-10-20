package dbinterface

import "app-server/infra/models"

type LabDBRepository interface {

	// APIs for Resource Type CRUD
	GetResourceType(resourceType string) (models.ResourceType, error)
	GetResourceTypes() ([]models.ResourceType, error)
	GetChildResourceTypes(resourceType string) ([]models.ResourceType, error)
	GetRecurseChildResourceTypes(resourceType string) ([]models.ResourceType, error)
	GetChildResourceTypesOfParentType(parentTypeArray []string) ([]models.ResourceType, error)
	AddResourceType(resourceType models.ResourceType) error
	DeleteAllResourceTypes() error

	// APIs for Resource CRUD
	GetResources() (models.Resources, error)
	GetResourcesInLab(labname string) (models.Resources, error)
	GetResourcesByType(resourceType string) (models.Resources, error)
	AddResource(models.Resource) error
	DeleteResourcesByName(resourceNames []string) (error)
	DeleteResourcesByID(resourceIDs []string) (error)
	DeleteAllResources () error
	UpdateResourceByName(resourceName string, resource models.Resource) error
	GetResourceReservationMatrix(resourceName string) ([]models.ReservationMatrix, error)
	ReserveResource(string, matrix models.ReservationMatrix) error
}
