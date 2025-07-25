# Coding Standards

## Core Standards

*   **Languages & Runtimes:** Go 1.17
*   **Style & Linting:** 遵循 Go 官方的 `go fmt` 和 `go vet` 工具，以及 `golint` 或 `staticcheck` 等静态分析工具的建议。
*   **Test Organization:** 测试文件应与被测试的源文件放在同一包下，并以 `_test.go` 结尾。

## Naming Conventions

| Element | Convention | Example |
|---|---|---|
| **包名** | 小写，单个单词，简洁 | `addresses`, `user`, `goods` |
| **函数名** | 驼峰命名法，导出函数首字母大写 | `AddAddress`, `GetAddressList`, `saveAddress` |
| **变量名** | 驼峰命名法，简洁，局部变量短名称 | `address`, `user`, `addrData`, `i` |
| **常量** | 全大写，下划线分隔 | `DEFAULT_ADDRESS_LIMIT` |
| **接口名** | 通常以 `er` 结尾，或描述其行为 | `AddressRepository`, `UserService` |
| **结构体名** | 驼峰命名法，名词 | `Address`, `User` |

## Critical Rules

*   **错误处理：** 始终检查函数返回的错误，并进行适当处理或向上传播。避免忽略错误。
*   **并发安全：** 对于共享数据，必须使用互斥锁（`sync.Mutex`）或其他并发原语来确保并发安全。
*   **日志记录：** 使用统一的日志库进行日志记录，避免直接使用 `fmt.Println`。日志级别应正确使用。
*   **配置管理：** 所有配置项必须通过 `viper` 或其他配置管理库进行加载，避免硬编码。
*   **数据库操作：** 数据库操作必须通过 Repository 层进行封装，避免在 Service 或 Handler 层直接操作 GORM。
*   **API 响应：** 所有 API 响应必须遵循统一的 JSON 格式（例如，包含 `code`, `msg`, `data` 字段）。

## Language-Specific Guidelines

#### Go Specifics

*   **接口：** 倾向于接受接口而不是结构体，以提高灵活性和可测试性。
*   **指针：** 仅在需要修改接收者或避免复制大结构体时使用指针。
*   **切片和映射：** 在传递切片和映射时，要意识到它们是引用类型。
*   **Context：** 在处理请求时，始终传递 `context.Context` 以便进行超时、取消和值传递。
