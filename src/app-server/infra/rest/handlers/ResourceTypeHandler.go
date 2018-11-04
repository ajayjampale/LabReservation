package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"app-server/infra/database/MongoDBRepository"
	"app-server/infra/models"
	"encoding/json"
	"app-server/infra/rest/generated/server/go"
)

//GetResourceTypes is REST handler for handling "GET" ResourceType requests.
func GetResourceTypes (w http.ResponseWriter, r *http.Request) {
	//Extract the Parameters
	vars := mux.Vars(r)
	parent_resource_type := vars["parent_resource_type"]

	labDB := MongoDBRepository.LabMongoDBRepository{Username: "dbuser", Password: "mith1234", IPAddress: "127.0.0.1", Port: "27017", Database: "dummyStore"}

	var resTypeOut []models.ResourceType
	var err error

	if len(parent_resource_type) == 0 {
		resTypeOut, err = labDB.GetResourceTypes()
	} else {
		resTypeOut, err = labDB.GetChildResourceTypes(parent_resource_type)
	}

	if err != nil {
		// Populate Error model.
	}
	var output swagger.ResourceTypesResponse

	for _,i := range resTypeOut {
		resType := swagger.ResourceTypeResponse{Resourcetypeid: i.ID, Resourcetypename:i.Type, Parentresourcetypeid:i.ParentTypeID}
		output.Items = append(output.Items, resType)
	}

	//Write Success Structure to the Body
	bytess, _ := json.Marshal(&output)
	w.Write(bytess)

}

