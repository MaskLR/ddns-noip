package web

import (
	"embed"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	//"path/filepath"
	"sync"

	"masklr-ddns/config"
)

//go:embed static/*
var staticFS embed.FS

var mu sync.Mutex

// configPath 可根据实际部署情况改为绝对路径
var configPath = "../usr/etc/masklr/config.json"

func StartConfigServer(addr string) {
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/api/config", handleConfig)

	log.Printf("[Web] 配置面板已启动: http://%s", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("[Web] 启动失败: %v", err)
	}
}

// serveIndex 提供静态页面
func serveIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	html, err := staticFS.ReadFile("static/index.html")
	if err != nil {
		http.Error(w, "无法加载前端界面", 500)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write(html)
}

// handleConfig 处理配置读写
func handleConfig(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	switch r.Method {
	case http.MethodGet:
		b, err := os.ReadFile(configPath)
		if err != nil {
			http.Error(w, "读取配置失败: "+err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)

	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "读取请求失败", 400)
			return
		}
		var cfg config.Config
		if err := json.Unmarshal(body, &cfg); err != nil {
			http.Error(w, "配置格式错误: "+err.Error(), 400)
			return
		}

		// 保存到磁盘
		b, _ := json.MarshalIndent(cfg, "", "  ")
		if err := os.WriteFile(configPath, b, 0600); err != nil {
			http.Error(w, "保存失败: "+err.Error(), 500)
			return
		}

		log.Printf("[Web] 配置已保存: %+v", cfg)
		w.Write([]byte("保存成功"))

	default:
		http.Error(w, "不支持的方法", 405)
	}
}
