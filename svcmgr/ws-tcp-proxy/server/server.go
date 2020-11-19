package server

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"

	"golang.org/x/net/websocket"
)

// Run starts up the websocket tcp proxy.
func Run(c *Config) error {
	err := validatePort(c.Port)
	if err != nil {
		return err
	}

	portString := fmt.Sprintf(":%d", c.Port)

	log.Printf("[INFO] Listening on %s\n", portString)

	var proxyFunc = func(ws *websocket.Conn) {
		fmt.Println("proxyFunc")
		proxyHandler(ws, c)
	}

	var proxyFuncVnc = func(ws *websocket.Conn) {
		fmt.Println("proxyFuncVnc")
		proxyHandlerVNC(ws, c)
	}

	mux := http.NewServeMux()
	mux.Handle("/", websocket.Handler(proxyFunc))
	mux.Handle("/vnc/", websocket.Handler(proxyFuncVnc))

	srv := &http.Server{
		Addr:      portString,
		Handler:   mux,
		TLSConfig: &tls.Config{},
	}

	if c.CertManager != nil {
		srv.TLSConfig.GetCertificate = c.CertManager.GetCertificate
	}

	if c.CertManager != nil || (c.Cert != "" && c.Key != "") {
		err = srv.ListenAndServeTLS(c.Cert, c.Key)
		if err != nil {
			return err
		}
	} else {
		err = srv.ListenAndServe()
		if err != nil {
			return err
		}
	}

	return nil
}

func proxyHandler(ws *websocket.Conn, c *Config) {
	log.Println("-:Config", ws.Config())
	log.Println("-:Requset", ws.Request())
	conn, err := getConn(c)
	if err != nil {
		log.Printf("[ERROR] %v\n", err)
		return
	}

	if !c.TextMode {
		ws.PayloadType = websocket.BinaryFrame
	}

	doneChan := make(chan bool)

	go copyData(conn, ws, doneChan)
	go copyData(ws, conn, doneChan)

	<-doneChan
	conn.Close()
	ws.Close()
	<-doneChan
}

func proxyHandlerVNC(ws *websocket.Conn, c *Config) {
	log.Println("VNC:Config", ws.Config())
	log.Println("VNC:Requset", ws.Request())
	arr := strings.Split(ws.Config().Location.Path, "/")
	if len(arr) != 4 {
		return
	}
	address := arr[2]
	port := arr[3]
	c.Address = fmt.Sprintf("%s:%s", address, port)
	log.Println(">>>>>>>> VNC: address", c.Address)
	conn, err := getConn(c)
	if err != nil {
		log.Printf("[ERROR] %v\n", err)
		return
	}

	if !c.TextMode {
		ws.PayloadType = websocket.BinaryFrame
	}

	doneChan := make(chan bool)

	go copyData(conn, ws, doneChan)
	go copyData(ws, conn, doneChan)

	<-doneChan
	conn.Close()
	ws.Close()
	<-doneChan
}

func getConn(c *Config) (io.ReadWriteCloser, error) {
	if c.TCPTLS {
		config, err := getTLSConfig(c)
		if err != nil {
			return nil, err
		}

		return tls.Dial("tcp", c.Address, config)
	}

	return net.Dial("tcp", c.Address)
}

func getTLSConfig(c *Config) (*tls.Config, error) {
	config := &tls.Config{
		ServerName: strings.Split(c.Address, ":")[0],
	}

	if c.TCPTLSRootCA != "" {
		root, err := ioutil.ReadFile(c.TCPTLSRootCA)
		if err != nil {
			return nil, err
		}

		certPool := x509.NewCertPool()
		if !certPool.AppendCertsFromPEM([]byte(root)) {
			return nil, errors.New("failed to parse root certificate")
		}

		config.RootCAs = certPool
	}

	if c.TCPTLSCert != "" && c.TCPTLSKey != "" {
		certificate, err := tls.LoadX509KeyPair(c.TCPTLSCert, c.TCPTLSKey)
		if err != nil {
			return nil, fmt.Errorf("could not load client key pair: %s", err)
		}

		config.Certificates = []tls.Certificate{certificate}
	}

	return config, nil
}

func copyData(dst io.Writer, src io.Reader, doneChan chan<- bool) {
	io.Copy(dst, src)
	doneChan <- true
}

func validatePort(port int) error {
	if port < 1 || port > 65535 {
		return fmt.Errorf("Invalid port requested. Valid values are 1-65535.")
	}

	return nil
}
