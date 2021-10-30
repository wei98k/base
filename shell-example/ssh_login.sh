#!/bin/sh
#basepath=$(cd `dirname $0`; pwd)
export LC_CTYPE=en_US
#expect脚本所在位置
#filepath=$basepath/XXX.sh
exec /usr/local/bin/auto_login.sh 22 root 192.168.2.27 root123
