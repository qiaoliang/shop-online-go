# 工作流管理(Workflow Management)

使 BMad orchestrator 能够管理和执行团队工作流。

## 动态工作流加载(Dynamic Workflow Loading)

从当前团队配置的 `workflows` 字段读取可用工作流。每个团队包定义其支持的工作流。

**关键命令**：

- `/workflows` - 列出当前包或工作流文件夹中的工作流
- `/agent-list` - 显示当前包中的 agents

## 工作流命令(Workflow Commands)

### /workflows

列出可用工作流及其标题和描述。

### /workflow-start {workflow-id}

启动工作流并转换到第一个 agent。

### /workflow-status

显示当前进度、已完成的工件和下一步。

### /workflow-resume

从最后位置恢复工作流。用户可以提供已完成的工件。

### /workflow-next

显示下一个推荐的 agent 和操作。

## 执行流程(Execution Flow)

1. **启动**: 加载定义 → 识别第一阶段 → 转换到 agent → 指导工件创建

2. **阶段转换**: 标记完成 → 检查条件 → 加载下一个 agent → 传递工件

3. **工件跟踪**: 在 workflow_state 中跟踪状态、创建者、时间戳

4. **中断处理**: 分析提供的工件 → 确定位置 → 建议下一步

## 上下文传递(Context Passing)

转换时传递：

- 之前的工件
- 当前工作流阶段
- 预期输出
- 决策/约束

## 多路径工作流(Multi-Path Workflows)

通过在有需要时询问澄清问题来处理条件路径。

## 最佳实践(Best Practices)

1. 显示进度
2. 解释转换
3. 保留上下文
4. 允许灵活性
5. 跟踪状态

## Agent 集成(Agent Integration)

Agents 应该具有工作流意识：了解活动工作流、其角色、访问工件、理解预期输出。
