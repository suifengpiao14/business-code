package config

import "github.com/gin-gonic/gin"

//Pagination 分页
type Pagination struct {
	Size   int64 `json:"size"`
	Page   int64 `json:"page"`
	Total  int64 `json:"total"`
	Offset int64 `json:"-"`
	Limit  int64 `json:"-"`
}

//PaginationResponse 分页返回
func PaginationResponse(c *gin.Context, pagination *Pagination, items interface{}) {
	response := &gin.H{
		"items":      items,
		"pagination": pagination,
	}
	c.JSON(200, response)

}

//ErrorResponse 异常返回
func ErrorResponse(c *gin.Context, httpCode int, errorCode int, err error) {
	response := &gin.H{
		"code": errorCode,
		"msg":  err.Error(),
	}
	c.AbortWithStatusJSON(httpCode, response)
}
