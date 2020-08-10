package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Data ...
var Data map[string][]interface{}

// Database ...
type Database struct {
	Name string
}

// Get ...
func (db *Database) Get() []byte {
	filePath := "./storage/" + db.Name + ".json"

	var file []byte
	var err error

	file, err = ioutil.ReadFile(filePath)
	if err != nil {
		content := []byte("[]")
		err = ioutil.WriteFile(filePath, []byte("[]"), 0644)
		file = content
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return file
}

// Set ...
func (db *Database) Set(payload interface{}) error {
	filePath := "./storage/" + db.Name + ".json"
	data, _ := json.Marshal(payload)
	err := ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
