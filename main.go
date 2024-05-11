package main

import (
	"encoding/base64"
	"fmt"
)

func main()  {
	s := "p9L8gm6ACe2_fIZjxPQH-w=="
	seed, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(seed),err)
}