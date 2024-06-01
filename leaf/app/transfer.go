package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlerfunc func(c *gin.Context) error

func TransferHandlerfunc(h Handlerfunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if err := h(ctx); err != nil {
			if e, ok := err.(APIError); ok {
				ctx.JSON(e.Status, &e)
				return
			}
			arr := APIError{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
			ctx.JSON(arr.Status, &arr)
		}
	}
}
