# Coding Standards

## 核心标准

-   **语言与运行环境**：Go 1.17 及以上
-   **代码风格与静态检查**：遵循 Go 官方的 `go fmt`、`go vet`，并结合 `golint` 或 `staticcheck` 等工具。
-   **测试组织**：测试文件与被测源文件同包，文件名以 `_test.go` 结尾。

---

## 项目结构规范

-   采用领域驱动设计(DDD)分层架构，分为：
    -   `router` 层：处理 CLI/HTTP 请求，路由转发
    -   `handler` 层：处理 CLI/HTTP 请求，参数校验，协调各组件和服务
    -   `service` 层：核心业务逻辑，协调各组件和服务
    -   `repository` 层：数据访问与持久化，
    -   `model` 层：定义数据结构和模型
    -   `framework` 层：基础设施与通用工具
-   **依赖方向**：严格 router → handler → service → repository → framework，禁止循环依赖。上层不能依赖下层实现细节，应通过接口依赖。

---

## 命名约定

-   **包名**：小写单词，无下划线/混合大小写，简短有描述性。如：`repository`、`handler`、`service`
-   **文件名**：小写字母，单词间下划线。如：`user_service.go`、`order_model.go`
-   **变量名**：驼峰命名。局部变量小驼峰（`userID`），全局变量大驼峰（`UserService`）。
-   **常量**：全大写，下划线分隔（`MAX_RETRY_COUNT`）。
-   **接口/结构体名**：大驼峰。接口名以 "er" 结尾（如 `Reader`、`Writer`），避免 "I" 前缀。

---

## 代码组织

-   **结构体字段顺序**：先导出字段，后非导出字段，相关字段分组。
-   **函数声明顺序**：先类型/常量，再变量，接着 init()，最后其他方法（按重要性/调用关系排序）。
-   **import 声明**：分为三组：标准库、第三方库、内部包，组间空行，推荐 goimports 自动排序。

```go
import (
    "context"
    "fmt"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/pkg/errors"

    "amap-aos-activity/framework/mlogger"
    "amap-aos-activity/model/base"
)
```

---

## 编码规范

-   **接口优先**：倾向于接受接口参数而非结构体，提高灵活性和可测试性。
-   **指针使用**：仅在需要修改接收者或避免复制大结构体时使用指针。
-   **切片/映射**：传递时注意它们为引用类型。
-   **Context**：处理请求时始终传递 `context.Context`，便于超时、取消、值传递。

---

## 错误处理规范

-   **统一错误类型**：用 `common.Error` 封装业务错误，保持错误码和信息一致。
-   **错误传播**：逻辑层用 `errors.Wrap` 保留堆栈信息，注意空指针和边界。
-   **日志记录**：只在错误源头记录日志，避免重复。用统一日志库（如 `mlogger`），日志应包含足够上下文。

```go
if err != nil {
    mlogger.XErrorf(ctx, "failed to query data: %v", err)
    return nil, common.NewError(common.ErrCodeDB, "数据查询失败")
}
```

---

## 并发处理规范

-   **超时控制**：所有长时间操作都应接受 context 参数，及时响应取消信号。
-   **并发安全**：用 mutex、atomic 或 channel 保护共享资源，避免 goroutine 泄漏，推荐用 errgroup 管理并发任务。
-   **资源控制**：用 worker pool 或信号量控制并发数量，推荐用项目框架组件处理异步任务。

---

## 性能优化规范

-   **内存分配**：预分配已知大小的切片/map，大结构体用指针传递，临时对象用 sync.Pool 复用。
-   **高效 IO**：用缓冲 IO，批量数据库操作，合理缓存策略（如 mcache、mredis）。
-   **高效 JSON**：用高性能 JSON 库（如 sonic），频繁结构体预定义字段标签。

---

## 测试规范

-   **单元测试**：所有关键业务逻辑需单元测试，使用 `testify/suite` 包组织测试套件，推荐表驱动测试。测试文件与被测文件同目录，文件名以 `_test.go` 结尾。
-   **模拟依赖**：用接口和依赖注入便于测试，不使用 mock 技术，直接使用真实依赖进行测试。
-   **基准测试**：为性能关键路径编写基准测试，结合性能分析工具。

---

## 项目标准组件使用指南

-   **日志**：统一用 `mlogger` 包，关键流程节点记录，避免过度。
-   **配置**：用 `mconfig` 包读取配置，避免硬编码。
-   **HTTP 客户端**：用 `mhttp` 包，设置合理超时和重试。
-   **缓存**：用 `mcache` 或 `mredis`，实现合适的失效策略。

---

## 代码示例

### Controller 层示例

