# Next Steps

## UX Expert Prompt

## Architect Prompt

### Story 11：收货地址数据模型与持久化基础 - 开发任务列表

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
7. 验证数据库迁移脚本对现有表无副作用，用户相关功能正常。
8. 切换数据库类型后，验证地址相关功能和测试均能正常运行。

#### 开发子任务

-   [x] 设计 addresses 表结构
    -   设计 SQL 表结构，包含所有指定字段及类型
    -   明确 user_id 外键与 users 表的关联关系
-   [x] 编写数据库迁移脚本
    -   编写 MySQL 版本的建表和回滚脚本
    -   编写 sqlite 版本的建表和回滚脚本
    -   脚本需通过本地测试验证
-   [x] 定义 Address 实体结构体
    -   在 `app/addresses/address.go` 中定义 Address 结构体
    -   字段与数据库表结构保持一致
-   [x] 定义 Repository 接口
    -   在 `app/addresses` 下定义 AddressRepository 接口
    -   包含 Create、Read、Update、Delete 等基本方法
-   [ ] 实现 Repository 的 MySQL 版本
    -   实现 MySQL 版 AddressRepository
    -   实现所有 CRUD 方法
    -   保证与数据库表结构一致
-   [x] 实现 Repository 的 sqlite 版本
    -   实现 sqlite 版 AddressRepository
    -   实现所有 CRUD 方法
    -   保证与数据库表结构一致
-   [ ] 实现数据库类型切换配置
    -   支持通过配置文件切换 MySQL 和 sqlite
    -   验证切换后功能正常
-   [x] 编写 Address 实体和 Repository 的单元测试
    -   覆盖 Address 结构体的基本操作
    -   覆盖 Repository 的所有 CRUD 方法
    -   测试需覆盖 MySQL 和 sqlite 两种实现
-   [ ] 集成验证
    -   运行数据库迁移脚本，确保对现有表无副作用
    -   切换数据库类型，验证地址相关功能和测试均能正常运行
