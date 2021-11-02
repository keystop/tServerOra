package models

import "context"

type CardTC struct {
	NumTC      string `json:"num_tc"`
	ModelTC    string `json:"model_tc"`
	DriverName string `json:"driver_name"`
	UserID     string
}

//Repository interface repo urls.
type Repository interface {
	SaveCard(context.Context, *CardTC) error
	CheckDBConnection(context.Context) error
	CreateUser(context.Context) (string, error)
}
