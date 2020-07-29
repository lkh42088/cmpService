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

type SearchParam struct {
	Type    string `json:"searchType"`
	Content string `json:"searchContent"`
}

type PageRequestMsg struct {
	RowsPerPage int         `json:"rows"`
	Offset      int         `json:"offset"`
	OrderBy     string      `json:"orderBy"`
	Order       string      `json:"order"`
	Param       SearchParam `json:"searchParam"`
}

type PageRequestForSearch struct {
	RowsPerPage int         `json:"rows"`
	Offset      int         `json:"offset"`
	OrderBy     string      `json:"orderBy"`
	Order       string      `json:"order"`
	SearchParam string 		`json:"searchParam"`
}

