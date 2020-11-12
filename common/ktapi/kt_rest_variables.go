package ktapi

import "time"

/** Global Variables */
var GlobalToken string
var GlobalAccountUrl string
var GlobalContainerName string

// KT Storage variables
const ExpiredTime = 60 * time.Minute
const ContentTypeJson = "application/json"
const ContentTypeBinary = "binary/octet-stream"
const Range4096 = "4096"
const MethodsPassword = "password"
const EconomyType = "ec"
const BackupFilePermission = 0644
const FILE_BLOCK_500M = 500 * 1024 * 1024

// DB
const StorageAccessKey = "iwhan@nubes-bridge.com"
const StorageSecretKey = "MTYwMTg2MzU1OTE2MDE4NjI5MTk1MTQ2"
const StorageProjectId = "fa632a4a0d04488c93b7184be92df4c8"
const StorageDomainId = "42a37f949dcd48a3a805fe0d2d3a7da5"

// VM API
const lbURL = "https://api.ucloudbiz.olleh.com/loadbalancer/v2/client/api"
const dbURL = "https://api.ucloudbiz.olleh.com/nas/v2/client/api"
const serverURL = "https://api.ucloudbiz.olleh.com/server/v2/client/api"
const watchURL = "https://api.ucloudbiz.olleh.com/watch/v2/client/api"

// STORAGE API
const storageBaseUrl = "https://ssproxy2.ucloudbiz.olleh.com"
const storageBaseUrlPort = "https://ssproxy2.ucloudbiz.olleh.com:5000"
const storageVerUrl = "/v1/"
const formatJsonUrl = "?format=json"
const storagePathUrl = "/v1/%s/%s/%s"
const storageAuthUrl = "/v3/auth"
const storageAuthTokenUrl = storageAuthUrl + "/tokens"
const getTempUrl = "v1/account/container/object?temp_url_sig=%s&temp_url_expires=%s"



