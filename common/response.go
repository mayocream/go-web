package common

import "github.com/gin-gonic/gin"

type ServiceResponse struct {
	Data		interface{}
	ApiError 	ApiError
}

type ControllerResponse struct {
	ServiceResponse   ServiceResponse
}

type JsonResponse struct {
	StatusCode  int          `json:"-"`
	Code 		string		 `json:"code"`
	Data    	interface{}  `json:"data"`
	Msg     	string		 `json:"msg"`
}

func (res *JsonResponse) Return(c *gin.Context) {
	if res.StatusCode == 0 {
		res.StatusCode = 200
	}
	c.JSON(res.StatusCode, res)
}