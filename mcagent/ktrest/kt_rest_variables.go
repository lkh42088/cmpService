package ktrest

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

// DB
const storageAccessKey = "iwhan@nubes-bridge.com"
const storageSecretKey = "MTYwMTg2MzU1OTE2MDE4NjI5MTk1MTQ2"
const storageProjectId = "fa632a4a0d04488c93b7184be92df4c8"
const storageDomainId = "42a37f949dcd48a3a805fe0d2d3a7da5"



