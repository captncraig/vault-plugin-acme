package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	flag.Parse()
	dat, err := ioutil.ReadFile(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	hsh := sha256.Sum256(dat)
	fmt.Println(hex.EncodeToString(hsh[:]))
}
