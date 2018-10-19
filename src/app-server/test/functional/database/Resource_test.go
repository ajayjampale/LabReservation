package database

import (
	"app-server/infra/database/MongoDBRepository"
	"app-server/infra/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResourceTypeCRUD(t *testing.T) {
	labDB := MongoDBRepository.LabMongoDBRepository{Username: "dbuser", Password: "mith1234", IPAddress: "127.0.0.1", Port: "27017", Database: "dummyStore"}

	// Delete any existing ResourceTypes
	err := labDB.DeleteAllResourceTypes()
	assert.Nil(t, err)

	// Add ResourceType
	resType1 := models.ResourceType{ID: "1", Type: "SLXOS Devices", ParentTypeID: ""}
	resType2 := models.ResourceType{ID: "2", Type: "NOS Devices", ParentTypeID: ""}
	err = labDB.AddResourceType(resType1)
	assert.Nil(t, err)
	err = labDB.AddResourceType(resType2)
	assert.Nil(t, err)

	// Get ResourceType and verify/assert.
	resTypeOut, err := labDB.GetResourceTypes()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(resTypeOut))
	if 2 == len(resTypeOut) {
		assert.Equal(t, "NOS Devices", resTypeOut[0].Type)
		assert.Equal(t, "SLXOS Devices", resTypeOut[1].Type)
	}

	// Cleanup all ResourceTypes
	err = labDB.DeleteAllResourceTypes()
	assert.Nil(t, err)

	resTypeOut, err = labDB.GetResourceTypes()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(resTypeOut))

}

func TestResourceTypeOneLevelCRUD(t *testing.T) {
	labDB := MongoDBRepository.LabMongoDBRepository{Username: "dbuser", Password: "mith1234", IPAddress: "127.0.0.1", Port: "27017", Database: "dummyStore"}

	// Delete any existing ResourceTypes
	err := labDB.DeleteAllResourceTypes()
	assert.Nil(t, err)

	// Add ResourceType
	resType1 := models.ResourceType{ID: "1", Type: "SLXOS Devices", ParentTypeID: ""}
	resType2 := models.ResourceType{ID: "2", Type: "NOS Devices", ParentTypeID: ""}
	err = labDB.AddResourceType(resType1)
	assert.Nil(t, err)
	err = labDB.AddResourceType(resType2)
	assert.Nil(t, err)

	// Get ResourceType and verify/assert.
	resTypeOut, err := labDB.GetResourceTypes()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(resTypeOut))
	if 2 == len(resTypeOut) {
		assert.Equal(t, "NOS Devices", resTypeOut[0].Type)
		assert.Equal(t, "SLXOS Devices", resTypeOut[1].Type)
	}

	// Add sub resource types.
	err = labDB.AddResourceType(models.ResourceType{ID: "3", Type: "Cedar", ParentTypeID: "1"})
	err = labDB.AddResourceType(models.ResourceType{ID: "4", Type: "Freedom", ParentTypeID: "1"})
	err = labDB.AddResourceType(models.ResourceType{ID: "5", Type: "Castor", ParentTypeID: "2"})
	err = labDB.AddResourceType(models.ResourceType{ID: "6", Type: "Rigel", ParentTypeID: "2"})

	resTypeOut, err = labDB.GetResourceTypes()
	assert.Nil(t, err)
	assert.Equal(t, 6, len(resTypeOut))
	if 6 == len(resTypeOut) {
		assert.Equal(t, "Castor", resTypeOut[0].Type)
		assert.Equal(t, "2", resTypeOut[0].ParentTypeID)

		assert.Equal(t, "Cedar", resTypeOut[1].Type)
		assert.Equal(t, "1", resTypeOut[1].ParentTypeID)

		assert.Equal(t, "Freedom", resTypeOut[2].Type)
		assert.Equal(t, "1", resTypeOut[2].ParentTypeID)

		assert.Equal(t, "NOS Devices", resTypeOut[3].Type)
		assert.Equal(t, "", resTypeOut[3].ParentTypeID)

		assert.Equal(t, "Rigel", resTypeOut[4].Type)
		assert.Equal(t, "2", resTypeOut[4].ParentTypeID)

		assert.Equal(t, "SLXOS Devices", resTypeOut[5].Type)
		assert.Equal(t, "", resTypeOut[5].ParentTypeID)

	}

	// Cleanup all ResourceTypes
	err = labDB.DeleteAllResourceTypes()
	assert.Nil(t, err)

	resTypeOut, err = labDB.GetResourceTypes()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(resTypeOut))

}

