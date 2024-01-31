package helpers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogRequest(ctx *gin.Context, requestString string, uuid string) string {
	logger := logrus.New()
	currentTime := time.Now()
	logger.Info("[Start][RequestId]= " + uuid + ", [Path]= " + ctx.FullPath() + ", [IP]= " + ctx.ClientIP() + ", [Time]= " + currentTime.Format("2006-01-02 15:04:05.000000") + ", [Request]= " + requestString)
	return "[Start][RequestId]= " + uuid + ", [Path]= " + ctx.FullPath() + ", [IP]= " + ctx.ClientIP() + ", [Time]= " + currentTime.Format("2006-01-02 15:04:05.000000") + ", [Request]= " + requestString
}

func LogResponse(ctx *gin.Context, responseString string, uuid string) string {
	logger := logrus.New()

	currentTime := time.Now()
	logger.Info("[Stop][RequestId]= " + uuid + ", [Path]= " + ctx.FullPath() + ", [IP]= " + ctx.ClientIP() + ", [Time]= " + currentTime.Format("2006-01-02 15:04:05.000000") + ", [Response]= " + responseString)
	return "[Stop][RequestId]= " + uuid + ", [Path]= " + ctx.FullPath() + ", [IP]= " + ctx.ClientIP() + ", [Time]= " + currentTime.Format("2006-01-02 15:04:05.000000") + ", [Response]= " + responseString
}

func LogError(ctx *gin.Context, responseString string, uuid string) {
	logger := logrus.New()

	currentTime := time.Now()
	logger.Error("[Stop][RequestId]= " + uuid + ", [Path]= " + ctx.FullPath() + ", [IP]= " + ctx.ClientIP() + ", [Time]= " + currentTime.Format("2006-01-02 15:04:05.000000") + ", [Error]= " + responseString)
}
