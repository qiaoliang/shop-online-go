# Story 1.4: 修改和删除用户收货地址 API

作为用户，我希望能够通过 API 修改我已保存的收货地址信息，并能删除不再需要的地址，以便我能灵活地管理我的收货地址。

## Acceptance Criteria

1.  存在 `POST /v1/user/shipping-address/modify` API 接口，接收地址 ID 和要修改的字段。
2.  存在 `POST /v1/user/shipping-address/delete` API 接口，接收地址 ID。
3.  成功修改或删除地址后，API 返回成功信息。
4.  修改后的地址信息正确反映在数据库中。
5.  删除地址后，该地址不再出现在用户地址列表中。

## Integration Verification

1.  IV1: 现有用户认证机制正常工作。
2.  IV2: 现有订单创建流程中选择地址的功能（如果存在）能够正确处理地址的修改和删除。
