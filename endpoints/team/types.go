package team

import "time"

// Either `Id` or `Slug` must be defined
type GetTeamRequest struct {
	// The team unique identifier. Always prepended by team_.
	// Required: false
	Id string

	// The team slug. A slugified version of the name.
	Slug string
}

type GetTeamResponse struct {
	ID         string    `json:"id"`
	Slug       string    `json:"slug"`
	Name       string    `json:"name"`
	CreatorID  string    `json:"creatorId"`
	Created    time.Time `json:"created"`
	CreatedAt  int64     `json:"createdAt"`
	UpdatedAt  int64     `json:"updatedAt"`
	Avatar     string    `json:"avatar"`
	Membership struct {
		Role      string `json:"role"`
		Confirmed bool   `json:"confirmed"`
		Created   int64  `json:"created"`
		CreatedAt int64  `json:"createdAt"`
		UID       string `json:"uid"`
		UpdatedAt int64  `json:"updatedAt"`
	} `json:"membership"`
	PlatformVersion interface{} `json:"platformVersion"`
	InviteCode      string      `json:"inviteCode"`
	Billing         struct {
		Address struct {
			State      string `json:"state"`
			Country    string `json:"country"`
			City       string `json:"city"`
			PostalCode string `json:"postalCode"`
			Line1      string `json:"line1"`
		} `json:"address"`
		Cancelation interface{} `json:"cancelation"`
		Email       string      `json:"email"`
		Language    interface{} `json:"language"`
		Name        string      `json:"name"`
		Period      struct {
			Start int64 `json:"start"`
			End   int64 `json:"end"`
		} `json:"period"`
		Plan         string      `json:"plan"`
		Tax          interface{} `json:"tax"`
		Currency     string      `json:"currency"`
		Trial        interface{} `json:"trial"`
		InvoiceItems struct {
			Pro struct {
				Price    int  `json:"price"`
				Quantity int  `json:"quantity"`
				Hidden   bool `json:"hidden"`
			} `json:"pro"`
			TeamSeats struct {
				Hidden   bool `json:"hidden"`
				Price    int  `json:"price"`
				Quantity int  `json:"quantity"`
			} `json:"teamSeats"`
			ConcurrentBuilds struct {
				Hidden   bool `json:"hidden"`
				Price    int  `json:"price"`
				Quantity int  `json:"quantity"`
			} `json:"concurrentBuilds"`
		} `json:"invoiceItems"`
		Addons        interface{} `json:"addons"`
		Platform      string      `json:"platform"`
		Overdue       interface{} `json:"overdue"`
		Subscriptions []struct {
			ID     string      `json:"id"`
			Trial  interface{} `json:"trial"`
			Period struct {
				Start int64 `json:"start"`
				End   int64 `json:"end"`
			} `json:"period"`
			Frequency struct {
				Interval      string `json:"interval"`
				IntervalCount int    `json:"intervalCount"`
			} `json:"frequency"`
			Discount interface{} `json:"discount"`
			Items    []struct {
				ID        string `json:"id"`
				PriceID   string `json:"priceId"`
				ProductID string `json:"productId"`
				Amount    int    `json:"amount"`
				Quantity  int    `json:"quantity"`
			} `json:"items"`
		} `json:"subscriptions"`
	} `json:"billing"`
	Description    interface{}   `json:"description"`
	Profiles       []interface{} `json:"profiles"`
	StagingPrefix  string        `json:"stagingPrefix"`
	ResourceConfig struct {
		ConcurrentBuilds int `json:"concurrentBuilds"`
	} `json:"resourceConfig"`
	PreviewDeploymentSuffix interface{} `json:"previewDeploymentSuffix"`
	SoftBlock               interface{} `json:"softBlock"`
	AllowProjectTransfers   bool        `json:"allowProjectTransfers"`
}
