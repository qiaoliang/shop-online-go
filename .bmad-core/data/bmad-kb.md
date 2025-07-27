# BMad知识库(BMad Knowledge Base)

## 概述(Overview)

BMad-Method(Breakthrough Method of Agile AI-driven Development)是一个将AI代理与敏捷开发方法论相结合的框架。v4系统引入了模块化架构，具有改进的依赖管理、包优化以及对Web和IDE环境的支持。

### 关键特性(Key Features)

- **模块化代理系统(Modular Agent System)**: 为每个敏捷角色提供专门的AI代理
- **构建系统(Build System)**: 自动化依赖解析和优化
- **双环境支持(Dual Environment Support)**: 针对Web UI和IDE进行优化
- **可重用资源(Reusable Resources)**: 可移植的模板、任务和检查清单
- **斜杠命令集成(Slash Command Integration)**: 快速代理切换和控制

### 何时使用BMad(When to Use BMad)

- **新项目(Greenfield)**: 完整的端到端开发
- **现有项目(Brownfield)**: 功能添加和增强
- **团队协作(Team Collaboration)**: 多个角色协同工作
- **质量保证(Quality Assurance)**: 结构化测试和验证
- **文档(Documentation)**: 专业的PRD、架构文档、用户故事

## BMad如何工作(How BMad Works)

### 核心方法(The Core Method)

BMad将您转变为"氛围CEO" - 通过结构化工作流程指导专门的AI代理团队。方法如下：

1. **您指导，AI执行(You Direct, AI Executes)**: 您提供愿景和决策；代理处理实施细节
2. **专门代理(Specialized Agents)**: 每个代理掌握一个角色(PM、Developer、Architect等)
3. **结构化工作流程(Structured Workflows)**: 经过验证的模式指导您从想法到部署代码
4. **清洁交接(Clean Handoffs)**: 新鲜的上下文窗口确保代理保持专注和有效

### 两阶段方法(The Two-Phase Approach)

#### 阶段1：规划(Web UI - 成本效益)(Phase 1: Planning (Web UI - Cost Effective))

- 使用大上下文窗口(Gemini的1M tokens)
- 生成综合文档(PRD、Architecture)
- 利用多个代理进行头脑风暴
- 创建一次，在整个开发过程中使用

#### 阶段2：开发(IDE - 实施)(Phase 2: Development (IDE - Implementation))

- 将文档分片为可管理的部分
- 执行专注的SM → Dev周期
- 一次一个故事，顺序进展
- 实时文件操作和测试

### 开发循环(The Development Loop)

```text
1. SM Agent (New Chat) → 从分片文档创建下一个故事
2. You → 审查并批准故事
3. Dev Agent (New Chat) → 实施批准的故事
4. QA Agent (New Chat) → 审查和重构代码
5. You → 验证完成
6. 重复直到epic完成
```

### 为什么这有效(Why This Works)

- **上下文优化(Context Optimization)**: 清洁聊天 = 更好的AI性能
- **角色清晰(Role Clarity)**: 代理不切换上下文 = 更高质量
- **增量进展(Incremental Progress)**: 小故事 = 可管理的复杂性
- **人工监督(Human Oversight)**: 您验证每个步骤 = 质量控制
- **文档驱动(Document-Driven)**: 规范指导一切 = 一致性

## 开始使用(Getting Started)

### 快速开始选项(Quick Start Options)

#### 选项1：Web UI(Option 1: Web UI)

**最适合**: ChatGPT、Claude、Gemini用户，希望立即开始

1. 导航到`dist/teams/`
2. 复制`team-fullstack.txt`内容
3. 创建新的Gemini Gem或CustomGPT
4. 上传文件，说明："您的重要操作说明已附加，请按指示不要打破角色"
5. 输入`/help`查看可用命令

#### 选项2：IDE集成(Option 2: IDE Integration)

**最适合**: Cursor、Claude Code、Windsurf、Trae、Cline、Roo Code、Github Copilot用户

```bash
# 交互式安装(推荐)
npx bmad-method install
```

**安装步骤(Installation Steps)**:

