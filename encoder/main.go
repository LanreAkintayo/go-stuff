package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	// "os"
)

type user struct {
	Name     string `json:"name" xml:"name"`
	Age      int    `json:"age" xml:"age"`
	Phone    string `json:"phone" xml:"phone_number"`
	Password string `json:"-" xml:"-"`
	IsActive bool   `json:"active" xml:"active"`
}

func main() {
	u := user{
		Name:  "John Smith",
		Age:   45,
		Phone: "13812345678",
	}

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	// enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(&u); err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf.Bytes()))

}