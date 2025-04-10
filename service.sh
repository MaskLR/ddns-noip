#!/system/bin/sh

LOG_DIR="/data/local/tmp/DDNS-no-ip"
SERVICE_LOG="$LOG_DIR/service.log"
PID_FILE="$LOG_DIR/DDNSv6.pid"
DDNSv6_LOG="$LOG_DIR/DDNSv6.log"
MODDIR=${0%/*}

mkdir -p "$LOG_DIR"

echo "$(date '+%Y-%m-%d %H:%M:%S') service.sh 已启动" >> "$SERVICE_LOG"

CONFIG_FILE="$MODDIR/system/DDNSv6.sh"

if [ -f "$PID_FILE" ] && kill -0 "$(cat "$PID_FILE")" 2>/dev/null; then
    echo "$(date '+%Y-%m-%d %H:%M:%S') WARN: DDNSv6.sh 已在运行中，PID=$(cat "$PID_FILE")" >> "$SERVICE_LOG"
else
    if [ -f "$CONFIG_FILE" ]; then
        echo "$(date '+%Y-%m-%d %H:%M:%S') INFO: 找到配置文件 $CONFIG_FILE，正在加载..." >> "$SERVICE_LOG"
        sh "$CONFIG_FILE" >> "$DDNSv6_LOG" 2>&1 &
        PID=$!
        echo $PID > "$PID_FILE"
        echo "$(date '+%Y-%m-%d %H:%M:%S') INFO: DDNSv6.sh 脚本已加载并运行，PID=$PID" >> "$SERVICE_LOG"
        tail -f "$DDNSv6_LOG" &
    else
        echo "$(date '+%Y-%m-%d %H:%M:%S') ERROR: 配置文件 $CONFIG_FILE 未找到！" >> "$SERVICE_LOG"
        exit 1
    fi
fi