package rest

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/globalsign/mgo/bson"
	"io/ioutil"
	"net/http"
	"net/url"
	"nubes/collector/config"
	"nubes/collector/collectdevice"
	"sort"
	"strings"
	"sync"
	"testing"
)

func TestRestRouter(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	//RestRouter(&wg)
	Start(&wg)
	wg.Wait()
}

func TestRestGet(t *testing.T) {
	req, err := http.NewRequest("GET",
		"http://localhost:8884" + apiPathPrefix + apiDevice, nil)
	if err != nil {
		fmt.Println("NewRequest err:", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("response err:", err)
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("resp:", string(data))
}

func TestRestGet2(t *testing.T) {
	resp, err := http.Get("http://localhost:8884" + apiPathPrefix + apiDevice)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("resp:", string(data))
}

// postform : old version
func TestRestPost(t *testing.T) {
	dev := collectdevice.ColletDevice{
		Id:            "1",
		Ip:            "192.168.122.19",
		Port:          161,
		SnmpCommunity: "nubes",
	}
	pbytes, _ := json.Marshal(dev)
	buff := bytes.NewBuffer(pbytes)
	url := "http://localhost:8884" + apiPathPrefix + apiDevice
	resp, err := http.Post(url, "application/json", buff)
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("resp:", string(data))
}

// json : new version
func TestRestPort2(t *testing.T) {
	dev := []collectdevice.ColletDevice{
		{
			Id:            "1",
			Ip:            "127.0.0.1",
			Port:          161,
			SnmpCommunity: "nubes",
		}, {
			Id:            "2",
			Ip:            "211.211.211.211",
			Port:          161,
			SnmpCommunity: "nubes",
		},
	}
	pbytes, _ := json.Marshal(dev)
	req, _ := http.NewRequest("POST",
		"http://127.0.0.1:7708" + apiPathPrefix + apiDevice,
		bytes.NewBuffer(pbytes))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("response err:", err)
		return
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(data)
}

func TestRestDelete(t *testing.T) {
	req, err := http.NewRequest("DELETE",
		"http://localhost:7708" + apiPathPrefix + apiDevice + "/all", nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("response err:", err)
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("resp:", string(data))
}

func TestId(t *testing.T) {
	objID := bson.NewObjectId()
	id := collectdevice.ID(fmt.Sprintf("%x",string(objID)))
	fmt.Printf("%s\n", id)
	fmt.Printf("%s\n", string(objID))
}

func TestInput(t *testing.T) {
	var conf config.CollectorConfig
	config.SetConfigByField("svcmgrip", "1.1.1.1", &conf)
	fmt.Println(conf)
}

const lbURL = "https://api.ucloudbiz.olleh.com/loadbalancer/v2/client/api"
const dbURL = "https://api.ucloudbiz.olleh.com/nas/v2/client/api"
const serverURL = "https://api.ucloudbiz.olleh.com/server/v2/client/api"
const watchURL = "https://api.ucloudbiz.olleh.com/watch/v2/client/api"
// CB KEY
const apiKey = "fYGnzisuTXlXVgxw9Des2me-CbQ-d2x1oFDAczUa2DknxtwbCXjlYb25CobtJWXpTbvtnhC3pujtZw-O4Qaq-Q"
const secretKey = "y8-kgAG1cBnunCZQy-SwnKC3m6nh4akXj1p3HGuFesnJB7speDBCZvhv6zjzz3n9LZ9797RnXwlBJ7MuwaM63w"
// NB KEY
//const apiKey = "zhY6AqhrBuxzBYahVleF57nXYia3wNg1iddLL0ElgwiKU9V76Iu-g2_Qvh2jE5QxSYT9n_z47nahFz0qI-Byug"
//const secretKey = "DGcIyrljdy28mKMgns9pEkPchMugzmmxnbB1cUU4fgcIvrjpDALFXqLVUhaweRQrM3PGuY1f6N1NO4Nw_etbqA"

// watch listMetricy15
const listMetrics = "met1ricname=CPUUtilization&command=listMetrics"
// Server VM TEST
const productypes = "command=listAvailableProductTypes"
const listIpAddr = "command=listPublicIpAddresses"
const listVMCharge = "command=listVirtualMachineForCharge"
const listAccount = "command=listAccounts"
const listVM = "command=listVirtualMachines"
const listZone = "command=listZones"
const listNetworkFlatRate = "command=listNetworkFlatRate"
const listNetwork = "command=listNetworks"
const listNetUsage = "command=listNetworkUsages&startdate=2020-04-01&enddate=2020-04-30"

func SortCommandLine(command string) string {
	var result string = ""

	// Convert to small letter command character
	// and sort command field alphabetically
	convert := strings.Split(strings.ToLower(command), "&")
	sort.Strings(convert)
	for i := range convert {
		result += convert[i] + "&"
	}
	return result[:len(result)-1]
}

func ComputeHmac(message string, secret string) string {
	// HMAC-SHA1 hashing
	h := hmac.New(sha1.New, []byte(secret))
	h.Write([]byte(message))
	message = base64.StdEncoding.EncodeToString(h.Sum(nil))
	// URL UTF-8 encoding
	return url.QueryEscape(message)
}

func KtRestGet(ktURL string, command string) string {
	// Add apiKey to command
	tmpStr := command + "&response=xml"
	tmpStr = tmpStr + "&apiKey=" + apiKey

	// command field Tolower() and sort()
	sortedStr := SortCommandLine(tmpStr)

	// Get signature
	sig := ComputeHmac(sortedStr, secretKey)

	// Make full api command
	command = ktURL + "?" + tmpStr + "&signature=" + sig
	fmt.Printf("URL : %s\n", command)

	// Send command
	resp, err := http.Get(command)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}

	return string(data)
}

func TestKtRestApi(t *testing.T) {
	response := KtRestGet(serverURL, listVM)
	fmt.Println(response)
}