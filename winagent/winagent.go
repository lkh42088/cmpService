package main

import (
	"cmpService/winagent/agent"
	"flag"
	"golang.org/x/sys/windows/svc"
	"io/ioutil"
	"time"
)

/**  global variable */
var CONFIG_PATH = "C:\\Program Files\\Nubes\\winagent.conf"

// 서비스 Type
type CMPWindowService struct {
}

// svc.Handler 인터페이스 구현
func (srv *CMPWindowService) Execute(args []string, req <-chan svc.ChangeRequest, stat chan<- svc.Status) (svcSpecificEC bool, exitCode uint32) {
	stat <- svc.Status{State: svc.StartPending}

	// 실제 서비스 내용
	stopChan := make(chan bool, 1)
	//go runBody(stopChan)
	go ServiceStart()

	stat <- svc.Status{State: svc.Running, Accepts: svc.AcceptStop | svc.AcceptShutdown}

LOOP:
	for {
		// 서비스 변경 요청에 대해 핸들링
		switch r := <-req; r.Cmd {
		case svc.Stop, svc.Shutdown:
			stopChan <- true
			break LOOP

		case svc.Interrogate:
			stat <- r.CurrentStatus
			time.Sleep(100 * time.Millisecond)
			stat <- r.CurrentStatus

		//case svc.Pause:
		//case svc.Continue:
		}
	}

	stat <- svc.Status{State: svc.StopPending}
	return
}

/*** 서비스에서 실제 하는 일 ***/
func ServiceStart() {
	conf := flag.String("file", CONFIG_PATH,
		"Input configuration file")
	flag.Parse()
	agent.Start(*conf)
}

func runBody(stopChan chan bool) {
	for {
		select {
		case <-stopChan:
			return
		default:
			// 10초 마다 현재시간 새로 쓰기
			time.Sleep(10 * time.Second)
			ioutil.WriteFile("C:/temp/winagent_log.txt", []byte(time.Now().String()), 0)
		}
	}
}

func main() {
	err := svc.Run("CMPWindowService", &CMPWindowService{})
	//err := debug.Run("DummyService", &dummyService{}) //콘솔출력 디버깅시
	if err != nil {
		ioutil.WriteFile("C:/temp/winagent_log.txt", []byte(err.Error()), 0)
		panic(err)
	} else {
		ioutil.WriteFile("C:/temp/winagent_log.txt", []byte("service run success"), 0)
	}
}

