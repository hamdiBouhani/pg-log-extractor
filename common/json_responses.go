package common

import "github.com/gin-gonic/gin"

//ResponseData returns a JSON Data using a gin.Context
func ResponseData(ctx *gin.Context, data interface{}) {
	rsp := gin.H{
		"success": true,
		"data":    data,
	}
	ctx.JSON(200, rsp)
}

//AbortWithError returns a JSON error using a gin.Context
func AbortWithError(
	g *gin.Context,
	status int,
	message string,
	err error,
	request interface{},
) {
	data := gin.H{
		"message": message,
		"success": false,
		"data":    nil,
	}

	if err != nil {
		data["error"] = err.Error()
	}
	g.AbortWithStatusJSON(status, data)
}

//BindJsonErr returns a JSON error using a gin.Context
func BindJsonErr(ctx *gin.Context, err error) {
	AbortWithError(ctx, 400, `request json invalid`, err, nil)
}
