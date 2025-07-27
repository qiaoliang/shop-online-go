# 记录现有项目(Document an Existing Project)

## 目的(Purpose)

为现有项目生成面向 AI 开发代理优化的全面文档。本任务创建结构化参考资料，使 AI 代理能够理解项目上下文、约定和模式，从而有效参与任何代码库。

## 任务指令(Task Instructions)

### 1. 初步项目分析(Initial Project Analysis)

**关键：** 首先检查上下文中是否存在 PRD, README.md 或需求文档。如果有，仅针对相关领域聚焦文档编写。

**如存在 PRD：**

- 审查 PRD，了解计划增强/特性的内容
- 确定将受影响的模块、服务或区域
- 仅针对这些相关领域编写文档
- 跳过与代码库无关的部分，保持文档精炼

**如无 PRD：**
请询问用户：

“我注意到你没有提供 PRD 或需求文档。为了创建更聚焦、更有用的文档，建议如下选项：

1. **先创建 PRD** - 需要我先帮你创建该遗留项目的 PRD 吗？这样文档能更聚焦相关领域。

2. **提供现有需求** - 你有需求文档、epic 或特性描述可以分享吗？

3. **描述关注点** - 能否简要描述你计划的增强或特性？例如：
    - '为用户服务添加支付处理'
    - '重构认证模块'
    - '集成新的第三方 API'

4. **全部文档化** - 还是需要我对整个代码库做全面文档？（注意：大项目可能会生成过多文档）

请告知你的偏好，或如需我直接全面文档也可。”

根据回复：

- 如选 1-3：用该上下文聚焦文档
- 如选 4 或拒绝：按下述全面分析继续

首先分析现有项目。用可用工具：

1. **项目结构发现**：检查根目录结构，识别主要文件夹，理解整体组织
2. **技术栈识别**：查找 package.json、requirements.txt、Cargo.toml、pom.xml 等，识别语言、框架和依赖
3. **构建系统分析**：查找构建脚本、CI/CD 配置和开发命令
4. **现有文档审查**：检查 README、docs 文件夹及其他文档
5. **代码模式分析**：抽样关键文件，了解编码模式、命名约定和架构方法

如果 `docs/startpoint`目录存在，则直接参考其中的所有内容。并开始执行下一步：Deep Codebase Analysis
否则，向用户提出以下引导性问题以更好理解需求：

- 该项目的主要目标是什么？
- 代码库中有哪些特别复杂或对代理理解尤为重要的区域？
- 你期望 AI 代理在本项目中执行哪些任务？（如修 bug、加特性、重构、测试）
- 有无现有文档标准或格式偏好？
- 文档应面向什么技术水平？（初级/高级/混合团队）
- 是否有具体计划的特性或增强？（有助于聚焦文档）

### 2. 深度代码库分析(Deep Codebase Analysis)

**关键：** 生成文档前，需对现有代码库做深入分析：

1. **探索关键区域**：
    - 入口点（主文件、index、app 初始化）
    - 配置文件和环境设置
    - 包依赖及版本
    - 构建与部署配置
    - 测试套件与覆盖率

2. **澄清提问**：
    - “我看到你用 [技术 X]，有无自定义模式或约定需记录？”
    - “系统中哪些部分最关键/复杂，开发者最难理解？”
    - “有无未文档化的‘部落知识’需补充？”
    - “有哪些技术债务或已知问题需记录？”
    - “哪些代码最常变动？”

3. **映射实际情况**：
    - 识别实际用到的模式（非理论最佳实践）
    - 找到关键业务逻辑位置
    - 定位集成点和外部依赖
    - 记录 workaround 和技术债务
    - 标注与标准模式不同的区域

**如有 PRD：** 还需分析增强所需变更内容

### 3. 核心文档生成(Core Documentation Generation)

