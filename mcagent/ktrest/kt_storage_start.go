package ktrest

import (
	"cmpService/common/ktapi"
	"cmpService/mcagent/config"
	"cmpService/mcagent/repo"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

func CheckKtAccount() error {
	// 1. KT account info check
	conf := config.GetMcGlobalConfig()
	server, err := repo.GetServerFromDbByIp(conf.ServerIp)
	if err != nil {
		return fmt.Errorf("! Error : Not found server info.\n")
	}
	//fmt.Printf("CheckKtAccount() GetServerFromDbByIp : %+v\n", server)

	// 2. KT ucloud user authorization
	token, resp := ktapi.PostAuthTokens()
	ktapi.GlobalToken = token
	ktapi.GlobalAccountUrl = resp.Token.Catalog[1].EndPoints[0].Url

	if ktapi.GlobalAccountUrl == "" {
		return fmt.Errorf("! Alarm : Not found the KT Ucloud information.\n")
	}

	// 3. Store auth_url
	_, err = repo.UpdateKtAuthUrl2Db(conf.ServerIp, ktapi.GlobalAccountUrl)
	if err != nil {
		return fmt.Errorf("! Error : %s\n", err)
	}

	// 4. Send to Svcmgr
	obj := KtAuthUrl{
		AuthUrl: ktapi.GlobalAccountUrl,
		CpIdx:   server.CompanyIdx,
		Ip:      conf.ServerIp,
	}
	if !SendUpdateAuthUrl2Svcmgr(obj ,conf.SvcmgrIp + ":" + conf.SvcmgrPort) {
		return fmt.Errorf("! Error : Failed to sync svcmgr DB.")
	}

	return nil
}

func ConfigurationForKtContainer() error {
	// get server info
	data := repo.GetMcServer()

	// data valid check
	if data == nil {
		return errors.Errorf("! Error: Server data is nil.\n")
	}

	// cronsch account check
	if data.UcloudAccessKey == "" {
		return nil
	}

	// container check
	ipNum := strings.Split(data.IpAddr, ".")

	// container name : serial_number + _ + last ip (ex: SN87_87)
	if len(ipNum) == 4 {
		ktapi.GlobalContainerName = data.SerialNumber + "_" + ipNum[len(ipNum)-1]
	} else {
		return errors.New("! Error: Server IP is invalid.\n")
	}

	code, err := GetStorageContainer(ktapi.GlobalContainerName)
	//fmt.Println("code : ", code)

	// create container
	if code != http.StatusOK &&
		code != http.StatusNoContent {
		err2 := PutStorageContainer(ktapi.GlobalToken, ktapi.GlobalContainerName)
		if err2 != nil {
			return err
		}
	}

	return nil
}
