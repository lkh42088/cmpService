package ktrest

// VM API
const lbURL = "https://api.ucloudbiz.olleh.com/loadbalancer/v2/client/api"
const dbURL = "https://api.ucloudbiz.olleh.com/nas/v2/client/api"
const serverURL = "https://api.ucloudbiz.olleh.com/server/v2/client/api"
const watchURL = "https://api.ucloudbiz.olleh.com/watch/v2/client/api"

// STORAGE API
const storageBaseUrl = "https://ssproxy2.ucloudbiz.olleh.com"
const storageBaseUrlPort = "https://ssproxy2.ucloudbiz.olleh.com:5000"
const storageAccountUrl = "/v1/"
const formatJsonUrl = "?format=json"
const storagePathUrl = "/v1/%s/%s/%s"
const storageAuthUrl = "/v3/auth"
const storageAuthTokenUrl = storageAuthUrl + "/tokens"
const getTempUrl = "v1/account/container/object?temp_url_sig=%s&temp_url_expires=%s"
