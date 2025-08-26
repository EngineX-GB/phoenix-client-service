package model

import "time"

type Client struct {
	Oid         int64     `json:"oid"`
	Username    string    `json:"username"`
	Nationality string    `json:"nationality"`
	Location    string    `json:"location"`
	Rating      int       `json:"rating"`
	Age         int       `json:"age"`
	R15         int       `json:"r15"`
	R30         int       `json:"r30"`
	R45         int       `json:"r45"`
	R100        int       `json:"r100"`
	R150        int       `json:"r150"`
	R200        int       `json:"r200"`
	R250        int       `json:"r250"`
	R300        int       `json:"r300"`
	R350        int       `json:"r350"`
	R400        int       `json:"r400"`
	RON         int       `json:"rOn"`
	Telephone   string    `json:"telephone"`
	Url         string    `json:"url"`
	RefreshTime time.Time `json:"refreshTime"`
	UserId      string    `json:"userId"`
	Region      string    `json:"region"`
	Gender      string    `json:"gender"`
	MemberSince time.Time `json:"memberSince"`
	Height      int       `json:"height"`
	DSize       int       `json:"dSize"`
	HairColor   string    `json:"hairColor"`
	EyeColor    string    `json:"eyeColor"`
}
