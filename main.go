package main

import (
<<<<<<< HEAD
	"log"                // 用于记录日志
	"masklr-ddns/config" // 导入自定义配置包
	"masklr-ddns/dns"    // 导入自定义的 DNS 更新包
	"masklr-ddns/ip"     // 导入自定义的获取 IP 地址包
	"masklr-ddns/web"
	"os"   // 导入自定义的 Web 服务器包
	"time" // 用于处理时间
=======
	"log"       // 用于记录日志
	"time"      // 用于处理时间
	"masklr-ddns/config"  // 导入自定义配置包
	"masklr-ddns/ip"      // 导入自定义的获取 IP 地址包
	"masklr-ddns/dns"     // 导入自定义的 DNS 更新包
    "masklr-ddns/web"
    "os"     // 导入自定义的 Web 服务器包
>>>>>>> a7585d8 (2025/7/12)
)

func main() {

<<<<<<< HEAD
	cwd, _ := os.Getwd()
	log.Printf("[DEBUG] 当前工作目录: %s", cwd)

	go web.StartConfigServer("0.0.0.0:6199")

	configPath := "../usr/etc/masklr/config.json"

	// 加载配置文件
	cfg, err := config.Ensure(configPath)
	if err != nil {
		log.Fatalf("配置加载失败: %v", err) // 如果加载配置失败，输出错误信息并退出
	}
	log.Printf("[DEBUG] 加载到的配置内容: %+v", cfg)

	// 创建一个新的 DNS 提供商实例（使用 No-IP）
	provider := dns.NewProvider("noip")

	// 初始化上次的 IPv4 和 IPv6 地址
	var lastIPv4, lastIPv6 string

	for {
		// 获取当前的 IPv6 地址
		ipv6, err := ip.GetIPv6()
		if err == nil && ipv6 != lastIPv6 { // 如果获取成功且 IPv6 地址变化
			// 更新 No-IP 上的 IPv6 地址
			err := provider.Update(cfg, ipv6)
			if err != nil {
				log.Printf("IPv6 更新失败: %v", err) // 如果更新失败，输出错误信息
			} else {
				log.Printf("IPv6 更新成功: %s", ipv6) // 更新成功，输出成功信息
				lastIPv6 = ipv6                   // 保存当前 IPv6 地址
			}
		} else if err == nil {
			log.Printf("IPv6 未变化: %s", ipv6) // 如果 IPv6 地址没有变化，输出未变化信息
		} else {
			log.Printf("IPv6 获取失败: %v", err)
		}

		// 获取当前的 IPv4 地址
		ipv4, err := ip.GetIPv4()
		if err == nil && ipv4 != lastIPv4 { // 如果获取成功且 IPv4 地址变化
			// 更新 No-IP 上的 IPv4 地址
			err := provider.Update(cfg, ipv4)
			if err != nil {
				log.Printf("IPv4 更新失败: %v", err) // 如果更新失败，输出错误信息
			} else {
				log.Printf("IPv4 更新成功: %s", ipv4) // 更新成功，输出成功信息
				lastIPv4 = ipv4                   // 保存当前 IPv4 地址
			}
		} else if err == nil {
			log.Printf("IPv4 未变化: %s", ipv4) // 如果 IPv4 地址没有变化，输出未变化信息
		}

		// 按照配置的间隔时间进行休眠
		time.Sleep(time.Duration(cfg.Interval) * time.Second)
	}
}
=======
    cwd, _ := os.Getwd()
    log.Printf("[DEBUG] 当前工作目录: %s", cwd)

    go web.StartConfigServer("0.0.0.0:6199")
    // 测试端口
  //  go web.StartConfigServer("0.0.0.0:9199")

    configPath := "/data/data/com.termux/files/usr/etc/masklr/config.json"

    // 加载配置文件
    cfg, err := config.Ensure(configPath)
    if err != nil {
        log.Fatalf("配置加载失败: %v", err) // 如果加载配置失败，输出错误信息并退出
    }
    log.Printf("[DEBUG] 加载到的配置内容: %+v", cfg)

    // 创建一个新的 DNS 提供商实例（使用 No-IP）
    provider := dns.NewProvider("noip")

    // 初始化上次的 IPv4 和 IPv6 地址
    var lastIPv4, lastIPv6 string

    for {
        // 获取当前的 IPv6 地址
        ipv6, err := ip.GetIPv6()
        if err == nil && ipv6 != lastIPv6 { // 如果获取成功且 IPv6 地址变化
            // 更新 No-IP 上的 IPv6 地址
            err := provider.Update(cfg, ipv6)
            if err != nil {
                log.Printf("IPv6 更新失败: %v", err) // 如果更新失败，输出错误信息
            } else {
                log.Printf("IPv6 更新成功: %s", ipv6) // 更新成功，输出成功信息
                lastIPv6 = ipv6 // 保存当前 IPv6 地址
            }
        } else if err == nil {
            log.Printf("IPv6 未变化: %s", ipv6) // 如果 IPv6 地址没有变化，输出未变化信息
        } else {
        log.Printf("IPv6 获取失败: %v", err)
        }

        // 获取当前的 IPv4 地址
        ipv4, err := ip.GetIPv4()
        if err == nil && ipv4 != lastIPv4 { // 如果获取成功且 IPv4 地址变化
            // 更新 No-IP 上的 IPv4 地址
            err := provider.Update(cfg, ipv4)
            if err != nil {
                log.Printf("IPv4 更新失败: %v", err) // 如果更新失败，输出错误信息
            } else {
                log.Printf("IPv4 更新成功: %s", ipv4) // 更新成功，输出成功信息
                lastIPv4 = ipv4 // 保存当前 IPv4 地址
            }
        } else if err == nil {
            log.Printf("IPv4 未变化: %s", ipv4) // 如果 IPv4 地址没有变化，输出未变化信息
        }
        
        // 按照配置的间隔时间进行休眠
        time.Sleep(time.Duration(cfg.Interval) * time.Second)
    }
}
>>>>>>> a7585d8 (2025/7/12)
