#!/system/bin/sh
# Magisk 模块卸载脚本

# 删除模块创建的文件或目录
rm -rf /data/local/tmp/DDNS-NO-IP
rm -rf /data/local/tmp/mymodule

# 输出日志（可选）
echo "模块已卸载"