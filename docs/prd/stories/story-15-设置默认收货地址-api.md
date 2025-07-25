# Story 1.5: 设置默认收货地址 API

作为用户，我希望能够通过 API 将我的某个收货地址设置为默认地址，以便在下单时自动选择。

## Acceptance Criteria

1.  存在 `POST /v1/user/shipping-address/set-default` API 接口，接收地址 ID。
2.  成功设置默认地址后，API 返回成功信息。
3.  该地址的 `is_default` 字段在数据库中设置为 `true`，同时该用户其他地址的 `is_default` 字段设置为 `false`。
4.  `GET /v1/user/shipping-address/default` API 返回新设置的默认地址。

## Integration Verification

1.  IV1: 现有用户认证机制正常工作。
2.  IV2: 现有订单创建流程能够正确获取并使用默认收货地址。

---
