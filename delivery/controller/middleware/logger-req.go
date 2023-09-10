package middleware

import (
	"final-project-enigma-clean/config"
	"final-project-enigma-clean/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func LogReqMiddle(log *logrus.Logger) gin.HandlerFunc {
	//init constructor config
	cfg, err := config.NewDbConfig()
	if err != nil {
		fmt.Printf("Failed to read config %v", err.Error())
		return nil
	}

	//open file
	file, err := os.OpenFile(cfg.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	//save output logger to file
	log.SetOutput(file)

	//start time for logger
	startTime := time.Now()

	//init middleware
	return func(c *gin.Context) {
		//next
		c.Next()

		endTime := time.Since(startTime)
		reqlog := model.LoggerReq{
			StartTime:  startTime,
			EndTime:    endTime,
			StatusCode: c.Writer.Status(),
			ClientIP:   c.ClientIP(),
			Method:     c.Request.Method,
			Path:       c.Request.URL.Path,
			UserAgent:  c.Request.UserAgent(),
		}

		//switch
		switch {
		case c.Writer.Status() >= 500: //if we got > 500 status code
			log.Error(reqlog)
		case c.Writer.Status() >= 400: //400-499
			log.Error(reqlog)
		default:
			log.Error(reqlog)
		}
	}
}