[[LLM: 生成反映代码库实际状态的全面 BROWNFIELD 架构文档。

**关键：** 这不是理想化的架构文档，要记录实际存在的内容，包括：

- 技术债务和 workaround
- 不同部分间的不一致模式
- 不能更改的遗留代码
- 集成约束
- 性能瓶颈

**文档结构：**

# [项目名] Brownfield 架构文档

## 简介(Introduction)

本文件记录 [项目名] 代码库的当前状态，包括技术债务、workaround 和真实模式。供 AI 代理做增强时参考。

### 文档范围(Document Scope)

[如有 PRD：“聚焦于：{enhancement description} 相关区域”]
[如无 PRD：“系统全量文档”]

### 变更日志(Change Log)

| Date   | Version | Description          | Author    |
| ------ | ------- | -------------------- | --------- |
| [Date] | 1.0     | 初始 brownfield 分析 | [Analyst] |

## 快速参考 - 关键文件与入口(Quick Reference - Key Files and Entry Points)

### 理解系统的关键文件(Critical Files for Understanding the System)

- **主入口**: `src/index.js`（或实际入口）
- **配置**: `config/app.config.js`、`.env.example`
- **核心业务逻辑**: `src/services/`、`src/domain/`
- **API 定义**: `src/routes/` 或 OpenAPI spec 链接
- **数据库模型**: `src/models/` 或 schema 文件链接
- **关键算法**: [列出具体复杂逻辑文件]

### 如有 PRD - 增强影响区域(Enhancement Impact Areas)

[突出哪些文件/模块将受增强影响]

## 高层架构(High Level Architecture)

### 技术摘要(Technical Summary)

### 实际技术栈(Actual Tech Stack)

| Category  | Technology | Version | Notes            |
| --------- | ---------- | ------- | ---------------- |
| Runtime   | Node.js    | 16.x    | [约束]           |
| Framework | Express    | 4.18.2  | [自定义中间件？] |
| Database  | PostgreSQL | 13      | [连接池设置]     |

等...

### 仓库结构现实检查(Repository Structure Reality Check)

- 类型: [Monorepo/Polyrepo/Hybrid]
- 包管理器: [npm/yarn/pnpm]
- 备注: [任何特殊结构决策]

## 源码树与模块组织(Source Tree and Module Organization)

### 项目结构（实际）(Project Structure (Actual))

```text
project-root/
├── src/
│   ├── controllers/     # HTTP 请求处理
│   ├── services/        # 业务逻辑（注意：user 与 payment service 模式不一致）
│   ├── models/          # 数据库模型（Sequelize）
│   ├── utils/           # 杂项 - 需重构
│   └── legacy/          # 不可修改 - 旧支付系统仍在用
├── tests/               # Jest 测试（60% 覆盖）
├── scripts/             # 构建与部署脚本
└── config/              # 环境配置
```

### 关键模块及其用途(Key Modules and Their Purpose)

- **用户管理**: `src/services/userService.js` - 处理所有用户操作
- **认证**: `src/middleware/auth.js` - JWT，自定义实现
- **支付处理**: `src/legacy/payment.js` - 关键：不可重构，强耦合
- **[其他关键模块及实际文件]**

## 数据模型与 API(Data Models and APIs)

### 数据模型(Data Models)

不重复，直接引用实际模型文件：

- **用户模型**: 见 `src/models/User.js`
- **订单模型**: 见 `src/models/Order.js`
- **相关类型**: TypeScript 定义在 `src/types/`

### API 规范(API Specifications)

- **OpenAPI Spec**: `docs/api/openapi.yaml`（如有）
- **Postman Collection**: `docs/api/postman-collection.json`
- **手动接口**: [列出发现的未文档化接口]

## 技术债务与已知问题(Technical Debt and Known Issues)

### 关键技术债务(Critical Technical Debt)

1. **支付服务**: `src/legacy/payment.js` 遗留代码 - 强耦合，无测试
2. **用户服务**: 与其他服务模式不同，使用回调而非 promise
3. **数据库迁移**: 手动跟踪，无正规迁移工具
4. **[其他重要债务]**

### Workaround 与注意事项(Workarounds and Gotchas)

- **环境变量**: staging 也必须设 `NODE_ENV=production`（历史原因）
- **数据库连接**: 连接池硬编码为 10，改动会导致支付服务异常
- **[其他开发需知的 workaround]**

## 集成点与外部依赖(Integration Points and External Dependencies)

### 外部服务(External Services)

| Service  | Purpose  | Integration Type | Key Files                      |
| -------- | -------- | ---------------- | ------------------------------ |
| Stripe   | Payments | REST API         | `src/integrations/stripe/`     |
| SendGrid | Emails   | SDK              | `src/services/emailService.js` |

等...

### 内部集成点(Internal Integration Points)

- **前端通信**: REST API 3000 端口，需特定 header
- **后台任务**: Redis 队列，见 `src/workers/`
- **[其他集成]**

## 开发与部署(Development and Deployment)

### 本地开发设置(Local Development Setup)

1. 实际可用步骤（非理想步骤）
2. 已知 setup 问题
3. 所需环境变量（见 `.env.example`）

### 构建与部署流程(Build and Deployment Process)

- **构建命令**: `npm run build`（webpack 配置见 `webpack.config.js`）
- **部署**: 手动执行 `scripts/deploy.sh`
- **环境**: Dev、Staging、Prod（见 `config/environments/`）

## 测试现状(Testing Reality)

### 当前测试覆盖(Current Test Coverage)

- 单元测试：60% 覆盖（Jest）
- 集成测试：极少，见 `tests/integration/`
- E2E 测试：无
- 手动测试：主 QA 方式

### 运行测试(Running Tests)

```bash
npm test           # 跑单元测试
npm run test:integration  # 跑集成测试（需本地 DB）
```

## 如有增强 PRD - 影响分析(If Enhancement PRD Provided - Impact Analysis)

### 需修改的文件(Files That Will Need Modification)

根据增强需求，这些文件将受影响：

- `src/services/userService.js` - 新增用户字段
- `src/models/User.js` - 更新 schema
- `src/routes/userRoutes.js` - 新接口
- [等...]

### 需新增的文件/模块(New Files/Modules Needed)

- `src/services/newFeatureService.js` - 新业务逻辑
- `src/models/NewFeature.js` - 新数据模型
- [等...]

### 集成注意事项(Integration Considerations)

- 需集成现有认证中间件
- 必须遵循 `src/utils/responseFormatter.js` 的响应格式
- [其他集成点]

## 附录 - 常用命令与脚本(Appendix - Useful Commands and Scripts)

### 常用命令(Frequently Used Commands)

```bash
npm run dev         # 启动开发服务器
npm run build       # 生产构建
npm run migrate     # 跑数据库迁移
npm run seed        # 填充测试数据
```

### 调试与故障排查(Debugging and Troubleshooting)

- **日志**: 查看 `logs/app.log`
- **调试模式**: 设 `DEBUG=app:*` 得到详细日志
- **常见问题**: 见 `docs/troubleshooting.md`]]

