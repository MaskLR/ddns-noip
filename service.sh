#!/system/bin/sh

# 定义日志路径
LOG_DIR="/data/local/tmp/DDNS-no-ip"
DEBUG_LOG="$LOG_DIR/service_debug.log"
PID_FILE="$LOG_DIR/ipv6_updater.pid"

# 动态获取脚本所在的目录
MODDIR=${0%/*}

# 创建日志目录
mkdir -p "$LOG_DIR"

# 记录启动日志
echo "$(date '+%Y-%m-%d %H:%M:%S') service.sh 已启动" >> "$DEBUG_LOG"

# 配置文件路径
CONFIG_FILE="$MODDIR/system/ipv6_updater.sh"

# 检查是否已有运行中的实例
if [ -f "$PID_FILE" ] && kill -0 "$(cat "$PID_FILE")" 2>/dev/null; then
    echo "$(date '+%Y-%m-%d %H:%M:%S') WARN: ipv6_updater.sh 已在运行中，PID=$(cat "$PID_FILE")" >> "$DEBUG_LOG"
else
    # 启动 ipv6_updater.sh 并记录 PID
    if [ -f "$CONFIG_FILE" ]; then
        echo "$(date '+%Y-%m-%d %H:%M:%S') INFO: 找到配置文件 $CONFIG_FILE，正在加载..." >> "$DEBUG_LOG"
        sh "$CONFIG_FILE" >> "$LOG_DIR/config_run.log" 2>&1 &
        PID=$!
        echo $PID > "$PID_FILE"
        echo "$(date '+%Y-%m-%d %H:%M:%S') INFO: ipv6_updater.sh 脚本已加载并运行，PID=$PID" >> "$DEBUG_LOG"

        # 实时查看日志（可选）
        tail -f "$LOG_DIR/config_run.log" &
    else
        # 输出错误日志并退出
        echo "$(date '+%Y-%m-%d %H:%M:%S') ERROR: 配置文件 $CONFIG_FILE 未找到！" >> "$DEBUG_LOG"
        echo "$(date '+%Y-%m-%d %H:%M:%S') ERROR: 配置文件 $CONFIG_FILE 未找到！" >> "$DEBUG_LOG"
        exit 1
    fi
fi

# 系统服务启动完成后执行，可用于启动常驻脚本或后台任务
SYSTEM_SCRIPT="/system/bin/my_script.sh"

if [ -f "$SYSTEM_SCRIPT" ]; then
    echo "$(date '+%Y-%m-%d %H:%M:%S') INFO: 正在运行 $SYSTEM_SCRIPT..." >> "$DEBUG_LOG"
    "$SYSTEM_SCRIPT" >> "$LOG_DIR/run.log" 2>&1 &
    echo "$(date '+%Y-%m-%d %H:%M:%S') INFO: $SYSTEM_SCRIPT 已启动。" >> "$DEBUG_LOG"
else
    echo "$(date '+%Y-%m-%d %H:%M:%S') ERROR: 系统脚本 $SYSTEM_SCRIPT 未找到！" >> "$DEBUG_LOG"
    echo "$(date '+%Y-%m-%d %H:%M:%S') ERROR: 系统脚本 $SYSTEM_SCRIPT 未找到！" >> "$DEBUG_LOG"
    exit 1
fi

