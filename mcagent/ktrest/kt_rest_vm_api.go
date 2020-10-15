package ktrest

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

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

type SecurityGroup struct {
	Data interface{}
}

type AffinityGroup struct {
	Data interface{}
}

type Tags struct {
	Tags string
}

type SecondaryIp struct {
	Ip string
}

type Nic struct {
	Id           string        `json:"id"`
	NetworkId    string        `json:"networkid"`
	NetworkName  string        `json:"networkname"`
	Netmask      string        `json:"netmask"`
	Gateway      string        `json:"gateway"`
	IpAddress    string        `json:"ipaddress"`
	IsolationUri string        `json:"isolationuri"`
	BroadcastUri string        `json:"broadcasturi"`
	TrafficType  string        `json:"traffictype"`
	NicType      string        `json:"type"`
	IsDefault    bool          `json:"isdefault"`
	MacAddress   string        `json:"macaddress"`
	SecondaryIp  []SecondaryIp `json:"secondaryip"`
}

type VirtualMachine struct {
	Id                    string          `json:"id"`
	Name                  string          `json:"name"`
	DisplayName           string          `json:"displayname"`
	Account               string          `json:"account"`
	UserId                string          `json:"userid"`
	Username              string          `json:"username"`
	DomainId              string          `json:"domainid"`
	Domain                string          `json:"domain"`
	Created               string          `json:"created"`
	State                 string          `json:"state"`
	HaEnable              bool            `json:"haenable"`
	ZoneId                string          `json:"zoneid"`
	ZoneName              string          `json:"zonename"`
	TemplateId            string          `json:"templateid"`
	TemplateName          string          `json:"templatename"`
	TemplateDisplayText   string          `json:"templatedisplaytext"`
	PasswordEnabled       bool            `json:"passwordenabled"`
	ServiceOfferingId     string          `json:"serviceofferingid"`
	ServiceOfferingName   string          `json:"serviceofferingname"`
	CpuNumber             int             `json:"cpunumber"`
	CpuSpeed              int             `json:"cpuspeed"`
	Memory                int             `json:"memory"`
	GuestOsId             string          `json:"guestosid"`
	RootDeviceId          int             `json:"rootdeviceid"`
	RootDeviceType        string          `json:"rootdevicetype"`
	SecurityGroup         []SecurityGroup `json:"securitygroup"`
	Nic                   []Nic           `json:"nic"`
	Hypervisor            string          `json:"hypervisor"`
	AffinityGroup         []AffinityGroup `json:"affinitygroup"`
	IsDynamicallyScalable bool            `json:"isdynamicallyscalable"`
	OsTypeId              int             `json:"ostypeid"`
	Tags                  []Tags          `json:"tags"`
}

type ListVirtualMachinesResponse struct {
	VirtualMachine []VirtualMachine `json:"virtualmachine"`
	Count          int              `json:"count"`
}

type ListVirtualMachine struct {
	List ListVirtualMachinesResponse `json:"listvirtualmachinesresponse"`
}

func KtRestGet(ktURL string, command string) string {
	// Add apiKey to command
	tmpStr := command + "&response=json"
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
	//json.Unmarshal(data, resp.Body)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}

	// data binding to struct
	var list ListVirtualMachine
	err = json.Unmarshal(data, &list)
	fmt.Printf("%+v\n", list)

	return string(data)
}


const resellerApiKey = "asfupvb9-abui-gaeu-z"
const resellerURL = "https://ucloudbiz.kt.com/jv_ssl_key_openapi.jsp"
const chargeVM = "startDate=2020-01&endDate=2020-03&type=serviceChargeInfoAccount&emailId=fin_bmetal1@vple.net"
const chargeListVM = "startDate=2020-01&endDate=2020-03&resellerKey=" + resellerApiKey + "&type=billingInfoListAccounts"

func KtChargeGet(ktURL string, command string) string {
	baseUrl, _ := url.Parse(resellerURL)
	params := url.Values{}

	// big
	params.Add("command", "listCharges")
	//params.Add("type", "billingInfoListAccounts")
	//params.Add("type", "useServiceListAccounts")
	params.Add("type", "serviceChargeInfoAccount")
	params.Add("emailId", "fin_bmetal1@vple.net")
	params.Add("startDate", "2020-04")
	params.Add("endDate", "2020-04")
	params.Add("resellerKey", resellerApiKey)
	params.Add("response", "json")

	baseUrl.RawQuery = params.Encode()

	fmt.Printf("URL : %s\n", baseUrl.String())

	//Send API Query
	resp, err := http.Get(baseUrl.String())
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

