package telegram

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"strconv"
	"strings"
)

type Update struct {
	UpdateID           int                 `json:"update_id"`
	Message            *Message            `json:"message"`
	EditedMessage      *Message            `json:"edited_message"`
	ChannelPost        *Message            `json:"channel_post"`
	EditedChannelPost  *Message            `json:"edited_channel_post"`
	InlineQuery        *InlineQuery        `json:"inline_query"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`
	CallbackQuery      *CallbackQuery      `json:"callback_query"`
	ShippingQuery      *ShippingQuery      `json:"shipping_query"`
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query"`
	Poll               *Poll               `json:"poll"`
}

type InlineQuery struct {
	ID       string    `json:"id"`
	From     *User     `json:"from"`
	Location *Location `json:"location"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            *User     `json:"from"`
	Location        *Location `json:"location"`
	InlineMessageID string    `json:"inline_message_id"`
	Query           string    `json:"query"`
}

type CallbackQuery struct {
	ID              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message"`
	InlineMessageID string   `json:"inline_message_id"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data"`
	GameShortName   string   `json:"game_short_name"`
}

type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line_1"`
	StreetLine2 string `json:"street_line_2"`
	PostCode    string `json:"post_code"`
}

type ShippingQuery struct {
	ID              string           `json:"id"`
	From            *User            `json:"from"`
	InvoicePayload  string           `json:"invoice_payload"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

type OrderInfo struct {
	Name            string           `json:"name"`
	PhoneNumber     string           `json:"phone_number"`
	Email           string           `json:"email"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

type PreCheckoutQuery struct {
	ID               string     `json:"id"`
	From             *User      `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int        `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionID string     `json:"shipping_option_id"`
	OrderInfo        *OrderInfo `json:"order_info"`
}

type User struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type ChatPhoto struct {
	SmallFileID string `json:"small_file_id"`
	BigFileID   string `json:"big_file_id"`
}

type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages"`
	CanSendMediaMessages  bool `json:"can_send_media_messages"`
	CanSendPolls          bool `json:"can_send_polls"`
	CanSendOtherMessages  bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	CanChangeInfo         bool `json:"can_change_info"`
	CanInviteUsers        bool `json:"can_invite_users"`
	CanPinMessages        bool `json:"can_pin_messages"`
}

type Chat struct {
	ID               int              `json:"id"`
	Type             string           `json:"type"`
	Title            string           `json:"title"`
	Username         string           `json:"username"`
	FirstName        string           `json:"first_name"`
	LastName         string           `json:"last_name"`
	Photo            *ChatPhoto       `json:"photo"`
	Description      string           `json:"description"`
	InviteLink       string           `json:"invite_link"`
	PinnedMessage    *Message         `json:"pinned_message"`
	Permissions      *ChatPermissions `json:"permissions"`
	StickerSetName   string           `json:"sticker_set_name"`
	CanSetStickerSet bool             `json:"can_set_sticker_set"`
}

type MessageEntity struct {
	Type   EntityType `json:"type"`
	Offset int        `json:"offset"`
	Length int        `json:"length"`
	Url    string     `json:"url"`
	User   *User      `json:"user"`
}

type Audio struct {
	FileID    string     `json:"file_id"`
	Duration  int        `json:"duration"`
	Performer string     `json:"performer"`
	Title     string     `json:"title"`
	MimeType  string     `json:"mime_type"`
	FileSize  int        `json:"file_size"`
	Thumb     *PhotoSize `json:"thumb"`
}

type Document struct {
	FileID   string    `json:"file_id"`
	Thumb    PhotoSize `json:"thumb"`
	FileName string    `json:"file_name"`
	MimeType string    `json:"mime_type"`
	FileSize int       `json:"file_size"`
}

type Video struct {
	FileID   string     `json:"file_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb"`
	MimeType string     `json:"mime_type"`
	FileSize int        `json:"file_size"`
}

type Voice struct {
	FileID   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"`
	FileSize int    `json:"file_size"`
}

type Animation struct {
	FileID   string    `json:"file_id"`
	Width    int       `json:"width"`
	Height   int       `json:"height"`
	Duration int       `json:"duration"`
	Thumb    PhotoSize `json:"thumb"`
	FileName string    `json:"file_name"`
	MimeType string    `json:"mime_type"`
	FileSize int       `json:"file_size"`
}

