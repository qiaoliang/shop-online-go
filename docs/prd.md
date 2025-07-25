# 用户收货地址管理 Brownfield Enhancement PRD

## Intro Project Analysis and Context

### Existing Project Overview

#### Analysis Source

IDE-based fresh analysis

#### Current Project State

本项目是一个基于 Go 语言开发的在线购物平台后端服务，旨在提供一个稳定、高效的电商解决方案，涵盖用户、商品、购物车、订单等核心业务功能。它采用单体应用架构，模块化设计，各业务模块（如用户、商品、购物车）职责清晰，并通过分层架构实现业务逻辑与数据访问的分离。

**技术栈：**
*   **后端语言**: Go
*   **Web框架**: Gin
*   **配置管理**: Viper
*   **数据库**: MySQL (支持内存数据库用于快速启动)
*   **ORM**: GORM
*   **数据库迁移**: Migrate
*   **全局唯一ID**: xid
*   **图形验证码**: afocus/captcha

核心业务逻辑代码位于 `app/` 目录下，按照功能领域进行模块划分，包括 `addresses` (收货地址管理), `banner` (广告管理), `configs` (项目配置管理), `goods` (商品管理), `order` (订单管理), `routers` (路由配置), `security` (安全相关功能), `shoppingcart` (购物车管理), `user` (用户管理), `utils` (通用工具函数) 和 `testutils` (测试工具函数)。

项目使用 MySQL 数据库进行数据持久化，并通过 `dbscripts/` 目录管理数据库迁移脚本。配置信息通过 `config.yaml` 文件和 `app/configs` 模块进行管理。

### Available Documentation Analysis

根据您提供的 `docs/old/` 目录下的文件和 `README.md`，我们有以下可用的文档：

*   **`README.md`**: 提供了项目的概览，包括技术栈、如何运行、如何验证（API 端点示例）、目录结构、应用初始化（数据持久化、默认用户、初始商品数据）和一些业务规则（购物车红点提示、自动化测试配置）。
*   **`project_api.md`**: 详细描述了项目的 API 接口，按照业务功能模块（安全、广告、用户、收货地址、订单、商品、购物车、书籍）进行划分，并提供了 API 名称、功能说明、参数、返回值和状态。
*   **`project_biz_workflow.md`**: 包含了“添加商品到购物车流程”和“获取商品详情流程”的核心流程图（Mermaid 格式）、调用链路和关键判断点，详细阐述了这两个业务流程的实现细节。
*   **`project_domain_model.md`**: 概述了项目的核心领域模型，包括核心实体关系图（Mermaid 格式）和用户、地址、Banner、书籍、商品分类、商品项、SKU、用户购物车项等核心实体的详细属性说明。
*   **`project_external_dependency.md`**: 列出了项目的外部依赖（Go Modules，如 Gin, Viper, GORM 等）和下游服务（MySQL 数据库），并推测了潜在的外部短信服务和支付服务，以及调用关系图。
*   **`project_intro.md`**: 提供了项目的背景、目标、功能概述、技术栈（Go, Gin, Viper, MySQL, GORM）和架构类型（单体应用架构，模块化设计）。
*   **`project_structure.md`**: 详细描述了项目的模块划分、代码组织结构（目录作用）和关键包说明，并阐述了项目采用的三层架构模式（表现层、业务逻辑层、数据访问层）。

### Enhancement Scope Definition

#### Enhancement Type

Major Feature Modification

#### Enhancement Description

本次增强旨在实现用户收货地址的完全持久化，并通过一套专用的 API 提供全面的增删改查（CRUD）功能。此修改将使用户能够有效地管理其收货地址，确保数据在会话间的完整性和可访问性。

#### Impact Assessment

Moderate Impact (some existing code changes)

### Goals and Background Context

#### Goals

*   用户可以维护自己的收货地址列表

#### Background Context

目前，系统中的用户收货地址功能受限，用户只能使用一个固定且无法修改的收货地址，这极大地限制了用户的灵活性和便利性。为了提升用户体验并满足实际购物场景的需求，本次增强旨在解决这一痛点。通过允许用户拥有并管理多个收货地址，并能够指定其中一个为默认地址，将显著简化用户的收件流程，提高购物效率。

### Change Log

