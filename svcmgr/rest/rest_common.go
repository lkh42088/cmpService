package rest

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
)

func JsonUnmarshal(body io.ReadCloser) (m map[string]interface{}, err error) {
	bodyByte, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, errors.New("Request body is invalid.")
	}
	mapData := make(map[string]interface{})
	err = json.Unmarshal(bodyByte,&mapData)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return mapData, nil
}