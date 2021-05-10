package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Employee struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	IPAddress string `json:"ip_address"`
}

func FiletoJSON(filename string) {
	employees := []Employee{}
	file, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("%s", err)
	}
	json.NewDecoder(file).Decode(&employees)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func FiletoJsonMarshal(file string) {
	var employees []Employee
	// jsonFile, err := os.OpenFile(file, os.O_RDONLY, 0666)
	// check(err)

	// defer jsonFile.Close()

	byteValue, err := ioutil.ReadFile(file)
	check(err)

	json.Unmarshal(byteValue, &employees)

	// fmt.Println(employees)

}
