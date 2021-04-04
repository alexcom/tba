package bot

import (
	"github.com/alexcom/tba/telegram"
)

type MessageServices interface {
	SendText(chatID int, msg string) (*telegram.Message, error)
	SendMarkdown(chatID int, msg string) (*telegram.Message, error)
	SendHTML(chatID int, msg string) (*telegram.Message, error)
	SendKeyboard(chatID int, msg string, kb *telegram.InlineKeyboardMarkup) (*telegram.Message, error)
	DeleteMessage(chatID, messageID int) error
	AnswerCallbackQuery(callbackQueryID, msg string, showAlert bool) error
	EditKeyboardMarkup(chatID, messageID int, kb *telegram.InlineKeyboardMarkup) (*telegram.Message, error)
	GetFile(fileID string) (*telegram.File, error)
	DownloadFile(filePath string) ([]byte, error)
}

func (bot Bot) SendText(chatID int, text string) (*telegram.Message, error) {
	req := telegram.SendMessageRequest{}
	req.ChatID = chatID
	req.Text = text
	return bot.Telegram.SendMessage(req)
}

func (bot Bot) SendMarkdown(chatID int, markdown string) (*telegram.Message, error) {
	req := telegram.SendMessageRequest{}
	req.ChatID = chatID
	req.Text = markdown
	req.ParseMode = telegram.ParseModeMarkdown
	return bot.Telegram.SendMessage(req)
}

func (bot Bot) SendHTML(chatID int, html string) (*telegram.Message, error) {
	req := telegram.SendMessageRequest{}
	req.ChatID = chatID
	req.Text = html
	req.ParseMode = telegram.ParseModeMarkdown
	return bot.Telegram.SendMessage(req)
}

func (bot Bot) SendKeyboard(chatID int, msg string, kb *telegram.InlineKeyboardMarkup) (*telegram.Message, error) {
	request := telegram.SendMessageRequest{}
	request.Text = msg
	request.ChatID = chatID
	request.ReplyMarkup = kb
	return bot.Telegram.SendMessage(request)
}

func (bot Bot) DeleteMessage(chatID, messageID int) error {
	req := telegram.DeleteMessageRequest{}
	req.ChatID = chatID
	req.MessageID = messageID
	_, err := bot.Telegram.DeleteMessage(req)
	return err
}

func (bot Bot) AnswerCallbackQuery(callbackQueryID, message string, showAlert bool) error {
	request := telegram.AnswerCallbackQueryRequest{}
	request.CallbackQueryID = callbackQueryID
	request.Text = message
	request.ShowAlert = showAlert
	return bot.Telegram.AnswerCallbackQuery(request)
}

func (bot Bot) EditKeyboardMarkup(chatID, messageID int, kb *telegram.InlineKeyboardMarkup) (*telegram.Message, error) {
	request := telegram.EditMessageReplyMarkupRequest{}
	request.ChatID = chatID
	request.MessageID = messageID
	request.ReplyMarkup = kb
	return bot.Telegram.EditMessageReplyMarkup(request)
}

func (bot Bot) GetFile(fileID string) (*telegram.File, error) {
	return bot.Telegram.GetFile(telegram.GetFileRequest{
		FileID: fileID,
	})
}

func (bot Bot) DownloadFile(filePath string) ([]byte, error) {
	return bot.Telegram.DownloadFile(filePath)
}