func TestResourceTypeMultiLevelCRUD(t *testing.T) {
	labDB := MongoDBRepository.LabMongoDBRepository{Username: "dbuser", Password: "mith1234", IPAddress: "127.0.0.1", Port: "27017", Database: "dummyStore"}

	// Delete any existing ResourceTypes
	err := labDB.DeleteAllResourceTypes()
	assert.Nil(t, err)

	// Add ResourceType
	resType1 := models.ResourceType{ID: "1", Type: "SLXOS Devices", ParentTypeID: ""}
	resType2 := models.ResourceType{ID: "2", Type: "NOS Devices", ParentTypeID: ""}
	err = labDB.AddResourceType(resType1)
	assert.Nil(t, err)
	err = labDB.AddResourceType(resType2)
	assert.Nil(t, err)

	// Get ResourceType and verify/assert.
	resTypeOut, err := labDB.GetResourceTypes()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(resTypeOut))
	if 2 == len(resTypeOut) {
		assert.Equal(t, "NOS Devices", resTypeOut[0].Type)
		assert.Equal(t, "SLXOS Devices", resTypeOut[1].Type)
	}

	// Add sub resource types.
	err = labDB.AddResourceType(models.ResourceType{ID: "3", Type: "Switches", ParentTypeID: "1"})
	err = labDB.AddResourceType(models.ResourceType{ID: "4", Type: "Routers", ParentTypeID: "1"})
	err = labDB.AddResourceType(models.ResourceType{ID: "5", Type: "Castor", ParentTypeID: "2"})
	err = labDB.AddResourceType(models.ResourceType{ID: "6", Type: "Rigel", ParentTypeID: "2"})

	resTypeOut, err = labDB.GetResourceTypes()
	assert.Nil(t, err)
	assert.Equal(t, 6, len(resTypeOut))
	if 6 == len(resTypeOut) {
		assert.Equal(t, "Castor", resTypeOut[0].Type)
		assert.Equal(t, "2", resTypeOut[0].ParentTypeID)

		assert.Equal(t, "NOS Devices", resTypeOut[1].Type)
		assert.Equal(t, "", resTypeOut[1].ParentTypeID)

		assert.Equal(t, "Rigel", resTypeOut[2].Type)
		assert.Equal(t, "2", resTypeOut[2].ParentTypeID)

		assert.Equal(t, "Routers", resTypeOut[3].Type)
		assert.Equal(t, "1", resTypeOut[3].ParentTypeID)

		assert.Equal(t, "SLXOS Devices", resTypeOut[4].Type)
		assert.Equal(t, "", resTypeOut[4].ParentTypeID)

		assert.Equal(t, "Switches", resTypeOut[5].Type)
		assert.Equal(t, "1", resTypeOut[5].ParentTypeID)

	}

	err = labDB.AddResourceType(models.ResourceType{ID: "7", Type: "Avalanche", ParentTypeID: "4"})
	err = labDB.AddResourceType(models.ResourceType{ID: "8", Type: "Fusion", ParentTypeID: "4"})
	err = labDB.AddResourceType(models.ResourceType{ID: "9", Type: "Cedar", ParentTypeID: "3"})
	err = labDB.AddResourceType(models.ResourceType{ID: "10", Type: "Freedom", ParentTypeID: "3"})

	resTypeOut, err = labDB.GetChildResourceTypes("Switches")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(resTypeOut))
	if 2 == len(resTypeOut) {
		assert.Equal(t, "Cedar", resTypeOut[0].Type)
		assert.Equal(t, "3", resTypeOut[0].ParentTypeID)

		assert.Equal(t, "Freedom", resTypeOut[1].Type)
		assert.Equal(t, "3", resTypeOut[1].ParentTypeID)
	}

	resTypeOut, err = labDB.GetChildResourceTypes("Routers")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(resTypeOut))
	if 2 == len(resTypeOut) {
		assert.Equal(t, "Avalanche", resTypeOut[0].Type)
		assert.Equal(t, "4", resTypeOut[0].ParentTypeID)

		assert.Equal(t, "Fusion", resTypeOut[1].Type)
		assert.Equal(t, "4", resTypeOut[1].ParentTypeID)
	}

	// Cleanup all ResourceTypes
	err = labDB.DeleteAllResourceTypes()
	assert.Nil(t, err)

	resTypeOut, err = labDB.GetResourceTypes()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(resTypeOut))

}

