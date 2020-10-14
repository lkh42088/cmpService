package ktrest

import "testing"

func TestGetUserAuth(t *testing.T) {
	_ = GetAuthTokens()
	GetStorageAccount(GlobalToken)
}

func TestGetKtStorageTempUrl(t *testing.T) {
	GetKtStorageTempUrl()
}