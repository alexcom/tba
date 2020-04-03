package telegram

type InlineQueryResultType string

const (
	InlineQueryResultTypeArticle  InlineQueryResultType = "article"
	InlineQueryResultTypePhoto    InlineQueryResultType = "photo"
	InlineQueryResultTypeGif      InlineQueryResultType = "gif"
	InlineQueryResultTypeMpeg4Gif InlineQueryResultType = "mpeg4_gif"
	InlineQueryResultTypeVideo    InlineQueryResultType = "video"
	InlineQueryResultTypeAudio    InlineQueryResultType = "audio"
	InlineQueryResultTypeVoice    InlineQueryResultType = "voice"
	InlineQueryResultTypeDocument InlineQueryResultType = "document"
	InlineQueryResultTypeLocation InlineQueryResultType = "location"
	InlineQueryResultTypeVenue    InlineQueryResultType = "venue"
	InlineQueryResultTypeContact  InlineQueryResultType = "contact"
	InlineQueryResultTypeGame     InlineQueryResultType = "game"
	InlineQueryResultTypeSticker  InlineQueryResultType = "sticker"
)

type IDAndTypeSource struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type TitleSource struct {
	Title string `json:"title"`
}

type DescriptionSource struct {
	Description string `json:"description"`
}

type InputMessageContentSource struct {
	InputMessageContent interface{} `json:"input_message_content"`
}

type InlineQueryResultCachedAudio struct {
	IDAndTypeSource
	InputMessageContentSource
	ReplyMarkupSource
	AudioFileID string `json:"audio_file_id"`
	CaptionSource
	ParseModeSource
}

type InlineQueryResultCachedDocument struct {
	IDAndTypeSource
	InputMessageContentSource
	ReplyMarkupSource
	TitleSource
	DocumentFileID string `json:"document_file_id"`
	DescriptionSource
	CaptionSource
	ParseModeSource
}

type InlineQueryResultCachedGif struct {
	IDAndTypeSource
	InputMessageContentSource
	ReplyMarkupSource
	GifFileID string `json:"gif_file_id"`
	TitleSource
	CaptionSource
	ParseModeSource
}

type InlineQueryResultCachedMpeg4Gif struct {
	IDAndTypeSource
	InputMessageContentSource
	ReplyMarkupSource
	Mpeg4FileID string `json:"mpeg4_file_id"`
	TitleSource
	CaptionSource
	ParseModeSource
}

type InlineQueryResultCachedPhoto struct {
	IDAndTypeSource
	InputMessageContentSource
	ReplyMarkupSource
	PhotoFileID string `json:"photo_file_id"`
	TitleSource
	DescriptionSource
	CaptionSource
	ParseModeSource
}

type InlineQueryResultCachedSticker struct {
	IDAndTypeSource
	InputMessageContentSource
	ReplyMarkupSource
	StickerFileID string `json:"sticker_file_id"`
}

type InlineQueryResultCachedVideo struct {
	ReplyMarkupSource
	IDAndTypeSource
	InputMessageContentSource
	VideoFileID string `json:"video_file_id"`
	TitleSource
	DescriptionSource
	CaptionSource
	ParseModeSource
}

type InlineQueryResultCachedVoice struct {
	ReplyMarkupSource
	IDAndTypeSource
	InputMessageContentSource
	VoiceFileID string `json:"voice_file_id"`
	TitleSource
	CaptionSource
	ParseModeSource
}

type InlineQueryResultArticle struct {
	IDAndTypeSource
	InputMessageContentSource
	ReplyMarkupSource
	TitleSource
	Url     string `json:"url"`
	HideUrl bool   `json:"hide_url"`
	DescriptionSource
	ThumbUrl    string `json:"thumb_url"`
	ThumbWidth  int    `json:"thumb_width"`
	ThumbHeight int    `json:"thumb_height"`
}

type InlineQueryResultAudio struct {
	IDAndTypeSource
	InputMessageContentSource
	ReplyMarkupSource
	AudioUrl string `json:"audio_url"`
	TitleSource
	CaptionSource
	ParseModeSource
	Performer     string `json:"performer"`
	AudioDuration int    `json:"audio_duration"`
}

type InlineQueryResultContact struct {
	IDAndTypeSource
	InputMessageContentSource
	ReplyMarkupSource
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Vcard       string `json:"vcard"`
	ThumbUrl    string `json:"thumb_url"`
	ThumbWidth  int    `json:"thumb_width"`
	ThumbHeight int    `json:"thumb_height"`
}

