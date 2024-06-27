package page

import (
	"github.com/bastean/dsgo/pkg/cmd/server/component/page/home"
	"github.com/gin-gonic/gin"
)

func Home() gin.HandlerFunc {
	return func(c *gin.Context) {
		home.Page().Render(c.Request.Context(), c.Writer)
	}
}
