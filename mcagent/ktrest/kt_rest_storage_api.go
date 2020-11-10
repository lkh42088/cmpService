package ktrest

import (
	"cmpService/common/lib"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

/**
 * STORAGE API
 */
//// Post Auth Token
//func PostAuthTokens() StorageAuthTokenResponse {
//	baseUrl, _ := url.Parse(storageBaseUrlPort + storageAuthTokenUrl)
//	req := StorageAuthRequest{}
//	response := StorageAuthTokenResponse{}
//
//	//Make request
//	req.Auth.Identity.Methods = append(req.Auth.Identity.Methods, MethodsPassword)
//	req.Auth.Identity.Password.User.Name = storageAccessKey
//	req.Auth.Identity.Password.User.Domain.Id = storageDomainId
//	req.Auth.Identity.Password.User.Password = storageSecretKey
//	req.Auth.Scope.Project.Id = storageProjectId
//	req.Auth.Scope.Project.Domain.Id = storageDomainId
//	pbytes, _ := json.Marshal(req)
//	body := bytes.NewBuffer(pbytes)
//	fmt.Println("# URL: ", req)
//
//	//Send API Query
//	resp, err := http.Post(baseUrl.String(), ContentTypeJson, body)
//	if err != nil {
//		fmt.Println("error:", err)
//	} else {
//		defer resp.Body.Close()
//	}
//
//	//Parsing data
//	data, err := ioutil.ReadAll(resp.Body)
//	err = json.Unmarshal(data, &response)
//
//	//tmp, _ := lib.PrettyPrint(data)
//	//fmt.Println("RESPONSE: ", string(tmp))
//	GlobalToken = resp.Header.Get("X-Subject-Token")
//	if len(response.Token.Catalog) < 1 ||
//		response.Token.Catalog[1].EndPoints == nil {
//		return response
//	}
//	GlobalAccountUrl = response.Token.Catalog[1].EndPoints[0].Url
//
//	return response
//}

// Get storage container name
func GetStorageAccount(auth StorageAuthTokenResponse) []StorageAccount {
	var response []StorageAccount
	if GlobalAccountUrl == "" {
		return response
	}

	// Request URL
	baseUrl := GlobalAccountUrl + formatJsonUrl
	req, _ := http.NewRequest("GET", baseUrl, nil)
	// Request HEADER
	req.Header.Add("X-Auth-Token", GlobalToken)
	req.Header.Add("Content-Type", ContentTypeJson)

	//fmt.Println("URL: ", req)

	//Send API Query
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		defer resp.Body.Close()
		//return response
	}

	//Parsing data
	data, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(data, &response)

	//fmt.Println("RES: ", response)
	return response
}

// Get storage container
func GetStorageContainer(containerName string) (code int, err error) {
	//var response []StorageContainer

	// Request URL
	baseUrl := GlobalAccountUrl + "/" + containerName
	req, _ := http.NewRequest("GET", baseUrl, nil)
	// Request HEADER
	req.Header.Add("X-Auth-Token", GlobalToken)
	req.Header.Add("Content-Type", ContentTypeJson)
	fmt.Println("URL: ", req)

	//Send API Query
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		defer resp.Body.Close()
	}

	//Parsing data
	if resp.StatusCode != http.StatusOK &&
		resp.StatusCode != http.StatusNoContent {
		return resp.StatusCode, fmt.Errorf("Error: %s\n", resp.Status)
	}

	//return fmt.Errorf("Success\n")
	return resp.StatusCode, nil
}

// Put storage container
func PutStorageContainer(token string, containerName string) (err error) {
	// Request URL
	baseUrl := GlobalAccountUrl + "/" + containerName
	req, _ := http.NewRequest("PUT", baseUrl, nil)
	// Request HEADER
	req.Header.Add("X-Auth-Token", token)
	req.Header.Add("Content-Type", ContentTypeJson)
	req.Header.Add("X-Storage-Policy", EconomyType)		// economy type
	fmt.Println("URL: ", req)

	//Send API Query
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		defer resp.Body.Close()
	}

	//Parsing data
	if resp.StatusCode != http.StatusCreated &&
		resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Error: %s\n", resp.Status)
	}

	return nil
}

// Delete storage container
func DeleteStorageContainer(containerName string) (err error) {
	// Request URL
	baseUrl := GlobalAccountUrl + "/" + containerName
	req, _ := http.NewRequest("DELETE", baseUrl, nil)
	// Request HEADER
	req.Header.Add("X-Auth-Token", GlobalToken)
	req.Header.Add("Content-Type", ContentTypeJson)
	fmt.Println("URL: ", req)

	//Send API Query
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		defer resp.Body.Close()
	}

	//Parsing data
	if resp.StatusCode != http.StatusOK &&
		resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("Error: %s\n", resp.Status)
	}

	return fmt.Errorf("Success\n")
}

// Make temp-url (Not complete - need to fix)
func GetStorageTempUrl() {
	method := "GET"
	path := fmt.Sprintf(storagePathUrl, "iwhan@nubes-bridge.com", "Nubes-HC", "")  // Storage db field : account url, filebox name, file name
	expired := int(time.Now().Add(ExpiredTime).Unix())
	baseUrl, _ := url.Parse(storageBaseUrl + path)
	params := url.Values{}

	// Get Signature
	hmacBody := fmt.Sprintf("%s\n%s\n%s", method, strconv.Itoa(expired), path)
	sig := ComputeHmac(hmacBody, storageSecretKey)  // Storage db field : secret key

	// TempUrl
	params.Add("temp_url_sig", sig)
	params.Add("temp_url_expires", strconv.Itoa(expired))  // Storage sb field : expired time
	baseUrl.RawQuery = params.Encode()

	fmt.Printf("URL : %s\n", baseUrl.String())

	//Send API Query
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println("error:", err)
	} else {
		defer resp.Body.Close()
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
	}

	b, _ := lib.PrettyPrint(data)
	fmt.Println("RESPONSE: ", string(b))
}

