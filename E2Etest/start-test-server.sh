#!/bin/bash

# 切换到项目根目录
cd ..

# 确保测试配置文件存在
echo "正在检查测试配置文件..."
if [ ! -f "config-test.yaml" ]; then
  echo "测试配置文件不存在，复制config.yaml为config-test.yaml..."
  cp config.yaml config-test.yaml
fi

# 设置环境变量，使用测试配置文件
export CONFIG_FILE=config-test.yaml

# 确保测试数据库存在
echo "正在准备测试数据库..."
touch test.db

# 确保迁移目录正确设置
echo "正在检查配置文件中的迁移目录设置..."
if ! grep -q "MIGRATION_DIR" config-test.yaml; then
  echo "在config-test.yaml中添加MIGRATION_DIR设置..."
  echo "MIGRATION_DIR: $(pwd)/dbscripts" >> config-test.yaml
fi

# 先运行一个简单的Go程序来执行数据库迁移
echo "正在执行数据库迁移..."
cat > migrate.go << EOF
package main

import (
	"bookstore/app/configs"
	_ "bookstore/app/migrations"
)

func main() {
	configs.GetConfigInstance("config-test.yaml")
	configs.Cfg.DBConnection()
	println("数据库迁移完成")
}
EOF

go run migrate.go

# 不需要导入测试用户，使用系统自带的admin用户
echo "使用系统自带的admin用户进行测试..."

# 检查端口9090是否被占用
echo "正在检查端口9090是否被占用..."
if command -v lsof >/dev/null 2>&1; then
  PORT_PID=$(lsof -t -i:9090 2>/dev/null)
  if [ ! -z "$PORT_PID" ]; then
    echo "端口9090被进程 $PORT_PID 占用，正在终止该进程..."
    kill -9 $PORT_PID
    sleep 1
  else
    echo "端口9090未被占用"
  fi
else
  echo "未找到lsof命令，无法检查端口占用情况"
fi

# 启动服务器
echo "正在使用测试配置启动服务器..."
go run main.go