- 选择"完整安装"
- 从支持的选项中选择您的IDE：
    - **Cursor**: 原生AI集成
    - **Claude Code**: Anthropic官方IDE
    - **Windsurf**: 内置AI功能
    - **Trae**: 内置AI功能
    - **Cline**: 带AI功能的VS Code扩展
    - **Roo Code**: 带代理支持的基于Web的IDE
    - **GitHub Copilot**: 带AI同行编程助手的VS Code扩展

**VS Code用户注意**: BMad-Method假设当您提到"VS Code"时，您正在使用带AI功能扩展的VS Code，如GitHub Copilot、Cline或Roo。没有AI功能的标准VS Code无法运行BMad代理。安装程序包含对Cline和Roo的内置支持。

**验证安装(Verify Installation)**:

- 创建了`.bmad-core/`文件夹，包含所有代理
- 创建了IDE特定的集成文件
- 所有代理命令/规则/模式可用

**记住**: 就其核心而言，BMad-Method是关于掌握和利用提示工程。任何具有AI代理支持的IDE都可以使用BMad - 框架提供使AI开发有效的结构化提示和工作流程

### 环境选择指南(Environment Selection Guide)

**使用Web UI用于**:

- 初始规划和文档(PRD、architecture)
- 成本效益的文档创建(特别是使用Gemini)
- 头脑风暴和分析阶段
- 多代理咨询和规划

**使用IDE用于**:

- 主动开发和编码
- 文件操作和项目集成
- 文档分片和故事管理
- 实施工作流程(SM/Dev周期)

**成本节省提示**: 在Web UI中创建大型文档(PRD、architecture)，然后在切换到IDE进行开发之前复制到项目的`docs/prd.md`和`docs/architecture.md`。

### 仅IDE工作流程考虑(IDE-Only Workflow Considerations)

**您能在IDE中做所有事情吗？** 是的，但要理解权衡：

**仅IDE的优点**:

- 单一环境工作流程
- 从一开始的直接文件操作
- 环境间无需复制/粘贴
- 立即项目集成

**仅IDE的缺点**:

- 大型文档创建的更高token成本
- 较小的上下文窗口(因IDE/模型而异)
- 在规划阶段可能达到限制
- 头脑风暴的成本效益较低

**在IDE中使用Web代理**:

- **不推荐**: Web代理(PM、Architect)具有为大型上下文设计的丰富依赖
- **为什么重要**: Dev代理保持精简以最大化编码上下文
- **原则**: "Dev代理编码，规划代理规划" - 混合破坏此优化

**关于bmad-master和bmad-orchestrator**:

- **bmad-master**: 可以在不切换代理的情况下执行任何任务，但是...
- **仍然使用专门代理进行规划**: PM、Architect和UX Expert具有产生更好结果的调优角色
- **为什么专业化重要**: 每个代理的个性和专注创造更高质量的输出
- **如果使用bmad-master/orchestrator**: 对规划阶段很好，但是...

**开发的关键规则**:

- **始终使用SM代理创建故事** - 永远不要使用bmad-master或bmad-orchestrator
- **始终使用Dev代理进行实施** - 永远不要使用bmad-master或bmad-orchestrator
- **为什么重要**: SM和Dev代理专门针对开发工作流程进行了优化
- **无例外**: 即使在其他方面使用bmad-master，也要切换到SM → Dev进行实施

**仅IDE的最佳实践**:

1. 使用PM/Architect/UX代理进行规划(比bmad-master更好)
2. 直接在项目中创建文档
3. 创建后立即分片
4. **必须切换到SM代理**进行故事创建
5. **必须切换到Dev代理**进行实施
6. 在单独的聊天会话中保持规划和编码

## 核心配置(core-config.yaml)(Core Configuration (core-config.yaml))

**V4新增**: `bmad-core/core-config.yaml`文件是一个关键创新，使BMad能够与任何项目结构无缝工作，提供最大灵活性和向后兼容性。

### 什么是core-config.yaml？(What is core-config.yaml?)

此配置文件充当BMad代理的地图，告诉它们确切在哪里找到您的项目文档以及它们的结构。它启用：

