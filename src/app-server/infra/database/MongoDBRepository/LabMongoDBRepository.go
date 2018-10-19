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

func (r *LabMongoDBRepository) GetChildResourceTypesOfParentType(parentTypeIDArray []string) ([]models.ResourceType, error) {
	DB, err := r.GetDatabaseHandle()

	results := []models.ResourceType{}
	if err1 := DB.C(models.CollectionResourceType).Find(bson.M{"parenttypeid": bson.M{"$in":parentTypeIDArray}}).Sort("+type").All(&results); err1 != nil {
		return nil, err1
	}
	return results, err
}

func (r *LabMongoDBRepository) GetRecurseChildResourceTypes(resourceType string) ([]models.ResourceType, error) {
	var err error

	var resType models.ResourceType
	if resType, err = r.GetResourceType(resourceType); err != nil {
		return nil, err
	}

	results := []models.ResourceType{}
	results = append(results, resType)

	parentTypeArray := []string{resType.ID}

	for (len(parentTypeArray) != 0) {
		resTypes, err := r.GetChildResourceTypesOfParentType(parentTypeArray)
		parentTypeArray = parentTypeArray[:0]
		if err == nil {
			for _,resType := range resTypes {
				results = append(results, resType)
				parentTypeArray = append(parentTypeArray, resType.ID)
			}
		}
	}
	return results, err
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

func (r *LabMongoDBRepository) GetResources() (models.Resources, error) {

	DB, err := r.GetDatabaseHandle()

	if err != nil {
		return nil, err
	}

	results := models.Resources{}
	if err1 := DB.C(models.CollectionResource).Find(nil).Sort("+name").All(&results); err1 != nil {
		return nil, err1
	}

	return results, nil

}
func (r *LabMongoDBRepository) GetResourcesInLab(labname string) (models.Resources, error) {

	DB, err := r.GetDatabaseHandle()

	if err != nil {
		return nil, err
	}

	results := models.Resources{}
	if err1 := DB.C(models.CollectionResource).Find( bson.M{"infos.lab" : labname} ).Sort("+name").All(&results); err1 != nil {
		return nil, err1
	}

	return results, nil

}
func (r *LabMongoDBRepository) GetResourcesByType (resourceType string) (models.Resources, error) {
	DB, err := r.GetDatabaseHandle()
	if err != nil {
		return nil, err
	}

	resourceTypeArr, err := r.GetRecurseChildResourceTypes(resourceType)
	if err != nil {
		return nil, err
	}

	var resTypeIDArr []string
	for _, resType := range resourceTypeArr {
		resTypeIDArr = append(resTypeIDArr, resType.ID)
	}
	results := models.Resources{}
	if err1 := DB.C(models.CollectionResource).Find( bson.M{"resourcetypeid" : bson.M{"$in" : resTypeIDArr}} ).Sort("+name").All(&results); err1 != nil {
		return nil, err1
	}

	return results, nil
}

func (r *LabMongoDBRepository) AddResource (Resource models.Resource) error {
	DB, err := r.GetDatabaseHandle()

	if err != nil {
		return err
	}

	err = DB.C(models.CollectionResource).Insert(Resource)
	return err
}

func (r *LabMongoDBRepository) DeleteResourcesByName(resourceName []string) error {
	return nil
}

func (r *LabMongoDBRepository) DeleteResourcesByID(resourceID []string) error {
	return nil
}

func (r *LabMongoDBRepository) DeleteAllResources () error {
	DB, err := r.GetDatabaseHandle()

	if err != nil {
		return err
	}

	_, err = DB.C(models.CollectionResource).RemoveAll(nil)

	return err
}

func (r *LabMongoDBRepository) UpdateResourceByName(resourceName string, resource models.Resource) error {
	return nil
}

func (r *LabMongoDBRepository) GetResourceReservationMatrix(resourceName string) ([]models.ReservationMatrix, error) {
	return nil, nil
}

func (r *LabMongoDBRepository) ReserveResource(string, matrix models.ReservationMatrix) error {
	return nil
}
