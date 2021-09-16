package handlers

type CreateSubscriberReq struct {
	UserID     uint64 `json:"userId"`
	Source     string `json:"source,omitempty"`
	IsTestUser bool   `json:"isTest,omitempty"`
}

type GetSubscriptionInfoReq struct {
	UserID uint64 `form:"userId"`
}

type GetReferralCodeReq struct {
	UserID uint64 `form:"userId"`
}
