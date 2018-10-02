package models

/*
Basic tasks in the LabReservation system to define RBAC.
TASK_Reserve
TASK_CREATE_ResourceType, TASK_READ_ResourceType, TASK_UPDATE_ResourceType, TASK_DELETE_ResourceType
TASK_CREATE_Resource, TASK_READ_Resource, TASK_UPDATE_Resource, TASK_DELETE_Resource
TASK_CREATE_User, TASK_READ_User, TASK_UPDATE_User, TASK_DELETE_User
TASK_CREATE_ReservationCapability, ASK_READ_ReservationCapability, TASK_UPDATE_ReservationCapability, TASK_DELETE_ReservationCapability
*/

// User structure contains details of registered user.
/* This will be a list
   addUser, GetUser, UpdateUser, DeleteUser */
type User struct {
	UserID        string `json:"userid"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	EmailID       string `json:"email"`
	Phone         string `json:"phone"`
	ProfilePicURL string `json:"profilepicurl"`
}

// ResourceType - Specify the type of the resource - server/switch/printer/router etc.
/* This will be a list
   addResourceType, GetResourceType, UpdateResourceType, DeleteResourceType */
type ResourceType struct {
	ID   string `json:"id"`
	Type string `json:"resourcetype"`
}

// ResourceSubType - Specify the type of the resource - server/switch/printer/router etc.
/* This will be a list
   addResourceSubType, GetResourceSubType, UpdateResourceSubType, DeleteResourceSubType */
type ResourceSubType struct {
	ID      string `json:"id"`
	Type    string `json:"resourctypeid"`
	SubType string `json:"resourcesubtype"`
}

// ResourceInfo contains more details about the Resource.
type ResourceInfo struct {
	IPAddress  []string
	IPConsole  string
	PowerTower string
	Rack       string
	Lab        string
	Building   string
}

// ResourcePortMatrix contains port matrix connectivity information.
type ResourcePortMatrix struct {
	SourcePort       string `json:"sourceport"`
	RemotePort       string `json:"remoteport"`
	RemoteResourceID string `json:"remoteresourceid"`
}

// Resource contains all details about the resource, its owner and reservation details.
/* This will be a list
   addResource, GetResource, UpdateResource, DeleteResource */
type Resource struct {
	ID                string               `json:"resourceid"`
	Name              string               `json:"resourcename"`
	Description       string               `json:"description"`
	ResourceTypeID    string               `json:"resourcetypeid"`
	ResourceSubTypeID string               `json:"resourcesubtypeid"`
	Infos             []ResourceInfo       `json:"resourceinfos"`
	PortMatrix        []ResourcePortMatrix `json:"portmatrix"`
	ResourceOwnerID   string               `json:"resourceownerid"`
	ReservationMatrix []ReservationMatrix  `json:"reservationmatrix"`
}

// ResourceReserveCapabilityMatrix contains list of resources that can be reserved by a user.
// Will be used for validation before a user reserves a resource.
/* This will be a list
   addReserveCapability - For a given user, API adds reserve capability for the given resource/s
   updateReserveCapability
   getReserveCapability
   deleteReserveCapability
*/
type ResourceReserveCapabilityMatrix struct {
	UserID     string `json:"userid"`
	ResourceID string `json:"resourceid"`
}

// ReservationMatrix stores start and end time of the reservation
/*
	resourceReserve API will populate the data and updates the Resource
 */
type ReservationMatrix struct {
	UserID    string `json:"userid"`
	StartTime string `json:"starttime"` // Need to think about the data type for time.
	EndTime   string `json:"endtime"`
}
