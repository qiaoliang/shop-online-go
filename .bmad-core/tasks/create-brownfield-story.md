# 创建 Brownfield 故事任务(Create Brownfield Story Task)

## 目的(Purpose)

为传统分片 PRD/架构文档可能不存在的 brownfield 项目创建详细、可实施的故事。此任务弥合各种文档格式（document-project 输出、brownfield PRD、epic 或用户文档）与开发代理可执行故事之间的差距。

## 何时使用此任务(When to Use This Task)

**在以下情况使用此任务：**

- 处理具有非标准文档的 brownfield 项目
- 需要从 document-project 输出创建故事
- 从没有完整 PRD/架构的 brownfield epic 工作
- 现有项目文档不遵循 BMad v4+ 结构
- 在故事创建期间需要从用户收集额外上下文

**在以下情况使用 create-next-story：**

- 使用正确分片的 PRD 和 v4 架构文档工作
- 遵循标准 greenfield 或文档完善的 brownfield 工作流程
- 所有技术上下文都以结构化格式提供

## 任务执行指令(Task Execution Instructions)

### 0. 文档上下文(Documentation Context)

按此顺序检查可用文档：

1. **分片 PRD/架构**（docs/prd/、docs/architecture/）
    - 如找到，建议使用 create-next-story 任务

2. **Brownfield 架构文档**（docs/brownfield-architecture.md 或类似）
    - 由 document-project 任务创建
    - 包含实际系统状态、技术债务、workaround

3. **Brownfield PRD**（docs/prd.md）
    - 可能包含嵌入的技术详情

4. **Epic 文件**（docs/epics/ 或类似）
    - 由 brownfield-create-epic 任务创建

5. **用户提供的文档**
    - 请用户指定位置和格式

### 1. 故事识别和上下文收集(Story Identification and Context Gathering)

#### 1.1 识别故事来源(Identify Story Source)

基于可用文档：

- **来自 Brownfield PRD**：从 epic 章节提取故事
- **来自 Epic 文件**：读取 epic 定义和故事列表
- **来自用户指导**：询问用户要实施哪个具体增强
- **无明确来源**：与用户合作定义故事范围

#### 1.2 收集基本上下文(Gather Essential Context)

**关键：** 对于 brownfield 故事，您必须收集足够的上下文以确保安全实施。准备向用户询问缺失信息。

**所需信息清单：**

- [ ] 哪些现有功能可能受影响？
- [ ] 与当前代码的集成点是什么？
- [ ] 应遵循哪些模式（带示例）？
- [ ] 存在哪些技术约束？
- [ ] 有哪些需要了解的"坑"或 workaround？

如任何必需信息缺失，列出缺失信息并请用户提供。

### 2. 从可用来源提取技术上下文(Extract Technical Context from Available Sources)

#### 2.1 从 Document-Project 输出(From Document-Project Output)

如使用来自 document-project 的 brownfield-architecture.md：

- **技术债务章节**：注意影响此故事的任何 workaround
- **关键文件章节**：识别需要修改的文件
- **集成点**：查找现有集成模式
- **已知问题**：检查故事是否触及问题区域
- **实际技术栈**：验证版本和约束

#### 2.2 从 Brownfield PRD(From Brownfield PRD)

如使用 brownfield PRD：

- **技术约束章节**：提取所有相关约束
- **集成要求**：记录兼容性要求
- **代码组织**：遵循指定模式
- **风险评估**：理解潜在影响

#### 2.3 从用户文档(From User Documentation)

请用户帮助识别：

- 相关技术规范
- 要遵循的现有代码示例
- 集成要求
- 项目中使用的测试方法

### 3. 带渐进详情收集的故事创建(Story Creation with Progressive Detail Gathering)

#### 3.1 创建初始故事结构(Create Initial Story Structure)

从故事模板开始，填写已知内容：

```markdown
# 故事 {{增强标题}}

## 状态：草稿

## 故事

作为 {{用户类型}}，
我想要 {{增强能力}}，
以便 {{交付价值}}。

## 上下文来源

- 源文档：{{文档名称/类型}}
- 增强类型：{{单一特性/bug 修复/集成/等}}
- 现有系统影响：{{简要评估}}
```

#### 3.2 开发验收标准(Develop Acceptance Criteria)

**关键：** 对于 brownfield，始终包括关于维护现有功能的标准

标准结构：

1. 新功能按指定工作
2. 现有 {{受影响功能}} 继续正常工作不变
3. 与 {{现有系统}} 的集成保持当前行为
4. {{相关区域}} 无回归
5. 性能保持在可接受范围内

#### 3.3 收集技术指导(Gather Technical Guidance)

**关键：** 这是您需要与用户互动的地方，如果信息缺失

