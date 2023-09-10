package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func ZapLogger(logger *zap.Logger) gin.HandlerFunc {
	// Configure log file path
	logFile := "log.log"

	// open file
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "level",
		TimeKey:      "timestamp",
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	// Create a Zap logger with console and file outputs
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(os.Stdout), zap.InfoLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(file), zap.InfoLevel),
	)
	logger = zap.New(core)

	// Init middleware
	return func(c *gin.Context) {
		// Start time for logger
		startTime := time.Now()

		// Next
		c.Next()

		endTime := time.Since(startTime)
		reqlog := struct {
			StartTime  time.Time
			EndTime    time.Duration
			StatusCode int
			ClientIP   string
			Method     string
			Path       string
			UserAgent  string
		}{
			StartTime:  startTime,
			EndTime:    endTime,
			StatusCode: c.Writer.Status(),
			ClientIP:   c.ClientIP(),
			Method:     c.Request.Method,
			Path:       c.Request.URL.Path,
			UserAgent:  c.Request.UserAgent(),
		}

		switch {
		case c.Writer.Status() >= 400: // If we got >= 500 status code
			logger.Error("Request handled with error", zap.Any("request", reqlog))
		default:
			logger.Info("Request handled", zap.Any("request", reqlog))
		}
	}
}

//func LogReqMiddle(log *logrus.Logger) gin.HandlerFunc {
//	//init constructor config
//	cfg, err := config.NewDbConfig()
//	if err != nil {
//		fmt.Printf("Failed to read config %v", err.Error())
//		return nil
//	}
//
//	//open file
//	file, err := os.OpenFile(cfg.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
//	if err != nil {
//		panic(err)
//	}
//
//	//save output logger to file
//	log.SetOutput(file)
//
//	//start time for logger
//	startTime := time.Now()
//
//	//init middleware
//	return func(c *gin.Context) {
//		//next
//		c.Next()
//
//		endTime := time.Since(startTime)
//		reqlog := model.LoggerReq{
//			StartTime:  startTime,
//			EndTime:    endTime,
//			StatusCode: c.Writer.Status(),
//			ClientIP:   c.ClientIP(),
//			Method:     c.Request.Method,
//			Path:       c.Request.URL.Path,
//			UserAgent:  c.Request.UserAgent(),
//		}
//
//		//switch
//		switch {
//		case c.Writer.Status() >= 500: //if we got > 500 status code
//			log.Error(reqlog)
//		case c.Writer.Status() >= 400: //400-499
//			log.Error(reqlog)
//		default:
//			log.Info(reqlog)
//		}
//	}
//}
