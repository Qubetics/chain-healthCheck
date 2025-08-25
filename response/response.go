package response

import "github.com/labstack/echo/v4"

// Response is the standard API response structure
// Status: HTTP status code
// Msg: message string
// Data: payload

type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

// New creates a new Response
func New(status int, msg string, data interface{}) *Response {
	return &Response{Status: status, Msg: msg, Data: data}
}

// JSON returns the response as JSON using echo.Context
func JSON(c echo.Context, status int, msg string, data interface{}) error {
	return c.JSON(status, New(status, msg, data))
}
