# ddns-noip

## 动态更新No-IP服务商的IPv6地址

### 1. 添加账号、密码、域名到 `system>config.sh` 文件
先将压缩包解压找到system文件夹下config.sh文件打开
将的账号、密码和域名信息添加到 `config.sh` 文件内，确保文件内容如下所示：

```bash
# No-IP 服务商配置
USER="your_username"     # No-IP 账号
PASS="your_password"     # No-IP 密码
DOMAIN="your_domain"     # No-IP 域名

日志文件目录/data/local/tmp/DDNS-no-ip
