# Data Models

## Address

**Purpose:** 存储用户的收货地址信息，支持用户拥有多个地址并指定默认地址。

**Key Attributes:**
- `Id`: string - 地址唯一标识
- `UserId`: string - 关联的用户 ID
- `LinkMan`: string - 联系人姓名
- `Mobile`: string - 联系手机号
- `IsDefault`: bool - 是否为默认地址
- `ProvinceStr`: string - 省份
- `CityStr`: string - 城市
- `AreaStr`: string - 区域
- `DetailAddress`: string - 详细地址

**Relationships:**
- `Address` 与 `User` 之间存在一对多关系，一个用户可以有多个收货地址。
