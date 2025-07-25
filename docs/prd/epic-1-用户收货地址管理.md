# Epic 1: 用户收货地址管理

**Epic Goal**: 实现用户收货地址的全面管理功能，包括持久化存储和通过 API 进行增删改查，同时确保与现有系统的无缝集成和兼容性。

**Integration Requirements**:
*   新的收货地址数据模型应与现有用户模型关联。
*   API 接口应遵循现有项目的 RESTful 规范和认证机制。
*   数据库模式变更应向后兼容。
*   确保现有用户相关功能（如订单创建）能够正确引用和使用新的收货地址数据。

## Story 1.1: 收货地址数据模型与持久化基础

作为系统管理员，我希望能够创建和管理收货地址的数据库表结构，以便为用户收货地址功能提供持久化存储基础。

### Acceptance Criteria

1.  数据库中存在 `addresses` 表，包含 `id`, `user_id`, `link_man`, `mobile`, `province_str`, `city_str`, `area_str`, `detail_address`, `is_default` 等字段。
2.  `addresses` 表与 `users` 表通过 `user_id` 建立关联。
3.  `app/addresses` 模块中包含 `address.go` 定义的地址实体结构。
4.  `app/addresses` 模块中包含基本的 Repository 接口和实现，用于地址数据的持久化操作。

### Integration Verification

1.  IV1: 运行数据库迁移脚本后，现有数据库表结构保持不变。
2.  IV2: 现有用户登录和详情查询功能正常工作。

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
