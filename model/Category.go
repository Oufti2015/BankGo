package model

type Category struct {
	Priority int    `json:"priority"`
	Name     string `json:"name"`
	Income   bool   `json:"income"`
	Saving   bool   `json:"saving"`

	Criteria []Criteria
}