type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         string          `json:"text"`
	TextEntities []MessageEntity `json:"text_entities"`
	Animation    *Animation      `json:"animation"`
}

type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size"`
}

type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

type Sticker struct {
	FileID       string        `json:"file_id"`
	Width        int           `json:"width"`
	Height       int           `json:"height"`
	IsAnimated   bool          `json:"is_animated"`
	Thumb        *PhotoSize    `json:"thumb"`
	Emoji        string        `json:"emoji"`
	SetName      *string       `json:"set_name"`
	MaskPosition *MaskPosition `json:"mask_position"`
	FileSize     int           `json:"file_size"`
}

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserID      int    `json:"user_id"`
	VCard       string `json:"v_card"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type VideoNote struct {
	FileID   string     `json:"file_id"`
	Length   int        `json:"length"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb"`
	FileSize int        `json:"file_size"`
}

type Venue struct {
	Location       *Location `json:"location"`
	Title          string    `json:"title"`
	Address        string    `json:"address"`
	FoursquareID   string    `json:"foursquare_id"`
	FoursquareType string    `json:"foursquare_type"`
}

type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

type Poll struct {
	ID       string       `json:"id"`
	Question string       `json:"question"`
	Options  []PollOption `json:"options"`
	IsClosed bool         `json:"is_closed"`
}

type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

type SuccessfulPayment struct {
	Currency                string     `json:"currency"`
	TotalAmount             int        `json:"total_amount"`
	InvoicePayload          string     `json:"invoice_payload"`
	ShippingOptionID        string     `json:"shipping_option_id"`
	OrderInfo               *OrderInfo `json:"order_info"`
	TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID string     `json:"provider_payment_charge_id"`
}

type PassportFile struct {
	FileID   string `json:"file_id"`
	FileSize int    `json:"file_size"`
	FileDate int    `json:"file_date"`
}

type EncryptedPassportElement struct {
	Type        string         `json:"type"`
	Data        string         `json:"data"`
	PhoneNumber string         `json:"phone_number"`
	Email       string         `json:"email"`
	Files       []PassportFile `json:"files"`
	FrontSide   *PassportFile  `json:"front_side"`
	ReverseSide *PassportFile  `json:"reverse_side"`
	Selfie      *PassportFile  `json:"selfie"`
	Translation []PassportFile `json:"translation"`
	Hash        string         `json:"hash"`
}

type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`
	Credentials EncryptedCredentials       `json:"credentials"`
}

type Message struct {
	MessageID             int                   `json:"message_id"`
	From                  *User                 `json:"from"`
	Date                  int                   `json:"date"`
	Chat                  *Chat                 `json:"chat"`
	ForwardFrom           *User                 `json:"forward_from"`
	ForwardFromChat       *Chat                 `json:"forward_from_chat"`
	ForwardFromMessageID  int                   `json:"forward_from_message_id"`
	ForwardSignature      string                `json:"forward_signature"`
	ForwardSenderName     string                `json:"forward_sender_name"`
	ForwardDate           int                   `json:"forward_date"`
	ReplyToMessage        *Message              `json:"reply_to_message"`
	EditDate              int                   `json:"edit_date"`
	MediaGroupID          string                `json:"media_group_id"`
	AuthorSignature       string                `json:"author_signature"`
	Text                  string                `json:"text"`
	Entities              []MessageEntity       `json:"entities"`
	CaptionEntities       []MessageEntity       `json:"caption_entities"`
	Audio                 *Audio                `json:"audio"`
	Document              *Document             `json:"document"`
	Animation             *Animation            `json:"animation"`
	Game                  *Game                 `json:"game"`
	Photo                 []PhotoSize           `json:"photo"`
	Sticker               *Sticker              `json:"sticker"`
	Video                 *Video                `json:"video"`
	Voice                 *Voice                `json:"voice"`
	VideoNote             *VideoNote            `json:"video_note"`
	Caption               string                `json:"caption"`
	Contact               *Contact              `json:"contact"`
	Location              *Location             `json:"location"`
	Venue                 *Venue                `json:"venue"`
	Poll                  *Poll                 `json:"poll"`
	NewChatMembers        []User                `json:"new_chat_members"`
	LeftChatMember        *User                 `json:"left_chat_member"`
	NewChatTitle          string                `json:"new_chat_title"`
	NewChatPhoto          *PhotoSize            `json:"new_chat_photo"`
	DeleteChatPhoto       bool                  `json:"delete_chat_photo"`
	GroupChatCreated      bool                  `json:"group_chat_created"`
	SupergroupChatCreated bool                  `json:"supergroup_chat_created"`
	ChannelChatCreated    bool                  `json:"channel_chat_created"`
	MigrateToChatID       int                   `json:"migrate_to_chat_id"`
	MigrateFromChatID     int                   `json:"migrate_from_chat_id"`
	PinnedMessage         *Message              `json:"pinned_message"`
	Invoice               *Invoice              `json:"invoice"`
	SuccessfulPayment     *SuccessfulPayment    `json:"successful_payment"`
	ConnectedWebsite      string                `json:"connected_website"`
	PassportData          *PassportData         `json:"passport_data"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup"`
}