- **版本灵活性**: 与V3、V4或自定义文档结构一起工作
- **自定义位置**: 定义您的文档和分片的位置
- **开发者上下文**: 指定dev代理应始终加载的文件
- **调试支持**: 内置日志记录用于故障排除

### 关键配置区域(Key Configuration Areas)

#### PRD配置(PRD Configuration)

- **prdVersion**: 告诉代理PRD是否遵循v3或v4约定
- **prdSharded**: epics是嵌入的(false)还是在单独文件中(true)
- **prdShardedLocation**: 在哪里找到分片epic文件
- **epicFilePattern**: epic文件名的模式(例如，`epic-{n}*.md`)

#### 架构配置(Architecture Configuration)

- **architectureVersion**: v3(整体)或v4(分片)
- **architectureSharded**: 架构是否拆分为组件
- **architectureShardedLocation**: 分片架构文件的位置

#### 开发者文件(Developer Files)

- **devLoadAlwaysFiles**: dev代理为每个任务加载的文件列表
- **devDebugLog**: dev代理记录重复失败的位置
- **agentCoreDump**: 聊天对话的导出位置

### 为什么重要(Why It Matters)

1. **无强制迁移**: 保持您现有的文档结构
2. **渐进采用**: 从V3开始，按您的节奏迁移到V4
3. **自定义工作流程**: 配置BMad以匹配您团队的过程
4. **智能代理**: 代理自动适应您的配置

### 常见配置(Common Configurations)

**遗留V3项目**:

```yaml
prdVersion: v3
prdSharded: false
architectureVersion: v3
architectureSharded: false
```

**V4优化项目**:

```yaml
prdVersion: v4
prdSharded: true
prdShardedLocation: docs/prd
architectureVersion: v4
architectureSharded: true
architectureShardedLocation: docs/architecture
```

## 核心哲学(Core Philosophy)

### 氛围CEO管理(Vibe CEO'ing)

您是"氛围CEO" - 像拥有无限资源和单一愿景的CEO一样思考。您的AI代理是您的高能团队，您的角色是：

- **指导(Direct)**: 提供清晰的指示和目标
- **精炼(Refine)**: 迭代输出以实现质量
- **监督(Oversee)**: 维护所有代理的战略一致性

### 核心原则(Core Principles)

1. **最大化AI杠杆(MAXIMIZE_AI_LEVERAGE)**: 推动AI交付更多。挑战输出并迭代。
2. **质量控制(QUALITY_CONTROL)**: 您是质量的最终仲裁者。审查所有输出。
3. **战略监督(STRATEGIC_OVERSIGHT)**: 维护高级愿景并确保一致性。
4. **迭代精炼(ITERATIVE_REFINEMENT)**: 期望重新访问步骤。这不是线性过程。
5. **清晰指示(CLEAR_INSTRUCTIONS)**: 精确的请求导致更好的输出。
6. **文档是关键(DOCUMENTATION_IS_KEY)**: 好的输入(简报、PRD)导致好的输出。
7. **从小规模快速开始(START_SMALL_SCALE_FAST)**: 测试概念，然后扩展。
8. **拥抱混乱(EMBRACE_THE_CHAOS)**: 适应并克服挑战。

### 关键工作流程原则(Key Workflow Principles)

1. **代理专业化**: 每个代理都有特定的专业知识和责任
2. **清洁交接**: 在代理间切换时始终重新开始
3. **状态跟踪**: 维护故事状态(Draft → Approved → InProgress → Done)
4. **迭代开发**: 在开始下一个之前完成一个故事
5. **文档优先**: 始终从坚实的PRD和架构开始

## 代理系统(Agent System)

### 核心开发团队(Core Development Team)

