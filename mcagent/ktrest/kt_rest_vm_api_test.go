package ktrest

import (
	"fmt"
	"testing"
)

func TestKtRestApi(t *testing.T) {
	response := KtRestGet(serverURL, listVM)
	fmt.Println(response)
}

func TestKtResellerRestApi(t *testing.T) {
	response := KtChargeGet(resellerURL, chargeListVM)
	fmt.Println(response)
}
