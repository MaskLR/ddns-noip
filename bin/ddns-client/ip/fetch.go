package ip

import (
    "io/ioutil" // 导入文件读取模块
    "net/http"  // 导入 HTTP 请求模块
    "strings"   // 导入字符串操作模块
)

// getIP 函数通过指定的 URL 获取 IP 地址
func getIP(url string) (string, error) {
    // 发送 HTTP 请求获取 IP 地址
    resp, err := http.Get(url)
    if err != nil {
        return "", err // 如果请求失败，返回错误信息
    }
    defer resp.Body.Close() // 确保函数退出时关闭响应体

    // 读取响应体内容
    b, _ := ioutil.ReadAll(resp.Body)
    // 返回去除多余空格的 IP 地址
    return strings.TrimSpace(string(b)), nil
}
