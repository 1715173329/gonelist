package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gonelist/onedrive"
	"gonelist/pkg/app"
	"gonelist/pkg/e"
	"net/http"
)

// 测试接口，从 MG 获取整个树结构
func MGGetFileTree(c *gin.Context) {
	root, err := onedrive.GetAllFiles()
	if err != nil {
		log.Warn("请求 graph.microsoft.com 错误")
		app.Response(c, http.StatusOK, e.MG_ERROR, nil)
		return
	}

	str, _ := json.Marshal(root)
	log.Info("*root", string(str))

	app.Response(c, http.StatusOK, e.SUCCESS, root)
}

// 获取对应路径的文件
func CacheGetPath(c *gin.Context) {
	oPath := c.Query("path")

	root, err := onedrive.CacheGetPathList(oPath)
	if err != nil {
		app.Response(c, http.StatusOK, e.ITEM_NOT_FOUND, nil)
	} else {
		app.Response(c, http.StatusOK, e.SUCCESS, root)
	}
}

// 分享文件下载链接
func Download(c *gin.Context) {
	filePath := c.Param("path")

	downloadURL, err := onedrive.GetDownloadUrl(filePath)
	if err != nil {
		app.Response(c, http.StatusOK, e.ITEM_NOT_FOUND, nil)
	} else {
		c.Redirect(http.StatusFound, downloadURL)
		c.Abort()
	}
}
