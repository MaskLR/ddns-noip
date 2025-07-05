package ip

// GetIPv4 函数通过指定的 URL 获取 IPv4 地址
func GetIPv4() (string, error) {
    // 调用 getIP 函数来获取 IPv4 地址
    return getIP("https://ipv4.icanhazip.com")
}

// GetIPv6 函数通过指定的 URL 获取 IPv6 地址
func GetIPv6() (string, error) {
    // 调用 getIP 函数来获取 IPv6 地址
    return getIP("https://6.ipw.cn")
}
