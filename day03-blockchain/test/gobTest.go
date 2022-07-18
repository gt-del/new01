package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func main() {
	var xiaoming Person
	xiaoming.Name = "xiaoming"
	xiaoming.Age = 20
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(&xiaoming)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("编码后的%v\n", buf.Bytes())
	decoder := gob.NewDecoder(bytes.NewReader(buf.Bytes()))
	var daming Person
	err = decoder.Decode(&daming)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%v\n", daming)
}
