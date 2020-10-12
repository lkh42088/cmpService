package ktrest

import "time"

// VM API
const lbURL = "https://api.ucloudbiz.olleh.com/loadbalancer/v2/client/api"
const dbURL = "https://api.ucloudbiz.olleh.com/nas/v2/client/api"
const serverURL = "https://api.ucloudbiz.olleh.com/server/v2/client/api"
const watchURL = "https://api.ucloudbiz.olleh.com/watch/v2/client/api"

// STORAGE API
const storageBaseUrl = "https://ssproxy2.ucloudbiz.olleh.com:5000"
const storagePathUrl = "/v1/%s/%s/%s"
const storageAuthUrl = "/v3/auth"
const getTempUrl = "v1/account/container/object?temp_url_sig=%s&temp_url_expires=%s"

const storageAccessKey = "iwhan$nubes-bridge.com"
const storageSecretKey = "MTYwMTg2MzU1OTE2MDE4NjI5MTk1MTQ2"
const storageProjectId = "fa632a4a0d04488c93b7184be92df4c8"
const storageDomainId = "42a37f949dcd48a3a805fe0d2d3a7da5"

const EXPIRED_TIME = 60 * time.Minute