type ResponseParameters struct {
	MigrateToChatID int `json:"migrate_to_chat_id"`
	RetryAfter      int `json:"retry_after"`
}

type GetFileResponse struct {
	Ok     bool `json:"ok"`
	Result File `json:"result"`
}
type File struct {
	FileID   string `json:"file_id"`
	FileSize int    `json:"file_size"`
	FilePath string `json:"file_path"`
}

type EntityType string

const (
	EntityTypeMention     EntityType = "mention"
	EntityTypeHashtag     EntityType = "hashtag"
	EntityTypeCashtag     EntityType = "cashtag"
	EntityTypeBotCommand  EntityType = "bot_command"
	EntityTypeUrl         EntityType = "url"
	EntityTypeEmail       EntityType = "email"
	EntityTypePhoneNumber EntityType = "phone_number"
	EntityTypeBold        EntityType = "bold"
	EntityTypeItalic      EntityType = "italic"
	EntityTypeCode        EntityType = "code"
	EntityTypePre         EntityType = "pre"
	EntityTypeTextLink    EntityType = "text_link"
	EntityTypeTextMention EntityType = "text_mention"
)

type FormDataFiller interface {
	FillFormData(m *map[string]io.Reader, reader io.Reader) *map[string]io.Reader
}

type ChatRequest struct {
	ChatID int `json:"chat_id"`
}

func (c ChatRequest) FillFormData(m *map[string]io.Reader, reader io.Reader) {
	if c.ChatID != 0 {
		(*m)["chat_id"] = strings.NewReader(strconv.Itoa(c.ChatID))
	}
}

type ParseModeSource struct {
	ParseMode ParseMode `json:"parse_mode,omitempty"`
}

func (c ParseModeSource) FillFormData(m *map[string]io.Reader, reader io.Reader) {
	if c.ParseMode != "" {
		(*m)["parse_mode"] = strings.NewReader(string(c.ParseMode))
	}
}

