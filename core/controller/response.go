package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Context *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	res := &Response{Context: ctx}
	data, err := ctx.GetRawData()
	if err != nil && data != nil && len(data) > 1 {
		err = json.Unmarshal(data, res)
		if err != nil {

		}
	}
	return res
}

func (c *Response) SuccResponse(data interface{}) {
	c.Context.JSON(http.StatusOK, data)
}
