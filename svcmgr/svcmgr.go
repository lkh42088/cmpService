package main

import (
	"cmpService/common/config"
	db2 "cmpService/common/db"
	"cmpService/common/mariadblayer"
	config2 "cmpService/svcmgr/config"
	"cmpService/svcmgr/rest"
	"cmpService/svcmgr/ws-tcp-proxy/server"
	"flag"
	"fmt"
	"sync"
)

func main() {
	configFile := flag.String("file", "svcmgr.conf",
		"Input configuration file")
	flag.Parse()
	config2.SetConfig(*configFile)
	cfg := config2.ReadConfig(config2.SvcmgrConfigPath)
	db, err := config2.SetMariaDB(cfg.MariaUser, cfg.MariaPassword, cfg.MariaDb,
		cfg.MariaIp, 3306)
	//db, err := SetMariaDB()
	if err != nil {
		fmt.Println("Main: ERROR - ", err)
		return
	}

	config2.SetInfluxDB()

	var wg sync.WaitGroup
	wg.Add(2)

	go server.SetWebsocketServer(&wg, "8083")

	go SetRestServer(&wg, db)

	wg.Wait()
}

func SetMariaDBOld() (db *mariadblayer.DBORM, err error) {
	cfg := config2.ReadConfig(config2.SvcmgrConfigPath)
	dbconfig, err := config.NewDBConfig("mysql",
		cfg.MariaUser, cfg.MariaPassword, cfg.MariaDb,
		cfg.MariaIp, 3306)
	if err != nil {
		fmt.Println("[SetMariaDB] Error:", err)
		return
	}
	config2.SvcmgrGlobalConfig.MariadbConfig = *dbconfig
	dataSource := db2.GetDataSourceName(dbconfig)
	db, err = mariadblayer.NewDBORM(dbconfig.DBDriver, dataSource)
	if err != nil {
		fmt.Println("[SetMariaDB] Error:", err)
		return
	}
	config2.SvcmgrGlobalConfig.Mariadb = *db
	return db, err
}

func SetRestServer(wgParent *sync.WaitGroup, db *mariadblayer.DBORM) {
	cfg := config2.ReadConfig(config2.SvcmgrConfigPath)
	restServer := fmt.Sprintf("%s:%s", cfg.RestServerIp, cfg.RestServerPort)
	config2.SvcmgrGlobalConfig.RestServer = restServer
	config2.SvcmgrConfigStore = &cfg
	webserver := fmt.Sprintf("%s:%s", cfg.WebServerIP, cfg.WebServerPort)
	rest.WebServerAddress = webserver
	rest.RunAPI(restServer, db)
	wgParent.Done()
}

/*
func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("...................... : ", r)
	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}


func setupRoutes() {
	http.HandleFunc("/v1/users/fileUpload", uploadFile)
	http.ListenAndServe(":8081", nil)
}*/
/*
func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.GET("/uploadpage", func(c *gin.Context) {
		title := "upload single file"
		c.HTML(http.StatusOK, "uploadfile.html", gin.H{
			"page": title,
		})
	})

	r.POST("/upload", uploadSingle)

	return r
}

func uploadSingle(c *gin.Context) {
	// single file
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	log.Println(file.Filename)

	// Upload the file to specific dst.
	filename := filepath.Base(file.Filename)
	uploadPath := "./example/upload/" + filename
	log.Println(filename)
	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.JSON(200, gin.H{
		"status":    "posted",
		"file name": file.Filename,
	})
}*/
