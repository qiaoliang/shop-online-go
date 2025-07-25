# Test Strategy and Standards

## Testing Philosophy

*   **Approach:** 优先采用测试驱动开发 (TDD) 或测试优先的方法，确保代码质量和功能正确性。
*   **Coverage Goals:** 目标是实现高代码覆盖率，特别是对于核心业务逻辑和关键路径。
*   **Test Pyramid:** 强调单元测试、集成测试和端到端测试的合理分层，以提高测试效率和可靠性。

## Test Types and Organization

#### Unit Tests

*   **Framework:** Go 的内置 `testing` 包。
*   **File Convention:** `_test.go` 文件，与被测试的源文件在同一包下。
*   **Location:** 与源文件相同的目录。
*   **Mocking Library:** gomock
*   **Coverage Requirement:** 核心业务逻辑单元测试覆盖率目标 80% 以上。

**AI Agent Requirements:**
*   为所有公共方法生成测试。
*   覆盖边缘情况和错误条件。
*   遵循 AAA 模式 (Arrange, Act, Assert)。
*   模拟所有外部依赖。

#### Integration Tests

*   **Scope:** 测试多个组件或服务之间的交互，例如 Service 层与 Repository 层、API 接口与数据库的集成。
*   **Location:** 独立的 `_test.go` 文件，或在 `testutils` 模块中提供辅助函数。
*   **Test Infrastructure:**
    *   **Database:** 使用真实的 MySQL 数据库实例（或测试专用的内存数据库如 SQLite），并在测试前进行数据清理和初始化。
    *   **API:** 使用 `httptest` 包模拟 HTTP 请求和响应，测试 API 端点。

#### E2E Tests

*   **Framework:** Postman
*   **Scope:** 模拟真实用户场景，测试整个系统端到端的流程。
*   **Environment:** 独立的测试环境，尽可能接近生产环境。
*   **Test Data:** 使用独立的测试数据，避免污染生产数据。

## Test Data Management

*   **Strategy:** 使用测试数据生成器或预定义的测试数据文件，确保测试数据的可重复性和隔离性。
*   **Fixtures:** 在测试设置阶段加载测试夹具，并在测试结束后清理。
*   **Factories:** (不适用，Go 语言中不常用)
*   **Cleanup:** 每个测试用例执行后，清理数据库或其他状态，确保测试的独立性。

## Continuous Testing

*   **CI Integration:** 将单元测试和集成测试集成到 CI/CD 流水线中，确保每次代码提交后自动运行测试。
