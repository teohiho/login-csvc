package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

type User struct {
    id   string
    username    string
    password    string
    fullname    string
    avatar 		string
    phone 		string
    id_donvi 	string
    id_role 	string
}

// Main function
// I realize this function is much too simple I am simply at a loss to

func main() {
    var file, e := ioutil.ReadFile("main/user.json")
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))
    //m := new(Dispatch)
    //var m interface{}
    // var users Users
    // json.Unmarshal(file, &users)
    // fmt.Printf("Results: %v\n", users)
    var data interface{}
	var err, _ := json.Unmarshal(file, &data)

	fmt.Println(data)
	fmt.Println(err)

}