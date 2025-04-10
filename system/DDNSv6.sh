#!/system/bin/sh

get_time() {
    date '+%Y-%m-%d %H:%M:%S'
}

MODDIR=${0%/*}

CONFIG_FILE="$MODDIR/config.sh" 

if [ -f "$CONFIG_FILE" ]; then
    echo "$(get_time) INFO: 找到配置文件 config.sh，正在加载..."
    source "$CONFIG_FILE"
    echo "$(get_time) INFO: 配置文件 config.sh 加载成功。"
else
    echo "$(get_time) ERROR: 配置文件 config.sh 未找到！"
    exit 1
fi

LAST_IPV6=""

get_ipv6() {
    echo "$(get_time) INFO: 正在获取本地 IPv6 地址..."
    IPV6=$(curl -6 --silent --max-time 10 https://6.ipw.cn/)
    if [ -z "$IPV6" ]; then
        echo "$(get_time) ERROR: 未能获取到 IPv6 地址，请检查网络连接。"
    else
        echo "$(get_time) INFO: 获取到的 IPv6 地址为 $IPV6"
    fi
}

while true; do
    get_ipv6
    if [ -z "$IPV6" ]; then
        echo "$(get_time) WARN: 未获取到 IPv6 地址，等待 1 分钟后重试..."
        sleep 60 # 等待 1 分钟
        continue
    fi

    if [ "$IPV6" = "$LAST_IPV6" ]; then
        echo "$(get_time) INFO: IPv6 地址未发生变化，无需更新。"
        sleep 666  # 等待11.1分钟后再次检查
        continue
    fi

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

    echo "$(get_time) INFO: 正在更新 No-IP 动态域名：$HOSTNAME"
    echo "$(get_time) INFO: 使用 IPv6 地址：$IPV6"

    echo "$(get_time) INFO: 发送更新请求到 No-IP..."
    UPDATE_RESULT=$(curl -s -u "$USER:$PASS" "https://dynupdate.no-ip.com/nic/update?hostname=$HOSTNAME&myip=$IPV6")

    if echo "$UPDATE_RESULT" | grep -q "good"; then
        echo "$(get_time) INFO: 更新请求成功：$UPDATE_RESULT"
        LAST_IPV6="$IPV6" 
    elif echo "$UPDATE_RESULT" | grep -q "nochg"; then
        echo "$(get_time) INFO: IPv6 地址未变化：$UPDATE_RESULT"
        LAST_IPV6="$IPV6" 
    else
        echo "$(get_time) ERROR: 更新请求失败，No-IP 返回信息：$UPDATE_RESULT"
    fi
    echo "$(get_time) INFO: 等待 16.65 分钟后再次检查..."
    sleep 999  # 每 16.65 分钟检查一次
done