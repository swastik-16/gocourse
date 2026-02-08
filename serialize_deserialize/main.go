package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	user := User{Name: "Alice", Email: "alice@example.com"}
	fmt.Println(user)
	jsonData, err := json.Marshal(user)
	if err!= nil{
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
	var newUser User
	err = json.Unmarshal(jsonData, &newUser)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("User created from json data",newUser)

	jsonData1 := `{"name":"Bob","email":"bob@example.com"}` 
	reader := strings.NewReader(jsonData1)
	decoder := json.NewDecoder(reader)

	var user2 User
	err = decoder.Decode(&user2)
	if err!= nil{
		log.Fatal(err)
	}
	fmt.Println("User created from json data",user2)

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	err = encoder.Encode(user)
	if err!= nil{
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}