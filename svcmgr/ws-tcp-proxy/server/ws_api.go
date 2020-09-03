package server

import (
	"crypto/tls"
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
)

func SetWebsocketServer(wgParent *sync.WaitGroup, port string) {
	RunWs(wgParent, port)
}

func RunWs(wgParent *sync.WaitGroup, port string) error {
	var proxyFuncVnc = func(ws *websocket.Conn) {
		fmt.Println("proxyFuncCustom")
		proxyHandlerCustom(ws)
	}

	mux := http.NewServeMux()
	mux.Handle("/vnc/", websocket.Handler(proxyFuncVnc))

	srv := &http.Server{
		Addr:      fmt.Sprintf(":%s", port),
		Handler:   mux,
		TLSConfig: &tls.Config{},
	}
	fmt.Println("RunWs: server ", port)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println("RunWs error: ", err)
		wgParent.Done()
		return err
	}

	wgParent.Done()
	return nil
}

func proxyHandlerCustom(ws *websocket.Conn) {
	log.Println("VNC:Config Path", ws.Config().Location.Path)
	//log.Println("VNC:Requset", ws.Request())
	arr := strings.Split(ws.Config().Location.Path, "/")
	if len(arr) != 4 {
		fmt.Printf("[ERROR] arr len %d\n", len(arr))
		return
	}
	address := arr[2]
	port := arr[3]
	addr := fmt.Sprintf("%s:%s", address, port)
	conn, err := getConnCustom(addr)
	if err != nil {
		//log.Printf("[ERROR] %v\n", err)
		fmt.Printf("[ERROR] %v\n", err)
		return
	}

	ws.PayloadType = websocket.BinaryFrame
	doneChan := make(chan bool)

	go copyData(conn, ws, doneChan)
	go copyData(ws, conn, doneChan)

	<-doneChan
	conn.Close()
	ws.Close()
	<-doneChan
}

func getConnCustom(addr string) (io.ReadWriteCloser, error) {
	fmt.Println(">>> getConnCustom Tcp addr:", addr)
	return net.Dial("tcp", addr)
}
