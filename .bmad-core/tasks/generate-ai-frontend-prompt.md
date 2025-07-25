# 创建 AI 前端提示任务

## 目的

生成一个精湛、全面且优化的提示，可用于任何 AI 驱动的前端开发工具（例如 Vercel v0、Lovable.ai 或类似工具），以搭建或生成前端应用程序的重要部分。

## 输入

- 已完成的 UI/UX 规范 (`front-end-spec.md`)
- 已完成的前端架构文档 (`front-end-architecture`) 或全栈组合架构（例如 `architecture.md`）
- 主系统架构文档 (`architecture` - 用于 API 契约和技术栈以提供进一步上下文)

## 关键活动和说明

### 1. 核心提示原则

在生成提示之前，您必须了解与生成式 AI 进行代码交互的这些核心原则。

- **明确和详细**：AI 无法读取您的思想。提供尽可能多的细节和上下文。模糊的请求会导致通用或不正确的输出。
- **迭代，不要期望完美**：一次性生成整个复杂应用程序的情况很少见。最有效的方法是每次提示一个组件或一个部分，然后在此基础上进行构建。
- **首先提供上下文**：始终首先向 AI 提供必要的上下文，例如技术栈、现有代码片段和整体项目目标。
- **移动优先方法**：以移动优先的设计理念来构建所有 UI 生成请求。首先描述移动布局，然后提供单独的说明，说明它应如何适应平板电脑和桌面。

### 2. 结构化提示框架

为确保最高质量的输出，您必须使用以下四部分框架来构建每个提示。

1. **高层目标**：以清晰、简洁的整体目标摘要开始。这使 AI 专注于主要任务。
   - _示例：“创建一个具有客户端验证和 API 集成的响应式用户注册表单。”_
2. **详细的、分步说明**：提供一个粒度化的、编号的 AI 应采取的操作列表。将复杂任务分解为更小、顺序的步骤。这是提示最关键的部分。
   - _示例：“1. 创建一个名为 `RegistrationForm.js` 的新文件。2. 使用 React hooks 进行状态管理。3. 为“姓名”、“电子邮件”和“密码”添加样式输入字段。4. 对于电子邮件字段，确保它是有效的电子邮件格式。5. 提交时，调用下面定义的 API 端点。”_
3. **代码示例、数据结构和约束**：包含任何相关的现有代码片段、数据结构或 API 契约。这为 AI 提供了具体的工作示例。至关重要的是，您还必须说明_不_要做什么。
   - _示例：“使用此 API 端点：`POST /api/register`。预期的 JSON 有效负载是 `{ "name": "string", "email": "string", "password": "string" }`。不要包含“确认密码”字段。所有样式都使用 Tailwind CSS。”_
4. **定义严格范围**：明确定义任务的边界。告诉 AI 它可以修改哪些文件，更重要的是，哪些文件要保持不变，以防止代码库中出现意外更改。
   - _示例：“您应该只创建 `RegistrationForm.js` 组件并将其添加到 `pages/register.js` 文件中。不要更改 `Navbar.js` 组件或任何其他现有页面或组件。”_

### 3. 组装主提示

您现在将综合输入和上述原则，形成最终的、全面的提示。

1. **收集基础上下文**：
   - 以描述整体项目目的、完整技术栈（例如 Next.js、TypeScript、Tailwind CSS）和正在使用的主要 UI 组件库的引言开始提示。
2. **描述视觉效果**：
   - 如果用户有设计文件（Figma 等），指示他们提供链接或屏幕截图。
   - 如果没有，描述视觉风格：调色板、排版、间距和整体美学（例如，“极简主义”、“企业”、“俏皮”）。
3. **使用结构化框架构建提示**：
   - 遵循第 2 节中的四部分框架来构建核心请求，无论是针对单个组件还是整个页面。
4. **呈现和完善**：
   - 以清晰、可复制粘贴的格式（例如，一个大的代码块）输出完整生成的提示。
   - 解释提示的结构以及为什么包含某些信息，并引用上述原则。
   - <important_note>最后提醒用户，所有 AI 生成的代码都需要仔细的人工审查、测试和完善，才能被视为可用于生产。</important_note>
