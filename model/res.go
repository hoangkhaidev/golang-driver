package model

type Response struct {
	StatusCode 	int 		`json:"status_code"`
	Message 	string 		`json:"message,omitempty"`
	Data 		interface{}	`json:"data,omitempty"`
	Err 		string		`json:"err,omitempty"`
}