#!/bin/bash

# 定义命令
COMMAND="go run main.go --env=prod"

# 检查是否已有运行中的进程
if pgrep -f "$COMMAND" > /dev/null; then
    echo "服务已经在运行中."
else
    echo "正在启动服务..."
    # 使用 nohup 在后台运行命令，并将输出重定向到 log 文件
    nohup $COMMAND > service.log 2>&1 &

    echo "服务已启动，PID: $!"
fi