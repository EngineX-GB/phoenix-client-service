package model

type QueryMarker struct {
	Min int64 `json: "min"`
	Max int64 `json: "max"`
}
