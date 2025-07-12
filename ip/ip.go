package ip

import (
    "errors"
    "net"
    "strings"
)

// GetIPv4 函数通过 操作系统 获取 IPv4 地址
func GetIPv4() (string, error) {
    iface, err := net.InterfaceByName("wlan0")
    if err != nil {
        return "", err
    }
    addrs, err := iface.Addrs()
    if err != nil {
        return "", err
    }
    for _, addr := range addrs {
        if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.To4() != nil {
            return ipnet.IP.String(), nil
        }
    }
    return "", errors.New("未找到 IPv4 地址")
}

// GetIPv6 函数通过 操作系统 获取 IPv6 地址
func GetIPv6() (string, error) {
    iface, err := net.InterfaceByName("wlan0")
    if err != nil {
        return "", err
    }
    addrs, err := iface.Addrs()
    if err != nil {
        return "", err
    }
    for _, addr := range addrs {
        if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.To16() != nil && ipnet.IP.To4() == nil {
            ip := ipnet.IP.String()
            // 过滤掉 link-local 地址（fe80::开头）
            if !strings.HasPrefix(ip, "fe80") {
                return ip, nil
            }
        }
    }
    return "", errors.New("未找到 global IPv6 地址")
}