| 代理(Agent) | 角色(Role)                         | 主要功能(Primary Functions) | 何时使用(When to Use)  |
| ----------- | ---------------------------------- | --------------------------- | ---------------------- |
| `analyst`   | 业务分析师(Business Analyst)       | 市场研究、需求收集          | 项目规划、竞争分析     |
| `pm`        | 产品经理(Product Manager)          | PRD创建、功能优先级排序     | 战略规划、路线图       |
| `architect` | 解决方案架构师(Solution Architect) | 系统设计、技术架构          | 复杂系统、可扩展性规划 |
| `dev`       | 开发者(Developer)                  | 代码实施、调试              | 所有开发任务           |
| `qa`        | QA专家(QA Specialist)              | 测试规划、质量保证          | 测试策略、bug验证      |
| `ux-expert` | UX设计师(UX Designer)              | UI/UX设计、原型             | 用户体验、界面设计     |
| `po`        | 产品负责人(Product Owner)          | 待办事项管理、故事验证      | 故事精炼、验收标准     |
| `sm`        | Scrum Master                       | 冲刺规划、故事创建          | 项目管理、工作流程     |

### 元代理(Meta Agents)

| 代理(Agent)         | 角色(Role)                   | 主要功能(Primary Functions) | 何时使用(When to Use) |
| ------------------- | ---------------------------- | --------------------------- | --------------------- |
| `bmad-orchestrator` | 团队协调员(Team Coordinator) | 多代理工作流程、角色切换    | 复杂的多角色任务      |
| `bmad-master`       | 通用专家(Universal Expert)   | 无需切换的所有功能          | 单会话综合工作        |

### 代理交互命令(Agent Interaction Commands)

#### IDE特定语法(IDE-Specific Syntax)

**按IDE加载代理**:

- **Claude Code**: `/agent-name` (例如，`/bmad-master`)
- **Cursor**: `@agent-name` (例如，`@bmad-master`)
- **Windsurf**: `@agent-name` (例如，`@bmad-master`)
- **Trae**: `@agent-name` (例如，`@bmad-master`)
- **Roo Code**: 从模式选择器选择模式 (例如，`bmad-master`)
- **GitHub Copilot**: 打开聊天视图(Mac上`⌃⌘I`，Windows/Linux上`Ctrl+Alt+I`)并从聊天模式选择器中选择**代理**。

**聊天管理指南**:

- **Claude Code、Cursor、Windsurf、Trae**: 切换代理时开始新聊天
- **Roo Code**: 在同一对话中切换模式

**常见任务命令**:

- `*help` - 显示可用命令
- `*status` - 显示当前上下文/进度
- `*exit` - 退出代理模式
- `*shard-doc docs/prd.md prd` - 将PRD分片为可管理部分
- `*shard-doc docs/architecture.md architecture` - 分片架构文档
- `*create` - 运行create-next-story任务(SM代理)

**在Web UI中**:

```text
/pm create-doc prd
/architect review system design
/dev implement story 1.2
/help - 显示可用命令
/switch agent-name - 更改活动代理(如果orchestrator可用)
```

## 团队配置(Team Configurations)

### 预构建团队(Pre-Built Teams)

#### 全团队(Team All)

- **包含**: 所有10个代理 + orchestrator
- **用例**: 需要所有角色的完整项目
- **包**: `team-all.txt`

#### 全栈团队(Team Fullstack)

- **包含**: PM、Architect、Developer、QA、UX Expert
- **用例**: 端到端Web/移动开发
- **包**: `team-fullstack.txt`

#### 无UI团队(Team No-UI)

- **包含**: PM、Architect、Developer、QA(无UX Expert)
- **用例**: 后端服务、API、系统开发
- **包**: `team-no-ui.txt`

## 核心架构(Core Architecture)

### 系统概述(System Overview)

BMad-Method围绕以`bmad-core`目录为中心的模块化架构构建，该目录作为整个系统的大脑。此设计使框架能够在IDE环境(如Cursor、VS Code)和基于Web的AI界面(如ChatGPT、Gemini)中有效运行。

### 关键架构组件(Key Architectural Components)

#### 1. 代理(`bmad-core/agents/`)

- **目的**: 每个markdown文件为特定的敏捷角色(PM、Dev、Architect等)定义专门的AI代理
- **结构**: 包含指定代理角色、功能和依赖的YAML头部
- **依赖**: 代理可以使用的任务、模板、检查清单和数据文件列表
- **启动说明**: 可以加载项目特定文档以立即获得上下文

#### 2. 代理团队(`bmad-core/agent-teams/`)

