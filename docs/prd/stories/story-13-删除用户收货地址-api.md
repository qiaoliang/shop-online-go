# Story 1.3: 删除用户收货地址 API

作为用户，我希望能够通过 API 删除不再需要的收货地址，以便我能灵活地管理我的收货地址。

## Acceptance Criteria

1.  存在 `DELETE /v1/user/shipping-address/delete` API 接口，接收地址 ID。
2.  成功删除地址后，API 返回成功。
3.  删除地址后，该地址不再出现在用户地址列表中。

## Integration Verification

1.  IV1: 现有用户认证机制正常工作。
2.  IV2: 现有订单创建流程中选择地址的功能（如果存在）能够正确处理地址的删除。

## 开发状态分析

### ✅ 已实现的部分

1. **数据层（Repository）**
    - `app/user/address.go` 中已实现 `Delete(id int) error` 方法
    - 数据模型 `Address` 结构体已定义完整

### ❌ 缺失的功能

1. **服务层（Service）**

    - `AddressService` 接口中缺少 `DeleteAddress` 方法定义
    - `addressService` 实现中缺少删除地址的业务逻辑

2. **控制器层（Handler）**

    - `AddressHandler` 中缺少 `DeleteAddress` 处理方法
    - 缺少删除地址的请求参数验证和错误处理

3. **路由配置**
    - 路由文件中缺少 `DELETE /v1/user/shipping-address/delete` 端点配置

### 📋 待实现任务

-   [] 1. 在 `AddressService` 接口中添加 `DeleteAddress(userID string, addressID int) error` 方法
    -   [ ] 在 `AddressService` 接口中添加 `DeleteAddress` 方法定义
    -   [ ] 在 `addressService` 中实现删除地址的业务逻辑，包括：
        -   验证地址是否属于当前用户
        -   检查是否为默认地址（如果是，需要处理默认地址逻辑）
            -   删除这个地址后，将该用户的第一个地址做为默认地址。
            -   如果这是最后一个地址，则不做任何处理。
        -   执行删除操作
    -   [ ] 创建完整的测试套件 `addressService_test.go`，包含所有边界情况测试
-   [ ] 2. 在 `AddressHandler` 中添加 `DeleteAddress` 处理方法，包括：
    -   请求参数验证
    -   用户认证检查
    -   错误处理和响应
-   [ ] 3. 在路由配置中添加删除地址的 API 端点
-   [ ] 4. 依据代码具体实现，更新 `./docs/architecture/rest-api-spec.md` 文档

### 🔍 技术要点

-   删除前需要验证地址所有权（确保用户只能删除自己的地址）
-   如果删除的是默认地址，需要考虑是否自动设置其他地址为默认
-   需要处理删除失败的情况（如地址不存在、数据库错误等）
-   API 应返回标准的 JSON 响应格式
