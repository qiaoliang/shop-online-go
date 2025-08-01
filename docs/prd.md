# Existing Project Overview

## 项目背景
本项目是一个基于 Go 语言开发的在线购物平台后端服务。旨在提供一个完整的电商解决方案，涵盖用户、商品、购物车、订单等核心业务功能。

## 项目目标
- 提供稳定、高效的电商后端API服务。
- 实现用户注册、登录、信息管理等功能。
- 实现商品分类、商品详情、商品搜索等功能。
- 实现购物车管理、订单创建与管理等功能。
- 确保系统安全性和数据一致性。

## 功能概述
- **用户管理**: 用户注册、登录、个人信息查询与修改、收货地址管理。
- **商品信息显示**: 商品分类展示、商品详情查询、商品评价、商品列表浏览。
- **购物车管理**: 商品加入购物车、修改购物车商品数量、查看购物车信息。
- **订单管理**: 订单统计、优惠券管理。
- **安全模块**: 图形验证码、短信验证码功能。
- **广告管理**: 首页Banner展示。

## 技术栈
- **后端语言**: Go
- **Web框架**: Gin
- **配置管理**: Viper
- **数据库**: SQLite3
- **ORM**: GORM
- **依赖管理**: Go Modules

## 功能详细说明


### 商品信息显示

#### 商品分类(已实现)
#### 商品详情页（已实现）
#### 商品列表浏览（已实现）
#### 商品评价（未实现）

### 购物车管理

#### 购物车红点提示（已实现）

-   购物车红点（reddot）在用户打开购物车页面之前始终显示。
-   红点数量表示用户购物车中 SKU 的总数量。
-   如果用户的购物车中存有商品，每次登录后红点都会显示。

#### 加入购物车（已实现）

#### 购物车信息持久化（未实现）

### 订单管理（未实现）

    -   订单统计
    -   优惠券管理
    -   发票管理

### 用户个人信息管理

#### 功能需求
    - FR1：用户必须登录才能浏览。用户可以logout
    - FR2: 用户可以维护自己的位置，昵称，密码，头像等

#### 非功能需求
    -   NFR1: 用户不能频繁更新密码
    -   NFR2: 头像不能太大
    -   NFR3: 昵称要合规，不能使用敏感词

#### 兼容性需求
    -   CR1: 信息变更后，不影响历史信息（如聊天记录，订单信息）

### 收货地址管理

#### 功能需求
    -   FR1: 用户应能添加新的收货地址，包括联系人、手机号、省份、城市、区域、详细地址和是否默认。
    -   FR2: 用户应能查看其所有收货地址列表。
    -   FR3: 用户应能修改现有收货地址的详细信息。
    -   FR4: 用户应能删除其收货地址。
    -   FR5: 用户应能指定一个收货地址为默认地址。
    -   FR6: 系统应能返回用户的默认收货地址。

#### 非功能需求

    -   NFR1: 收货地址的增删改查操作应在合理的时间内响应（例如，小于 500ms，支持 100qps）。
    -   NFR2: 收货地址数据应安全存储，并符合数据隐私规范。
    -   NFR3: 新增的收货地址功能不应影响现有用户管理和订单流程的性能。

#### 兼容性需求

    -   CR1: 现有用户相关 API（如用户详情、订单创建）应能无缝使用新的收货地址功能，无需修改。
    -   CR2: 数据库模式的更改应是向后兼容的，不影响现有数据。
    -   CR3: 新增的收货地址管理界面应与现有 UI/UX 保持一致。
    -   CR4: 收货地址的集成应遵循现有系统的集成模式。

### 登录管理（已实现）

### 订单管理（未实现）
    - 订单统计
    - 优惠券管理
    - 发票管理

## 架构类型
本项目采用单体应用架构，模块化设计。各业务模块（如用户、商品、购物车）职责清晰，通过分层架构实现业务逻辑与数据访问的分离。



IDE-based fresh analysis

#### Current Project State

本项目是一个基于 Go 语言开发的在线购物平台后端服务，旨在提供一个稳定、高效的电商解决方案，涵盖用户、商品、购物车、订单等核心业务功能。它采用单体应用架构，模块化设计，各业务模块（如用户、商品、购物车）职责清晰，并通过分层架构实现业务逻辑与数据访问的分离。

**技术栈：**

-   **后端语言**: Go
-   **Web 框架**: Gin
-   **配置管理**: Viper
-   **数据库**: MySQL (支持内存数据库用于快速启动)
-   **ORM**: GORM
-   **数据库迁移**: Migrate
-   **全局唯一 ID**: xid
-   **图形验证码**: afocus/captcha

核心业务逻辑代码位于 `app/` 目录下，按照功能领域进行模块划分，包括 `addresses` (收货地址管理), `banner` (广告管理), `configs` (项目配置管理), `goods` (商品管理), `order` (订单管理), `routers` (路由配置), `security` (安全相关功能), `shoppingcart` (购物车管理), `user` (用户管理), `utils` (通用工具函数) 和 `testutils` (测试工具函数)。

项目使用 MySQL 数据库进行数据持久化，并通过 `dbscripts/` 目录管理数据库迁移脚本。配置信息通过 `config.yaml` 文件和 `app/configs` 模块进行管理。

### Enhancement Scope Definition

#### Enhancement Type

Major Feature Modification

#### Enhancement Description

本次增强旨在实现用户收货地址的完全持久化，并通过一套专用的 API 提供全面的增删改查（CRUD）功能。此修改将使用户能够有效地管理其收货地址，确保数据在会话间的完整性和可访问性。

#### Impact Assessment

Moderate Impact (some existing code changes)

### Goals and Background Context



## Technical Constraints and Integration Requirements

### Existing Technology Stack

**Languages**: Go
**Frameworks**: Gin
**Database**: MySQL (支持内存数据库用于快速启动)
**Infrastructure**: (未明确，但推测为基于 Linux 的服务器环境)
**External Dependencies**: GORM, Migrate, Viper, xid, afocus/captcha

### Integration Approach

**Database Integration Strategy**: 通过 GORM 库与 MySQL 数据库进行交互，实现收货地址数据的持久化。
**API Integration Strategy**: 提供 RESTful API 接口，用于收货地址的增删改查操作。这些 API 将遵循现有项目的 API 设计规范。
**Frontend Integration Strategy**: (不适用，因为是纯后端 API 服务)
**Testing Integration Strategy**: 新增功能将通过单元测试和集成测试进行验证，确保其功能正确性并与现有系统无缝集成。