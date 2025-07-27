# 在 Brownfield 中工作：完整指南(Working in the Brownfield: A Complete Guide)

> **强烈推荐：使用 Gemini Web 或 Gemini CLI 进行 Brownfield 文档生成！**
>
> Gemini Web 的 100万+ token 上下文窗口或 Gemini CLI（当它工作时）可以一次性分析您的整个代码库或其中的关键部分（显然在合理范围内）：
>
> - 通过 GitHub URL 上传或在项目文件夹中使用 gemini cli
> - 如果在 web 中：上传最多 1000 个文件或压缩项目或直接提供 github url

## 什么是 Brownfield 开发？(What is Brownfield Development?)

Brownfield 开发是指向现有软件项目添加功能、修复错误或现代化。与 greenfield（新）项目不同，brownfield 工作需要理解现有代码、尊重约束，并确保新更改无缝集成而不破坏现有功能。

## 何时使用 BMad 进行 Brownfield(When to Use BMad for Brownfield)

- 向现有应用程序添加重要新功能
- 现代化遗留代码库
- 集成新技术或服务
- 重构复杂系统
- 修复需要架构理解的错误
- 记录未记录的系统

## 何时不使用 Brownfield 流程(When NOT to use a Brownfield Flow)

如果您刚刚使用 BMad 完成了 MVP，并且想要继续 post-MVP，更容易只是与 PM 交谈，要求他与您合作创建一个新的 epic 添加到 PRD 中，分片 epic，与 architect 更新任何架构文档，然后从那里开始。

## 完整的 Brownfield 工作流(The Complete Brownfield Workflow)

### 选择您的方法(Choose Your Approach)

#### 方法 A：PRD 优先（如果添加非常大和复杂的新功能、单个或多个 epics 或大规模更改，推荐）(Approach A: PRD-First (Recommended if adding very large and complex new features, single or multiple epics or massive changes))

**最适合**：大型代码库、monorepos，或当您确切知道要构建什么时

1. **首先创建 PRD** 来定义需求
2. **仅记录相关区域** 基于 PRD 需求
3. **更高效** - 避免记录未使用的代码

#### 方法 B：文档优先（适合较小项目）(Approach B: Document-First (Good for Smaller Projects))

**最适合**：较小的代码库、未知系统或探索性更改

1. **首先记录整个系统**
2. **创建具有完整上下文的 PRD**
3. **更彻底** - 捕获一切

### 方法 A：PRD 优先工作流（推荐）(Approach A: PRD-First Workflow (Recommended))

#### 阶段 1：首先定义需求(Phase 1: Define Requirements First)

**在 Gemini Web 中（上传您的代码库）：**

```bash
@pm
*create-doc brownfield-prd
```

PM 将：

- **询问您的增强**需求
- **探索代码库**以了解当前状态
- **识别需要文档的受影响区域**
- **创建具有明确范围的专注 PRD**

**关键优势**：PRD 识别您的 monorepo/大型代码库中实际需要文档的部分！

#### 阶段 2：专注文档(Phase 2: Focused Documentation)

**仍在 Gemini Web 中，现在具有 PRD 上下文：**

```bash
@analyst
*document-project
```

分析师将：

- **如果没有提供 PRD，询问您的重点**
- **提供选项**：创建 PRD、提供需求或描述增强
- **参考 PRD/描述**以理解范围
- **专注于 PRD 或您的描述中识别的相关模块**
- **跳过不相关区域**以保持文档精简
- **为所有环境生成一个架构文档**

分析师创建：

- **一个综合架构文档**，遵循 fullstack-architecture 模板
- **涵盖所有系统方面**在单个文件中
- **易于复制和保存**为 `docs/project-architecture.md`
- **如果需要在 IDE 中稍后可以分片**

例如，如果您说"向用户服务添加支付处理"：

- 仅记录：用户服务、API 端点、数据库模式、支付集成
- 创建专注的源树，仅显示支付相关代码路径
- 跳过：管理面板、报告模块、不相关的微服务

### 方法 B：文档优先工作流(Approach B: Document-First Workflow)

#### 阶段 1：记录现有系统(Phase 1: Document the Existing System)

**最佳方法 - 具有 100万+ 上下文的 Gemini Web**：

1. **转到 Gemini Web**（gemini.google.com）
2. **上传您的项目**：
    - **选项 A**：直接粘贴您的 GitHub 仓库 URL
    - **选项 B**：从您的 src/project 文件夹上传最多 1000 个文件
    - **选项 C**：压缩您的项目并上传存档
3. **加载分析师 agent**：上传 `dist/agents/analyst.txt`
4. **运行文档**：输入 `*document-project`

分析师将生成一切的全面文档。

#### 阶段 2：规划您的增强(Phase 2: Plan Your Enhancement)

#### 选项 A：完整 Brownfield 工作流（推荐用于重大更改）(Option A: Full Brownfield Workflow (Recommended for Major Changes))

**1. 创建 Brownfield PRD**：

```bash
@pm
*create-doc brownfield-prd
```

PM agent 将：

- **分析阶段 1 的现有文档**
- **向您请求具体的增强细节**
- **评估复杂性**并推荐方法
- **为增强创建 epic/story 结构**
- **识别风险和集成点**

**PM Agent 如何获取项目上下文**：

- 在 Gemini Web 中：已经从阶段 1 文档中具有完整的项目上下文
- 在 IDE 中：将询问"请提供您现有项目文档的路径"

**您将遇到的关键提示**：