| Change | Date | Version | Description | Author |
|---|---|---|---|---|
| Initial Draft | 2025-07-25 | 1.0 | Initial draft of Brownfield Enhancement PRD for User Shipping Address Management. | Gemini |

## Requirements

### Functional

*   FR1: 用户应能添加新的收货地址，包括联系人、手机号、省份、城市、区域、详细地址和是否默认。
*   FR2: 用户应能查看其所有收货地址列表。
*   FR3: 用户应能修改现有收货地址的详细信息。
*   FR4: 用户应能删除其收货地址。
*   FR5: 用户应能指定一个收货地址为默认地址。
*   FR6: 系统应能返回用户的默认收货地址。

### Non Functional

*   NFR1: 收货地址的增删改查操作应在合理的时间内响应（例如，小于 500ms，支持 100qps）。
*   NFR2: 收货地址数据应安全存储，并符合数据隐私规范。
*   NFR3: 新增的收货地址功能不应影响现有用户管理和订单流程的性能。

### Compatibility Requirements

*   CR1: 现有用户相关 API（如用户详情、订单创建）应能无缝使用新的收货地址功能，无需修改。
*   CR2: 数据库模式的更改应是向后兼容的，不影响现有数据。
*   CR3: 新增的收货地址管理界面应与现有 UI/UX 保持一致。
*   CR4: 收货地址的集成应遵循现有系统的集成模式。

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

### Code Organization and Standards

**File Structure Approach**: 新增代码将遵循现有 `app/addresses` 模块的结构，包括 handler、service 和 repository 层。
**Naming Conventions**: 遵循 Go 语言的命名规范和现有项目的命名约定。
**Coding Standards**: 遵循现有项目的编码规范和最佳实践。
**Documentation Standards**: 新增 API 将在 `project_api.md` 中进行更新，并根据需要更新其他相关文档。

### Deployment and Operations

**Build Process Integration**: 新增代码将集成到现有的 Go Modules 构建流程中。
**Deployment Strategy**: 遵循现有项目的部署策略，作为单体应用的一部分进行部署。
**Monitoring and Logging**: 新增功能将集成到现有项目的监控和日志系统中，以便于问题排查和性能监控。
**Configuration Management**: 任何新的配置项将通过 `config.yaml` 和 `app/configs` 模块进行管理。

### Risk Assessment and Mitigation

**Technical Risks**:
*   **数据库模式变更**: 确保数据库模式的更改是向后兼容的，不影响现有数据和功能。
*   **并发操作**: 考虑多用户同时操作收货地址可能导致的并发问题。
**Integration Risks**:
*   **现有 API 兼容性**: 确保新的收货地址 API 与现有用户和订单相关 API 的兼容性。
*   **数据一致性**: 确保收货地址数据与用户数据保持一致。
**Deployment Risks**:
*   **部署回滚**: 确保在出现问题时能够快速回滚到之前的版本。
**Mitigation Strategies**:
*   **数据库迁移工具**: 使用 `golang-migrate` 进行数据库模式管理，确保平滑升级和回滚。
*   **并发控制**: 在代码中实现适当的并发控制机制（如锁或事务）。
*   **API 版本控制**: 如果未来需要，考虑 API 版本控制策略。
*   **充分测试**: 在部署前进行全面的单元测试、集成测试和回归测试。

## Epic and Story Structure

### Epic Approach

**Epic Structure Decision**: 鉴于本次增强是针对现有收货地址功能的“主要功能修改”，并且范围相对集中，我建议将其结构化为一个**单一的综合史诗**。这个史诗将涵盖收货地址持久化以及增删改查 API 的所有相关工作。

**理由**:
*   **集中性**: 所有工作都围绕着一个核心功能（收货地址管理）展开，将其分解为多个史诗会增加不必要的复杂性。
*   **可管理性**: 单一史诗有助于保持焦点，确保所有相关任务都在一个统一的框架下进行。
*   **风险最小化**: 在棕地项目中，将相关更改集中在一个史诗中，可以更好地管理对现有系统的影响和风险。

## Epic 1: 用户收货地址管理

**Epic Goal**: 实现用户收货地址的全面管理功能，包括持久化存储和通过 API 进行增删改查，同时确保与现有系统的无缝集成和兼容性。

