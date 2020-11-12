package ktapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Post Auth Token
func PostAuthTokens() (token string, response StorageAuthTokenResponse) {
	baseUrl, _ := url.Parse(storageBaseUrlPort + storageAuthTokenUrl)
	req := StorageAuthRequest{}

	//Make request
	req.Auth.Identity.Methods = append(req.Auth.Identity.Methods, MethodsPassword)
	req.Auth.Identity.Password.User.Name = StorageAccessKey
	req.Auth.Identity.Password.User.Domain.Id = StorageDomainId
	req.Auth.Identity.Password.User.Password = StorageSecretKey
	req.Auth.Scope.Project.Id = StorageProjectId
	req.Auth.Scope.Project.Domain.Id = StorageDomainId
	pbytes, _ := json.Marshal(req)
	body := bytes.NewBuffer(pbytes)
	fmt.Println("# URL: ", req)

	//Send API Query
	resp, err := http.Post(baseUrl.String(), ContentTypeJson, body)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		defer resp.Body.Close()
	}

	//Parsing data
	data, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(data, &response)

	//tmp, _ := lib.PrettyPrint(data)
	//fmt.Println("RESPONSE: ", string(tmp))
	//GlobalToken := resp.Header.Get("X-Subject-Token")
	globalToken := resp.Header.Get("X-Subject-Token")
	if len(response.Token.Catalog) < 1 ||
		response.Token.Catalog[1].EndPoints == nil {
		return globalToken, response
	}
	//GlobalAccountUrl = response.Token.Catalog[1].EndPoints[0].Url

	return globalToken, response
}
