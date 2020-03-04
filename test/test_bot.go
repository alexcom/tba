package main

import (
	"fmt"
	"github.com/alexcom/tba/telegram"
	"io/ioutil"
)

func main() {
	bytes, _ := ioutil.ReadFile("token.txt")
	client := telegram.NewClient(string(bytes), 15)
	msg := telegram.SendPhotoRequest{}
	msg.ChatID = 139455782
	msg.Caption = "<b>ololo</b>"
	msg.ParseMode = "HTML"
	msg.Photo = "C:\\projects\\SVG\\bitmap.png"
	message, err := client.SendPhoto(msg)
	if err != nil {
		fmt.Printf("%v\n%v", message, err)
	}
	fmt.Printf("%v", message)
}
