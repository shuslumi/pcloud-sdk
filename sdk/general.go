package sdk

import (
	"context"
	"net/url"
)

type UserInfo struct {
	result
	CryptoSetup           bool
	Plan                  int
	CryptoSubscription    bool
	UserID                int
	HasPassword           bool
	PublicLinkQuota       uint64
	CryptoLifetime        bool
	PremiumExpires        APITime
	Email                 string
	TrashRevRetentionDays int
	Auth                  string
	EmailVerified         bool
	UsedPublinkBranding   bool
	Currency              string
	AgreedWithPP          bool // pp: privacy policy
	Quota                 uint64
	CryptoExpires         APITime
	Premium               bool
	PremiumLifetime       bool
	Business              bool
	UsedQuota             uint64
	Language              string
	HasPaidRelocation     bool
	Registered            APITime
	RegistrationInfo      RegistrationInfo
	Journey               Journey
	APIServer             APIServer
}

type RegistrationInfo struct {
	Provider int
	Device   string
	Country  string
	Ref      int
}

type Journey struct {
	Steps Steps
}

type Steps struct {
	VerifyMail     bool
	UploadFile     bool
	AutoUpload     bool
	DownloadApp    bool
	DownloadDrive  bool
	SentInvitation bool
}

type APIServer struct {
	BinAPI []string
	API    []string
}

func toQuery(opts ...ClientOption) url.Values {
	q := url.Values{}

	for _, opt := range opts {
		opt(&q)
	}

	return q
}

// UserInfo returns information about the current user.
// As there is no specific login method as credentials can be passed to any method,
// this is an especially good place for logging in with no particular action in mind.
// https://docs.pcloud.com/methods/general/userinfo.html
func (c *Client) UserInfo(ctx context.Context, opts ...ClientOption) (*UserInfo, error) {
	q := toQuery(opts...)
	q.Add("getregistrationinfo", "1")
	q.Add("getapiserver", "1")

	ui := &UserInfo{}

	err := parseAPIOutput(ui)(c.request(ctx, "userinfo", q))
	if err != nil {
		return nil, err
	}

	return ui, nil
}
