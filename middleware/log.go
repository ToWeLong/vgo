package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"math"
	"strings"
	"time"
)

type Log struct {
	log *zap.Logger
}

func New(log *zap.Logger) *Log {
	return &Log{log: log}
}

func (l *Log) Log(ctx *gin.Context) {
	start := time.Now()
	data, err := ctx.GetRawData()
	if err != nil {
		l.log.Error("[middleware Log: ]" + err.Error())
		ctx.Abort()
		return
	}
	query := ctx.Copy().Request.URL.RawQuery
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // 重新赋值
	contentType := ctx.Copy().ContentType()
	ctx.Next()
	costs := time.Since(start)
	latency := int(math.Ceil(float64(costs.Nanoseconds()) / 1000000.0))
	newData := strings.TrimSpace(string(data))
	body := strings.ReplaceAll(newData, "\n", "")
	body = strings.ReplaceAll(body, " ", "")
	if body == "" || strings.HasPrefix(contentType, "multipart") {
		body = "{}"
	}
	var token string
	if token = ctx.Copy().Request.Header.Get("Authorization"); token == "" {
		token = ""
	}
	msg := fmt.Sprintf(`[%s] -> [%s] from: %s costs: %vms User-Agent: [%s] token: [%s] 
data: { 
	params: %s, 
	body: %s
}`,
		ctx.Request.Method,
		ctx.Request.RequestURI,
		ctx.ClientIP(),
		latency,
		ctx.Request.UserAgent(),
		token,
		formatQuery(query),
		body,
	)
	l.log.Info(msg)
}

// FormatQuery 将url的query变成JSON的键值对模式, 例如：query=123&u=456 ==> {"query":"123","u":"456"}
func formatQuery(query string) string {
	if query != "" {
		queryParts := strings.Split(query, "&")
		var temp strings.Builder
		temp.Grow(20)
		for _, part := range queryParts {
			s := strings.Split(part, "=")
			temp.WriteString("\"")
			temp.WriteString(s[0])
			temp.WriteString("\":\"")
			temp.WriteString(s[1])
			temp.WriteString("\"")
			temp.WriteString(",")
			//temp = temp + fmt.Sprintf(`"%s":"%s",`, s[0], s[1])
		}
		return "{" + strings.TrimRight(temp.String(), ",") + "}"
	}
	return "{}"
}