type InlineQueryResultGame struct {
	IDAndTypeSource
	ReplyMarkupSource
	GameShortName string `json:"game_short_name"`
}

type InlineQueryResultDocument struct {
	IDAndTypeSource
	InputMessageContentSource
	ReplyMarkupSource
	TitleSource
	CaptionSource
	ParseModeSource
	DocumentUrl string `json:"document_url"`
	MimeType    string `json:"mime_type"`
	DescriptionSource
	ThumbUrl    string `json:"thumb_url"`
	ThumbWidth  int    `json:"thumb_width"`
	ThumbHeight int    `json:"thumb_height"`
}

type InlineQueryResultGif struct {
	IDAndTypeSource
	InputMessageContentSource
	ReplyMarkupSource
	GifUrl      string `json:"gif_url"`
	GifWidth    int    `json:"gif_width"`
	GifHeight   int    `json:"gif_height"`
	GifDuration int    `json:"gif_duration"`
	ThumbUrl    string `json:"thumb_url"`
	TitleSource
	CaptionSource
	ParseModeSource
}

type InlineQueryResultLocation struct {
	IDAndTypeSource
	InputMessageContentSource
	ReplyMarkupSource
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Title       string  `json:"title"`
	LivePeriod  int     `json:"live_period"`
	ThumbUrl    string  `json:"thumb_url"`
	ThumbWidth  int     `json:"thumb_width"`
	ThumbHeight int     `json:"thumb_height"`
}

type InlineQueryResultMpeg4Gif struct {
	IDAndTypeSource
	InputMessageContentSource
	Mpeg4Url      string `json:"mpeg4_url"`
	Mpeg4Width    int    `json:"mpeg4_width"`
	Mpeg4Height   int    `json:"mpeg4_height"`
	Mpeg4Duration int    `json:"mpeg4_duration"`
	ThumbUrl      string `json:"thumb_url"`
	TitleSource
	CaptionSource
	ParseModeSource
	ReplyMarkupSource
}

type InlineQueryResultPhoto struct {
	IDAndTypeSource
	InputMessageContentSource
	ReplyMarkupSource
	PhotoUrl    string `json:"photo_url"`
	ThumbUrl    string `json:"thumb_url"`
	PhotoWidth  int    `json:"photo_width"`
	PhotoHeight int    `json:"photo_height"`
	TitleSource
	DescriptionSource
	CaptionSource
	ParseModeSource
}

type InlineQueryResultVenue struct {
	IDAndTypeSource
	InputMessageContentSource
	ReplyMarkupSource
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	TitleSource
	Address        string `json:"address"`
	FoursquareID   string `json:"foursquare_id"`
	FoursquareType string `json:"foursquare_type"`
	ThumbUrl       string `json:"thumb_url"`
	ThumbWidth     int    `json:"thumb_width"`
	ThumbHeight    int    `json:"thumb_height"`
}

type InlineQueryResultVideo struct {
	IDAndTypeSource
	InputMessageContentSource
	ReplyMarkupSource
	VideoUrl string `json:"video_url"`
	MimeType string `json:"mime_type"`
	ThumbUrl string `json:"thumb_url"`
	TitleSource
	CaptionSource
	ParseModeSource
	VideoWidth    int `json:"video_width"`
	VideoHeight   int `json:"video_height"`
	VideoDuration int `json:"video_duration"`
	DescriptionSource
}

type InlineQueryResultVoice struct {
	IDAndTypeSource
	InputMessageContentSource
	ReplyMarkupSource
	VoiceUrl string `json:"voice_url"`
	TitleSource
	CaptionSource
	ParseModeSource
	VoiceDuration int `json:"voice_duration"`
}

type InputTextMessageContent struct {
	MessageText string `json:"message_text"`
	ParseModeSource
	DisableWebPagePreview bool `json:"disable_web_page_preview"`
}
type InputLocationMessageContent struct {
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	LivePeriod int     `json:"live_period"`
}
type InputVenueMessageContent struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	TitleSource
	Address        string `json:"address"`
	FoursquareId   string `json:"foursquare_id"`
	FoursquareType string `json:"foursquare_type"`
}
type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Vcard       string `json:"vcard"`
}
