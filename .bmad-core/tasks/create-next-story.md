# 创建下一个故事任务(Create Next Story Task)

## 目的(Purpose)

基于项目进度和 epic 定义识别下一个逻辑故事，然后使用 `故事模板` 准备全面、自包含和可操作的故事文件。此任务确保故事包含所有必要的技术上下文、要求和验收标准，使其准备好由开发代理高效实施，无需额外研究或自行查找上下文。

## 顺序任务执行（在当前任务完成之前不要继续）(SEQUENTIAL Task Execution (Do not proceed until current Task is complete))

### 0. 加载核心配置并检查工作流程(Load Core Configuration and Check Workflow)

- 从项目根目录加载 `.bmad-core/core-config.yaml`
- 如文件不存在，停止并告知用户："未找到 core-config.yaml。故事创建需要此文件。您可以：1) 从 GITHUB bmad-core/core-config.yaml 复制并为您的项目配置它，或 2) 对您的项目运行 BMad 安装程序以自动升级和添加文件。请在继续前添加和配置 core-config.yaml。"
- 提取关键配置：`devStoryLocation`、`prd.*`、`architecture.*`、`workflow.*`

### 1. 识别下一个要准备的故事(Identify Next Story for Preparation)

#### 1.1 定位 Epic 文件并审查现有故事(Locate Epic Files and Review Existing Stories)

- 基于配置中的 `prdSharded`，定位 epic 文件（分片位置/模式或整体 PRD 章节）
- 如 `devStoryLocation` 有故事文件，加载最高的 `{epicNum}.{storyNum}.story.md` 文件
- **如最高故事存在：**
    - 验证状态为'Done'。如不是，提醒用户："警告：发现未完成故事！文件：{lastEpicNum}.{lastStoryNum}.story.md 状态：[当前状态] 您应该先修复此故事，但您是否愿意接受风险并覆盖以创建草稿中的下一个故事？"
    - 如继续，选择当前 epic 中的下一个顺序故事
    - 如 epic 完成，提示用户："Epic {epicNum} 完成：Epic {epicNum} 中的所有故事已完成。您是否想要：1) 以故事 1 开始 Epic {epicNum + 1} 2) 选择特定故事处理 3) 取消故事创建"
    - **关键**：绝不自动跳到另一个 epic。用户必须明确指示创建哪个故事。
- **如无故事文件存在：** 下一个故事始终是 1.1（第一个 epic 的第一个故事）
- 向用户宣布识别的故事："识别下一个要准备的故事：{epicNum}.{storyNum} - {故事标题}"

### 2. 收集故事要求和前一个故事上下文(Gather Story Requirements and Previous Story Context)

- 从识别的 epic 文件提取故事要求
- 如前一个故事存在，审查开发代理记录章节的：
    - 完成说明和调试日志引用
    - 实施偏差和技术决策
    - 遇到的挑战和学到的经验
- 提取为当前故事准备提供信息的相关洞见

### 3. 收集架构上下文(Gather Architecture Context)

#### 3.1 确定架构阅读策略(Determine Architecture Reading Strategy)

- **如 `architectureVersion: >= v4` 且 `architectureSharded: true`**：读取 `{architectureShardedLocation}/index.md` 然后按下面结构化阅读顺序
- **否则**：使用整体 `architectureFile` 的类似章节

#### 3.2 基于故事类型读取架构文档(Read Architecture Documents Based on Story Type)

**对于所有故事：** tech-stack.md、unified-project-structure.md、coding-standards.md、testing-strategy.md

**对于后端/API 故事，另外：** data-models.md、database-schema.md、backend-architecture.md、rest-api-spec.md、external-apis.md

**对于前端/UI 故事，另外：** frontend-architecture.md、components.md、core-workflows.md、data-models.md

**对于全栈故事：** 读取上述后端和前端章节

#### 3.3 提取故事特定技术详情(Extract Story-Specific Technical Details)

仅提取与实施当前故事直接相关的信息。不要发明源文档中没有的新库、模式或标准。

提取：

- 故事将使用的具体数据模型、schema 或结构
- 故事必须实施或消费的 API 端点
- 故事中 UI 元素的组件规范
- 新代码的文件路径和命名约定
- 故事特性特定的测试要求
- 影响故事的安全或性能考虑

始终引用源文档：`[来源：architecture/{filename}.md#{section}]`

### 4. 验证项目结构对齐(Verify Project Structure Alignment)

- 将故事要求与 `docs/architecture/unified-project-structure.md` 中的项目结构指南交叉引用
- 确保文件路径、组件位置或模块名称与定义的结构对齐
- 在故事草稿的"项目结构说明"章节中记录任何结构冲突

### 5. 用完整上下文填充故事模板(Populate Story Template with Full Context)

- 使用故事模板创建新故事文件：`{devStoryLocation}/{epicNum}.{storyNum}.story.md`
- 填写基本故事信息：标题、状态（草稿）、故事陈述、来自 Epic 的验收标准
- **`开发说明`章节（关键）：**
    - **关键**：此章节必须仅包含从架构文档提取的信息。绝不发明或假设技术详情。
    - 包括来自步骤 2-3 的所有相关技术详情，按类别组织：
        - **前一个故事洞见**：前一个故事的关键学习
        - **数据模型**：具体 schema、验证规则、关系[带源引用]
        - **API 规范**：端点详情、请求/响应格式、认证要求[带源引用]
        - **组件规范**：UI 组件详情、props、状态管理[带源引用]
        - **文件位置**：基于项目结构应创建新代码的确切路径
        - **测试要求**：来自 testing-strategy.md 的具体测试案例或策略
        - **技术约束**：版本要求、性能考虑、安全规则
    - 每个技术详情必须包括其源引用：`[来源：architecture/{filename}.md#{section}]`
    - 如架构文档中未找到类别的信息，明确说明："架构文档中未找到具体指导"
- **`任务/子任务`章节：**
    - 基于以下内容生成详细、顺序的技术任务列表：Epic 要求、故事 AC、审查的架构信息
    - 每个任务必须引用相关架构文档
    - 基于测试策略将单元测试作为明确子任务包括
    - 如适用，将任务链接到 AC（如 `任务 1（AC：1、3）`）
- 添加步骤 4 中发现的项目结构对齐或差异说明

### 6. 故事草稿完成和审查(Story Draft Completion and Review)

- 审查所有章节的完整性和准确性
- 验证所有源引用都包括技术详情
- 确保任务与 epic 要求和架构约束都对齐
- 更新状态为"草稿"并保存故事文件
- 执行 `.bmad-core/tasks/execute-checklist` `.bmad-core/checklists/story-draft-checklist`
- 向用户提供摘要，包括：
    - 创建的故事：`{devStoryLocation}/{epicNum}.{storyNum}.story.md`
    - 状态：草稿
    - 从架构文档包含的关键技术组件
    - epic 和架构间注意到的任何偏差或冲突
    - 清单结果
    - 下一步：对于复杂故事，建议用户仔细审查故事草稿，也可选择让 PO 运行任务 `.bmad-core/tasks/validate-next-story`
