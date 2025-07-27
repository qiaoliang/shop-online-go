# Intro Project Analysis and Context

## Existing Project Overview

### Analysis Source

IDE-based fresh analysis

### Current Project State

本项目是一个基于 Go 语言开发的在线购物平台后端服务，旨在提供一个稳定、高效的电商解决方案，涵盖用户、商品、购物车、订单等核心业务功能。它采用单体应用架构，模块化设计，各业务模块（如用户、商品、购物车）职责清晰，并通过分层架构实现业务逻辑与数据访问的分离。

**技术栈：**

-   **后端语言**: Go
-   **Web 框架**: Gin
-   **配置管理**: Viper
-   **数据库**: SQLite3 (文件模式)
-   **ORM**: GORM
-   **数据库迁移**: Migrate
-   **全局唯一 ID**: xid
-   **图形验证码**: afocus/captcha

核心业务逻辑代码位于 `app/` 目录下，按照功能领域进行模块划分，包括 `addresses` (收货地址管理), `banner` (广告管理), `configs` (项目配置管理), `goods` (商品管理), `order` (订单管理), `routers` (路由配置), `security` (安全相关功能), `shoppingcart` (购物车管理), `user` (用户管理), `utils` (通用工具函数) 和 `testutils` (测试工具函数)。

项目使用 SQLite3 数据库进行数据持久化，并通过 `dbscripts/` 目录管理数据库迁移脚本。配置信息通过 `config.yaml` 文件和 `app/configs` 模块进行管理。

## Available Documentation Analysis

根据您提供的 `docs/old/` 目录下的文件和 `README.md`，我们有以下可用的文档：

-   **`README.md`**: 提供了项目的概览，包括技术栈、如何运行、如何验证（API 端点示例）、目录结构、应用初始化（数据持久化、默认用户、初始商品数据）和一些业务规则（购物车红点提示、自动化测试配置）。
-   **`project_api.md`**: 详细描述了项目的 API 接口，按照业务功能模块（安全、广告、用户、收货地址、订单、商品、购物车、书籍）进行划分，并提供了 API 名称、功能说明、参数、返回值和状态。
-   **`project_biz_workflow.md`**: 包含了“添加商品到购物车流程”和“获取商品详情流程”的核心流程图（Mermaid 格式）、调用链路和关键判断点，详细阐述了这两个业务流程的实现细节。
-   **`project_domain_model.md`**: 概述了项目的核心领域模型，包括核心实体关系图（Mermaid 格式）和用户、地址、Banner、书籍、商品分类、商品项、SKU、用户购物车项等核心实体的详细属性说明。
-   **`project_external_dependency.md`**: 列出了项目的外部依赖（Go Modules，如 Gin, Viper, GORM 等）和下游服务（SQLite3 数据库），并推测了潜在的外部短信服务和支付服务，以及调用关系图。
-   **`project_intro.md`**: 提供了项目的背景、目标、功能概述、技术栈（Go, Gin, Viper, SQLite3, GORM）和架构类型（单体应用架构，模块化设计）。
-   **`project_structure.md`**: 详细描述了项目的模块划分、代码组织结构（目录作用）和关键包说明，并阐述了项目采用的三层架构模式（表现层、业务逻辑层、数据访问层）。

## Enhancement Scope Definition

### Enhancement Type

Major Feature Modification

### Enhancement Description

本次增强旨在实现用户收货地址的完全持久化，并通过一套专用的 API 提供全面的增删改查（CRUD）功能。此修改将使用户能够有效地管理其收货地址，确保数据在会话间的完整性和可访问性。

### Impact Assessment

Moderate Impact (some existing code changes)

## Goals and Background Context

### Goals

-   用户可以维护自己的收货地址列表

### Background Context

目前，系统中的用户收货地址功能受限，用户只能使用一个固定且无法修改的收货地址，这极大地限制了用户的灵活性和便利性。为了提升用户体验并满足实际购物场景的需求，本次增强旨在解决这一痛点。通过允许用户拥有并管理多个收货地址，并能够指定其中一个为默认地址，将显著简化用户的收件流程，提高购物效率。

## Change Log

| Change        | Date       | Version | Description                                                                       | Author |
| ------------- | ---------- | ------- | --------------------------------------------------------------------------------- | ------ |
| Initial Draft | 2025-07-25 | 1.0     | Initial draft of Brownfield Enhancement PRD for User Shipping Address Management. | Gemini |
