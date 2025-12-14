package model

type Order struct {
	Oid          uint64 `json:"oid"`
	UserId       string `json:"userId"`
	Location     string `json:"location"`
	Region       string `json:"region"`
	DateOfEvent  string `json:"dateOfEvent"`
	TimeOfEvent  string `json:"timeOfEvent"`
	CreationDate string `json:"creationDate"`
	ModifiedDate string `json:"modifiedDate"`
	Duration     uint64 `json:"duration"`
	Rate         uint64 `json:"rate"`
	Deductions   uint64 `json:"deductions"`
	Surplus      uint64 `json:"surplus"`
	Price        uint64 `json:"price"`
	Status       string `json:"status"`
	Notes        string `json:"notes"`
}
