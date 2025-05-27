#!/system/bin/sh

# 这是将会映射到 /system/bin 的自定义脚本
# 你可以在任何地方调用它，例如开机自动执行

echo "Hello from my Magisk module!" > /data/local/tmp/mymodule/hello.txt
