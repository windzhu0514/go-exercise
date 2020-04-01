package main

import (
	"encoding/json"
	"fmt"
)

type AccountInfo struct {
	AccountType int               `json:"accountType,omitempty"` // 帐号类别(1：自购，0：代购)
	UserName    string            `json:"userName,omitempty"`
	UserPass    bool              `json:"userPassword,omitempty"`
	Mobile      []string          `json:"mobile,omitempty"`
	Email       []byte            `json:"email,omitempty"`
	Sex         map[string]string `json:"sex,omitempty"`
}

func main() {
	var accountInfo AccountInfo
	accountInfo.AccountType = -1
	data, _ := json.Marshal(accountInfo)
	fmt.Println(string(data))
}
