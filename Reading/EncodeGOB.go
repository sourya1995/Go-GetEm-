package main
import (
	"fmt"
	"bufio"
	"encoding/gob"
	"os"
	"log"
)

type Address struct {
	Type string
	City string
	Country string
  }
  
  type VCard struct {
	FirstName string
	LastName string
	Addresses []*Address
	Remark string
  }

  var content string
  var vc VCard

  func main() {
	pa := &Address{"private", "Aartselaar","Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa,wa}, "none"}

	file, _ := os.OpenFile("output/vcard.gob", os.O_CREATE|os.O_WRONLY, 0)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
	  log.Println("Error in encoding gob")
	}

	//using a decoder

	file1, _ := os.Open("vcard.gob")
	defer file1.Close()
	inReader := bufio.NewReader(file1)
	dec := gob.NewDecoder(inReader)
	err1 := dec.Decode(&vc) 
	if err1 != nil {
		log.Println("Error in decoding gob")
	}
	fmt.Println(vc)
  }