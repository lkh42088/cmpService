package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Post Auth Token
func PostAuthTokens() StorageAuthTokenResponse {
	baseUrl, _ := url.Parse(storageBaseUrlPort + storageAuthTokenUrl)
	req := StorageAuthRequest{}
	response := StorageAuthTokenResponse{}

	//Make request
	req.Auth.Identity.Methods = append(req.Auth.Identity.Methods, MethodsPassword)
	req.Auth.Identity.Password.User.Name = storageAccessKey
	req.Auth.Identity.Password.User.Domain.Id = storageDomainId
	req.Auth.Identity.Password.User.Password = storageSecretKey
	req.Auth.Scope.Project.Id = storageProjectId
	req.Auth.Scope.Project.Domain.Id = storageDomainId
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
	GlobalToken = resp.Header.Get("X-Subject-Token")
	if len(response.Token.Catalog) < 1 ||
		response.Token.Catalog[1].EndPoints == nil {
		return response
	}
	GlobalAccountUrl = response.Token.Catalog[1].EndPoints[0].Url

	return response
}
