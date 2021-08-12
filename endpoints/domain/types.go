package domain

type Domain struct {
	// If it was purchased through Vercel, the date when it was purchased.
	BoughtAt         interface{} `json:"boughtAt"`
	ConfigVerifiedAt interface{} `json:"configVerifiedAt"`

	// The date when the domain was created in the registry.
	CreatedAt int64 `json:"createdAt"`

	// The date at which the domain is set to expire. null if not bought with Vercel.
	ExpiresAt interface{} `json:"expiresAt"`
	// The unique ID of the domain.
	ID string `json:"id"`

	// The domain name.
	Name string `json:"name"`

	// The date at which the domain's nameservers were verified based on the intended set.
	NsVerifiedAt interface{} `json:"nsVerifiedAt"`

	// The type of service the domain is handled by. external if the DNS is externally handled,
	// or zeit.world if handled with Vercel.
	ServiceType   string      `json:"serviceType"`
	TransferredAt interface{} `json:"transferredAt"`

	// The date at which the domain's TXT DNS record was verified.
	TxtVerifiedAt interface{} `json:"txtVerifiedAt"`

	// The ID of the verification record in the registry.
	VerificationRecord string `json:"verificationRecord"`

	// Whether the domain has the Vercel Edge Network enabled or not.
	CdnEnabled bool `json:"cdnEnabled"`

	// If the domain has the ownership verified.
	Verified bool `json:"verified"`

	// A list of the current nameservers of the domain.
	Nameservers []string `json:"nameservers"`

	// If the domain will be auto renewed. undefined is equivalent to true.
	Renew bool `json:"renew"`

	// A list of the intended nameservers for the domain to point to Vercel DNS.
	IntendedNameservers []string `json:"intendedNameservers"`

	// An object containing information of the domain creator, including the user's id, username, and email.
	Creator struct {
		ID         string      `json:"id"`
		CustomerID interface{} `json:"customerId"`
		Email      string      `json:"email"`
		Username   string      `json:"username"`
		Name       string      `json:"name"`
	} `json:"creator"`
	Zone bool `json:"zone"`
}

// The response of this API can be controlled with the following optional query parameters.
type ListDomainsRequest struct {
	// Maximum number of domains to list from a request.
	// Required: No
	Limit int `json:"limit,omitempty"`

	// Get domains created after this JavaScript timestamp.
	// Required: No
	Since int `json:"since,omitempty"`

	// Get domains created before this JavaScript timestamp.
	// Required: No
	Until int `json:"until,omitempty"`
}

type ListDomainsResponse struct {
	Domains    []Domain `json:"domains"`
	Pagination struct {
		Count int   `json:"count"`
		Next  int64 `json:"next"`
		Prev  int64 `json:"prev"`
	} `json:"pagination"`
}

type AddDomainRequest struct {
	// The domain name you want to add.
	// Required: Yes
	Name string `json:"name"`
}

type AddDomainResponse struct {
	Domain Domain `json:"domain"`
}

type TransferDomainRequest struct {
	// The domain add operation to perform. When transferring in, use transfer-in.
	// Required: Yes
	Method string `json:"method"`

	// The domain name you want to add.
	// Required: Yes
	Name string `json:"name"`

	// The authorization code assigned to the domain.
	// Required: Yes
	AuthCode string `json:"authCode"`

	// The price you expect to be charged for the required 1 year renewal.
	// Required: Yes
	ExpectedPrice int `json:"expectedPrice"`
}

type TransferDomainResponse struct {
	Domain Domain `json:"domain"`
}

type VerifyDomainRequest struct {
	// The domain name you want to verify.
	// Required: Yes
	Name string `json:"name"`
}

type VerifyDomainResponse struct {
	Domain Domain `json:"domain"`
}

type GetDomainRequest struct {
	// The domain name you want to get.
	// Required: Yes
	Name string `json:"name"`
}
type GetDomainResponse struct {
	Domain Domain `json:"domain"`
}
type GetAuthCodeRequest struct {
	// The domain name you want to authenticate.
	// Required: Yes
	Name string `json:"name"`
}

type GetAuthCodeResponse struct {
	AuthCode string `json:"authCode"`
}

type RemoveDomainRequest struct {
	// The domain name you want to remove.
	// Required: Yes
	Name string `json:"name"`
}

type RemoveDomainResponse struct {
	// The unique ID of the removed domain.
	Uid string `json:"uid"`
}

type CheckAvailabilityRequest struct {
	// The domain name you want to check.
	// Required: Yes
	Name string `json:"name"`
}

type CheckAvailabilityResponse struct {
	// If the domain is available or not.
	Available bool `json:"available"`
}

type CheckPriceRequest struct {
	// The domain name you want to check.
	// Required: Yes
	Name string `json:"name"`
}

type CheckPriceResponse struct {
	// The domain price.
	Price int `json:"price"`

	// The time period by which the domain is purchased.
	Period int `json:"period"`
}
type PurchaseRequest struct {
	// The domain name to purchase.
	// Required: Yes
	Name string `json:"name"`

	// The price you expect to be charged for the purchase.
	// Required: Yes
	ExpectedPrice int `json:"expectedPrice"`
}

// In case of a successful purchase the returns with code 200 and an empty body.
type PurchaseResponse struct{}
