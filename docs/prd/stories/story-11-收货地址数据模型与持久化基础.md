# Story 1.1 收货地址数据模型与持久化基础

**标题**：收货地址数据模型与持久化基础

**背景与目标**：
为实现用户收货地址管理功能，需建立与用户表关联的收货地址数据模型，并实现持久化存储，兼容 MySQL 和内存数据库 sqlite，作为后续 API 功能的基础。

**业务价值**：
为用户提供可靠的收货地址存储能力，支持后续的地址增删改查等操作。

**技术约束**：

-   数据库需支持 MySQL 和 sqlite（内存数据库），并通过配置切换。
-   表结构需与现有用户表通过 user_id 字段建立外键关联。

**任务拆解**：

1. 设计 `addresses` 表结构，字段包括：
    - id（主键，自增）
    - user_id（外键，关联 users 表）
    - link_man（收件人姓名）
    - mobile（手机号）
    - province_str（省份）
    - city_str（城市）
    - area_str（区/县）
    - detail_address（详细地址）
    - is_default（是否为默认地址，布尔型）
2. 编写数据库迁移脚本，支持 MySQL 和 sqlite 两种数据库。
3. 在 `app/addresses/address.go` 中定义 Address 实体结构体，字段与表结构一致。
4. 在 `app/addresses` 下定义 Repository 接口及实现，支持地址的增删改查（CRUD）操作。
5. 实现 Repository 的 MySQL 版本和 sqlite 版本，确保通过配置可切换。
6. 编写单元测试，覆盖 Address 实体和 Repository 的基本操作。

**验收标准**：

-   [ ] 数据库中成功创建 `addresses` 表，字段和类型符合设计。
-   [ ] `addresses` 表与 `users` 表通过 `user_id` 正确关联。
-   [ ] `app/addresses/address.go` 中存在 Address 结构体，字段齐全。
-   [ ] `app/addresses` 下有 Repository 接口及 MySQL、sqlite 实现，支持基本 CRUD。
-   [ ] 通过配置可切换 MySQL 和 sqlite，单元测试均通过。

**集成校验点**：

-   [ ] 运行数据库迁移脚本后，现有数据库表结构保持不变，用户相关功能正常。
-   [ ] 切换数据库类型后，地址相关功能和测试均能正常运行。

**备注**：

-   字段命名、类型需与现有项目风格保持一致。
-   如有历史数据迁移需求，需单独评审。
