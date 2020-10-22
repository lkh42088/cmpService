package cron

import (
	"cmpService/common/messages"
	"cmpService/mcagent/config"
	"cmpService/mcagent/kvm"
	"cmpService/mcagent/repo"
	"cmpService/mcagent/svcmgrapi"
	"fmt"
	"github.com/robfig/cron"
)

var RegularSendToSvcmgrCronId cron.EntryID

func RegisterRegularMsg() {
	cronTime := fmt.Sprintf("@every 0h0m%ss", "30")
	id, err := CronSch.Cr.AddFunc(cronTime, func() {
		var msg messages.ServerRegularMsg
		cfg := config.GetMcGlobalConfig()
		addr := fmt.Sprintf("%s:%s", cfg.SvcmgrIp, cfg.SvcmgrPort)
		msg.SerialNumber = cfg.SerialNumber
		msg.Enable = repo.GlobalServerRepo.Enable
		msg.PrivateIp = cfg.ServerIp
		msg.PublicIp = cfg.ServerPublicIp
		msg.Port = cfg.McagentPort
		if repo.GlobalServerRepo.Enable {
			// Case 1: Send KeepAlive
			fmt.Printf("** RegisterRegularMsg(Cron ID:%d): Send keepalive msg to svcmgr\n", RegularSendToSvcmgrCronId)
			svcmgrapi.SendRegularMsg2Svcmgr(msg, addr, repo.GlobalServerRepo.Enable)
		} else {
			// Case 2: Send Registration
			fmt.Printf("** RegisterRegularMsg(Cron ID:%d): Send Registration msg to svcmgr\n", RegularSendToSvcmgrCronId)
			res := svcmgrapi.SendRegularMsg2Svcmgr(msg, addr, repo.GlobalServerRepo.Enable)
			if res == true {
				// Send ServerMsg
				svcmgrRestAddr := fmt.Sprintf("%s:%s", cfg.SvcmgrIp, cfg.SvcmgrPort)
				serverInfo := kvm.GetMcServerInfo()
				svcmgrapi.SendUpdateServer2Svcmgr(serverInfo, svcmgrRestAddr)
			}
		}
	})
	if err != nil {
		fmt.Println("RegisterRegularMsg: error ", err)
		return
	}
	RegularSendToSvcmgrCronId = id
	fmt.Println(">>> RegisterRegularMsg: Cron Id -", RegularSendToSvcmgrCronId)
}
