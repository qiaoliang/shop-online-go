# BMad 文档模板规范(BMad Document Template Specification)

## 概述(Overview)

BMad 文档模板以 YAML 格式定义，用于驱动交互式文档生成和 agent 交互。模板将结构定义与内容生成分离，使其既适合人类阅读，也适合 LLM-agent 使用。

## 模板结构(Template Structure)

```yaml
template:
    id: template-identifier
    name: Human Readable Template Name
    version: 1.0
    output:
        format: markdown
        filename: default-path/to/{{filename}}.md
        title: "{{variable}} Document Title"

workflow:
    mode: interactive
    elicitation: advanced-elicitation

sections:
    - id: section-id
      title: Section Title
      instruction: |
          Detailed instructions for the LLM on how to handle this section
      # ... additional section properties
```

## 核心字段(Core Fields)

### 模板元数据(Template Metadata)

- **id**: 模板的唯一标识符
- **name**: 在 UI 中显示的人类可读名称
- **version**: 用于跟踪变更的模板版本
- **output.format**: 文档模板默认为 "markdown"
- **output.filename**: 默认输出文件路径（可包含变量）
- **output.title**: 文档标题（在 markdown 中成为 H1）

### 工作流配置(Workflow Configuration)

- **workflow.mode**: 默认交互模式（"interactive" 或 "yolo"）
- **workflow.elicitation**: 要使用的启发式任务（"advanced-elicitation"）

## 章节属性(Section Properties)

### 必需字段(Required Fields)

- **id**: 唯一的章节标识符
- **title**: 章节标题文本
- **instruction**: 为 LLM 处理此章节提供的详细指导

### 可选字段(Optional Fields)

#### 内容控制(Content Control)

- **type**: 结构化章节的内容类型提示
- **template**: 章节内容的固定模板文本
- **item_template**: 章节内可重复项目的模板
- **prefix**: 编号项目的前缀（例如，"FR"，"NFR"）

#### 行为标志(Behavior Flags)

- **elicit**: 布尔值 - 在章节渲染后应用启发式
- **repeatable**: 布尔值 - 章节可以重复多次
- **condition**: 字符串 - 包含章节的条件（例如，"has ui requirements"）

#### Agent 权限(Agent Permissions)

- **owner**: 字符串 - 最初创建/填充此章节的 agent 角色
- **editors**: 数组 - 允许修改此章节的 agent 角色列表
- **readonly**: 布尔值 - 章节在初始创建后无法修改

#### 内容指导(Content Guidance)

- **examples**: 示例内容数组（不包含在输出中）
- **choices**: 包含常见决策选择选项的对象
- **placeholder**: 默认占位符文本

#### 结构(Structure)

- **sections**: 嵌套子章节数组

## 支持的类型(Supported Types)

### 内容类型(Content Types)

- **bullet-list**: 无序列表项目
- **numbered-list**: 有序列表，带可选前缀
- **paragraphs**: 自由格式段落文本
- **table**: 结构化表格数据
- **code-block**: 代码或配置块
- **template-text**: 带变量替换的固定模板
- **mermaid**: 具有指定类型和详情的 Mermaid 图表

### 特殊类型(Special Types)

- **repeatable-container**: 多个实例的容器
- **conditional-block**: 基于条件显示的内容
- **choice-selector**: 向用户呈现选择

## 高级功能(Advanced Features)

### 变量替换(Variable Substitution)

在标题、模板和内容中使用 `{{variable_name}}`：

```yaml
title: "Epic {{epic_number}} {{epic_title}}"
template: "As a {{user_type}}, I want {{action}}, so that {{benefit}}."
```

### 条件章节(Conditional Sections)

```yaml
- id: ui-section
  title: User Interface Design
  condition: Project has UX/UI Requirements
  instruction: Only include if project has UI components
```

### 选择集成(Choice Integration)

```yaml
choices:
    architecture: [Monolith, Microservices, Serverless]
    testing: [Unit Only, Unit + Integration, Full Pyramid]
```

### Mermaid 图表(Mermaid Diagrams)

```yaml
- id: system-architecture
  title: System Architecture Diagram
  type: mermaid
  instruction: Create a system architecture diagram showing key components and data flow
  mermaid_type: flowchart
  details: |
      Show the following components:
      - User interface layer
      - API gateway
      - Core services
      - Database layer
      - External integrations
```

**支持的 mermaid_type 值：**

**核心图表类型：**

