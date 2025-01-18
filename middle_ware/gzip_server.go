package middleware

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func GzipServer(staticDir string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/h5/") {
			// 去掉 /h5 前缀后的请求路径
			requestedPath := strings.TrimPrefix(c.Request.URL.Path, "/h5/")
			filePath := filepath.Join(staticDir, requestedPath)
			gzFilePath := filePath + ".gz"

			// 检查是否有 .gz 文件
			if fileExists(gzFilePath) {
				c.Header("Content-Encoding", "gzip") // 告诉浏览器这是 gzip 文件
				c.Header("Content-Type", getContentType(filePath))
				c.File(gzFilePath)
				c.Abort()
				return
			}
		}
		// 如果不是 /h5 或没有 .gz 文件，则继续执行其他处理
		c.Next()
	}
}

// 判断文件是否存在
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// 获取 MIME 类型
func getContentType(filePath string) string {
	ext := filepath.Ext(filePath)
	switch ext {
	case ".html":
		return "text/html"
	case ".css":
		return "text/css"
	case ".js":
		return "application/javascript"
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	case ".svg":
		return "image/svg+xml"
	default:
		return "application/octet-stream"
	}
}