- **目的**: 定义为特定目的捆绑在一起的代理集合
- **示例**: `team-all.yaml`(综合包)、`team-fullstack.yaml`(全栈开发)
- **用法**: 为Web UI环境创建预打包上下文

#### 3. 工作流程(`bmad-core/workflows/`)

- **目的**: 为特定项目类型定义规定步骤序列的YAML文件
- **类型**: Greenfield(新项目)和Brownfield(现有项目)用于UI、服务和全栈开发
- **结构**: 定义代理交互、创建的工件和转换条件

#### 4. 可重用资源(Reusable Resources)

- **模板**(`bmad-core/templates/`): PRD、架构规范、用户故事的Markdown模板
- **任务**(`bmad-core/tasks/`): 特定可重复操作的说明，如"shard-doc"或"create-next-story"
- **检查清单**(`bmad-core/checklists/`): 用于验证和审查的质量保证检查清单
- **数据**(`bmad-core/data/`): 核心知识库和技术偏好

### 双环境架构(Dual Environment Architecture)

#### IDE环境(IDE Environment)

- 用户直接与代理markdown文件交互
- 代理可以动态访问所有依赖
- 支持实时文件操作和项目集成
- 针对开发工作流程执行进行优化

#### Web UI环境(Web UI Environment)

- 使用来自`dist/teams`的预构建包作为独立的1个上传文件，包含所有代理及其资产，以及一个协调代理
- `dist/agents/`中的单个文本文件包含所有代理依赖 - 除非您想创建一个只是单个代理而不是团队的Web代理，否则这些是不必要的
- 由web-builder工具创建，用于上传到Web界面
- 在一个包中提供完整上下文

### 模板处理系统(Template Processing System)

BMad采用复杂的模板系统，具有三个关键组件：

1. **模板格式**(`utils/bmad-doc-template.md`): 定义用于变量替换和AI处理指令的标记语言，来自yaml模板
2. **文档创建**(`tasks/create-doc.md`): 协调模板选择和用户交互，将yaml规范转换为最终markdown输出
3. **高级需求获取**(`tasks/advanced-elicitation.md`): 通过结构化头脑风暴提供交互式精炼

### 技术偏好集成(Technical Preferences Integration)

`technical-preferences.md`文件作为持久技术配置文件：

- 确保所有代理和项目的一致性
- 消除重复的技术规范
- 提供与用户偏好一致的个人化建议
- 随着经验教训而发展

### 构建和交付过程(Build and Delivery Process)

`web-builder.js`工具通过以下方式创建Web就绪包：

1. 读取代理或团队定义文件
2. 递归解析所有依赖
3. 将内容连接成具有清晰分隔符的单个文本文件
4. 输出准备上传到Web AI界面的包

此架构使跨环境无缝操作成为可能，同时维护使BMad强大的丰富、相互关联的代理生态系统。

## 完整开发工作流程(Complete Development Workflow)

### 规划阶段(推荐Web UI - 特别是Gemini！)(Planning Phase (Web UI Recommended - Especially Gemini!))

**特别适合Gemini巨大上下文的成本效益**:

**对于Brownfield项目 - 从这里开始！**:

1. **将整个项目上传到Gemini Web**(GitHub URL、文件或zip)
2. **记录现有系统**: `/analyst` → `*document-project`
3. **从整个代码库分析创建综合文档**

**对于所有项目**:

1. **可选分析**: `/analyst` - 市场研究、竞争分析
2. **项目简报**: 创建基础文档(Analyst或用户)
3. **PRD创建**: `/pm create-doc prd` - 综合产品需求
4. **架构设计**: `/architect create-doc architecture` - 技术基础
5. **验证和一致性**: `/po`运行主检查清单以确保文档一致性
6. **文档准备**: 将最终文档复制到项目作为`docs/prd.md`和`docs/architecture.md`

#### 规划提示示例(Example Planning Prompts)

**对于PRD创建**:

```text
"我想构建一个[类型]应用程序，[核心目的]。
帮助我头脑风暴功能并创建综合PRD。"
```

**对于架构设计**:

```text
"基于这个PRD，设计一个可扩展的技术架构，
可以处理[特定需求]。"
```

### 关键转换：Web UI到IDE(Critical Transition: Web UI to IDE)

