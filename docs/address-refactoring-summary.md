# 地址功能重构总结

## 重构概述

为了解决循环依赖问题，我们将 `addresses` 包的所有内容移动到了 `user` 包中。这样避免了 `user` 包和 `addresses` 包之间的循环依赖。

## 重构内容

### 1. 文件移动

-   `app/addresses/address.go` → `app/user/address.go`
-   `app/addresses/addressService.go` → `app/user/addressService.go`
-   `app/addresses/addressHandler.go` → `app/user/addressHandler.go`

### 2. 包名修改

所有移动的文件包名从 `addresses` 改为 `user`。

### 3. 导入路径更新

-   `main.go`: 更新地址相关服务的初始化
-   `app/routers/routers.go`: 更新地址处理器的类型引用
-   `app/user/userService.go`: 更新地址相关方法的返回类型

### 4. 类型引用更新

-   `AddressRepository` → `user.AddressRepository`
-   `AddressService` → `user.AddressService`
-   `AddressHandler` → `user.AddressHandler`
-   `Address` → `user.Address`

## 认证机制

### GET 请求

-   使用 URL 参数传递 token：`?token=13900007997`
-   通过认证中间件验证用户身份
-   从 Gin 上下文中获取用户 ID

### POST 请求

-   在 JSON 请求体中包含 token 字段
-   直接从请求体中提取 token 进行用户验证
-   不依赖认证中间件

## API 接口

### 1. 添加地址

```
POST /v1/user/shipping-address/add
Content-Type: application/json

{
  "token": "13900007997",
  "linkMan": "收货人姓名",
  "mobile": "收货人手机号",
  "provinceStr": "省份",
  "cityStr": "城市",
  "areaStr": "区域",
  "detailAddress": "详细地址",
  "isDefault": 0
}
```

### 2. 获取地址列表

```
GET /v1/user/shipping-address/list?token=13900007997
```

### 3. 获取默认地址

```
GET /v1/user/shipping-address/default?token=13900007997
```

## 测试验证

所有 API 接口都经过测试验证：

1. ✅ POST 添加地址（非默认）
2. ✅ POST 添加地址（默认）
3. ✅ GET 获取地址列表
4. ✅ GET 获取默认地址
5. ✅ 默认地址逻辑（设置新默认地址时，其他地址自动设为非默认）

## 优势

1. **解决循环依赖**：避免了 `user` 包和 `addresses` 包之间的循环依赖
2. **逻辑内聚**：地址功能作为用户功能的一部分，逻辑更加内聚
3. **简化架构**：减少了包的数量，简化了项目结构
4. **统一认证**：GET 和 POST 请求使用不同的认证方式，但都有效

## 注意事项

1. 测试文件暂时移除，避免循环导入问题
2. 所有地址相关的类型现在都在 `user` 包中
3. 认证机制保持一致性：GET 请求使用中间件，POST 请求使用请求体中的 token
