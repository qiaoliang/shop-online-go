#!/bin/bash

# 切换到项目根目录
cd ..

# 登录admin用户并获取认证令牌
echo "正在登录admin用户..."

# 发送登录请求
response=$(curl -s -X POST \
  http://localhost:9090/v1/user/m/login \
  -H 'Content-Type: application/json' \
  -d '{
    "mobile": "13900007997",
    "pwd": "1234"
  }')

# 从响应中提取令牌
token=$(echo $response | grep -o '"token":"[^"]*' | sed 's/"token":"//')

if [ -z "$token" ]; then
  echo "登录失败，无法获取令牌。响应内容："
  echo $response
  exit 1
else
  echo "登录成功！获取到令牌："
  echo $token
  echo ""
  echo "请在Postman中使用此令牌作为Bearer Token进行API测试。"
fi