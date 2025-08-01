# 对 Brownfield 项目进行初步分析

## 目的(Purpose)

为现有项目生成面向 AI Dev-Agent 的全面文档。

## 任务指令(Task Instructions)

对项目文件与代码进行全面扫描与分析，并生成项目梳理文档 `docs/startpoint/project_explain.md`

## 基本要求

如果有`docs/spec/project_specs.md`，你生成的项目梳理文档 `docs/startpoint/project_explain.md` 必须严格按照该文档的要求生成。

## `docs/startpoint/project_explain.md`文档规范

**文档受众明确为软件开发人员**，目的是帮助开发团队快速理解系统架构、业务逻辑和技术实现细节，便于代码维护、功能扩展和知识传递。

## 关键规则

- 项目文档必须包含四个核心部分：项目简介、核心领域模型、项目结构和外部依赖, 接口文档,业务流程.
- 接口文档必须按照 `docs/spec/api_spec.md` 进行编写和维护
- 业务流程文档必须按照 `docs/spec/bizflow-spec.md` 进行编写和维护
- 文档应保持客观性，基于现有代码而非理想状态
- 文档中使用的术语必须与代码中的术语保持一致
- 文档应使用 Markdown 格式，支持图表、表格和代码块
    - 代码示例必须是从实际代码中提取的，而非虚构的
    - 图表应使用 Mermaid 语法，以确保可维护性
    - 如有 SQL 文件，使用 Mermaid 中的 ER图 语法生成对应的 ER 图
    - 文档应当引用具体的代码文件路径，便于读者查找相关实现

- 首先判断项目是否使用流行架构（例如 BBF，GBF，SPA），并根据实际架构选择适合的文档结构和内容
- 所有新生成的文档必须统一放置在`docs/startpoint` 目录下，并使用规定的英文名称
- **文档生成过程中必须确保完整覆盖所有内容，不允许任何遗漏**

## 文档优化与结构指南

- **主索引文档**：每个核心部分创建一个主索引文档，包含子文档链接和简要说明
- **文档内导航**：超过 500 行的文档必须在开头提供目录
- **分层结构**：按照"金字塔结构"组织（顶层：核心概念；中层：主要功能模块；底层：具体实现细节）
- **文档拆分**：接口超过20个时按业务域拆分；核心实体超过10个时按业务领域拆分. 拆分后，文档以连续的“xxx\_{数字序号}.md" 命名。“业务域”可按子系统或业务功能模块的划分。

## 文档结构和内容要求

### 1. 项目简介 - docs/startpoint/project_intro.md

必须包含：项目背景、项目目标、功能概述、技术栈和架构类型

### 2. 核心领域模型 - docs/startpoint/project_domain_model.md

必须包含：

- 领域模型概述：核心业务概念的定义和边界
- 核心实体关系图：使用E-R图或类图表示
- 关键业务场景下的模型交互
- 数据流转关系

**强制性领域模型扫描规则**：

- **按目录结构识别**：位于项目目录下的类文件
- **完整提取**：实体属性和业务含义、实体关系、聚合结构、生命周期和状态流转
- **识别规则**：属性约束、实体关系约束、状态转换规则

**领域模型分析策略**：

- 全域扫描实体类和值对象，支持多种 ORM 框架
- 提取关联关系（通过字段类型、泛型参数和ORM注解）
- 识别聚合根和聚合边界（通过包结构和类间关系）
- 分析继承结构（包括抽象类、接口和实现类）
- 提取业务方法和状态转换逻辑
- 生成完整属性表和业务规则说明

### 3. 接口文档 - docs/startpoint/project_api.md

- 如有`docs/spec/api_spec.md`，接口文档应遵循该文件进行创建和维护，以确保API接口的完整记录和更新。
- 接口文档应该按业务功能模块划分编写，每个模块的API 列表如下

```markdown
the following is a example.

| API名称                 | 功能说明     | 参数数据说明                   | 返回值说明                                       |
| ----------------------- | ------------ | ------------------------------ | ------------------------------------------------ |
| get_article_list(param) | 获得文章列表 | param: int, 指定所查的文章状态 | 返回值类型，及含义，例如：list，查找到的文章列表 |
```

### 4. 业务流程 - docs/startpoint/project_biz_workflow.md

