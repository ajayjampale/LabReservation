package MongoDBRepository

import (
	"app-server/infra/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type LabMongoDBRepository struct {
	Username  string
	Password  string
	IPAddress string
	Port      string
	Database  string
}

func (r *LabMongoDBRepository) GetDatabaseHandle() (*mgo.Database, error) {
	serverURL := "mongodb://" + r.Username + ":" + r.Password + "@" + r.IPAddress + ":" + r.Port + "/" + r.Database
	session, err := mgo.Dial(serverURL)
	if err != nil {
		return nil, err
	} else {
		DB := session.DB(r.Database)
		return DB, nil
	}
}

func (r *LabMongoDBRepository) GetResources() models.Resources {

	return nil
}

func (r *LabMongoDBRepository) GetResourceDetail(ResourceName string) models.Resource {
	var res models.Resource
	return res
}

func (r *LabMongoDBRepository) AddResource(Resource models.Resource) error {
	return nil
}

func (r *LabMongoDBRepository) GetResourceType(resourceType string) (models.ResourceType, error) {
	DB, err := r.GetDatabaseHandle()

	if err != nil {
		return models.ResourceType{}, err
	}

	result := models.ResourceType{}

	err = DB.C(models.CollectionResourceType).Find(bson.M{"type": resourceType}).One(&result)

	return result, err
}

func (r *LabMongoDBRepository) GetResourceTypes() ([]models.ResourceType, error) {
	DB, err := r.GetDatabaseHandle()

	if err != nil {
		return nil, err
	}

	results := []models.ResourceType{}
	if err1 := DB.C(models.CollectionResourceType).Find(nil).Sort("+type").All(&results); err1 != nil {
		return nil, err1
	}

	return results, nil
}


func (r *LabMongoDBRepository) GetChildResourceTypes(resourceType string) ([]models.ResourceType, error) {
	var err error

	DB, err := r.GetDatabaseHandle()

	if err != nil {
		return nil, err
	}

	var resType models.ResourceType
	if resType, err = r.GetResourceType(resourceType); err != nil {
		return nil, err
	}

	results := []models.ResourceType{}
	if err1 := DB.C(models.CollectionResourceType).Find(bson.M{"parenttypeid": resType.ID}).Sort("+type").All(&results); err1 != nil {
		return nil, err1
	}
	return results, nil

}


func (r *LabMongoDBRepository) AddResourceType(resourceType models.ResourceType) error {
	DB, err := r.GetDatabaseHandle()

	if err != nil {
		return err
	}

	err = DB.C(models.CollectionResourceType).Insert(resourceType)
	return err
}

func (r *LabMongoDBRepository) DeleteAllResourceTypes() error {
	DB, err := r.GetDatabaseHandle()

	if err != nil {
		return err
	}

	_, err = DB.C(models.CollectionResourceType).RemoveAll(nil)

	return err
}

func (r *LabMongoDBRepository) ReserveResource(string, matrix models.ReservationMatrix) error {
	return nil
}
