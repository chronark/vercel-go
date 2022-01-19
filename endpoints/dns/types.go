package dns

type CreateDnsRequest struct {
	Domain string `json:"-"`
	Name   string `json:"name"`
	Value  string `json:"value"`
	Type   string `json:"type"`
}

type CreateDnsResponse struct {
	Uid string `json:"uid"`
}

type ListDnsRequest struct {
	Domain string
	// Maximum number of records to list from a request.
	// Required: No
	Limit int

	// Get records created after this JavaScript timestamp.
	// Required: No
	Since int

	// Get records created before this JavaScript timestamp.
	// Required: No
	Until int
}

type Record struct {
	Id         string `json:"id"`
	Slug       string `json:"slug"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Value      string `json:"value"`
	MxPriority int    `json:"mxPriority,omitempty"`
	Priority   int    `json:"priority,omitempty"`
	Creator    string `json:"creator"`
	Created    int    `json:"created,omitempty"`
	Updated    int    `json:"updated,omitempty"`
	CreatedAt  int    `json:"createdAt,omitempty"`
	UpdatedAt  int    `json:"updatedAt,omitempty"`
}
type ListDnsResponse struct {
	Records []Record `json:"records"`

	Pagination struct {
		Count int   `json:"count"`
		Next  int64 `json:"next"`
		Prev  int64 `json:"prev"`
	} `json:"pagination,omitempty"`
}

type DeleteDnsRequest struct {
	Domain string
	Record string
}

type DeleteDnsResponse struct{}