**一旦规划完成，您必须切换到IDE进行开发**:

- **为什么**: 开发工作流程需要文件操作、实时项目集成和文档分片
- **成本效益**: Web UI对大型文档创建更具成本效益；IDE针对开发任务进行优化
- **必需文件**: 确保`docs/prd.md`和`docs/architecture.md`存在于您的项目中

### IDE开发工作流程(IDE Development Workflow)

**先决条件**: 规划文档必须存在于`docs/`文件夹中

1. **文档分片(关键步骤)**:
    - 由PM/Architect创建的文档(在Web或IDE中)必须为开发进行分片
    - 两种分片方法：
      a) **手动**: 将`shard-doc`任务+文档文件拖入聊天
      b) **代理**: 要求`@bmad-master`或`@po`分片文档
    - 分片`docs/prd.md` → `docs/prd/`文件夹
    - 分片`docs/architecture.md` → `docs/architecture/`文件夹
    - **警告**: 不要在Web UI中分片 - 复制许多小文件很痛苦！

2. **验证分片内容**:
    - `docs/prd/`中至少有一个`epic-n.md`文件，故事按开发顺序排列
    - 源代码树文档和编码标准供dev代理参考
    - 分片文档供SM代理故事创建

结果文件夹结构：

- `docs/prd/` - 分解的PRD部分
- `docs/architecture/` - 分解的架构部分
- `docs/stories/` - 生成的用户故事

1. **开发周期**(顺序，一次一个故事):

    **关键上下文管理**:
    - **上下文窗口重要！** 始终使用新鲜、清洁的上下文窗口
    - **模型选择重要！** 为SM故事创建使用最强大的思考模型
    - **始终在SM、Dev和QA工作之间开始新聊天**

    **步骤1 - 故事创建**:
    - **新的清洁聊天** → 选择强大模型 → `@sm` → `*create`
    - SM执行create-next-story任务
    - 审查`docs/stories/`中生成的故事
    - 将状态从"Draft"更新为"Approved"

    **步骤2 - 故事实施**:
    - **新的清洁聊天** → `@dev`
    - 代理询问要实施哪个故事
    - 包含故事文件内容以节省dev代理查找时间
    - Dev遵循任务/子任务，标记完成
    - Dev维护所有更改的文件列表
    - Dev在所有测试通过时标记故事为"Review"

    **步骤3 - 高级QA审查**:
    - **新的清洁聊天** → `@qa` → 执行review-story任务
    - QA执行高级开发者代码审查
    - QA可以直接重构和改进代码
    - QA将结果附加到故事的QA结果部分
    - 如果批准：状态 → "Done"
    - 如果需要更改：状态保持"Review"，dev有未检查项目

    **步骤4 - 重复**: 继续SM → Dev → QA周期，直到所有epic故事完成

**重要**: 一次只有1个故事进行中，顺序工作直到所有epic故事完成。

### 状态跟踪工作流程(Status Tracking Workflow)

故事通过定义的状态进展：

- **Draft** → **Approved** → **InProgress** → **Done**

每个状态更改需要用户验证和批准才能继续。

### 工作流程类型(Workflow Types)

#### Greenfield开发

- 业务分析和市场研究
- 产品需求和功能定义
- 系统架构和设计
- 开发执行
- 测试和部署

#### Brownfield增强(现有项目)

**关键概念**: Brownfield开发需要全面记录您的现有项目，以便AI代理理解上下文、模式和约束。

**完整Brownfield工作流程选项**:

**选项1：PRD优先(推荐用于大型代码库/单体仓库)**:

1. **将项目上传到Gemini Web**(GitHub URL、文件或zip)
2. **首先创建PRD**: `@pm` → `*create-doc brownfield-prd`
3. **专注文档**: `@analyst` → `*document-project`
    - 如果没有提供PRD，Analyst会要求专注
    - 为Web UI选择"单一文档"格式
    - 使用PRD仅记录相关区域
    - 创建一个综合markdown文件
    - 避免用未使用的代码膨胀文档

**选项2：文档优先(适合较小项目)**:

