package main

import (
	"fmt"
	"github.com/alexcom/tba/telegram"
)

func main() {
	client := telegram.NewClient("", 15)
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
