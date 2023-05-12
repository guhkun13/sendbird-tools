package sendbird

const (
	PushMessageTemplate string = "alternative"
	SendPush            bool   = true
	IsSilent            bool   = false
	MarkAsRead          bool   = false
)

type ChannelResource struct {
	Name                 string `json:"name"`
	ChannelURL           string `json:"channel_url"`
	CoverURL             string `json:"cover_url"`
	CustomType           string `json:"custom_type"`
	UnreadMessageCount   int    `json:"unread_message_count"`
	Data                 string `json:"data"`
	IsDistinct           bool   `json:"is_distinct"`
	IsPublic             bool   `json:"is_public"`
	IsSuper              bool   `json:"is_super"`
	IsEphemeral          bool   `json:"is_ephemeral"`
	IsAccessCodeRequired bool   `json:"is_access_code_required"`
	HiddenState          string `json:"hidden_state"`
	MemberCount          int    `json:"member_count"`
	JoinedMemberCount    int    `json:"joined_member_count"`
	Members              []struct {
		UserID             string   `json:"user_id"`
		Nickname           string   `json:"nickname"`
		ProfileURL         string   `json:"profile_url"`
		IsActive           bool     `json:"is_active"`
		IsOnline           bool     `json:"is_online"`
		FriendDiscoveryKey []string `json:"friend_discovery_key"`
		LastSeenAt         int64    `json:"last_seen_at"`
		State              string   `json:"state"`
		Role               string   `json:"role"`
		Metadata           struct {
			Location string `json:"location"`
			Marriage string `json:"marriage"`
		} `json:"metadata"`
	} `json:"members"`
	Operators []struct {
		UserID     string `json:"user_id"`
		Nickname   string `json:"nickname"`
		ProfileURL string `json:"profile_url"`
		Metadata   struct {
			Location string `json:"location"`
			Marriage string `json:"marriage"`
		} `json:"metadata"`
	} `json:"operators"`
	MaxLengthMessage int         `json:"max_length_message"`
	LastMessage      interface{} `json:"last_message"`
	CreatedAt        int         `json:"created_at"`
	Freeze           bool        `json:"freeze"`
}

type CreateGroupChannelRequest struct {
	Name       string  `json:"name"`
	ChannelURL string  `json:"channel_url"`
	CoverURL   string  `json:"cover_url"`
	CustomType string  `json:"custom_type"`
	IsDistinct bool    `json:"is_distinct"`
	UserIDs    []int64 `json:"user_ids,omitempty"`
	Data       string  `json:"data"`
	IsSuper    bool    `json:"is_super,omitempty"`
}

type CreateGroupChannelResponse ChannelResource

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