### 4. 文档交付(Document Delivery)

1. **Web UI（Gemini、ChatGPT、Claude）**：
    - 一次性输出完整文档（如太长可分多次）
    - 告知用户复制保存为 `docs/brownfield-architecture.md` 或 `docs/project-architecture.md`
    - 可在 IDE 后续分片

2. **IDE 环境**：
    - 创建文档为 `docs/brownfield-architecture.md`
    - 告知用户该文档包含所有架构信息
    - 可后续用 PO agent 分片

文档应足够全面，使后续代理能理解：

- 系统实际状态（非理想化）
- 关键文件和逻辑位置
- 存在的技术债务
- 必须遵守的约束
- 如有 PRD：需变更内容

### 5. 质量保证(Quality Assurance)

**关键：** 定稿前：

1. **准确性检查**：核实所有技术细节与实际代码库一致
2. **完整性审查**：确保所有主要系统组件均有文档
3. **聚焦验证**：如用户提供范围，确保相关领域突出
4. **清晰度评估**：确保 AI 代理易于理解
5. **导航性**：确保文档结构清晰便于查阅

主章节后可用高级引导任务根据用户反馈优化。

## 成功标准(Success Criteria)

- 生成单一全面的 brownfield 架构文档
- 文档反映现实，包括技术债务和 workaround
- 关键文件和模块引用实际路径
- 模型/API 直接引用源文件而非重复内容
- 如有 PRD：有清晰影响分析
- 文档使 AI 代理能导航和理解实际代码库
- 技术约束和“坑”有清晰记录

## 备注(Notes)

- 本任务只生成一份反映系统真实状态的文档
- 能引用实际文件时不重复内容
- 诚实记录技术债务、workaround 和约束
- 对 brownfield 项目有 PRD 时：提供清晰增强影响分析
- 目标是为 AI 代理做实际工作提供实用文档
