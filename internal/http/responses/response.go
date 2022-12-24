package responses

import (
	"github.com/gin-gonic/gin"
)

type JSONResponse interface {
	Response(c *gin.Context, code int)
}

type response struct {
	Status baseStatus  `json:"status"`
	Data   interface{} `json:"data"`
}

type baseStatus struct {
	Message string `json:"message,omitempty"`
}

func NewResponse(data interface{}) JSONResponse {
	return response{
		Status: baseStatus{
			Message: "OK",
		},
		Data: data,
	}
}
func NewErrorResponse(err error) JSONResponse {
	return response{
		Status: baseStatus{
			Message: err.Error(),
		},
	}
}

// Response make a http writer to write json body and http status code
func (r response) Response(c *gin.Context, httpStatus int) {
	c.AbortWithStatusJSON(httpStatus, r)
}