1. **将项目上传到Gemini Web**
2. **记录一切**: `@analyst` → `*document-project`
3. **然后创建PRD**: `@pm` → `*create-doc brownfield-prd`
    - 更彻底但可能创建过多文档

4. **需求收集**:
    - **Brownfield PRD**: 使用PM代理与`brownfield-prd-tmpl`
    - **分析**: 现有系统、约束、集成点
    - **定义**: 增强范围、兼容性要求、风险评估
    - **创建**: 变更的Epic和故事结构

5. **架构规划**:
    - **Brownfield Architecture**: 使用Architect代理与`brownfield-architecture-tmpl`
    - **集成策略**: 新功能如何与现有系统集成
    - **迁移规划**: 逐步推出和向后兼容性
    - **风险缓解**: 解决潜在的破坏性变更

**Brownfield特定资源**:

**模板**:

- `brownfield-prd-tmpl.md`: 具有现有系统分析的综合增强规划
- `brownfield-architecture-tmpl.md`: 针对现有系统的集成聚焦架构

**任务**:

- `document-project`: 从现有代码库生成综合文档
- `brownfield-create-epic`: 为专注增强创建单个epic(当完整PRD过度时)
- `brownfield-create-story`: 为小的、孤立的变更创建单个故事

**何时使用每种方法**:

**完整Brownfield工作流程**(推荐用于):

- 主要功能添加
- 系统现代化
- 复杂集成
- 多个相关变更

**快速Epic/故事创建**(当使用):

- 单一、专注的增强
- 孤立的bug修复
- 小功能添加
- 文档完善的现有系统

**关键成功因素**:

1. **文档优先**: 如果文档过时/缺失，始终运行`document-project`
2. **上下文重要**: 为代理提供相关代码部分的访问权限
3. **集成聚焦**: 强调兼容性和非破坏性变更
4. **增量方法**: 计划逐步推出和测试

**详细指南**: 参见`docs/working-in-the-brownfield.md`

## 文档创建最佳实践(Document Creation Best Practices)

### 框架集成的必需文件命名(Required File Naming for Framework Integration)

- `docs/prd.md` - 产品需求文档
- `docs/architecture.md` - 系统架构文档

**为什么这些名称重要**:

- 代理在开发过程中自动引用这些文件
- 分片任务期望这些特定文件名
- 工作流程自动化依赖于标准命名

### 成本效益文档创建工作流程(Cost-Effective Document Creation Workflow)

**推荐用于大型文档(PRD、Architecture):**

1. **使用Web UI**: 在Web界面中创建文档以获得成本效益
2. **复制最终输出**: 将完整markdown保存到您的项目
3. **标准名称**: 保存为`docs/prd.md`和`docs/architecture.md`
4. **切换到IDE**: 使用IDE代理进行开发和较小文档

### 文档分片(Document Sharding)

具有Level 2标题(`##`)的模板可以自动分片：

**原始PRD**:

```markdown
## Goals and Background Context

## Requirements

## User Interface Design Goals

## Success Metrics
```

**分片后**:

- `docs/prd/goals-and-background-context.md`
- `docs/prd/requirements.md`
- `docs/prd/user-interface-design-goals.md`
- `docs/prd/success-metrics.md`

使用`shard-doc`任务或`@kayvan/markdown-tree-parser`工具进行自动分片。

## 使用模式和最佳实践(Usage Patterns and Best Practices)

### 环境特定使用(Environment-Specific Usage)

**Web UI最适合**:

- 初始规划和文档阶段
- 成本效益的大型文档创建
- 代理咨询和头脑风暴
- 与orchestrator的多代理工作流程

**IDE最适合**:

- 主动开发和实施
- 文件操作和项目集成
- 故事管理和开发周期
- 代码审查和调试

### 质量保证(Quality Assurance)

- 为专门任务使用适当的代理
- 遵循敏捷仪式和审查过程
- 与PO代理维护文档一致性
- 使用检查清单和模板进行定期验证

### 性能优化(Performance Optimization)

- 对专注任务使用特定代理vs `bmad-master`
- 为项目需求选择适当的团队规模
- 利用技术偏好实现一致性
- 定期上下文管理和缓存清理

## 成功提示(Success Tips)

