package models

import (
	"encoding/json"
	"fmt"
)

type Pagination struct {
	TotalCount  int    `json:"count"`
	RowsPerPage int    `json:"rows"`
	Offset      int    `json:"offset"`
	OrderBy     string `json:"orderBy"`
	Order       string `json:"order"`
}

func (p Pagination) String() {
	fmt.Printf("%v", &p)
	//fmt.Printf("%+v", &p)
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("%s\n", data)
}
