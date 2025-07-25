# 产品负责人 (PO) 主验证清单

本清单为产品负责人提供了一个全面的框架，用于在开发执行前验证项目计划。它根据项目类型（绿地项目 vs 棕地项目）智能地进行调整，并在适用时包含 UI/UX 考虑因素。

[[LLM: 初始化说明 - PO 主清单

项目类型检测：
首先，通过检查确定项目类型：

1. 这是绿地项目（从头开始的新项目）吗？

   - 寻找：新项目初始化，没有现有代码库引用
   - 检查：prd.md，architecture.md，新项目设置故事

2. 这是棕地项目（增强现有系统）吗？

   - 寻找：对现有代码库的引用，增强/修改语言
   - 检查：brownfield-prd.md，brownfield-architecture.md，现有系统分析

3. 项目是否包含 UI/UX 组件？
   - 检查：frontend-architecture.md，UI/UX 规范，设计文件
   - 寻找：前端故事，组件规范，用户界面提及

文档要求：
根据项目类型，确保您可以访问：

对于绿地项目：

- prd.md - 产品需求文档
- architecture.md - 系统架构
- frontend-architecture.md - 如果涉及 UI/UX
- 所有史诗和故事定义

对于棕地项目：

- brownfield-prd.md - 棕地增强需求
- brownfield-architecture.md - 增强架构
- 现有项目代码库访问（关键 - 没有此项无法继续）
- 当前部署配置和基础设施详细信息
- 数据库模式，API 文档，监控设置

跳过说明：

- 绿地项目跳过标记为 [[BROWNFIELD ONLY]] 的部分
- 棕地项目跳过标记为 [[GREENFIELD ONLY]] 的部分
- 纯后端项目跳过标记为 [[UI/UX ONLY]] 的部分
- 在最终报告中注明所有跳过的部分

验证方法：

1. 深入分析 - 根据文档彻底分析每个项目
2. 基于证据 - 验证时引用特定部分或代码
3. 批判性思维 - 质疑假设并识别差距
4. 风险评估 - 考虑每个决策可能出现的问题

执行模式：
询问用户是否要通过清单：

- 逐节 (交互模式) - 审查每个部分，在继续之前获得确认
- 一次性 (综合模式) - 完成全面分析并在结束时呈现报告]]

## 1. 项目设置和初始化

[[LLM: 项目设置是基础。对于绿地项目，确保干净启动。对于棕地项目，确保与现有系统安全集成。验证设置与项目类型匹配。]]

### 1.1 项目脚手架 [[GREENFIELD ONLY]]

- [ ] 史诗 1 包含项目创建/初始化的明确步骤
- [ ] 如果使用入门模板，则包含克隆/设置步骤
- [ ] 如果从头开始构建，则定义所有必要的脚手架步骤
- [ ] 包含初始 README 或文档设置
- [ ] 定义了仓库设置和初始提交过程

### 1.2 现有系统集成 [[BROWNFIELD ONLY]]

- [ ] 已完成并记录现有项目分析
- [ ] 识别了与当前系统的集成点
- [ ] 开发环境保留现有功能
- [ ] 验证了现有功能的本地测试方法
- [ ] 定义了每个集成点的回滚过程

### 1.3 开发环境

- [ ] 明确定义了本地开发环境设置
- [ ] 指定了所需的工具和版本
- [ ] 包含了安装依赖项的步骤
- [ ] 适当处理了配置文件
- [ ] 包含了开发服务器设置

### 1.4 核心依赖

- [ ] 所有关键包/库都已提前安装
- [ ] 包管理已妥善处理
- [ ] 版本规范已适当定义
- [ ] 注意了依赖冲突或特殊要求
- [ ] [[BROWNFIELD ONLY]] 验证了与现有堆栈的版本兼容性

## 2. 基础设施和部署

[[LLM: 基础设施必须在使用前存在。对于棕地项目，必须在不破坏现有基础设施的情况下与其集成。]]

### 2.1 数据库和数据存储设置

