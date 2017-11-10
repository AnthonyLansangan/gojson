package gojson

import (
	"testing"
)

type Configuration struct {
	MongoDB mongoDB `json:"mongodb" bson:"mongodb"`
}

type mongoDB struct {
	Ip     string `json:"ip" bson:"ip"`
	Dbname string `json:"dbname" bson:"dbname"`
}

func TestLoadJsonConfig(t *testing.T) {
	var conf Configuration
	err := FileToStruct("conf.json", &conf)
	if err != nil {
		t.Error("Error:", err)
		t.Error("Error on loading configuration file")
		t.Fail()
	}
	if conf.MongoDB.Ip != "127.0.0.1" {
		t.Error("Config not loaded")
		t.Fail()
	}
}

func TestStringToStruct(t *testing.T) {
	var mongoDB mongoDB
	jsonString := `{"mongodb": {
    "ip": "127.0.0.1",
    "dbname": "iwpro_monitoring_test"
    }}`
	err := StringToStruct(jsonString, &mongoDB)
	if err != nil {
		t.Error("Error:", err)
		t.Fail()
	}
}

func BenchmarkLoadJsonFileToStruct(b *testing.B) {
	var conf Configuration
	for i := 0; i < b.N; i++ {
		FileToStruct("conf.json", &conf)
	}
}

func BenchmarkStringToStruct(b *testing.B) {
	var mongoDB mongoDB
	jsonString := `{"mongodb": {
    "ip": "127.0.0.1",
    "dbname": "iwpro_monitoring_test"
    }}`
	for i := 0; i < b.N; i++ {
		StringToStruct(jsonString, &mongoDB)
	}
}
