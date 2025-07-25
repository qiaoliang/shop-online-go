# Source Tree

```plaintext
shop-online-go/
├── app/
│   ├── addresses/          # 新增：用户收货地址模块
│   │   ├── address.go      # 地址实体定义
│   │   ├── addressHandler.go # 表现层：处理 HTTP 请求
│   │   ├── addressService.go # 业务逻辑层：处理业务逻辑
│   │   └── addressRepo.go  # 数据访问层：与数据库交互
│   ├── banner/             # 广告模块
│   ├── configs/            # 配置模块
│   ├── goods/              # 商品模块
│   ├── order/              # 订单模块
│   ├── routers/            # 路由配置
│   │   └── router.go       # 路由定义，将包含 addresses 模块的路由
│   ├── security/           # 安全模块
│   ├── shoppingcart/       # 购物车模块
│   ├── user/               # 用户管理模块
│   │   └── user.go         # 用户实体，可能需要更新以关联地址
│   ├── utils/              # 通用工具
│   └── testutils/          # 测试工具
├── dbscripts/              # 数据库迁移脚本
│   └── XX_create_table_addresses.up.sql # 新增：addresses 表创建脚本
│   └── XX_create_table_addresses.down.sql # 新增：addresses 表回滚脚本
├── docs/                   # 项目文档
│   ├── prd.md              # 产品需求文档
│   └── architecture.md     # 架构文档 (当前正在生成)
├── initData/               # 初始化数据文件
├── static/                 # 静态资源文件
├── main.go                 # 项目入口文件
├── go.mod                  # Go 模块定义文件
├── config.yaml             # 配置文件
└── README.md               # 项目说明
```
