package telegram

import (
	"io"
	"os"
)

type MeGetter interface {
	GetMe() (*User, error)
}

type MessageSender interface {
	SendMessage(SendMessageRequest) (*Message, error)
}

type MessageForwarder interface {
	ForwardMessage(ForwardMessageRequest) (*Message, error)
}

type PhotoSender interface {
	SendPhoto(SendPhotoRequest) (*Message, error)
}

type AudioSender interface {
	SendAudio(SendAudioRequest) (*Message, error)
}

type MessageReplyMarkupEditor interface {
	EditMessageReplyMarkup(EditMessageReplyMarkupRequest) (*Message, error)
}

type UpdatesGetter interface {
	GetUpdates(GetUpdatesRequest) (*[]Update, error)
}

type FileGetter interface {
	GetFile(GetFileRequest) (*File, error)
}

type ChatOperations interface {
	SetChatPhoto(SetChatPhotoRequest) bool
	DeleteChatPhoto(ChatRequest) bool
	SetChatTitle(SetChatTitleRequest) bool
	SetChatDescription(SetChatDescriptionRequest) bool
	PinChatMessage(PinChatMessageRequest) bool
	UnpinChatMessage(ChatRequest) bool
	LeaveChat(ChatRequest) bool
	GetChat(ChatRequest) Chat
	GetChatAdministrators(ChatRequest) []ChatMember
	GetChatMembersCount(ChatRequest) int
	GetChatMember(GetChatMemberRequest) ChatMember
	SetChatStickerSet(SetChatStickerSetRequest) bool
	DeleteChatStickerSet(ChatRequest) bool
}

type MessageDeleter interface {
	DeleteMessage(DeleteMessageRequest) (bool, error)
}

type Client interface {
	MeGetter
	MessageSender
	MessageForwarder
	MessageDeleter
	PhotoSender
	AudioSender
	FileGetter
	SendDocument(SendDocumentRequest) (*Message, error)
	SendVideo(SendVideoRequest) (*Message, error)
	SendAnimation(SendAnimationRequest) (*Message, error)
	SendVoice(SendVoiceRequest) (*Message, error)
	SendVideoNote(SendVideoNoteRequest) (*Message, error)
	SendMediaGroup(SendMediaGroupRequest) (*Message, error)
	SendLocation(SendLocationRequest) (*Message, error)
	SendContact(SendContactRequest) (*Message, error)
	SendPoll(SendPollRequest) (*Message, error)
	SendChatAction(SendChatActionRequest) bool
	GetUserProfilePhotos(GetUserProfilePhotosRequest) UserProfilePhotos
	ChatOperations
	AnswerCallbackQuery(AnswerCallbackQueryRequest) bool
}

func (c BaseClient) GetMe() (*User, error) {
	resp, err := c.makeRequest("getMe", nil, &User{})
	if err != nil {
		return nil, err
	}
	return resp.(*User), err
}

func (c BaseClient) SendMessage(request SendMessageRequest) (*Message, error) {
	resp, err := c.makeRequest("sendMessage", request, &Message{})
	if err != nil {
		return nil, err
	}
	return resp.(*Message), err
}

func (c BaseClient) ForwardMessage(request ForwardMessageRequest) (*Message, error) {
	resp, err := c.makeRequest("forwardMessage", request, &Message{})
	if err != nil {
		return nil, err
	}
	return resp.(*Message), err
}

func (c BaseClient) SendPhoto(request SendPhotoRequest) (interface{}, error) {
	file, err := os.Open(request.Photo)
	if err != nil {
		return nil, err
	}
	defer closeOrWarn(file)
	m := request.FillFormData(&map[string]io.Reader{}, file)
	resp, err := c.doFormRequest("sendPhoto", m, &Message{})
	if err != nil {
		return nil, err
	}
	return resp.(*Message), err
}

func (c BaseClient) AnswerCallbackQuery(request AnswerCallbackQueryRequest) error {
	var b bool
	_, err := c.makeRequest("answerCallbackQuery", request, &b)
	return err
}

func (c BaseClient) EditMessageReplyMarkup(request EditMessageReplyMarkupRequest) (*Message, error) {
	resp, err := c.makeRequest("editMessageReplyMarkup", request, &Message{})
	if err != nil {
		return nil, err
	}
	return resp.(*Message), err
}

func (c BaseClient) GetUpdates(request GetUpdatesRequest) (*[]Update, error) {
	var updates = make([]Update, 0)
	resp, err := c.makeRequest("getUpdates", request, &updates)
	if err != nil {
		return nil, err
	}
	return resp.(*[]Update), nil
}

func (c BaseClient) GetFile(request GetFileRequest) (*File, error) {
	resp, err := c.makeRequest("getFile", request, &File{})
	if err != nil {
		return nil, err
	}
	return resp.(*File), err
}

func (c BaseClient) DeleteMessage(request DeleteMessageRequest) (bool, error) {
	var b bool
	resp, err := c.makeRequest("deleteMessage", request, &b)
	if err != nil {
		return false, err
	}
	return *resp.(*bool), err
}
