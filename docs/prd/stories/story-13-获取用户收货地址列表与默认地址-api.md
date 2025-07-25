# Story 1.3: 获取用户收货地址列表与默认地址 API

作为用户，我希望能够通过 API 查看我的所有收货地址列表，并能获取我的默认收货地址，以便我能方便地选择或管理我的地址。

## Acceptance Criteria

1.  存在 `GET /v1/user/shipping-address/list` API 接口，返回当前用户的所有收货地址列表。
2.  存在 `GET /v1/user/shipping-address/default` API 接口，返回当前用户的默认收货地址。
3.  如果用户没有设置默认地址，`default` API 返回空或适当的提示。
4.  列表和默认地址的返回数据结构与 `Address` 实体一致。

## Integration Verification

1.  IV1: 现有用户认证机制正常工作。
2.  IV2: 现有商品列表和详情查询 API 正常工作。
