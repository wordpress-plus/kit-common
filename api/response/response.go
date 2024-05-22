package response

import (
	"github.com/micro-services-roadmap/oneid-core/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, model.Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(c *gin.Context) {
	Result(model.SUCCESS, map[string]interface{}{}, "Success", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(model.SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(model.SUCCESS, data, "Success", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(model.SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(model.ERROR, map[string]interface{}{}, "Fail", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(model.ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(model.ERROR, data, message, c)
}

func FailWithError(data error) *model.Response {
	return &model.Response{
		Code: model.SUCCESS,
		Data: data.Error(),
		Msg:  "Failed",
	}
}