```go
// controller/example/example.go
package example

import (
    "context"
    "github.com/gin-gonic/gin"
    "amap-aos-activity/framework/mlogger"
    "amap-aos-activity/logic/example"
)

// Handler 处理HTTP请求
func Handler(c *gin.Context) {
    ctx := c.Request.Context()
    // 参数解析与验证
    var req struct {
        UserID string `json:"userId" binding:"required"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        mlogger.XWarnf(ctx, "invalid request: %v", err)
        c.JSON(400, gin.H{"code": "INVALID_PARAM", "message": "参数不合法"})
        return
    }
    // 调用逻辑层
    resp, err := example.ProcessRequest(ctx, req)
    if err != nil {
        c.JSON(500, gin.H{"code": err.Code, "message": err.Message})
        return
    }
    c.JSON(200, resp)
}
```

### Logic 层示例

```go
// logic/example/processor.go
package example

import (
    "context"
    "github.com/pkg/errors"
    "amap-aos-activity/basic/common"
    "amap-aos-activity/framework/mlogger"
    "amap-aos-activity/model/example"
)

// ProcessRequest 处理业务逻辑
func ProcessRequest(ctx context.Context, req interface{}) (interface{}, *common.Error) {
    // 类型断言
    request, ok := req.(struct{ UserID string })
    if !ok {
        return nil, common.NewError(common.ErrCodeParam, "请求参数类型错误")
    }
    // 业务逻辑处理
    data, err := example.GetUserData(ctx, request.UserID)
    if err != nil {
        mlogger.XErrorf(ctx, "failed to get user data: %+v", errors.WithStack(err))
        return nil, common.NewError(common.ErrCodeService, "获取用户数据失败")
    }
    // 返回结果
    return map[string]interface{}{
        "userId": request.UserID,
        "data":   data,
    }, nil
}
```

### Model 层示例

```go
// model/example/user_model.go
package example

import (
    "context"
    "time"
    "github.com/pkg/errors"
    "amap-aos-activity/framework/mdb"
    "amap-aos-activity/framework/mcache"
)

// UserData 表示用户数据
type UserData struct {
    UserID    string    `json:"userId" db:"user_id"`
    Name      string    `json:"name" db:"name"`
    CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

// GetUserData 从数据库获取用户数据
func GetUserData(ctx context.Context, userID string) (*UserData, error) {
    // 尝试从缓存获取
    cacheKey := "user_data:" + userID
    var userData UserData
    cached, err := mcache.Get(ctx, cacheKey, &userData)
    if err == nil && cached {
        return &userData, nil
    }
    // 从数据库查询
    query := "SELECT user_id, name, created_at FROM user_table WHERE user_id = ?"
    err = mdb.GetDB().GetContext(ctx, &userData, query, userID)
    if err != nil {
        return nil, errors.Wrap(err, "query database failed")
    }
    // 更新缓存
    _ = mcache.Set(ctx, cacheKey, userData, 5*time.Minute)
    return &userData, nil
}
```

---

## 良好与不良实践对比

```go
// 良好实践：正确的错误处理和日志记录
package service

import (
    "context"
    "github.com/pkg/errors"
    "amap-aos-activity/basic/common"
    "amap-aos-activity/framework/mlogger"
    "amap-aos-activity/model/user"
)

func GetUserProfile(ctx context.Context, userID string) (*user.Profile, *common.Error) {
    // 参数验证
    if userID == "" {
        return nil, common.NewError(common.ErrCodeParam, "用户ID不能为空")
    }
    // 调用模型层
    profile, err := user.GetProfileByID(ctx, userID)
    if err != nil {
        // 包装错误并记录日志
        wrappedErr := errors.Wrap(err, "获取用户资料失败")
        mlogger.XErrorf(ctx, "%+v", wrappedErr) // 记录堆栈信息
        // 返回适当的业务错误
        if errors.Is(err, user.ErrUserNotFound) {
            return nil, common.NewError(common.ErrCodeNotFound, "用户不存在")
        }
        return nil, common.NewError(common.ErrCodeService, "获取用户资料失败")
    }
    return profile, nil
}
```

```go
// 不良实践：不规范的错误处理和日志记录
package service

import (
    "context"
    "fmt"
    "log"
    "amap-aos-activity/model/user"
)

func GetUserProfile(ctx context.Context, userID string) (*user.Profile, error) {
    // 缺少参数验证
    // 直接调用模型层
    profile, err := user.GetProfileByID(ctx, userID)
    if err != nil {
        // 错误：使用fmt直接打印错误
        fmt.Println("Error getting user profile:", err)
        // 错误：使用标准log包而非项目日志框架
        log.Printf("Failed to get profile for user %s: %v", userID, err)
        // 错误：直接返回底层错误，没有包装或分类
        return nil, err
    }
    return profile, nil
}
```
