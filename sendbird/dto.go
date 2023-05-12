package sendbird

var EndpointSendMessage string = "/v3/group_channels/{channel_url}/messages"

const (
	PushMessageTemplate string = "alternative"
	SendPush            bool   = true
	IsSilent            bool   = false
	MarkAsRead          bool   = false
)

type SendMessageRequest struct {
	MessageType         string `json:"message_type"`
	ChannelURL          string `json:"channel_url"`
	CustomType          string `json:"custom_type"`
	UserId              string `json:"user_id"`
	Message             string `json:"message"`
	Data                string `json:"data,omitempty"`
	SendPush            string `json:"send_push,omitempty"`
	PushMessageTemplate string `json:"push_message_template,omitempty"`
	IsSilent            string `json:"is_silent,omitempty"`
	MarkAsRead          string `json:"mark_as_read,omitempty"`
}

type SendMessageResponse struct {
	Status string `json:"status"`
	// Type       string `json:"type"`
	// ChannelURL string `json:"channel_url"`
	// MessageID  int64  `json:"message_id"`
	// CreatedAt  int64  `json:"created_at"`
}
