# Shop Online Go

一个基于 [Gin](https://github.com/gin-gonic/gin) 框架构建的简单电商后端应用。

# 技术栈 (Stack)

-   [GO](https://go.dev/)
-   [MySQL](https://www.mysql.com/)
-   [GORM](https://gorm.io/index.html)
-   [Gin](https://github.com/gin-gonic/gin)
-   [Migrate](https://github.com/golang-migrate)
-   [Viper](https://github.com/spf13/viper)

# 如何运行 (How to Run)

1.  **配置应用程序**: 
    编辑 `config.yaml` 文件，根据您的MySQL数据库配置更新以下信息：
    ```yaml
    MYSQL:
      DB_USERNAME: your_username
      DB_PASSWORD: your_password
      BASE_URL: 127.0.0.1
      DB_PORT: 3306
      DB_NAME: shop_online_go # 建议使用更具描述性的名称，例如 shop_online_go
      DB_MIG_PROTO: file
      DB_MIG_DIR: dbscripts
    ```
    确保 `DB_NAME` 与您将要创建的数据库名称一致。

2.  **创建数据库**: 
    在运行应用程序之前，请确保您的MySQL服务器已启动，并创建一个与 `config.yaml` 中 `DB_NAME` 对应的数据库。例如：
    ```sql
    CREATE DATABASE shop_online_go CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
    ```

3.  **编译应用程序**: 
    在项目根目录下执行以下命令编译应用程序：
    ```bash
    go build -o shop-online-go main.go
    ```
    这将在当前目录生成一个名为 `shop-online-go` 的可执行文件。

4.  **运行应用程序**: 
    执行编译后的可执行文件来启动应用程序。应用程序启动时会自动执行数据库迁移脚本 (`dbscripts` 目录下的SQL文件)，初始化数据库结构和数据。
    ```bash
    ./shop-online-go
    ```
    应用程序将监听 `config.yaml` 中 `PORT` 配置的端口（默认为 `8080`）。

# 如何验证 (How to Verify)

应用程序成功启动后，您可以通过访问以下API端点来验证其功能：

1.  **获取所有商品分类**:
    访问 `http://localhost:8080/v1/shop/goods/category/all` (请将 `8080` 替换为 `config.yaml` 中配置的实际端口)。
    预期响应示例：
    ```json
    {"code":0,"data":[{"id":0,"name":"DevOps"},{"id":1,"name":"大数据"}],"msg":"OK"}
    ```

2.  **获取Banner列表**:
    访问 `http://localhost:8080/v1/banner/list`。
    预期响应示例：
    ```json
    {"code":0,"data":[{"businessId":0,"dateAdd":"2023-01-01","id":1,"linkUrl":"http://example.com/banner1","paixu":1,"picUrl":"/pic/banners/b0001.jpeg","remark":"Banner 1","status":1,"statusStr":"正常","title":"Banner Title 1","type":"首页","userId":0}],"msg":"OK"}
    ```

# 目录结构 (Directory Structure)

*   `app`: 核心业务逻辑代码
    *   `addresses`: 用户收货地址管理模块
    *   `banner`: 广告（Banner）管理模块
    *   `configs`: 系统配置管理模块
    *   `goods`: 商品及分类管理模块
    *   `order`: 订单管理模块
    *   `routers`: 应用程序路由配置
    *   `security`: 安全相关功能（如验证码）
    *   `shoppingcart`: 购物车管理模块
    *   `user`: 用户信息管理模块
    *   `utils`: 通用工具函数
    *   `testutils`: 测试工具函数
*   `dbscripts`: 数据库迁移脚本，由 `main.go` 在启动时调用执行。
*   `initData`: 应用程序初始化数据文件（如 `banners.json`, `categories.json`, `SKUs.json`）。
*   `resources`: 应用程序资源文件（如验证码字体）。
*   `static`: 静态资源文件，包含商品图片、用户头像和广告Banner图片。
*   `config.yaml`: 应用程序主配置文件，主要用于数据库连接和端口设置。
*   `main.go`: 应用程序的入口文件。
*   `go.mod`: Go模块定义文件，管理项目依赖。

# 应用初始化 (Application Initialization)

*   **数据持久化**:
    应用程序支持数据持久化到MySQL数据库。您可以通过修改 `config.yaml` 中的 `PERSISTANCE` 字段来控制是否启用持久化。
    *   `PERSISTANCE: true` (默认): 数据将存储在MySQL数据库中，并在应用程序重启后保留。
    *   `PERSISTANCE: false`: 应用程序将使用内存存储数据，每次重启后数据会重置。

*   **默认用户**:
    *   用户名: `13900007997`
    *   密码: `1234`

*   **初始商品数据**:
    应用程序启动时会加载初始商品分类和SKU数据。
    *   **商品分类**: `DevOps` 和 `大数据`。
    *   **SKU ID 范围**:
        *   `DevOps` 分类包含 SKU ID 从 `g7225946` 到 `g7225949` 的商品。
        *   `大数据` 分类包含 SKU ID 从 `g1872110` 到 `g1872113` 的商品。
    *   **图片资源**: 所有商品图片都存储在 `static` 目录下，并以SKU ID命名。

# 一些业务规则 (Business Rules)

*   **购物车红点提示**:
    *   购物车红点（reddot）在用户打开购物车页面之前始终显示。
    *   红点数量表示用户购物车中SKU的总数量。
    *   如果用户的购物车中存有商品，每次登录后红点都会显示。

# 自动化测试 (Automated Testing)

*   **测试环境配置**:
    为了将测试环境与生产环境分离，自动化测试会加载 `config-test.yaml` 文件作为测试配置，而不是 `config.yaml`。
*   **数据库准备**:
    在运行测试之前，您需要在MySQL中创建一个名为 `bookstore` 的数据库实例。
*   **运行测试**:
    在项目根目录下执行以下命令来运行所有测试：
    ```bash
    go test ./...
    ```
