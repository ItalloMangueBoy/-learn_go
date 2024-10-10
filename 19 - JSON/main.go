package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type People struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Age    uint   `json:"age"`
}

func main() {
	// Marshal
	peoples := map[string]People{
		"itallo": People{"Itallo", "male", 20},
		"jenny":  People{"Jennyfer", "female", 18},
	}

	peoplesJSON, err := json.Marshal(peoples)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(peoplesJSON))

	// Unmarshal
	userFile, err := ioutil.ReadFile("./user.json")
	if err != nil {
		log.Fatal(err)
	}

	var peopleMap map[string]People

	err = json.Unmarshal(userFile, &peopleMap)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(peopleMap)
}