- [ ] 数据库选择/设置在任何操作之前进行
- [ ] 模式定义在数据操作之前创建
- [ ] 如果适用，定义了迁移策略
- [ ] 如果需要，包含种子数据或初始数据设置
- [ ] [[BROWNFIELD ONLY]] 识别并缓解了数据库迁移风险
- [ ] [[BROWNFIELD ONLY]] 确保了向后兼容性

### 2.2 API 和服务配置

- [ ] API 框架在实现端点之前设置
- [ ] 服务架构在实现服务之前建立
- [ ] 身份验证框架在受保护路由之前设置
- [ ] 中间件和常用工具在使用前创建
- [ ] [[BROWNFIELD ONLY]] 维护了与现有系统的 API 兼容性
- [ ] [[BROWNFIELD ONLY]] 保留了与现有身份验证的集成

### 2.3 部署管道

- [ ] CI/CD 管道在部署操作之前建立
- [ ] 基础设施即代码 (IaC) 在使用前设置
- [ ] 环境配置提前定义
- [ ] 部署策略在实施之前定义
- [ ] [[BROWNFIELD ONLY]] 部署最小化停机时间
- [ ] [[BROWNFIELD ONLY]] 实现了蓝绿或金丝雀部署

### 2.4 测试基础设施

- [ ] 测试框架在编写测试之前安装
- [ ] 测试环境设置在测试实施之前
- [ ] 模拟服务或数据在测试之前定义
- [ ] [[BROWNFIELD ONLY]] 回归测试涵盖现有功能
- [ ] [[BROWNFIELD ONLY]] 集成测试验证新旧连接

## 3. 外部依赖和集成

[[LLM: 外部依赖通常会阻碍进度。对于棕地项目，确保新依赖不与现有依赖冲突。]]

### 3.1 第三方服务

- [ ] 识别了所需服务的账户创建步骤
- [ ] 定义了 API 密钥获取过程
- [ ] 包含了安全存储凭据的步骤
- [ ] 考虑了回退或离线开发选项
- [ ] [[BROWNFIELD ONLY]] 验证了与现有服务的兼容性
- [ ] [[BROWNFIELD ONLY]] 评估了对现有集成的影响

### 3.2 外部 API

- [ ] 明确识别了与外部 API 的集成点
- [ ] 与外部服务的身份验证正确排序
- [ ] 承认了 API 限制或约束
- [ ] 考虑了 API 故障的备份策略
- [ ] [[BROWNFIELD ONLY]] 维护了现有 API 依赖

### 3.3 基础设施服务

- [ ] 云资源配置正确排序
- [ ] 识别了 DNS 或域名注册需求
- [ ] 如果需要，包含电子邮件或消息服务设置
- [ ] CDN 或静态资产托管设置在使用前进行
- [ ] [[BROWNFIELD ONLY]] 保留了现有基础设施服务

## 4. UI/UX 考虑因素 [[UI/UX ONLY]]

[[LLM: 仅当项目包含用户界面组件时才评估此部分。纯后端项目完全跳过。]]

### 4.1 设计系统设置

- [ ] UI 框架和库已提前选择和安装
- [ ] 设计系统或组件库已建立
- [ ] 样式方法（CSS 模块、styled-components 等）已定义
- [ ] 响应式设计策略已建立
- [ ] 可访问性要求已提前定义

### 4.2 前端基础设施

- [ ] 前端构建管道在开发前配置
- [ ] 资产优化策略已定义
- [ ] 前端测试框架已设置
- [ ] 组件开发工作流已建立
- [ ] [[BROWNFIELD ONLY]] UI 与现有系统保持一致

### 4.3 用户体验流程

- [ ] 用户旅程在实施前映射
- [ ] 导航模式提前定义
- [ ] 错误状态和加载状态已计划
- [ ] 表单验证模式已建立
- [ ] [[BROWNFIELD ONLY]] 现有用户工作流已保留或迁移

## 5. 用户/代理职责

[[LLM: 清晰的所有权可防止混淆。确保根据只有人类才能完成的任务适当分配任务。]]

