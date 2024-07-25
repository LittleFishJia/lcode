package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	//"time"
)



func main() {
	ttl := 3600
	time := time.Now().Unix() + int64(ttl)
	// Generate the username string
	username := strconv.FormatInt(time, 10) + ":789"
	// Secret key used for HMAC
	//secret := "gFEuDwhl/DsppWqUMf1D8xsbfb/DtL68" //正式
	secret := "K2fx6uNh3Z0Q+7KukW4b7+tFgS3okNQn"
	// Generate HMAC-SHA1 hash
	hash := hmac.New(sha1.New, []byte(secret))
	hash.Write([]byte(username))
	passwordBytes := hash.Sum(nil)
	fmt.Println(hex.EncodeToString(passwordBytes))
	// Encode the hash to Base64
	password := base64.StdEncoding.EncodeToString(passwordBytes)

	// Print the results
	fmt.Println("Username:", username)
	fmt.Println("Password:", password)



}

