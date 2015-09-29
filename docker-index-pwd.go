package main

import (
	"encoding/base64"
	"flag"
	"fmt"
)

func main() {
	pwd := flag.String("user", "", "Your user name for docker index")
	user := flag.String("pwd", "", "Your password for docker index")
	flag.Parse()

	if *pwd == "" || *user == "" {
		fmt.Println("Please provide your user name and password for docker index")
		return
	}

	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", *user, *pwd)))
	fmt.Println("Docker index auth:", auth)
}
