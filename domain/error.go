package domain

import "errors"

const (
	//InternalServerError indicates that the server encountered an unexpected condition that prevented it from fulfilling the request
	InternalServerError = "Internal Server Error"
	//BadRequest indicates that the server cannot or will not process the request due to something that is perceived to be a client error
	BadRequest = "Bad Request"
	//InvalidTodoID represents an error when the user ID is invalid
	InvalidTodoID = "Todo ID is invalid"
	//MIssingTodoID represents an error when the user ID is missing from the request
	MIssingTodoID = "Todo ID is missing"
)

//ErrTodoNotFound represents an error when the user ID is not found in the request
var ErrTodoNotFound = errors.New("Requested Todo ID was not found")
