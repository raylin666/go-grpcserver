package params

type GetTokenAuthReq struct {
	Key    string `json:"key" form:"key" validate:"required" label:"颁布标识 Key"`
	Secret string `json:"secret" form:"secret" validate:"required" label:"颁布标识 Secret"`
	UserID string `json:"user_id" form:"user_id" validate:"required" label:"用户标识 ID"`
	TTL    uint32 `json:"ttl" form:"ttl"`
}

type GetTokenAuthResp struct {

}