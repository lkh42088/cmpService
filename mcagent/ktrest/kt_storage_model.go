package ktrest

import "time"

// KT Auth Request
type StorageUser struct {
	Name 		string  		`json:"name"`
	Domain  	StorageDomain 	`json:"domain"`
	Password 	string 			`json:"password"`
}

type StoragePass struct {
	User 		StorageUser `json:"user"`
}

type StorageIdentity struct {
	Methods 	[]string 	`json:"methods"`
	Password 	StoragePass	`json:"password"`
}

type StorageAuth struct {
	Identity 	StorageIdentity	`json:"identity"`
}

type StorageDomain struct {
	Id 			string		`json:"id"`
}

type StorageProject struct {
	Id 			string			`json:"id"`
	Domain 		StorageDomain	`json:"domain"`
}

type StorageScope struct {
	Project		StorageProject	`json:"project"`
}

type StorageAuthRequest struct {
	Auth 		StorageAuth		`json:"auth"`
	Scope 		StorageScope	`json:"scope"`
}

// KT Auth Response
type StorageAuthMediaTypes struct {
	Base        string		`json:"base"`
	Type 		string 		`json:"type"`
}

type StorageAuthLinks struct {
	Href 		string 		`json:"href"`
	Type 		string 		`json:"type"`
	Rel 		string 		`json:"rel"`
}

type StorageAuthValues struct {
	Status 		string					`json:"status"`
	Updated 	time.Time				`json:"updated`
	MediaTypes 	[]StorageAuthMediaTypes 	`json:"media-types"`
	Id 			string 					`json:"id"`
	Links 		[]StorageAuthLinks 		`json:"links"`
}

type StorageAuthVersions struct {
	Values 		[]StorageAuthValues		`json:"values"`
}

type StorageAuthResponse struct {
	Versions 	StorageAuthVersions		`json:"versions"`
}