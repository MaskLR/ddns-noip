#!/system/bin/sh
MODDIR=${0%/*}
exec "$MODDIR/bin/ddns-client" >> "$MODDIR/ddns.log" 2>&1 &