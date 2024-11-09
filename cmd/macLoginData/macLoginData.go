package main

import (
	"encoding/base64"
	"flag"
	"strings"

	"github.com/kawakatz/macLoginData/pkg/decrypt"
	"github.com/kawakatz/macLoginData/pkg/parser"
	"github.com/kawakatz/macLoginData/pkg/types"
)

func main() {
	osPtr := flag.Bool("win", false, "decrypt Windows' Login Data")
	flag.Parse()

	cmdArgs := flag.Args()
	browserName := cmdArgs[0]
	loginDataFile := cmdArgs[1]

	if strings.EqualFold(browserName, "Chrome") || strings.EqualFold(browserName, "Edge") {
		var secretKey []byte
		var decryptedLoginData []types.LoginData
		if !*osPtr {
			//fmt.Println("mac")
			browserPassword := cmdArgs[2]
			secretKey = decrypt.MacPassword2SecretKey(browserPassword)
			decryptedLoginData = decrypt.ChromeLoginData(loginDataFile, secretKey, "mac")
		} else {
			//fmt.Println("win")
			secretKeyBase64 := cmdArgs[2]
			secretKey, _ = base64.StdEncoding.DecodeString(secretKeyBase64)
			decryptedLoginData = decrypt.ChromeLoginData(loginDataFile, secretKey, "win")
		}

		parser.PrintLoginData(decryptedLoginData)
	}
}
