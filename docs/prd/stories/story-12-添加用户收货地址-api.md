# Story 1.2: 添加用户收货地址 API

作为用户，我希望能够通过 API 添加新的收货地址，以便我可以保存我的常用收货地点。

## Acceptance Criteria

1.  存在 `POST /v1/user/shipping-address/add` API 接口，接收 `linkMan`, `mobile`, `provinceStr`, `cityStr`, `areaStr`, `detailAddress`, `isDefault` 等参数。
2.  成功添加地址后，API 返回成功信息。
3.  地址数据正确持久化到数据库。
4.  `isDefault` 字段能够正确处理，如果设置为 `true`，则将该用户其他地址的 `is_default` 设置为 `false`。

## Integration Verification

1.  IV1: 现有用户认证机制正常工作，能够通过 Token 识别用户。
2.  IV2: 现有用户详情查询 API 正常工作。