- 业务流程文档应遵循专门的 `docs/startpoint/biz_workflow_doc_spec.md` 进行创建和维护，以确保业务流程的完整记录和更新。
- 业务流程应该按功能不同进行业务流程分组，每个组下为一组相关的业务流程。流程使用 mermaid 格式输出。

### 5. 项目结构 - docs/startpoint/project_structure.md

必须包含：项目模块划分、代码组织结构、关键包说明、分层架构说明

### 6. 外部依赖与下游服务 - docs/startpoint/project_external_dependency.md

必须包含：

- 下游服务概述：依赖的所有外部服务列表和用途
- 调用关系图：系统与外部服务的调用关系

## 文档生成工作流程

1. **架构识别**：确定项目架构类型、识别关键组件和分层结构
2. **代码分析**：识别核心业务包和类、分析领域模型、提取接口定义、理解调用链路
3. **内容整理**：按文档结构组织信息、提取代码示例、绘制图表
4. **审核完善**：验证文档与代码一致性、补充关键信息、完善图表和示例
    - **接口覆盖性验证**：确认总览文档中的所有接口都在详细文档中有完整描述
    - **文档完整性检查**：确保没有遗漏任何必要的接口和服务描述
5. **定期更新**：与代码审查流程集成、重大变更更新文档、每季度全面审核

## 以下为部分示例文档

---

### 核心实体关系图

下面是这部分内容的示例，使用 mermaid 格式

```mermaid
classDiagram
    class Item {
        +Long id
        +String name
        +BigDecimal price
        +String status
        +validatePrice()
        +changeStatus(String)
    }
    class TyingRule {
        +Long id
        +Long mainItemId
        +List<Long> subItemIds
        +Date startTime
        +Date endTime
        +enable()
        +disable()
    }
    Item "1" -- "n" TyingRule: 被定义为主商品
    TyingRule "1" -- "n" Item: 关联搭售商品
```

#### Item 商品实体

下面是这部分内容的示例，使用 markdown 格式

```markdown
| 属性名 | 类型       | 说明                                                             |
| ------ | ---------- | ---------------------------------------------------------------- |
| id     | Long       | 商品唯一标识                                                     |
| name   | String     | 商品名称，长度限制：2-50个字符                                   |
| price  | BigDecimal | 商品价格，精确到小数点后2位，最小值：0.01                        |
| status | String     | 商品状态，枚举值：ON_SHELF(上架)、OFF_SHELF(下架)、DELETED(删除) |
```

#### 业务规则

下面是这部分内容的示例，使用 markdown 格式

```markdown
- 商品价格必须大于0
- 商品状态只能按特定流程转换（上架->下架->删除）
```

### 业务流程示例

## 搭售规则创建流程

### 核心流程图

下面是这部分内容的示例，使用 mermaid 格式

```mermaid
flowchart TD
    A[创建请求] --> B{校验参数}
    B -->|无效| C[返回错误]
    B -->|有效| D[查询主商品]
    D --> E{商品存在?}
    E -->|否| F[返回错误]
    E -->|是| G[查询搭售商品]
    G --> H{商品存在?}
    H -->|否| I[返回错误]
    H -->|是| J[保存规则]
    J --> K[返回成功]
```

### 调用链路

下面是这部分内容的示例，使用 markdown 格式

```markdown
**入口点**: `ItemTyingController.createTyingRule()`
**调用流程**:

1. 请求参数校验 - `validateTyingRequest(request)`
2. 查询主商品信息 - `itemService.getItemById()`
3. 校验主商品状态 - `validateItemStatus(item)`
4. 查询并校验搭售商品列表 - `validateSubItems()`
5. 构建并保存搭售规则 - `tyingRuleRepository.save()`
6. 发送规则创建事件 - `eventPublisher.publishEvent()`
```

### 关键判断点

下面是这部分内容的示例，使用 markdown 格式

```markdown
the following is a example.

| 判断点       | 条件         | 处理路径           |
| ------------ | ------------ | ------------------ |
| 参数校验     | 主商品ID为空 | 返回参数错误       |
| 主商品校验   | 主商品不存在 | 返回商品不存在错误 |
| 搭售商品校验 | 存在无效商品 | 返回商品无效错误   |
```