- "您想添加什么具体的增强或功能？"
- "这需要与任何现有系统或 API 集成吗？"
- "我们必须尊重的关键约束是什么？"
- "您的时间线和团队规模是什么？"

**2. 创建 Brownfield 架构**：

```bash
@architect
*create-doc brownfield-architecture
```

架构师将：

- **审查 brownfield PRD**
- **设计集成策略**
- **规划迁移方法**（如果需要）
- **识别技术风险**
- **定义兼容性要求**

#### 选项 B：快速增强（用于专注更改）(Option B: Quick Enhancement (For Focused Changes))

**对于没有完整 PRD 的单个 Epic**：

```bash
@pm
*brownfield-create-epic
```

当以下情况时使用：

- 增强定义明确且隔离
- 现有文档全面
- 更改不影响多个系统
- 您需要快速周转

**对于单个 Story**：

```bash
@pm
*brownfield-create-story
```

当以下情况时使用：

- 错误修复或小功能
- 非常隔离的更改
- 没有架构影响
- 明确的实现路径

### 阶段 3：验证规划工件(Phase 3: Validate Planning Artifacts)

```bash
@po
*execute-checklist po-master-checklist
```

PO 确保：

- 与现有系统兼容
- 没有计划破坏性更改
- 风险缓解策略到位
- 明确的集成方法

### 阶段 4：过渡到开发(Phase 4: Transition to Development)

遵循增强的 IDE 开发工作流：

1. **确保文档在项目中**：
    - 复制 `docs/prd.md`（或 brownfield-prd.md）
    - 复制 `docs/architecture.md`（或 brownfield-architecture.md）

2. **分片文档**：

    ```bash
    @po
    # 请求分片 docs/prd.md
    ```

3. **开发周期**：
    - **SM** 创建具有集成意识的故事
    - **Dev** 实现时尊重现有代码
    - **QA** 审查兼容性和改进

## Brownfield 最佳实践(Brownfield Best Practices)

### 1. 始终先记录(Always Document First)

即使您认为您了解代码库：

- 运行 `document-project` 捕获当前状态
- AI agents 需要这个上下文
- 发现未记录的模式

### 2. 尊重现有模式(Respect Existing Patterns)

brownfield 模板专门寻找：

- 当前编码约定
- 现有架构模式
- 技术约束
- 团队偏好

### 3. 规划逐步推出(Plan for Gradual Rollout)

Brownfield 更改应该：

- 支持功能标志
- 规划回滚策略
- 包括迁移脚本
- 保持向后兼容性

### 4. 彻底测试集成(Test Integration Thoroughly)

专注于测试：

- 集成点
- 现有功能（回归）
- 性能影响
- 数据迁移

### 5. 沟通更改(Communicate Changes)

记录：

- 更改了什么以及为什么
- 迁移说明
- 引入的新模式
- 弃用通知

## 常见 Brownfield 场景(Common Brownfield Scenarios)

### 场景 1：添加新功能(Scenario 1: Adding a New Feature)

1. 记录现有系统
2. 创建专注于集成的 brownfield PRD
3. 架构强调兼容性
4. 故事包括集成任务

### 场景 2：现代化遗留代码(Scenario 2: Modernizing Legacy Code)

1. 广泛的文档阶段
2. PRD 包括迁移策略
3. 架构规划逐步过渡
4. 故事遵循绞杀无花果模式

### 场景 3：复杂系统中的错误修复(Scenario 3: Bug Fix in Complex System)

1. 记录相关子系统
2. 使用 `brownfield-create-story` 进行专注修复
3. 包括回归测试要求
4. QA 验证无副作用

### 场景 4：API 集成(Scenario 4: API Integration)

1. 记录现有 API 模式
2. PRD 定义集成需求
3. 架构确保一致模式
4. 故事包括 API 文档更新

## 故障排除(Troubleshooting)

### "AI 不理解我的代码库"

**解决方案**：使用更具体的路径到关键文件重新运行 `document-project`

### "生成的计划不适合我们的模式"

**解决方案**：在规划阶段之前用您的特定约定更新生成的文档

### "小更改的样板太多"

**解决方案**：使用 `brownfield-create-story` 而不是完整工作流

### "集成点不明确"

**解决方案**：在 PRD 创建期间提供更多上下文，特别突出集成系统

## 快速参考(Quick Reference)

### Brownfield 特定命令(Brownfield-Specific Commands)

```bash
# 记录现有项目
@analyst → *document-project

# 创建增强 PRD
@pm → *create-doc brownfield-prd

# 创建具有集成重点的架构
@architect → *create-doc brownfield-architecture

# 快速 epic 创建
@pm → *brownfield-create-epic

# 单个 story 创建
@pm → *brownfield-create-story
```

### 决策树(Decision Tree)

```text
您有大型代码库或 monorepo 吗？
├─ 是 → PRD 优先方法
│   └─ 创建 PRD → 仅记录受影响区域
└─ 否 → 代码库对您来说是否熟悉？
    ├─ 是 → PRD 优先方法
    └─ 否 → 文档优先方法

这是影响多个系统的重大增强吗？
├─ 是 → 完整 Brownfield 工作流
└─ 否 → 这比简单的错误修复更多吗？
    ├─ 是 → brownfield-create-epic
    └─ 否 → brownfield-create-story
```

## 结论(Conclusion)

使用 BMad-Method 的 Brownfield 开发在修改现有系统时提供结构和安全性。关键是通过文档提供全面的上下文，使用考虑集成需求的专门模板，并遵循尊重现有约束同时实现进展的工作流。

记住：**先记录，仔细规划，安全集成**
