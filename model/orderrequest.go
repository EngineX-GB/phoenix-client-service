package model

type OrderRequest struct {
	UserId    string `json:"userId"`
	UserName  string `json:"username"`
	Location  string `json:"location"`
	Region    string `json:"region"`
	Telephone string `json:"telephone"`
	R15       int    `json:"r15"`
	R30       int    `json:"r30"`
	R45       int    `json:"r45"`
	R100      int    `json:"r100"`
	R150      int    `json:"r150"`
	R200      int    `json:"r200"`
	R250      int    `json:"r250"`
	R300      int    `json:"r300"`
	R350      int    `json:"r350"`
	R400      int    `json:"r400"`
	RON       int    `json:"rOn"`
}