type ReplyMarkupSource struct {
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

func (c ReplyMarkupSource) FillFormData(m *map[string]io.Reader, reader io.Reader) *map[string]io.Reader {
	if c.ReplyMarkup != nil {
		jsonBytes, err := json.Marshal(c.ReplyMarkup)
		if err != nil {
			logrus.WithError(err).Warn("cannot serialize reply markup to json")
		} else {
			(*m)["reply_markup"] = bytes.NewReader(jsonBytes)
		}
	}
	return m
}

type ReplyToMessageIDSource struct {
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`
}

func (c ReplyToMessageIDSource) FillFormData(m *map[string]io.Reader, reader io.Reader) *map[string]io.Reader {
	if c.ReplyToMessageID != 0 {
		(*m)["reply_to_message_id"] = strings.NewReader(strconv.Itoa(c.ReplyToMessageID))
	}
	return m
}

type CaptionSource struct {
	Caption string `json:"caption,omitempty"`
}

func (c CaptionSource) FillFormData(m *map[string]io.Reader, reader io.Reader) *map[string]io.Reader {
	if c.Caption != "" {
		(*m)["reply_to_message_id"] = strings.NewReader(c.Caption)
	}
	return m
}

type DisableNotificationsSource struct {
	DisableNotification bool `json:"disable_notification,omitempty"`
}

func (c DisableNotificationsSource) FillFormData(m *map[string]io.Reader, reader io.Reader) *map[string]io.Reader {
	if c.DisableNotification != false {
		(*m)["disable_notifications"] = strings.NewReader(strconv.FormatBool(c.DisableNotification))
	}
	return m
}

type ThumbSource struct {
	Thumb string `json:"thumb,omitempty"`
}

func (c ThumbSource) FillFormData(m *map[string]io.Reader, reader io.Reader) *map[string]io.Reader {
	if c.Thumb != "" {
		(*m)["thumb"] = strings.NewReader(c.Thumb)
	}
	return m
}

type SendMessageRequest struct {
	ChatRequest
	Text                  string `json:"text"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview,omitempty"`
	ParseModeSource
	DisableNotificationsSource
	ReplyToMessageIDSource
	ReplyMarkupSource
}

type ForwardMessageRequest struct {
	ChatRequest
	DisableNotificationsSource
	FromChatID int `json:"from_chat_id"`
	MessageID  int `json:"message_id"`
}
type DeleteMessageRequest struct {
	ChatRequest
	MessageID int `json:"message_id,omitempty"`
}

type SendPhotoRequest struct {
	Photo string `json:"photo"`
	ChatRequest
	CaptionSource
	ParseModeSource
	DisableNotificationsSource
	ReplyToMessageIDSource
	ReplyMarkupSource
}

func (c SendPhotoRequest) FillFormData(m *map[string]io.Reader, reader io.Reader) *map[string]io.Reader {
	c.ChatRequest.FillFormData(m, nil)
	c.CaptionSource.FillFormData(m, nil)
	c.ParseModeSource.FillFormData(m, nil)
	c.DisableNotificationsSource.FillFormData(m, nil)
	c.ReplyToMessageIDSource.FillFormData(m, nil)
	c.ReplyMarkupSource.FillFormData(m, nil)
	(*m)["photo"] = reader
	return m
}

type SendAudioRequest struct {
	Audio     string `json:"audio"`
	Duration  int    `json:"duration"`
	Performer string `json:"performer"`
	Title     string `json:"title"`
	ChatRequest
	ThumbSource
	CaptionSource
	ParseModeSource
	DisableNotificationsSource
	ReplyToMessageIDSource
	ReplyMarkupSource
}

func (c SendAudioRequest) FillFormData(m *map[string]io.Reader, reader io.Reader) *map[string]io.Reader {
	c.ChatRequest.FillFormData(m, nil)
	c.ThumbSource.FillFormData(m, nil)
	c.CaptionSource.FillFormData(m, nil)
	c.ParseModeSource.FillFormData(m, nil)
	c.DisableNotificationsSource.FillFormData(m, nil)
	c.ReplyToMessageIDSource.FillFormData(m, nil)
	c.ReplyMarkupSource.FillFormData(m, nil)
	(*m)["duration"] = strings.NewReader(strconv.Itoa(c.Duration))
	(*m)["performer"] = strings.NewReader(c.Performer)
	(*m)["title"] = strings.NewReader(c.Title)
	(*m)["audio"] = reader
	return m
}

type SendDocumentRequest struct {
	Document string `json:"document"`
	ChatRequest
	ThumbSource
	CaptionSource
	ParseModeSource
	DisableNotificationsSource
	ReplyToMessageIDSource
	ReplyMarkupSource
}

func (c SendDocumentRequest) FillFormData(m *map[string]io.Reader, reader io.Reader) *map[string]io.Reader {
	c.ChatRequest.FillFormData(m, nil)
	c.ThumbSource.FillFormData(m, nil)
	c.CaptionSource.FillFormData(m, nil)
	c.ParseModeSource.FillFormData(m, nil)
	c.DisableNotificationsSource.FillFormData(m, nil)
	c.ReplyToMessageIDSource.FillFormData(m, nil)
	c.ReplyMarkupSource.FillFormData(m, nil)
	(*m)["document"] = reader
	return m
}

type SendVideoRequest struct {
	Video             string `json:"video"`
	Duration          int    `json:"duration"`
	Width             int    `json:"width"`
	Height            int    `json:"height"`
	SupportsStreaming bool   `json:"supports_streaming"`
	ChatRequest
	ThumbSource
	CaptionSource
	ParseModeSource
	DisableNotificationsSource
	ReplyToMessageIDSource
	ReplyMarkupSource
}

func (c SendVideoRequest) FillFormData(m *map[string]io.Reader, reader io.Reader) *map[string]io.Reader {
	c.ChatRequest.FillFormData(m, nil)
	c.ThumbSource.FillFormData(m, nil)
	c.CaptionSource.FillFormData(m, nil)
	c.ParseModeSource.FillFormData(m, nil)
	c.DisableNotificationsSource.FillFormData(m, nil)
	c.ReplyToMessageIDSource.FillFormData(m, nil)
	c.ReplyMarkupSource.FillFormData(m, nil)
	(*m)["video"] = reader
	(*m)["duration"] = strings.NewReader(strconv.Itoa(c.Duration))
	(*m)["width"] = strings.NewReader(strconv.Itoa(c.Width))
	(*m)["height"] = strings.NewReader(strconv.Itoa(c.Height))
	(*m)["supports_streaming"] = strings.NewReader(strconv.FormatBool(c.SupportsStreaming))
	return m
}

type SendAnimationRequest struct {
	Animation string `json:"animation"`
	Duration  int    `json:"duration"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	ChatRequest
	ThumbSource
	CaptionSource
	ParseModeSource
	DisableNotificationsSource
	ReplyToMessageIDSource
	ReplyMarkupSource
}

func (c SendAnimationRequest) FillFormData(m *map[string]io.Reader, reader io.Reader) *map[string]io.Reader {
	c.ChatRequest.FillFormData(m, nil)
	c.ThumbSource.FillFormData(m, nil)
	c.CaptionSource.FillFormData(m, nil)
	c.ParseModeSource.FillFormData(m, nil)
	c.DisableNotificationsSource.FillFormData(m, nil)
	c.ReplyToMessageIDSource.FillFormData(m, nil)
	c.ReplyMarkupSource.FillFormData(m, nil)
	(*m)["animation"] = reader
	(*m)["duration"] = strings.NewReader(strconv.Itoa(c.Duration))
	(*m)["width"] = strings.NewReader(strconv.Itoa(c.Width))
	(*m)["height"] = strings.NewReader(strconv.Itoa(c.Height))
	return m
}

type SendVoiceRequest struct {
	Voice    string `json:"voice"`
	Duration int    `json:"duration"`
	ChatRequest
	CaptionSource
	ParseModeSource
	DisableNotificationsSource
	ReplyToMessageIDSource
	ReplyMarkupSource
}

func (c SendVoiceRequest) FillFormData(m *map[string]io.Reader, reader io.Reader) *map[string]io.Reader {
	(*m)["voice"] = reader
	(*m)["duration"] = strings.NewReader(strconv.Itoa(c.Duration))
	return m
}

type SendVideoNoteRequest struct {
	ChatRequest
	VideoNote string `json:"video_note"`
	Duration  int    `json:"duration"`
	Length    int    `json:"length"`
	ThumbSource
	DisableNotificationsSource
	ReplyToMessageIDSource
	ReplyMarkupSource
}

func (c SendVideoNoteRequest) FillFormData(m *map[string]io.Reader, reader io.Reader) *map[string]io.Reader {
	(*m)["video_note"] = reader
	(*m)["duration"] = strings.NewReader(strconv.Itoa(c.Duration))
	(*m)["length"] = strings.NewReader(strconv.Itoa(c.Length))
	return m
}

type SendMediaGroupRequest struct {
	ChatRequest
	Media []interface{} `json:"media"`
	DisableNotificationsSource
	ReplyToMessageIDSource
}

type SendLocationRequest struct {
	ChatRequest
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	LivePeriod int     `json:"live_period"`
	DisableNotificationsSource
	ReplyToMessageIDSource
	ReplyMarkupSource
}

type SendContactRequest struct {
	ChatRequest
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	VCard       string `json:"v_card"`
	DisableNotificationsSource
	ReplyToMessageIDSource
	ReplyMarkupSource
}

type SendPollRequest struct {
	ChatRequest
	Question string   `json:"question"`
	Options  []string `json:"options"`
	DisableNotificationsSource
	ReplyToMessageIDSource
	ReplyMarkupSource
}

type SendChatActionRequest struct {
	ChatRequest
	Action ChatAction `json:"action"`
}

type ChatAction string

const (
	ChatActionTyping          ChatAction = "typing"
	ChatActionUploadPhoto     ChatAction = "upload_photo"
	ChatActionRecordVideo     ChatAction = "record_video"
	ChatActionUploadVideo     ChatAction = "upload_video"
	ChatActionRecordAudio     ChatAction = "record_audio"
	ChatActionUploadAudio     ChatAction = "upload_audio"
	ChatActionUploadDocument  ChatAction = "upload_document"
	ChatActionFindLocation    ChatAction = "find_location"
	ChatActionRecordVideoNote ChatAction = "record_video_note"
	ChatActionUploadVideoNote ChatAction = "upload_video_note"
)

type GetUserProfilePhotosRequest struct {
	UserID int `json:"user_id"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

type GetFileRequest struct {
	FileID string `json:"file_id"`
}

type SetChatPhotoRequest struct {
	ChatRequest
	Photo string `json:"photo"`
}

func (c SetChatPhotoRequest) FillFormData(m *map[string]io.Reader, reader io.Reader) {
	(*m)["photo"] = reader
}

type SetChatTitleRequest struct {
	ChatRequest
	Title string `json:"title"`
}

type SetChatDescriptionRequest struct {
	ChatRequest
	Description string `json:"description"`
}

type PinChatMessageRequest struct {
	ChatRequest
	MessageID            int  `json:"message_id"`
	DisableNotifications bool `json:"disable_notifications"`
}

type ChatMember struct {
	ChatPermissions
	User               *User  `json:"user"`
	Status             string `json:"status"`
	UntilDate          int    `json:"until_date"`
	CanBeEdited        bool   `json:"can_be_edited"`
	CanPostMessages    bool   `json:"can_post_messages"`
	CanEditMessages    bool   `json:"can_edit_messages"`
	CanDeleteMessages  bool   `json:"can_delete_messages"`
	CanRestrictMembers bool   `json:"can_restrict_members"`
	CanPromoteMembers  bool   `json:"can_promote_members"`
	IsMember           bool   `json:"is_member"`
}

type GetChatMemberRequest struct {
	ChatRequest
	UserID int `json:"user_id"`
}

type SetChatStickerSetRequest struct {
	ChatRequest
	StickerSetName string `json:"sticker_set_name"`
}

type AnswerCallbackQueryRequest struct {
	CallbackQueryID string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	Url             string `json:"url,omitempty"`
	CacheTime       int    `json:"cache_time,omitempty"`
}

type ParseMode string

const (
	ParseModeMarkdown ParseMode = "Markdown"
	ParseModeHTML     ParseMode = "HTML"
)

type EditMessageReplyMarkupRequest struct {
	ChatRequest
	MessageID       int         `json:"message_id"`
	InlineMessageID int         `json:"inline_message_id"`
	ReplyMarkup     interface{} `json:"reply_markup"`
}

type GetUpdatesRequest struct {
	Offset         int          `json:"offset"`
	Limit          int          `json:"limit"`
	Timeout        int          `json:"timeout"`
	AllowedUpdates []UpdateType `json:"allowed_updates"`
}

type UpdateType string

const (
	UpdateTypeMessage            UpdateType = "message"
	UpdateTypeEditedMessage      UpdateType = "edited_message"
	UpdateTypeChannelPost        UpdateType = "channel_post"
	UpdateTypeEditedChannelPost  UpdateType = "edited_channel_post"
	UpdateTypeInlineResult       UpdateType = "inline_query"
	UpdateTypeChosenInlineResult UpdateType = "chosen_inline_result"
	UpdateTypeCallbackQuery      UpdateType = "callback_query"
	UpdateTypeShippingQuery      UpdateType = "shipping_query"
	UpdateTypePreCheckoutQuery   UpdateType = "pre_checkout_query"
	UpdateTypePoll               UpdateType = "poll"
)
