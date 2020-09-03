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

type DocumentSender interface {
	SendDocument(SendDocumentRequest) (*Message, error)
}

type VideoSender interface {
	SendVideo(SendVideoRequest) (*Message, error)
}
type AnimationSender interface {
	SendAnimation(SendAnimationRequest) (*Message, error)
}
type VoiceSender interface {
	SendVoice(SendVoiceRequest) (*Message, error)
}

type VideoNoteSender interface {
	SendVideoNote(SendVideoNoteRequest) (*Message, error)
}

type MediaGroupSender interface {
	SendMediaGroup(SendMediaGroupRequest) (*Message, error)
}

type LocationSender interface {
	SendLocation(SendLocationRequest) (*Message, error)
}

type ContactSender interface {
	SendContact(SendContactRequest) (*Message, error)
}

type PollSender interface {
	SendPoll(SendPollRequest) (*Message, error)
}

type MessageTextEditor interface {
	EditMessageText(EditMessageTextRequest) (*Message, error)
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

type InlineQueryAnswerer interface {
	AnswerInlineQuery(AnswerInlineQueryRequest) (bool, error)
}

type Client interface {
	MeGetter
	MessageSender
	MessageForwarder
	MessageDeleter
	PhotoSender
	AudioSender
	FileGetter
	DocumentSender
	VideoSender
	AnimationSender
	VoiceSender
	VideoNoteSender
	MediaGroupSender
	LocationSender
	ContactSender
	PollSender
	SendChatAction(SendChatActionRequest) bool
	GetUserProfilePhotos(GetUserProfilePhotosRequest) UserProfilePhotos
	ChatOperations
	AnswerCallbackQuery(AnswerCallbackQueryRequest) bool
	InlineQueryAnswerer
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

func (c BaseClient) SendAudio(request SendAudioRequest) (interface{}, error) {
	file, err := os.Open(request.Audio)
	if err != nil {
		return nil, err
	}
	defer closeOrWarn(file)
	m := request.FillFormData(&map[string]io.Reader{}, file)
	resp, err := c.doFormRequest("sendAudio", m, &Message{})
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

func (c BaseClient) EditMessageText(request EditMessageTextRequest) (*Message, error) {
	resp, err := c.makeRequest("editMessageText", request, &Message{})
	if err != nil {
		return nil, err
	}
	return resp.(*Message), err
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

func (c BaseClient) AnswerInlineQuery(request AnswerInlineQueryRequest) (bool, error) {
	var b bool
	resp, err := c.makeRequest("answerInlineQuery", request, &b)
	if err != nil {
		return false, err
	}
	return *resp.(*bool), err
}
