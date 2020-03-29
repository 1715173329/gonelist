package onedrive

import (
	"GOIndex/conf"
	log "github.com/sirupsen/logrus"
	"time"
)

// 定时刷新缓存
func SetAutoRefresh() {
	timer(AutoRefresh)
}

func AutoRefresh() {
	log.WithFields(log.Fields{
		"time": time.Now(),
	}).Info("自动刷新所有缓存")
	GetAllFiles()  // 获取所有文件并且刷新树结构
	log.Debug(FileTree)
}

func timer(timer func()) {
	ticker := time.NewTicker(conf.GetRefreshTime())
	for {
		select {
		case <-ticker.C:
			timer()
		}
	}
}
