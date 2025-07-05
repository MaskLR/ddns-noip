#!/system/bin/sh
SKIPMOUNT=false
PROPFILE=false
POSTFSDATA=false
LATESTARTSERVICE=true

print_modname() {
  ui_print "*************************************"
  ui_print "    MaskLR DDNS (Go版) 安装中...    "
  ui_print "*************************************"
}

on_install() {
  ui_print "- 拷贝模块文件..."
  cp -af "$MODPATH/bin" /data/adb/modules_update/$MODID/

  if [ ! -f /data/adb/modules_update/$MODID/config.json ]; then
    ui_print "- 初始化配置文件..."
    cp -f "$MODPATH/config.json" /data/adb/modules_update/$MODID/
  fi
}

set_permissions() {
  set_perm_recursive "$MODPATH" 0 0 0755 0644
  set_perm "$MODPATH/bin/ddns-client" 0 0 0755
  set_perm "$MODPATH/service.sh" 0 0 0755
}