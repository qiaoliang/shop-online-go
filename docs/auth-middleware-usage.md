# 认证中间件使用指南

## 概述

本项目引入了一个新的认证中间件系统，用于统一处理 token 解析和用户认证，避免循环依赖问题。

## 架构设计

### 核心组件

1. **AuthMiddleware** (`app/security/auth.go`)

    - 主要的认证中间件
    - 支持可选认证和强制认证两种模式
    - 自动从多种来源提取 token

2. **TokenExtractor** 接口 (`app/security/token_extractor.go`)

    - 定义 token 解析接口
    - 支持多种实现方式
    - 避免直接依赖 UserService

3. **Utils 包** (`app/utils/context.go`)
    - 提供从 gin 上下文获取用户信息的工具函数
    - 避免循环依赖

### 依赖关系

```
main.go
├── security.AuthMiddleware
│   └── security.TokenExtractor (接口)
│       └── security.UserTokenExtractor (实现)
│           └── user.UserRepo (接口)
└── handlers
    └── utils (获取用户信息)
```

## 使用方法

### 1. 在 main.go 中初始化

```go
func main() {
    // 创建用户仓库
    userRepo := user.NewUserRepoDB(db)

    // 创建token提取器
    tokenExtractor := security.NewUserTokenExtractor(userRepo)

    // 创建认证中间件
    authMiddleware := security.NewAuthMiddleware(tokenExtractor)

    // 在SetupRouter中传入认证中间件
    routers.SetupRouter(r, ..., authMiddleware)
}
```

### 2. 在路由中使用

#### 可选认证（Authenticate）

```go
// 使用可选认证中间件
authenticated := v1.Group("")
authenticated.Use(authMiddleware.Authenticate())

// 这些接口可以带token也可以不带token
authenticated.GET("/user/detail", userHandler.GetUserDetail)
```

#### 强制认证（RequireAuth）

```go
// 使用强制认证中间件
protected := v1.Group("")
protected.Use(authMiddleware.RequireAuth())

// 这些接口必须提供有效的token
protected.POST("/user/shipping-address/add", addressHandler.AddAddress)
```

### 3. 在 Handler 中获取用户信息

```go
func (h *AddressHandler) AddAddress(c *gin.Context) {
    // 从认证中间件获取用户ID
    userID := utils.GetUserIDFromContext(c)
    if userID == "" {
        c.JSON(http.StatusUnauthorized, gin.H{
            "code": 401,
            "msg":  "User not authenticated",
        })
        return
    }

    // 使用userID进行业务逻辑处理
    err := h.addressService.AddAddress(userID, req)
    // ...
}
```

## Token 提取策略

中间件支持从以下位置提取 token：

1. **Authorization Header**: `Bearer <token>`
2. **Query 参数**: `?token=<token>`
3. **Form 参数**: `token=<token>`
4. **JSON Body**: `{"token": "<token>"}`

## 当前实现

### Token 格式

-   当前使用手机号作为 token（临时方案）
-   未来可以升级为 JWT 或其他 token 机制

### 用户信息注入

中间件会在 gin 上下文中注入以下信息：

-   `userID`: 用户 ID
-   `mobile`: 手机号
-   `authContext`: 完整的认证上下文

## 优势

1. **避免循环依赖**: 通过接口和依赖注入解决循环依赖问题
2. **统一认证逻辑**: 所有 handler 使用相同的认证机制
3. **灵活配置**: 支持可选认证和强制认证
4. **易于测试**: 提供 mock 实现用于测试
5. **易于扩展**: 可以轻松添加新的 token 提取器实现

## 迁移指南

### 从旧的手动 token 解析迁移

**旧方式**:

```go
func (h *CartHandler) GetShopingCart(c *gin.Context) {
    token, err := c.GetQuery("token")
    if !err {
        // 处理错误
    }
    cart := h.service.GetCartByToken(token)
    // ...
}
```

**新方式**:

```go
func (h *CartHandler) GetShopingCart(c *gin.Context) {
    mobile := utils.GetMobileFromContext(c)
    if mobile == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "User not authenticated"})
        return
    }
    cart := h.service.GetCartByToken(mobile)
    // ...
}
```

## 测试

运行认证中间件测试：

```bash
go test ./app/security -v
```

## 未来改进

1. **JWT 支持**: 实现 JWT token 解析
2. **Redis 缓存**: 添加 token 缓存机制
3. **权限控制**: 集成 RBAC 权限系统
4. **Token 刷新**: 支持 token 自动刷新机制
