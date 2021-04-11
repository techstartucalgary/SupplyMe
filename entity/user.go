/*

 */
package entity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var DB *sql.DB

//The basics of what is collected by the user.
type User struct {
	ID        string
	Email     string
	Password  string
	Username  string
	PublicKey string
}

//Our understanding of what is kept between organizations, and the structures of the required details.
type Profile struct {
	Org  string
	PKey string
}

func create_profile() (*Profile, error) {
	p := &Profile{
		Org:  "University of Calgary",
		PKey: "hash",
	}
	return p, nil
}

//Our Method for making a new user in our Entity Database. Will take inputs of User information in the form of a JSON file, which we unmarshal and provide
func create_user() (*User, error) {
	k := &User{
		ID:        "Test",
		Email:     "Darryl.huet1@ucalgary.ca",
		Password:  "Hash",
		Username:  "Donk",
		PublicKey: "12345",
	}
	return k, nil
}

func decode_json(file_name string, filter func(map[string]interface{}) bool) []map[string]interface{} {
	file, _ := os.Open(file_name)
	defer file.Close()

	decoder := json.NewDecover(file)
	filtered_data := []map[string]interface{}{}
	decoder.Token()

	data := map[string]interface{}{}
	for decoder.More() {
		decoder.Decode(&data)
		if filter(data) {
			filtered_data = append(filtered_data, data)
		}
	}
	return filtered_data
}

//Our Method to verify our users upon login. Our process will be the unmarshal the data, and determine whether the password is correct
func (u *User) verify_user(key *User, token string) bool {
	json := decode_json(token, func(data map[string]interface{}))
}

func main() {
	a, err := create_user()
	b, err2 := create_profile()
	if err != nil || err2 != nil {
		panic("idk man")
	}
	fmt.Printf("%+v\n", *a)
	fmt.Printf("%+v\n", *b)
}
