package evmchat

const (
	ChannelEvermosPromo string = "evm_promo"
)

type JoinSuperGroupRequest struct {
	CustomType string `json:"custom_type"`
	UserID     string `json:"user_id"`
}

type JoinSuperGroupResponse struct {
	ChannelName string `json:"channel_name"`
	ChannelURL  string `json:"channel_url"`
	CustomType  string `json:"custom_type"`
	MemberCount int    `json:"member_count"`
}