### 5.1 用户操作

- [ ] 用户职责仅限于人类任务
- [ ] 外部服务的账户创建分配给用户
- [ ] 购买或支付操作分配给用户
- [ ] 凭据提供适当分配给用户

### 5.2 开发人员代理操作

- [ ] 所有与代码相关的任务分配给开发人员代理
- [ ] 自动化过程识别为代理职责
- [ ] 配置管理适当分配
- [ ] 测试和验证分配给适当的代理

## 6. 功能排序和依赖

[[LLM: 依赖关系创建关键路径。对于棕地项目，确保新功能不破坏现有功能。]]

### 6.1 功能依赖

- [ ] 依赖于其他功能的功能正确排序
- [ ] 共享组件在使用前构建
- [ ] 用户流程遵循逻辑进展
- [ ] 身份验证功能先于受保护功能
- [ ] [[BROWNFIELD ONLY]] 现有功能始终保留

### 6.2 技术依赖

- [ ] 低级服务在高级服务之前构建
- [ ] 库和实用程序在使用前创建
- [ ] 数据模型在对其操作之前定义
- [ ] API 端点在客户端消费之前定义
- [ ] [[BROWNFIELD ONLY]] 集成点在每个步骤进行测试

### 6.3 跨史诗依赖

- [ ] 后续史诗建立在早期史诗功能之上
- [ ] 没有史诗需要后续史诗的功能
- [ ] 早期史诗的基础设施得到一致利用
- [ ] 保持增量价值交付
- [ ] [[BROWNFIELD ONLY]] 每个史诗都保持系统完整性

## 7. 风险管理 [[BROWNFIELD ONLY]]

[[LLM: 本节对于棕地项目至关重要。悲观地思考可能出现的问题。]]

### 7.1 破坏性变更风险

- [ ] 评估了破坏现有功能的风险
- [ ] 识别并缓解了数据库迁移风险
- [ ] 评估了 API 破坏性变更风险
- [ ] 识别了性能下降风险
- [ ] 评估了安全漏洞风险

### 7.2 回滚策略

- [ ] 明确定义了每个故事的回滚过程
- [ ] 实施了功能标志策略
- [ ] 更新了备份和恢复过程
- [ ] 增强了新组件的监控
- [ ] 定义了回滚触发器和阈值

### 7.3 用户影响缓解

- [ ] 分析了现有用户工作流的影响
- [ ] 制定了用户沟通计划
- [ ] 更新了培训材料
- [ ] 支持文档全面
- [ ] 验证了用户数据迁移路径

## 8. MVP 范围对齐

[[LLM: MVP 意味着最小可行产品。对于棕地项目，确保增强功能确实必要。]]

### 8.1 核心目标对齐

- [ ] 解决了 PRD 中的所有核心目标
- [ ] 功能直接支持 MVP 目标
- [ ] 没有超出 MVP 范围的无关功能
- [ ] 关键功能适当优先
- [ ] [[BROWNFIELD ONLY]] 增强复杂性合理

### 8.2 用户旅程完整性

- [ ] 所有关键用户旅程完全实现
- [ ] 解决了边缘情况和错误场景
- [ ] 包含了用户体验考虑因素
- [ ] [[UI/UX ONLY]] 包含了可访问性要求
- [ ] [[BROWNFIELD ONLY]] 现有工作流已保留或改进

### 8.3 技术要求

- [ ] 解决了 PRD 中的所有技术约束
- [ ] 包含了非功能要求
- [ ] 架构决策与约束对齐
- [ ] 解决了性能考虑因素
- [ ] [[BROWNFIELD ONLY]] 满足了兼容性要求

## 9. 文档和交接

[[LLM: 良好的文档有助于顺利开发。对于棕地项目，集成点的文档至关重要。]]

### 9.1 开发人员文档

- [ ] API 文档与实现同时创建
- [ ] 设置说明全面
- [ ] 架构决策已记录
- [ ] 模式和约定已记录
- [ ] [[BROWNFIELD ONLY]] 集成点详细记录