使用可用信息创建开发技术指导章节：

```markdown
## 开发技术指导

### 现有系统上下文

[从可用文档提取]

### 集成方法

[基于找到的模式或询问用户]

### 技术约束

[来自文档或用户输入]

### 缺失信息

关键：列出开发需要的任何找不到的内容，并询问缺失信息
```

### 4. 带安全检查的任务生成(Task Generation with Safety Checks)

#### 4.1 生成实施任务(Generate Implementation Tasks)

基于收集的上下文，创建任务：

- 如系统理解不完整，包括探索任务
- 为现有功能添加验证任务
- 包括回滚考虑
- 如已知，引用具体文件/模式

Brownfield 示例任务结构：

```markdown
## 任务/子任务

- [ ] 任务 1：分析现有 {{组件/功能}} 实施
    - [ ] 审查 {{具体文件}} 的当前模式
    - [ ] 记录集成点
    - [ ] 识别潜在影响

- [ ] 任务 2：实施 {{新功能}}
    - [ ] 遵循 {{示例文件}} 的模式
    - [ ] 与 {{现有组件}} 集成
    - [ ] 保持与 {{约束}} 的兼容性

- [ ] 任务 3：验证现有功能
    - [ ] 测试 {{现有功能 1}} 仍工作
    - [ ] 验证 {{集成点}} 行为不变
    - [ ] 检查性能影响

- [ ] 任务 4：添加测试
    - [ ] 遵循 {{项目测试模式}} 的单元测试
    - [ ] {{集成点}} 的集成测试
    - [ ] 如需要更新现有测试
```

### 5. 风险评估和缓解(Risk Assessment and Mitigation)

**关键：** 对于 brownfield - 始终包括风险评估

为 brownfield 特定风险添加章节：

```markdown
## 风险评估

### 实施风险

- **主要风险**：{{对现有系统的主要风险}}
- **缓解**：{{如何解决}}
- **验证**：{{如何确认安全}}

### 回滚计划

- {{如需要撤销变更的简单步骤}}

### 安全检查

- [ ] 变更前测试现有 {{功能}}
- [ ] 变更可功能标记或隔离
- [ ] 回滚程序已记录
```

### 6. 最终故事验证(Final Story Validation)

最终确定前：

1. **完整性检查**：
    - [ ] 故事有清晰范围和验收标准
    - [ ] 技术上下文足以实施
    - [ ] 集成方法已定义
    - [ ] 风险已识别并有缓解

2. **安全检查**：
    - [ ] 包括现有功能保护
    - [ ] 回滚计划可行
    - [ ] 测试涵盖新功能和现有功能

3. **信息差距**：
    - [ ] 从用户收集了所有关键缺失信息
    - [ ] 剩余未知内容为开发代理记录
    - [ ] 如需要添加探索任务

### 7. 故事输出格式(Story Output Format)

使用适当命名保存故事：

- 如来自 epic：`docs/stories/epic-{n}-story-{m}.md`
- 如独立：`docs/stories/brownfield-{feature-name}.md`
- 如顺序：遵循现有故事编号

包括注明文档上下文的标题：

```markdown
# 故事：{{标题}}

<!-- 来源：{{使用的文档类型}} -->
<!-- 上下文：对 {{现有系统}} 的 Brownfield 增强 -->

## 状态：草稿

[故事内容其余部分...]
```

### 8. 交接沟通(Handoff Communication)

向用户提供清晰交接：

```text
Brownfield 故事已创建：{{故事标题}}

源文档：{{使用的内容}}
故事位置：{{文件路径}}

识别的关键集成点：
- {{集成点 1}}
- {{集成点 2}}

注意的风险：
- {{主要风险}}

{{如缺失信息}}：
注意：一些技术详情不清楚。故事包括探索任务，在实施期间收集所需信息。

下一步：
1. 审查故事准确性
2. 验证集成方法与您的系统一致
3. 批准故事或请求调整
4. 开发代理然后可实施安全检查
```

## 成功标准(Success Criteria)

Brownfield 故事创建成功时：

1. 故事可实施，无需开发搜索多个文档
2. 集成方法清晰且对现有系统安全
3. 所有可用技术上下文已提取并组织
4. 缺失信息已识别并解决
5. 风险已记录并有缓解策略
6. 故事包括现有功能验证
7. 回滚方法已定义

## 重要说明(Important Notes)

- 此任务专门用于具有非标准文档的 brownfield 项目
- 始终优先考虑现有系统稳定性而非新功能
- 有疑问时，添加探索和验证任务
- 询问用户澄清比做假设更好
- 每个故事对开发代理都应是自包含的
- 如可用，包括对现有代码模式的引用
