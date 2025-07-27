# 创建 AI 前端提示任务(Create AI Frontend Prompt Task)

## 目的(Purpose)

生成可用于任何 AI 驱动前端开发工具（如 Vercel v0、Lovable.ai 或类似工具）以搭建或生成前端应用主要部分的高质量、全面且优化的提示。

## 输入(Inputs)

- 已完成的 UI/UX 规范（`front-end-spec.md`）
- 已完成的前端架构文档（`front-end-architecture`）或全栈合并架构如 `architecture.md`
- 主系统架构文档（`architecture` - 用于 API 合同和技术栈以提供更多上下文）

## 关键活动与指令(Key Activities & Instructions)

### 1. 核心提示原则(Core Prompting Principles)

在生成提示前，必须理解与生成式 AI 交互的这些核心原则。

- **明确且详细**：AI 无法读心。请尽可能提供详细和丰富的上下文。模糊的请求会导致通用或错误的输出。
- **迭代，不要期望一次完美**：一次性生成完整复杂应用很少见。最有效的方法是每次提示一个组件或一个部分，然后在结果基础上迭代。
- **先提供上下文**：始终先向 AI 提供必要的上下文，如技术栈、已有代码片段和整体项目目标。
- **移动优先**：所有 UI 生成请求都应以移动端优先的设计思路描述。先描述移动端布局，再分别说明如何适配平板和桌面。

### 2. 结构化提示框架(The Structured Prompting Framework)

为确保最高质量输出，必须使用以下四部分结构化框架组织每个提示。

1. **高层目标(High-Level Goal)**：以清晰简明的总结开头，说明整体目标。这有助于 AI 明确主要任务。
    - _示例："创建一个带有客户端校验和 API 集成的响应式用户注册表单。"_
2. **详细分步指令(Detailed, Step-by-Step Instructions)**：提供细致、编号的操作步骤。将复杂任务拆解为更小的顺序步骤。这是提示中最关键的部分。
    - _示例："1. 新建文件 `RegistrationForm.js`。2. 使用 React hooks 管理状态。3. 添加 'Name'、'Email' 和 'Password' 的样式输入框。4. 对 email 字段进行格式校验。5. 提交时调用下方定义的 API 接口。"_
3. **代码示例、数据结构与约束(Code Examples, Data Structures & Constraints)**：包含相关的已有代码片段、数据结构或 API 合同。这为 AI 提供具体参考。务必说明不允许做什么。
    - _示例："使用此 API 接口：`POST /api/register`。期望的 JSON 负载为 `{ "name": "string", "email": "string", "password": "string" }`。不要包含 'confirm password' 字段。所有样式使用 Tailwind CSS。"_
4. **严格限定范围(Define a Strict Scope)**：明确界定任务边界。告诉 AI 只能修改哪些文件，更重要的是哪些文件不能动，以防止对代码库造成意外更改。
    - _示例："你只应创建 `RegistrationForm.js` 组件并将其添加到 `pages/register.js` 文件。不要更改 `Navbar.js` 组件或任何其他现有页面或组件。"_

### 3. 组装主提示(Assembling the Master Prompt)

现在将综合输入和上述原则，生成最终全面的提示。

1. **收集基础上下文(Gather Foundational Context)**：
    - 以前言开头，描述整体项目目标、完整技术栈（如 Next.js、TypeScript、Tailwind CSS）及主要 UI 组件库。
2. **描述视觉效果(Describe the Visuals)**：
    - 如果用户有设计文件（Figma 等），请指示其提供链接或截图。
    - 如果没有，请描述视觉风格：色彩方案、字体、间距和整体美学（如“极简”、“企业”、“活泼”）。
3. **使用结构化框架构建提示(Build the Prompt using the Structured Framework)**：
    - 按第 2 节的四部分框架构建核心请求，无论是单个组件还是完整页面。
4. **展示与完善(Present and Refine)**：
    - 以清晰、可复制粘贴的格式（如大代码块）输出完整生成的提示。
    - 解释提示结构及为何包含某些信息，并引用上述原则。
    - <important_note>最后提醒用户，所有 AI 生成的代码都需要经过仔细人工审查、测试和完善，才能用于生产环境。</important_note>
