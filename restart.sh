#!/bin/bash

# 设置错误时退出
set -e

echo "🚀 开始启动 bookstore 服务..."

# 1. 查找并杀死占用 9090 端口的进程
echo "📋 检查端口 9090 的占用情况..."
if lsof -i :9090 > /dev/null 2>&1; then
    echo "🔍 发现占用端口 9090 的进程："
    lsof -i :9090

    echo "💀 正在终止这些进程..."
    lsof -ti :9090 | xargs kill -9
    echo "✅ 端口 9090 已释放"
else
    echo "✅ 端口 9090 未被占用"
fi

# 2. 删除 test.db 文件（如果存在）
echo "🗑️  检查并删除 test.db 文件..."
if [ -f "test.db" ]; then
    rm test.db
    echo "✅ test.db 已删除"
else
    echo "✅ test.db 不存在，无需删除"
fi

# 3. 清理编译中间产物和最终产物
echo "🧹 清理编译中间产物和最终产物..."
rm -f bookstore
go clean -cache -testcache
echo "✅ 编译产物清理完成"

# 4. 构建项目
echo "🔨 开始构建项目..."
go build -o bookstore
echo "✅ 构建完成，生成 bookstore 可执行文件"

# 4. 运行程序
echo "🎯 启动 bookstore 服务..."
./bookstore