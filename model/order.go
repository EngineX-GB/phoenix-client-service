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

// create table if not exists tbl_order (
// 	oid integer not null auto_increment,
// 	user_id varchar(10),
// 	location varchar(200),
//     region varchar(100),
// 	date_of_event date,
// 	time_of_event time,
// 	creation_date datetime,
// 	modification_date datetime,
// 	duration integer,
// 	rate integer,
// 	deductions integer,
// 	surplus integer,
// 	price integer, 			-- the final price
// 	status varchar(20),
// 	notes text,
// 	primary key (oid)
// );
