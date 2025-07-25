# Error Handling Strategy

## General Approach

*   **Error Model:** Go 语言的错误处理遵循“显式错误处理”模式，通过返回 `error` 类型来表示错误。
*   **Exception Hierarchy:** Go 中没有传统的异常层次结构。错误通常通过自定义错误类型或包装标准库错误来表示。
*   **Error Propagation:** 错误通过函数返回值向上传播，直到被适当处理或记录。

## Logging Standards

*   **Library:** (未明确，不在本次修改范围内)
*   **Format:** (未明确，建议使用结构化日志，如 JSON 格式)
*   **Levels:** `DEBUG`, `INFO`, `WARN`, `ERROR`, `FATAL` (常见日志级别)
*   **Required Context:**
    *   **Correlation ID:** 建议在请求进入系统时生成一个唯一的 ID，并在整个请求生命周期中传递，以便于追踪。
    *   **Service Context:** 记录当前服务或模块的名称。
    *   **User Context:** 记录与请求相关的用户 ID (如果可用且符合隐私规范)。

## Error Handling Patterns

*   **External API Errors:** (不适用，本次增强不涉及外部 API 调用)
*   **Business Logic Errors:**
    *   **Custom Exceptions:** 定义业务相关的自定义错误类型，例如 `ErrAddressNotFound`, `ErrInvalidAddressData`。
    *   **User-Facing Errors:** 将内部错误转换为用户友好的错误消息，避免暴露敏感信息。
    *   **Error Codes:** (不在本次开发范围)
*   **Data Consistency:**
    *   **Transaction Strategy:** 对于涉及多个数据库操作的业务逻辑，使用数据库事务来确保数据一致性。
    *   **Compensation Logic:** (不适用，本次增强不涉及复杂分布式事务)
    *   **Idempotency:** 对于幂等操作（如添加收货地址），确保重复执行不会产生副作用。
