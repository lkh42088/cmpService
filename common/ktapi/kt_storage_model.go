package ktapi

import (
	"time"
)

// KT Auth Request
type StorageUser struct {
	Id       string        `json:"id"`
	Name     string        `json:"name"`
	Domain   StorageDomain `json:"domain"`
	Password string        `json:"password"`
}

type StoragePass struct {
	User StorageUser `json:"user"`
}

type StorageIdentity struct {
	Methods  []string    `json:"methods"`
	Password StoragePass `json:"password"`
}

type StorageAuth struct {
	Identity StorageIdentity `json:"identity"`
	Scope    StorageScope    `json:"scope"`
}

type StorageDomain struct {
	Id 			string		`json:"id"`
	Name 		string 		`json:"name"`
}

type StorageProject struct {
	Id     string        `json:"id"`
	Domain StorageDomain `json:"domain"`
	Name   string        `json:"name"`
}

type StorageScope struct {
	Project StorageProject `json:"project"`
}

type StorageAuthRequest struct {
	Auth StorageAuth `json:"auth"`
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
	Status 		string                  `json:"status"`
	Updated 	time.Time                  `json:"updated`
	MediaTypes 	[]StorageAuthMediaTypes `json:"media-types"`
	Id 			string                  `json:"id"`
	Links 		[]StorageAuthLinks       `json:"links"`
}

type StorageAuthVersions struct {
	Values 		[]StorageAuthValues `json:"values"`
}

type StorageAuthResponse struct {
	Versions StorageAuthVersions  `json:"versions"`
	Error    StorageResponseError `json:"error"`
}

/** RESPONSE TOKEN */
type StorageAuthRole struct {
	Id 			string 				`json:"id"`
	Name 		string				`json:"name"`
}

type StorageAuthToken struct {
	Methods []string             `json:"methods"`
	Roles   []StorageAuthRole    `json:"roles"`
	Expires time.Time            `json:"expires_at"`
	Project StorageProject       `json:"project"`
	Catalog []StorageAuthCatalog `json:"catalog"`
	User    StorageUser          `json:"user"`
	Audit   []string             `json:"audit_ids"`
	Issued  time.Time            `json:"issued_at"`
}

type StorageEndpoint struct {
	RegionId 	string 				`json:"region_id"`
	Url 		string 				`json:"url"`
	Region 		string 				`json:"region"`
	Interface 	string 				`json:"interface"`
	Id 			string 				`json:"id"`
}

type StorageAuthCatalog struct {
	EndPoints 	[]StorageEndpoint `json:"endpoints"`
	Type 		string             `json:"type"`
	Id 			string           `json:"id"`
	Name 		string             `json:"name"`
}

type StorageAuthTokenResponse struct {
	Token StorageAuthToken     `json:"token"`
	Error StorageResponseError `json:"error"`
}

/** Account */
type StorageAccount struct {
	Count 		int 			`json:"count"`
	Bytes 		int 			`json:"bytes"`
	Name 		string 			`json:"name"`
}

/** Auth Url */
type KtAuthUrl struct {
	AuthUrl 	string 			`json:"authUrl"`
	CpIdx 		int 			`json:"cpIdx"`
	Ip 			string			`json:"ip"`
}

/** ERROR */
type StorageResponseError struct {
	Message 	string 			`json:"message"`
	Code 		int 			`json:"code"`
	Title 		string 			`json:"title"`
}