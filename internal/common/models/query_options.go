package models

// QueryOption --
type QueryOption struct {
	Condition   string        `json:"condition"`   // ex : SELECT * FROM `table`  WHERE (field='field')
	Page        int           `json:"page"`        // ex : SELECT * FROM `table`  LIMIT 10 OFFSET 10
	Limit       int           `json:"limit"`       // ex : SELECT * FROM `table`  LIMIT 10 OFFSET 0
	Order       string        `json:"order"`       // ex : SELECT * FROM `table`  ORDER BY ID ASC
	Filter      []Filter      `json:"filter"`      // filtering key on process filter
	ExtraParams []ExtraParams `json:"extraParams"` // Ex: /company/{compID}/address/billing/{billingAddressID}
}

// Filter --
type Filter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// ExtraParams --
type ExtraParams struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

// Pagination --
type Pagination struct {
	Page        int  `json:"page"`
	TotalPages  int  `json:"total_pages"`
	TotalItems  int  `json:"total_items"`
	PerPage     int  `json:"limit"`
	HasNext     bool `json:"has_next"`
	HasPrevious bool `json:"has_previous"`
}
