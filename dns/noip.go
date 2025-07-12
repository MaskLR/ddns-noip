package dns

import (
    "fmt"        // 导入格式化字符串的包
    "net/http"   // 导入 HTTP 请求模块
    "net/url"    // 导入 URL 处理模块
    "strings"    // 导入字符串操作模块

    "masklr-ddns/config" // 导入配置包
)

// 设置默认的 User-Agent，用于 HTTP 请求，按照官方建议标注客户端名称+联系方式
var userAgent = "MaskLR-DDNS/1.0 lingran2023@gmail.com"

// NoIPProvider 结构体表示 No-IP 提供商
type NoIPProvider struct{}

// Update 方法用于更新 No-IP 提供商的 DNS 记录
func (p *NoIPProvider) Update(cfg *config.Config, ip string) error {
    client := &http.Client{} // 创建 HTTP 客户端

    // 准备请求参数
    params := url.Values{}
    params.Set("hostname", cfg.Hostname) // 设置主机名
    params.Set("myip", ip)               // 设置新的 IP 地址

    // 构造请求 URL
    req, err := http.NewRequest("GET", "https://dynupdate.no-ip.com/nic/update?"+params.Encode(), nil)
    if err != nil {
        return fmt.Errorf("构造请求失败: %w", err) // 如果请求构造失败，返回错误信息
    }
    // 输出构造的请求 URL
    fmt.Printf("[DEBUG] No-IP 请求URL: %s\n", req.URL.String())

    // 设置请求的基本认证和 User-Agent
    req.SetBasicAuth(cfg.Username, cfg.Password)
    req.Header.Set("User-Agent", userAgent)

    // 发送请求
    resp, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("HTTP 请求失败: %w", err) // 如果请求失败，返回错误信息
    }
    defer resp.Body.Close() // 确保函数退出时关闭响应体

    // 读取响应体
    var buf [512]byte
    n, _ := resp.Body.Read(buf[:])
    response := strings.TrimSpace(string(buf[:n])) // 去除响应中的多余空格

    // 检查响应的 HTTP 状态码
    if resp.StatusCode != 200 {
        return fmt.Errorf("No-IP 状态码: %d，响应: %s", resp.StatusCode, response) // 如果状态码不是 200，返回错误
    }

    // 根据响应内容判断更新是否成功
    if strings.HasPrefix(response, "good") || strings.HasPrefix(response, "nochg") {
        return nil // 如果响应表示成功或无变化，返回 nil
    }

    // 如果响应内容异常，返回错误信息
    return fmt.Errorf("No-IP 响应异常: %s", response)
}
