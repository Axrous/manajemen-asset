package middleware

import (
	"final-project-enigma-clean/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()
        for _, err := range c.Errors {
            switch e := err.Err.(type) {
            case *exception.Http :
                c.AbortWithStatusJSON(e.StatusCode, e)
            default:
                c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"message": e.Error()})
            }
        }
    }
}