- **使用Gemini进行大图规划** - team-fullstack包提供协作专业知识
- **使用bmad-master进行文档组织** - 分片创建可管理的块
- **严格遵循SM → Dev周期** - 这确保系统性进展
- **保持对话专注** - 一个代理，每次对话一个任务
- **审查一切** - 在标记完成之前始终审查和批准

## 为BMad-Method贡献(Contributing to BMad-Method)

### 快速贡献指南(Quick Contribution Guidelines)

完整详情，参见`CONTRIBUTING.md`。要点：

**Fork工作流程**:

1. Fork仓库
2. 创建功能分支
3. 向`next`分支(默认)或`main`提交PR(仅限关键修复)
4. 保持PR小：200-400行理想，800行最大
5. 每个PR一个功能/修复

**PR要求**:

- 清晰描述(最多200字)，包含What/Why/How/Testing
- 使用约定提交(feat:、fix:、docs:)
- 原子提交 - 每次提交一个逻辑变更
- 必须与指导原则一致

**核心原则**(来自docs/GUIDING-PRINCIPLES.md):

- **Dev代理必须精简**: 最小化依赖，为代码保存上下文
- **自然语言优先**: 核心中一切都是markdown，无代码
- **核心vs扩展包**: 核心用于通用需求，包用于专门领域
- **设计哲学**: "Dev代理编码，规划代理规划"

## 扩展包(Expansion Packs)

### 什么是扩展包？(What Are Expansion Packs?)

扩展包将BMad-Method扩展到传统软件开发之外的任何领域。它们提供专门的代理团队、模板和工作流程，同时保持核心框架精简并专注于开发。

### 为什么使用扩展包？(Why Use Expansion Packs?)

1. **保持核心精简**: Dev代理保持最大编码上下文
2. **领域专业知识**: 深度、专门知识而不膨胀核心
3. **社区创新**: 任何人都可以创建和分享包
4. **模块化设计**: 仅安装您需要的

### 可用扩展包(Available Expansion Packs)

**技术包**:

- **基础设施/DevOps**: 云架构师、SRE专家、安全专家
- **游戏开发**: 游戏设计师、关卡设计师、叙事作家
- **移动开发**: iOS/Android专家、移动UX专家
- **数据科学**: ML工程师、数据科学家、可视化专家

**非技术包**:

- **商业策略**: 顾问、财务分析师、营销策略师
- **创意写作**: 情节架构师、角色开发者、世界构建者
- **健康与福祉**: 健身教练、营养师、习惯工程师
- **教育**: 课程设计师、评估专家
- **法律支持**: 合同分析师、合规检查员

**专业包**:

- **扩展创建者**: 构建您自己的扩展包的工具
- **RPG游戏大师**: 桌面游戏协助
- **生活事件规划**: 婚礼策划师、活动协调员
- **科学研究**: 文献审查员、方法论设计师

### 使用扩展包(Using Expansion Packs)

1. **浏览可用包**: 检查`expansion-packs/`目录
2. **获得灵感**: 参见`docs/expansion-packs.md`获取详细示例和想法
3. **通过CLI安装**:

    ```bash
    npx bmad-method install
    # 选择"安装扩展包"选项
    ```

4. **在您的工作流程中使用**: 安装的包与现有代理无缝集成

### 创建自定义扩展包(Creating Custom Expansion Packs)

使用**expansion-creator**包构建您自己的：

1. **定义领域**: 您要捕获什么专业知识？
2. **设计代理**: 创建具有清晰边界的专门角色
3. **构建资源**: 为您的领域的任务、模板、检查清单
4. **测试和分享**: 用真实用例验证，与社区分享

**关键原则**: 扩展包通过使专门知识通过AI代理可访问来民主化专业知识。

## 获取帮助(Getting Help)

- **命令**: 在任何环境中使用`*/*help`查看可用命令
- **代理切换**: 使用`*/*switch agent-name`与orchestrator进行角色更改
- **文档**: 检查`docs/`文件夹获取项目特定上下文
- **社区**: Discord和GitHub资源可用于支持
- **贡献**: 参见`CONTRIBUTING.md`获取完整指南