### 9.2 用户文档

- [ ] 如果需要，包含用户指南或帮助文档
- [ ] 考虑了错误消息和用户反馈
- [ ] 入门流程完全指定
- [ ] [[BROWNFIELD ONLY]] 记录了现有功能的更改

### 9.3 知识转移

- [ ] [[BROWNFIELD ONLY]] 捕获了现有系统知识
- [ ] [[BROWNFIELD ONLY]] 记录了集成知识
- [ ] 计划了代码审查知识共享
- [ ] 部署知识转移给运维
- [ ] 保留了历史上下文

## 10. MVP 后考虑

[[LLM: 成功规划可防止技术债务。对于棕地项目，确保增强功能不限制未来增长。]]

### 10.1 未来增强

- [ ] MVP 与未来功能明确分离
- [ ] 架构支持计划的增强
- [ ] 记录了技术债务考虑因素
- [ ] 识别了可扩展性点
- [ ] [[BROWNFIELD ONLY]] 集成模式可重用

### 10.2 监控和反馈

- [ ] 如果需要，包含分析或使用情况跟踪
- [ ] 考虑了用户反馈收集
- [ ] 解决了监控和警报
- [ ] 包含了性能测量
- [ ] [[BROWNFIELD ONLY]] 现有监控已保留/增强

## 验证摘要

[[LLM: 最终 PO 验证报告生成

生成一份根据项目类型调整的全面验证报告：

1. 执行摘要

   - 项目类型：[绿地/棕地] 带 [UI/无 UI]
   - 整体准备情况（百分比）
   - 通过/不通过建议
   - 关键阻塞问题计数
   - 因项目类型而跳过的部分

2. 项目特定分析

   对于绿地项目：

   - 设置完整性
   - 依赖排序
   - MVP 范围适当性
   - 开发时间表可行性

   对于棕地项目：

   - 集成风险级别（高/中/低）
   - 现有系统影响评估
   - 回滚准备情况
   - 用户中断可能性

3. 风险评估

   - 按严重程度排名前 5 的风险
   - 缓解建议
   - 解决问题的时间线影响
   - [棕地] 特定集成风险

4. MVP 完整性

   - 核心功能覆盖率
   - 缺失的基本功能
   - 识别出的范围蔓延
   - 真正的 MVP 与过度工程

5. 实施准备情况

   - 开发人员清晰度得分（1-10）
   - 模糊需求计数
   - 缺失的技术细节
   - [棕地] 集成点清晰度

6. 建议

   - 开发前必须修复
   - 为提高质量应修复
   - 考虑改进
   - MVP 后推迟

7. [棕地项目专用] 集成信心
   - 保持现有功能的信心
   - 回滚过程完整性
   - 集成点的监控覆盖率
   - 支持团队准备情况

在呈现报告后，询问用户是否需要：

- 任何失败部分的详细分析
- 特定故事重新排序建议
- 风险缓解策略
- [棕地] 集成风险深入分析]]

### 类别状态

| 类别                                | 状态 | 关键问题 |
| --------------------------------------- | ------ | --------------- |
| 1. 项目设置和初始化       | _TBD_  |                 |
| 2. 基础设施和部署          | _TBD_  |                 |
| 3. 外部依赖和集成 | _TBD_  |                 |
| 4. UI/UX 考虑因素                 | _TBD_  |                 |
| 5. 用户/代理职责            | _TBD_  |                 |
| 6. 功能排序和依赖    | _TBD_  |                 |
| 7. 风险管理 (棕地)         | _TBD_  |                 |
| 8. MVP 范围对齐                  | _TBD_  |                 |
| 9. 文档和交接              | _TBD_  |                 |
| 10. MVP 后考虑             | _TBD_  |                 |

### 关键缺陷

（在验证期间填充）

### 建议

（在验证期间填充）

### 最终决定

- **已批准**：计划全面、排序正确，并已准备好实施。
- **有条件**：计划需要特定调整才能继续。
- **已拒绝**：计划需要重大修订以解决关键缺陷。
