package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/kawakatz/macLoginData/pkg/types"
)

func CookieQuickManager(decryptedCookies []types.Cookie) {
	var cookieQuickManagers []types.CookieQuickManager
	for _, Cookie := range decryptedCookies {
		cookieQuickManager := types.CookieQuickManager{
			PathRaw:           Cookie.Path,
			HostRaw:           "https://" + Cookie.Host + "/",
			ExpiresRaw:        strconv.FormatInt(time.Now().AddDate(1, 0, 0).Unix(), 10),
			ContentRaw:        Cookie.Value,
			NameRaw:           Cookie.KeyName,
			SameSiteRaw:       "no_restriction",
			ThisDomainOnlyRaw: "false",
			StoreRaw:          "firefox-default",
			FirstPartyDomain:  "",
			HTTPOnlyRaw:       "",
		}

		cookieQuickManagers = append(cookieQuickManagers, cookieQuickManager)
	}

	file, _ := json.MarshalIndent(cookieQuickManagers, "", "\t")
	_ = ioutil.WriteFile("macCookies.json", file, 0644)

	fmt.Println("\x1b[32m[+] successfully exported!\x1b[0m")
	fmt.Println("\x1b[32m[+] import macCookies.json to Firefox with CookieQuickManager\x1b[0m")
}

func PrintLoginData(decryptedLoginData []types.LoginData) {
	for _, LoginData := range decryptedLoginData {
		fmt.Println("--------------------------------")
		fmt.Println("origin_url:", LoginData.OriginURL)
		fmt.Println("username:", LoginData.Username)
		fmt.Println("password:", LoginData.Password)

		formattedTime := LoginData.DateCreated.Format("2006/01/02 15:04:05") // 2006/01/02 15:04:05
		fmt.Println("date_created:", formattedTime)
		formattedTime = LoginData.DateLastUsed.Format("2006/01/02 15:04:05")
		fmt.Println("date_last_used:", formattedTime)
		formattedTime = LoginData.DatePasswordModified.Format("2006/01/02 15:04:05")
		fmt.Println("date_password_modified:", formattedTime)

	}
	fmt.Println("--------------------------------")
}
