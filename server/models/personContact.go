package models

// PersonContactDetails -PersonContactDetails
type PersonContactDetails struct {
	Name    string `xlsx:"1"`
	Mobile  string `xlsx:"0"`
	Address string `xlsx:"2"`
	Tags    string `xlsx:"3"`
	Dob     string `xlsx:"4"`
}

type ContactResponse struct {
	PersonContactDetails []PersonContactDetails `json:"PersonContactDetails"`
	TotalContacts        int                    `json:"TotalContacts"`
}
