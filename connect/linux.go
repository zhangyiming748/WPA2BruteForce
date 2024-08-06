package connect

import (
	"fmt"
	"os/exec"
)

func Linux(ssid, password string) error {
	// 替换为你的WiFi名称和密码
	//ssid := "YourSSID"
	//password := "YourPassword"

	// 连接WiFi
	cmd := exec.Command("nmcli", "dev", "wifi", "connect", ssid, "password", password)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error connecting to WiFi:", err)
		return err
	}
	fmt.Println(string(output))
	return nil
}
