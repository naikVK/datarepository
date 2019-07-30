package models

// PersonContactDetails -PersonContactDetails
type PersonContactDetails struct {
	Name    string `xlsx:"1"`
	Mobile  string `xlsx:"0"`
	Address string `xlsx:"2"`
	Tags    string `xlsx:"3"`
}
