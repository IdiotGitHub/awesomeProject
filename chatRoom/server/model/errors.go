package model

import "errors"

var (
	//UserNotExists
	UserNotExists    = errors.New("UserNotExists")
	UserIsExists     = errors.New("UserIsExists")
	JsonMarshalError = errors.New("JsonMarshalError")
)
