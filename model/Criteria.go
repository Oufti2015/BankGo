package model

type Criteria struct {
	Criteria      string `json:"criteria"`
	ImpactedField string `json:"field"`
	Comment       string `json:"comment"`
	Period        Period `json:"period"`
}
