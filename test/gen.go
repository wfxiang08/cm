package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type SuitInfo struct {
	SuitName   string `json:"suit_name"`
	TableName  string `json:"table_name"`
	DataType   string `json:"data_type"`
	DataTypeGo string `json:"data_type_go"`

	Id          string `json:"id"`
	IdType      string `json:"id_type"`
	Data        string `json:"data"`
	DataUpdated string `json:"data_updated"`
}

func main() {
	tmpl := os.Args[1]
	b, err := ioutil.ReadFile(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	t := template.Must(template.New("type_test").Parse(string(b)))

	suitInfoFile := os.Args[2]
	b, err = ioutil.ReadFile(suitInfoFile)
	if err != nil {
		log.Fatal(err)
	}

	var info SuitInfo
	err = json.Unmarshal(b, &info)
	if err != nil {
		log.Fatal(err)
	}

	if len(info.IdType) == 0 {
		info.IdType = "INT"
	}

	t.Execute(os.Stdout, &info)
}
