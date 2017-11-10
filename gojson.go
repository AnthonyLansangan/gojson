/*
This package makes use of encoding/json
to convert json file and string to struct
*/

package gojson

import (
	"encoding/json"
	"os"
)

//Converts json file to struct
func FileToStruct(filePath string, structure interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&structure)
	return err
}

//Converts json string to struct
func StringToStruct(jsonString string, structure interface{}) error {
	err := json.Unmarshal([]byte(jsonString), &structure)
	return err
}
