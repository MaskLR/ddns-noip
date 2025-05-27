#!/system/bin/sh

# 定义一个函数，用于获取当前时间
get_time() {
    date '+%Y-%m-%d %H:%M:%S'
}

# 动态获取脚本所在的目录
MODDIR=${0%/*}

# 配置文件路径
CONFIG_FILE="$MODDIR/config.sh" 

# 加载配置文件
if [ -f "$CONFIG_FILE" ]; then
    echo "$(get_time) INFO: 找到配置文件 config.sh，正在加载..."
    source "$CONFIG_FILE"
    echo "$(get_time) INFO: 配置文件 config.sh 加载成功。"
else
    echo "$(get_time) ERROR: 配置文件 config.sh 未找到！"
    exit 1
fi

# 用于存储上一次成功更新的 IPv6 地址
LAST_IPV6=""

# 获取本地设备的 IPv6 地址
get_ipv6() {
    echo "$(get_time) INFO: 正在获取本地 IPv6 地址..."
    IPV6=$(curl -6 --silent --max-time 10 https://6.ipw.cn/)
    if [ -z "$IPV6" ]; then
        echo "$(get_time) ERROR: 未能获取到 IPv6 地址，请检查网络连接。"
    else
        echo "$(get_time) INFO: 获取到的 IPv6 地址为 $IPV6"
    fi
}

# 无限循环
while true; do
    # 获取 IPv6 地址
    get_ipv6

    # 如果未获取到 IPv6 地址，等待 1 分钟后重试
    if [ -z "$IPV6" ]; then
        echo "$(get_time) WARN: 未获取到 IPv6 地址，等待 1 分钟后重试..."
        sleep 60 # 等待 1 分钟
        continue
    fi

    # 如果当前 IPv6 地址与上一次成功更新的地址相同，则跳过更新请求
    if [ "$IPV6" = "$LAST_IPV6" ]; then
        echo "$(get_time) INFO: IPv6 地址未发生变化，无需更新。"
        sleep 666  # 等待11.1分钟后再次检查
        continue
    fi

    # 确保用户名、密码和主机名已设置
    if [ -z "$USER" ]; then
        echo "$(get_time) ERROR: 用户名未设置！"
        exit 1
    fi

    if [ -z "$PASS" ]; then
        echo "$(get_time) ERROR: 密码未设置！"
        exit 1
    fi

    if [ -z "$HOSTNAME" ]; then
        echo "$(get_time) ERROR: 主机名未设置！"
        exit 1
    fi

    # 输出调试信息
    echo "$(get_time) INFO: 正在更新 No-IP 动态域名：$HOSTNAME"
    echo "$(get_time) INFO: 使用 IPv6 地址：$IPV6"

    # 发送更新请求到 No-IP
    echo "$(get_time) INFO: 发送更新请求到 No-IP..."
    UPDATE_RESULT=$(curl -s -u "$USER:$PASS" "https://dynupdate.no-ip.com/nic/update?hostname=$HOSTNAME&myip=$IPV6")

    # 检查 curl 命令执行结果
    if echo "$UPDATE_RESULT" | grep -q "good"; then
        echo "$(get_time) INFO: 更新请求成功：$UPDATE_RESULT"
        LAST_IPV6="$IPV6"  # 更新成功后记录当前 IPv6 地址
    elif echo "$UPDATE_RESULT" | grep -q "nochg"; then
        echo "$(get_time) INFO: IPv6 地址未变化：$UPDATE_RESULT"
        LAST_IPV6="$IPV6"  # 即使返回 "nochg"，也更新 LAST_IPV6
    else
        echo "$(get_time) ERROR: 更新请求失败，No-IP 返回信息：$UPDATE_RESULT"
    fi

    # 输出完成信息
    echo "$(get_time) INFO: 更新请求已发送。请检查 No-IP 网站上的主机状态以确认更新是否成功。"

    # 每隔 16.65 分钟检查并更新一次（可以根据需要修改间隔时间）
    echo "$(get_time) INFO: 等待 16.65 分钟后再次检查..."
    sleep 999  # 每 16.65 分钟检查一次
done