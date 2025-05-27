#!/system/bin/sh

# 系统挂载后、开机早期阶段运行的脚本
# 可以在这里创建目录或做一些初始化工作

log_dir="/data/local/tmp/mymodule"
mkdir -p "$log_dir"
echo "post-fs-data 阶段执行成功" > "$log_dir/init.log"