**Integration Requirements**:
*   新的收货地址数据模型应与现有用户模型关联。
*   API 接口应遵循现有项目的 RESTful 规范和认证机制。
*   数据库模式变更应向后兼容。
*   确保现有用户相关功能（如订单创建）能够正确引用和使用新的收货地址数据。

### Story 1.1: 收货地址数据模型与持久化基础

作为系统管理员，我希望能够创建和管理收货地址的数据库表结构，以便为用户收货地址功能提供持久化存储基础。

#### Acceptance Criteria

1.  数据库中存在 `addresses` 表，包含 `id`, `user_id`, `link_man`, `mobile`, `province_str`, `city_str`, `area_str`, `detail_address`, `is_default` 等字段。
2.  `addresses` 表与 `users` 表通过 `user_id` 建立关联。
3.  `app/addresses` 模块中包含 `address.go` 定义的地址实体结构。
4.  `app/addresses` 模块中包含基本的 Repository 接口和实现，用于地址数据的持久化操作。

#### Integration Verification

1.  IV1: 运行数据库迁移脚本后，现有数据库表结构保持不变。
2.  IV2: 现有用户登录和详情查询功能正常工作。

### Story 1.2: 添加用户收货地址 API

作为用户，我希望能够通过 API 添加新的收货地址，以便我可以保存我的常用收货地点。

#### Acceptance Criteria

1.  存在 `POST /v1/user/shipping-address/add` API 接口，接收 `linkMan`, `mobile`, `provinceStr`, `cityStr`, `areaStr`, `detailAddress`, `isDefault` 等参数。
2.  成功添加地址后，API 返回成功信息。
3.  地址数据正确持久化到数据库。
4.  `isDefault` 字段能够正确处理，如果设置为 `true`，则将该用户其他地址的 `is_default` 设置为 `false`。

#### Integration Verification

1.  IV1: 现有用户认证机制正常工作，能够通过 Token 识别用户。
2.  IV2: 现有用户详情查询 API 正常工作。

### Story 1.3: 获取用户收货地址列表与默认地址 API

作为用户，我希望能够通过 API 查看我的所有收货地址列表，并能获取我的默认收货地址，以便我能方便地选择或管理我的地址。

#### Acceptance Criteria

1.  存在 `GET /v1/user/shipping-address/list` API 接口，返回当前用户的所有收货地址列表。
2.  存在 `GET /v1/user/shipping-address/default` API 接口，返回当前用户的默认收货地址。
3.  如果用户没有设置默认地址，`default` API 返回空或适当的提示。
4.  列表和默认地址的返回数据结构与 `Address` 实体一致。

#### Integration Verification

1.  IV1: 现有用户认证机制正常工作。
2.  IV2: 现有商品列表和详情查询 API 正常工作。

### Story 1.4: 修改和删除用户收货地址 API

作为用户，我希望能够通过 API 修改我已保存的收货地址信息，并能删除不再需要的地址，以便我能灵活地管理我的收货地址。

#### Acceptance Criteria

1.  存在 `POST /v1/user/shipping-address/modify` API 接口，接收地址 ID 和要修改的字段。
2.  存在 `POST /v1/user/shipping-address/delete` API 接口，接收地址 ID。
3.  成功修改或删除地址后，API 返回成功信息。
4.  修改后的地址信息正确反映在数据库中。
5.  删除地址后，该地址不再出现在用户地址列表中。

#### Integration Verification

1.  IV1: 现有用户认证机制正常工作。
2.  IV2: 现有订单创建流程中选择地址的功能（如果存在）能够正确处理地址的修改和删除。

### Story 1.5: 设置默认收货地址 API

作为用户，我希望能够通过 API 将我的某个收货地址设置为默认地址，以便在下单时自动选择。

#### Acceptance Criteria

1.  存在 `POST /v1/user/shipping-address/set-default` API 接口，接收地址 ID。
2.  成功设置默认地址后，API 返回成功信息。
3.  该地址的 `is_default` 字段在数据库中设置为 `true`，同时该用户其他地址的 `is_default` 字段设置为 `false`。
4.  `GET /v1/user/shipping-address/default` API 返回新设置的默认地址。

#### Integration Verification

1.  IV1: 现有用户认证机制正常工作。
2.  IV2: 现有订单创建流程能够正确获取并使用默认收货地址。

## Checklist Results Report

## Next Steps

### UX Expert Prompt

### Architect Prompt
