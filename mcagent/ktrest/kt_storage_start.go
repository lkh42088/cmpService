package ktrest

import (
	"cmpService/mcagent/config"
	"cmpService/mcagent/repo"
	//"cmpService/mcagent/svcmgrapi"
	"fmt"
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
	PostAuthTokens()
	if GlobalAccountUrl == "" {
		return fmt.Errorf("! Alarm : Not found the KT Ucloud information.\n")
	}

	// 3. Store auth_url
	_, err = repo.UpdateKtAuthUrl2Db(conf.ServerIp, GlobalAccountUrl)
	if err != nil {
		return fmt.Errorf("! Error : %s\n", err)
	}

	// 4. Send to Svcmgr
	obj := KtAuthUrl{
		AuthUrl: GlobalAccountUrl,
		CpIdx:   server.CompanyIdx,
		Ip:      conf.ServerIp,
	}
	SendUpdateAuthUrl2Svcmgr(obj ,conf.SvcmgrIp + ":" + conf.SvcmgrPort)

	return nil
}
