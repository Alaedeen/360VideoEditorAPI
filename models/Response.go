package models



// Response Struct
type Response struct {
	Code	int 		`json:"code"`
	Status	string 		`json:"status"`
	Data	interface{}	`json:"data"`
}