# Epic 1: 用户收货地址管理

**Epic Goal**: 实现用户收货地址的全面管理功能，包括持久化存储和通过 API 进行增删改查，同时确保与现有系统的无缝集成和兼容性。

**Integration Requirements**:

-   新的收货地址数据模型应与现有用户模型关联。
-   API 接口应遵循现有项目的 RESTful 规范和认证机制。
-   数据库模式变更应向后兼容。
-   确保现有用户相关功能（如订单创建）能够正确引用和使用新的收货地址数据。

## Story 1.1 收货地址数据模型与持久化基础

**标题**：收货地址数据模型与持久化基础

**背景与目标**：
为实现用户收货地址管理功能，需建立与用户表关联的收货地址数据模型，并实现持久化存储，使用 SQLite3 文件模式，作为后续 API 功能的基础。

**业务价值**：
为用户提供可靠的收货地址存储能力，支持后续的地址增删改查等操作。

**技术约束**：

-   数据库使用 SQLite3 文件模式。
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
2. 编写数据库迁移脚本，支持 SQLite3 数据库。
3. 在 `app/addresses/address.go` 中定义 Address 实体结构体，字段与表结构一致。
4. 在 `app/addresses` 下定义 Repository 接口及实现，支持地址的增删改查（CRUD）操作。
5. 实现 Repository 的 SQLite3 版本。
6. 编写单元测试，覆盖 Address 实体和 Repository 的基本操作。

**验收标准**：

-   [ ] 数据库中成功创建 `addresses` 表，字段和类型符合设计。
-   [ ] `addresses` 表与 `users` 表通过 `user_id` 正确关联。
-   [ ] `app/addresses/address.go` 中存在 Address 结构体，字段齐全。
-   [ ] `app/addresses` 下有 Repository 接口及 SQLite3 实现，支持基本 CRUD。
-   [ ] 单元测试均通过。

**集成校验点**：

-   [ ] 运行数据库迁移脚本后，现有数据库表结构保持不变，用户相关功能正常。
-   [ ] 地址相关功能和测试均能正常运行。

**备注**：

-   字段命名、类型需与现有项目风格保持一致。
-   如有历史数据迁移需求，需单独评审。

## Story 1.2: 添加用户收货地址 API

作为用户，我希望能够通过 API 添加新的收货地址，以便我可以保存我的常用收货地点。

### Acceptance Criteria

1.  存在 `POST /v1/user/shipping-address/add` API 接口，接收 `linkMan`, `mobile`, `provinceStr`, `cityStr`, `areaStr`, `detailAddress`, `isDefault` 等参数。
2.  成功添加地址后，API 返回成功信息。
3.  地址数据正确持久化到数据库。
4.  `isDefault` 字段能够正确处理，如果设置为 `true`，则将该用户其他地址的 `is_default` 设置为 `false`。

### Integration Verification

1.  IV1: 现有用户认证机制正常工作，能够通过 Token 识别用户。
2.  IV2: 现有用户详情查询 API 正常工作。

## Story 1.3: 获取用户收货地址列表与默认地址 API

作为用户，我希望能够通过 API 查看我的所有收货地址列表，并能获取我的默认收货地址，以便我能方便地选择或管理我的地址。

### Acceptance Criteria

1.  存在 `GET /v1/user/shipping-address/list` API 接口，返回当前用户的所有收货地址列表。
2.  存在 `GET /v1/user/shipping-address/default` API 接口，返回当前用户的默认收货地址。
3.  如果用户没有设置默认地址，`default` API 返回空或适当的提示。
4.  列表和默认地址的返回数据结构与 `Address` 实体一致。

### Integration Verification

1.  IV1: 现有用户认证机制正常工作。
2.  IV2: 现有商品列表和详情查询 API 正常工作。

## Story 1.4: 修改和删除用户收货地址 API

作为用户，我希望能够通过 API 修改我已保存的收货地址信息，并能删除不再需要的地址，以便我能灵活地管理我的收货地址。

### Acceptance Criteria

1.  存在 `POST /v1/user/shipping-address/modify` API 接口，接收地址 ID 和要修改的字段。
2.  存在 `POST /v1/user/shipping-address/delete` API 接口，接收地址 ID。
3.  成功修改或删除地址后，API 返回成功信息。
4.  修改后的地址信息正确反映在数据库中。
5.  删除地址后，该地址不再出现在用户地址列表中。

### Integration Verification

1.  IV1: 现有用户认证机制正常工作。
2.  IV2: 现有订单创建流程中选择地址的功能（如果存在）能够正确处理地址的修改和删除。

## Story 1.5: 设置默认收货地址 API

作为用户，我希望能够通过 API 将我的某个收货地址设置为默认地址，以便在下单时自动选择。

### Acceptance Criteria

1.  存在 `POST /v1/user/shipping-address/set-default` API 接口，接收地址 ID。
2.  成功设置默认地址后，API 返回成功信息。
3.  该地址的 `is_default` 字段在数据库中设置为 `true`，同时该用户其他地址的 `is_default` 字段设置为 `false`。
4.  `GET /v1/user/shipping-address/default` API 返回新设置的默认地址。

### Integration Verification

1.  IV1: 现有用户认证机制正常工作。
2.  IV2: 现有订单创建流程能够正确获取并使用默认收货地址。