func TestResourceCRUD(t *testing.T) {
	labDB := MongoDBRepository.LabMongoDBRepository{Username: "dbuser", Password: "mith1234", IPAddress: "127.0.0.1", Port: "27017", Database: "dummyStore"}

	// Delete any existing ResourceTypes
	err := labDB.DeleteAllResourceTypes()
	assert.Nil(t, err)

	// Delete any existing Resources
	err = labDB.DeleteAllResources()
	assert.Nil(t, err)

	// Add ResourceType
	err = labDB.AddResourceType(models.ResourceType{ID: "1", Type: "SLXOS Devices", ParentTypeID: ""});
	assert.Nil(t, err)
	err = labDB.AddResourceType(models.ResourceType{ID: "2", Type: "NOS Devices", ParentTypeID: ""});
	assert.Nil(t, err)
	err = labDB.AddResourceType(models.ResourceType{ID: "3", Type: "Switches", ParentTypeID: "1"});
	assert.Nil(t, err)
	err = labDB.AddResourceType(models.ResourceType{ID: "4", Type: "Routers", ParentTypeID: "1"});
	assert.Nil(t, err)
	err = labDB.AddResourceType(models.ResourceType{ID: "5", Type: "Castor", ParentTypeID: "2"});
	assert.Nil(t, err)
	err = labDB.AddResourceType(models.ResourceType{ID: "6", Type: "Rigel", ParentTypeID: "2"});
	assert.Nil(t, err)
	err = labDB.AddResourceType(models.ResourceType{ID: "7", Type: "Avalanche", ParentTypeID: "4"});
	assert.Nil(t, err)
	err = labDB.AddResourceType(models.ResourceType{ID: "8", Type: "Fusion", ParentTypeID: "4"});
	assert.Nil(t, err)
	err = labDB.AddResourceType(models.ResourceType{ID: "9", Type: "Cedar", ParentTypeID: "3"});
	assert.Nil(t, err)
	err = labDB.AddResourceType(models.ResourceType{ID: "10", Type: "Freedom", ParentTypeID: "3"});
	assert.Nil(t, err)

	err = labDB.AddResource(models.Resource{
		ID:    "1", Name: "Team1_Castor1", Description: "Castor for xyz Team", ResourceTypeID: "5",
		Infos: []models.ResourceInfo{models.ResourceInfo{IPAddress: []string{"10.10.10.10"}, Lab: "Bangalore Campus"}},
		PortMatrix: []models.ResourcePortMatrix{models.ResourcePortMatrix{SourcePort: "1", RemotePort: "2", RemoteResourceID: "2"},
			models.ResourcePortMatrix{SourcePort: "2", RemotePort: "2", RemoteResourceID: "3"},},
		ResourceOwnerID: "1",
	});
	assert.Nil(t, err)

	err = labDB.AddResource(models.Resource{
		ID:              "2", Name: "Team1_Castor2", Description: "Castor for xyz Team", ResourceTypeID: "5",
		Infos:           []models.ResourceInfo{models.ResourceInfo{IPAddress: []string{"10.10.10.11"}, Lab: "Bangalore Campus"}},
		PortMatrix:      []models.ResourcePortMatrix{models.ResourcePortMatrix{SourcePort: "2", RemotePort: "1", RemoteResourceID: "1"},},
		ResourceOwnerID: "1",
	});
	assert.Nil(t, err)

	err = labDB.AddResource(models.Resource{
		ID:              "3", Name: "Team1_Rigel", Description: "Rigel for xyz Team", ResourceTypeID: "6",
		Infos:           []models.ResourceInfo{models.ResourceInfo{IPAddress: []string{"10.10.10.12"}, Lab: "Bangalore Campus"}},
		PortMatrix:      []models.ResourcePortMatrix{models.ResourcePortMatrix{SourcePort: "2", RemotePort: "2", RemoteResourceID: "1"},},
		ResourceOwnerID: "1",
	});
	assert.Nil(t, err)


	err = labDB.AddResource(models.Resource{
		ID:              "4", Name: "Team2_AVA1", Description: "Avalanche for abc Team", ResourceTypeID: "7",
		Infos:           []models.ResourceInfo{models.ResourceInfo{IPAddress: []string{"10.10.11.10"}, Lab: "Sanjose Campus"}},
		PortMatrix:      []models.ResourcePortMatrix{},
		ResourceOwnerID: "1",
	});
	assert.Nil(t, err)

	resourceOut, err := labDB.GetResources();
	assert.Nil(t, err)

	assert.Equal(t, 4, len(resourceOut))


	resourceOut, err = labDB.GetResourcesInLab("Bangalore Campus");
	assert.Nil(t, err)

	assert.Equal(t, 3, len(resourceOut))

	resourceOut, err = labDB.GetResourcesInLab("Sanjose Campus");
	assert.Nil(t, err)

	assert.Equal(t, 1, len(resourceOut))

	resourceOut, err = labDB.GetResourcesInLab("Mysore Campus");
	assert.Nil(t, err)

	assert.Equal(t, 0, len(resourceOut))

	resourceOut, err = labDB.GetResourcesByType("NOS Devices");
	assert.Nil(t, err)

	assert.Equal(t, 3, len(resourceOut))

	resourceOut, err = labDB.GetResourcesByType("SLXOS Devices");
	assert.Nil(t, err)

	assert.Equal(t, 1, len(resourceOut))

}