- `flowchart` - 流程图和过程图
- `sequenceDiagram` - 交互的序列图
- `classDiagram` - 类关系图（UML）
- `stateDiagram` - 状态转换图
- `erDiagram` - 实体关系图
- `gantt` - 时间线的甘特图
- `pie` - 数据可视化的饼图

**高级图表类型：**

- `journey` - 用户旅程图
- `mindmap` - 头脑风暴的思维导图
- `timeline` - 时间事件的时序图
- `quadrantChart` - 数据分类的四象限图
- `xyChart` - XY 图表（条形图、折线图）
- `sankey` - 流程可视化的桑基图

**专业类型：**

- `c4Context` - C4 上下文图（实验性）
- `requirement` - 需求图
- `packet` - 网络数据包图
- `block` - 块图
- `kanban` - 看板

### Agent 权限示例(Agent Permissions Example)

```yaml
- id: story-details
  title: Story
  owner: scrum-master
  editors: [scrum-master]
  readonly: false
  sections:
      - id: dev-notes
        title: Dev Notes
        owner: dev-agent
        editors: [dev-agent]
        readonly: false
        instruction: Implementation notes and technical details
      - id: qa-results
        title: QA Results
        owner: qa-agent
        editors: [qa-agent]
        readonly: true
        instruction: Quality assurance test results
```

### 可重复章节(Repeatable Sections)

```yaml
- id: epic-details
  title: Epic {{epic_number}} {{epic_title}}
  repeatable: true
  sections:
      - id: story
        title: Story {{epic_number}}.{{story_number}} {{story_title}}
        repeatable: true
        sections:
            - id: criteria
              title: Acceptance Criteria
              type: numbered-list
              item_template: "{{criterion_number}}: {{criteria}}"
              repeatable: true
```

### 代码块示例(Examples with Code Blocks)

````yaml
examples:
    - "FR6: The system must authenticate users within 2 seconds"
    - |
        ```mermaid
        sequenceDiagram
            participant User
            participant API
            participant DB
            User->>API: POST /login
            API->>DB: Validate credentials
            DB-->>API: User data
            API-->>User: JWT token
        ```
    - |
        **Architecture Decision Record**

        **Decision**: Use PostgreSQL for primary database
        **Rationale**: ACID compliance and JSON support needed
        **Consequences**: Requires database management expertise
````

## 章节层次结构(Section Hierarchy)

模板定义从第一个 H2 开始的完整文档结构 - 每个级别都是下一个 H#：

```yaml
sections:
    - id: overview
      title: Project Overview
      sections:
          - id: goals
            title: Goals
          - id: scope
            title: Scope
            sections:
                - id: in-scope
                  title: In Scope
                - id: out-scope
                  title: Out of Scope
```

## 处理流程(Processing Flow)

1. **解析模板**: 加载并验证 YAML 结构
2. **初始化工作流**: 设置交互模式和启发式
3. **处理章节**: 按顺序处理每个章节：
    - 检查条件
    - 应用指令
    - 生成内容
    - 处理选择和变量
    - 如果指定则应用启发式
    - 处理嵌套章节
4. **生成输出**: 创建干净的 markdown 文档

## 最佳实践(Best Practices)

### 模板设计(Template Design)

- 保持指令清晰具体
- 为复杂内容使用示例
- 逻辑地构建章节
- 包含 LLM 所需的所有必要指导

### 内容指令(Content Instructions)

- 明确说明预期格式
- 包含决策的理由
- 指定交互模式
- 在需要时引用其他文档

### 变量命名(Variable Naming)

- 使用描述性变量名
- 遵循一致的命名约定
- 记录预期的变量值

### 示例使用(Examples Usage)

- 为复杂章节提供具体示例
- 包含简单和复杂的情况
- 使用现实的项目场景
- 在有用时包含代码块和图表

## 验证(Validation)

模板应验证：

- 有效的 YAML 语法
- 存在必需字段
- 一致的章节 ID
- 正确的嵌套结构
- 有效的变量引用

## 从旧版本迁移(Migration from Legacy)

从 markdown+frontmatter 模板转换时：

1. 将嵌入的 `[[LLM:]]` 指令提取到 `instruction` 字段
2. 将 `<<REPEAT>>` 块转换为 `repeatable: true` 章节
3. 将 `^^CONDITIONS^^` 提取到 `condition` 字段
4. 将 `@{examples}` 移动到 `examples` 数组
5. 将 `{{placeholders}}` 转换为正确的变量语法

此规范确保模板既适合人类阅读又适合机器处理，同时保持复杂文档生成所需的灵活性。
