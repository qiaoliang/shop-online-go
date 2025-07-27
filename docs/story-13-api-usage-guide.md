# Story 13 API 使用指南

## 概述

Story 13 实现了用户收货地址的查询功能，包括获取地址列表和获取默认地址。所有 API 都需要用户认证。

## 认证方式

当前系统使用简化的 token 机制：

-   **Token 格式**: 用户的手机号码
-   **传递方式**: 通过 URL 查询参数 `?token=手机号`

例如：`?token=13900007997`

## API 接口

### 1. 添加收货地址

**接口**: `POST /v1/user/shipping-address/add`

**认证**: 需要 token 参数

**请求示例**:

```bash
curl -X POST "http://localhost:9090/v1/user/shipping-address/add?token=13900007997" \
  -H "Content-Type: application/json" \
  -d '{
    "linkMan": "张三",
    "mobile": "13800138000",
    "provinceStr": "广东省",
    "cityStr": "深圳市",
    "areaStr": "南山区",
    "detailAddress": "科技园路1号",
    "isDefault": 1
  }'
```

**响应示例**:

```json
{
    "code": "200",
    "msg": "Address added successfully"
}
```

### 2. 获取地址列表

**接口**: `GET /v1/user/shipping-address/list`

**认证**: 需要 token 参数

**请求示例**:

```bash
curl -X GET "http://localhost:9090/v1/user/shipping-address/list?token=13900007997"
```

**响应示例**:

```json
{
    "code": "200",
    "data": [
        {
            "id": 1,
            "userId": "admin",
            "linkMan": "张三",
            "mobile": "13800138000",
            "provinceStr": "广东省",
            "cityStr": "深圳市",
            "areaStr": "南山区",
            "detailAddress": "科技园路1号",
            "isDefault": 1
        }
    ],
    "msg": "Address list retrieved successfully"
}
```

### 3. 获取默认地址

**接口**: `GET /v1/user/shipping-address/default`

**认证**: 需要 token 参数

**请求示例**:

```bash
curl -X GET "http://localhost:9090/v1/user/shipping-address/default?token=13900007997"
```

**响应示例**:

```json
{
    "code": "200",
    "data": {
        "id": 1,
        "userId": "admin",
        "linkMan": "张三",
        "mobile": "13800138000",
        "provinceStr": "广东省",
        "cityStr": "深圳市",
        "areaStr": "南山区",
        "detailAddress": "科技园路1号",
        "isDefault": 1
    },
    "msg": "Default address retrieved successfully"
}
```

**无默认地址时的响应**:

```json
{
    "code": "200",
    "data": null,
    "msg": "No default address found"
}
```

## 错误处理

### 认证错误 (401)

```json
{
    "code": 401,
    "msg": "无效的认证token"
}
```

**原因**:

-   没有提供 token 参数
-   提供的 token（手机号）对应的用户不存在

### 参数错误 (400)

```json
{
    "code": "400",
    "msg": "Invalid request parameters"
}
```

**原因**:

-   请求体格式不正确
-   缺少必需的字段

### 服务器错误 (500)

```json
{
    "code": "500",
    "msg": "Failed to add address"
}
```

## 测试用户

当前系统中可用的测试用户：

-   **手机号**: `13900007997` (admin 用户)

## 注意事项

1. **Token 机制**: 当前使用手机号作为 token，未来会升级为更安全的 JWT 机制
2. **认证方式**: 所有接口都通过 URL 查询参数传递 token，使用统一的认证中间件
3. **默认地址**: 每个用户只能有一个默认地址，设置新默认地址时会自动取消其他地址的默认状态
4. **数据验证**: 所有地址字段都是必填的
5. **用户隔离**: 用户只能访问自己的地址数据
6. **请求体字段**:
    - `linkMan`: 收件人姓名
    - `mobile`: 收件人手机号（不是当前用户的手机号）
    - `provinceStr`, `cityStr`, `areaStr`: 省市区
    - `detailAddress`: 详细地址
    - `isDefault`: 是否为默认地址（1=是，0=否）

## 完整测试流程

1. **添加地址**:

```bash
curl -X POST "http://localhost:9090/v1/user/shipping-address/add?token=13900007997" \
  -H "Content-Type: application/json" \
  -d '{"linkMan":"张三","mobile":"13800138000","provinceStr":"广东省","cityStr":"深圳市","areaStr":"南山区","detailAddress":"科技园路1号","isDefault":1}'
```

2. **获取地址列表**:

```bash
curl -X GET "http://localhost:9090/v1/user/shipping-address/list?token=13900007997"
```

3. **获取默认地址**:

```bash
curl -X GET "http://localhost:9090/v1/user/shipping-address/default?token=13900007997"
```

## 架构说明

-   **认证中间件**: 所有需要认证的接口都通过统一的认证中间件处理
-   **用户识别**: 通过 token（手机号）在 userService 中查找用户
-   **数据隔离**: 每个用户只能访问自己的地址数据
-   **错误处理**: 统一的错误响应格式和状态码
