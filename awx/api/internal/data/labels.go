package data

type Labels struct {
	Count   int      `json:"count,omitempty"`
	Results []*Label `json:"results,omitempty"`
}
