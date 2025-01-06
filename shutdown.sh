#!/bin/bash

# 查找进程ID
PID=$(ps aux | grep 'go run main.go --env=prod' | grep -v grep | awk '{print $2}')

# 检查是否找到了进程
if [ -z "$PID" ]; then
    echo "没有找到运行中的进程."
    PID=$(ps aux | grep 'env=prod' | grep -v grep | awk '{print $2}')
fi

# 再次检查是否找到了进程
if [ -z "$PID" ]; then
    echo "没有找到运行中的进程."
else
    echo "找到进程 ID: $PID，正在关闭..."
    kill "$PID"

    # 可选：如果进程没有响应，可以强制关闭
    sleep 2  # 等待2秒
    if ps -p "$PID" > /dev/null; then
        echo "进程未响应，正在强制关闭..."
        kill -9 "$PID"
    else
        echo "进程已成功关闭."
    fi
fi