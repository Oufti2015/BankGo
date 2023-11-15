package model

type Category struct {
	Priority int32  `json:"priority"`
	Name     string `json:"name"`
	Income   bool   `json:"income"`
	Saving   bool   `json:"saving"`
}
