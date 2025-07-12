package config

import (
<<<<<<< HEAD
	"encoding/json" // 导入 JSON 编解码包
	"fmt"           // 导入格式化输出包，用于错误处理
	"os"            // 导入操作系统包，用于文件操作
=======
    "encoding/json" // 导入 JSON 编解码包
    "os"            // 导入操作系统包，用于文件操作
    "fmt"           // 导入格式化输出包，用于错误处理
>>>>>>> a7585d8 (2025/7/12)
)

// 配置结构体，包含了需要的配置信息
type Config struct {
<<<<<<< HEAD
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Hostname string `json:"hostname"` // 主机名
	Interval int    `json:"interval"` // 更新间隔（秒）
=======
    Username string `json:"username"` // 用户名
    Password string `json:"password"` // 密码
    Hostname string `json:"hostname"` // 主机名
    Interval int    `json:"interval"` // 更新间隔（秒）
>>>>>>> a7585d8 (2025/7/12)
}

// 确保配置文件存在，如果不存在则创建一个默认配置文件
func Ensure(path string) (*Config, error) {
<<<<<<< HEAD
	// 检查文件是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 如果配置文件不存在，创建一个默认配置
		def := Config{
			Username: "uname",        // 默认用户名
			Password: "password",     // 默认密码
			Hostname: "you.ddns.net", // 默认主机名
			Interval: 999,            // 默认更新间隔：300秒
		}
		// 将默认配置转换为 JSON 格式
		data, _ := json.MarshalIndent(def, "", "  ")
		// 将默认配置写入文件
		err := os.WriteFile(path, data, 0600)
		if err != nil {
			return nil, fmt.Errorf("无法创建配置文件: %v", err)
		}
	}
	// 加载配置文件并返回
	return Load(path)
=======
    // 检查文件是否存在
    if _, err := os.Stat(path); os.IsNotExist(err) {
        // 如果配置文件不存在，创建一个默认配置
        def := Config{
            Username: "gyjd4yv",    // 默认用户名
            Password: "ihdMsPazryc1",    // 默认密码
            Hostname: "all.ddnskey.com", // 默认主机名
            Interval: 999,             // 默认更新间隔：300秒
        }
        // 将默认配置转换为 JSON 格式
        data, _ := json.MarshalIndent(def, "", "  ")
        // 将默认配置写入文件
        err := os.WriteFile(path, data, 0600)
         if err != nil {
            return nil, fmt.Errorf("无法创建配置文件: %v", err)
        }
    }
    // 加载配置文件并返回
    return Load(path)
>>>>>>> a7585d8 (2025/7/12)
}

// 加载配置文件
func Load(path string) (*Config, error) {
<<<<<<< HEAD
	// 读取配置文件内容
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err // 如果读取文件出错，返回错误
	}
	// 解析 JSON 数据
	var cfg Config
	err = json.Unmarshal(b, &cfg)
	return &cfg, err // 返回解析后的配置
=======
    // 读取配置文件内容
    b, err := os.ReadFile(path)
    if err != nil {
        return nil, err // 如果读取文件出错，返回错误
    }
    // 解析 JSON 数据
    var cfg Config
    err = json.Unmarshal(b, &cfg)
    return &cfg, err // 返回解析后的配置
>>>>>>> a7585d8 (2025/7/12)
}
