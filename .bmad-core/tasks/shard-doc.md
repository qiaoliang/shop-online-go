# 文档分片任务(Document Sharding Task)

## 目的(Purpose)

- 基于二级章节将大型文档拆分为多个较小的文档
- 创建文件夹结构来组织分片文档
- 保持所有内容完整性，包括代码块、图表和 markdown 格式

## 主要方法：使用 markdown-tree 自动分片(Primary Method: Automatic with markdown-tree)

[[LLM: 首先，检查 .bmad-core/core-config.yaml 中是否将 markdownExploder 设置为 true。如果是，尝试运行命令：`md-tree explode {input file} {output path}`。

如果命令成功，通知用户文档已成功分片并停止 - 不要继续。

如果命令失败（特别是出现命令未找到或不可用的错误），通知用户："markdownExploder 设置已启用但 md-tree 命令不可用。请：

1. 全局安装 @kayvan/markdown-tree-parser：`npm install -g @kayvan/markdown-tree-parser`
2. 或在 .bmad-core/core-config.yaml 中将 markdownExploder 设置为 false

**重要：在此停止 - 在采取上述操作之一之前不要继续手动分片。**"

如果 markdownExploder 设置为 false，通知用户："markdownExploder 设置当前为 false。为了更好的性能和可靠性，您应该：

1. 在 .bmad-core/core-config.yaml 中将 markdownExploder 设置为 true
2. 全局安装 @kayvan/markdown-tree-parser：`npm install -g @kayvan/markdown-tree-parser`

我现在将继续手动分片过程。"

然后仅在 markdownExploder 为 false 时继续下面的手动方法。]]

### 安装和使用(Installation and Usage)

1. **全局安装**：

    ```bash
    npm install -g @kayvan/markdown-tree-parser
    ```

2. **使用 explode 命令**：

    ```bash
    # 对于 PRD
    md-tree explode docs/prd.md docs/prd

    # 对于架构
    md-tree explode docs/architecture.md docs/architecture

    # 对于任何文档
    md-tree explode [源文档] [目标文件夹]
    ```

3. **它的作用**：
    - 自动按二级章节拆分文档
    - 创建正确命名的文件
    - 适当调整标题级别
    - 处理代码块和特殊 markdown 的所有边缘情况

如果用户已安装 @kayvan/markdown-tree-parser，请使用它并跳过下面的手动过程。

---

## 手动方法(Manual Method)（如果 @kayvan/markdown-tree-parser 不可用或用户指示手动方法）

### 任务指令(Task Instructions)

1. 识别文档和目标位置(Identify Document and Target Location)

- 确定要分片的文档（用户提供的路径）
- 在 `docs/` 下创建一个与文档同名的新文件夹（不带扩展名）
- 示例：`docs/prd.md` → 创建文件夹 `docs/prd/`

2. 解析和提取章节(Parse and Extract Sections)

关键代理分片规则(CRITICAL AEGNT SHARDING RULES)：

1. 读取整个文档内容
2. 识别所有二级章节（## 标题）
3. 对于每个二级章节：
    - 提取章节标题和直到下一个二级章节的所有内容
    - 包括所有子章节、代码块、图表、列表、表格等
    - 对以下内容要极其小心：
        - 围栏代码块（```）- 确保捕获完整块，包括结束反引号，并考虑可能误导的二级标题，这些实际上是围栏部分示例的一部分
        - Mermaid 图表 - 保持完整的图表语法
        - 嵌套 markdown 元素
        - 可能包含代码块内 ## 的多行内容

关键：使用理解 markdown 上下文的正确解析。代码块内的 ## 不是章节标题。]]

### 3. 创建单独文件(Create Individual Files)

对于每个提取的章节：

1. **生成文件名**：将章节标题转换为小写连字符格式
    - 删除特殊字符
    - 用连字符替换空格
    - 示例："## 技术栈" → `tech-stack.md`

2. **调整标题级别**：
    - 二级标题在分片的新文档中成为一级标题（# 而不是 ##）
    - 所有子章节级别减少 1：

    ```txt
      - ### → ##
      - #### → ###
      - ##### → ####
      - 等等
    ```

3. **写入内容**：将调整后的内容保存到新文件

### 4. 创建索引文件(Create Index File)

在分片文件夹中创建一个 `index.md` 文件，该文件：

1. 包含原始一级标题和第一个二级章节之前的任何内容
2. 列出所有分片文件并带有链接：

```markdown
# 原始文档标题

[原始介绍内容（如果有）]

## 章节

- [章节名称 1](./section-name-1.md)
- [章节名称 2](./section-name-2.md)
- [章节名称 3](./section-name-3.md)
  ...
```

### 5. 保持特殊内容(Preserve Special Content)

1. **代码块**：必须捕获完整块，包括：

    ```language
    content
    ```

2. **Mermaid 图表**：保持完整语法：

    ```mermaid
    graph TD
    ...
    ```

3. **表格**：保持正确的 markdown 表格格式

4. **列表**：保持缩进和嵌套

5. **内联代码**：保持反引号

6. **链接和引用**：保持所有 markdown 链接完整

7. **模板标记**：如果文档包含 {{placeholders}}，精确保持

### 6. 验证(Validation)

分片后：

1. 验证所有章节都已提取
2. 检查没有内容丢失
3. 确保标题级别已正确调整
4. 确认所有文件都已成功创建

### 7. 报告结果(Report Results)

提供摘要：

```text
文档分片成功：
- 源： [原始文档路径]
- 目标： docs/[文件夹名称]/
- 创建的文件： [数量]
- 章节：
  - section-name-1.md: "章节标题 1"
  - section-name-2.md: "章节标题 2"
  ...
```

## 重要说明(Important Notes)

- 永远不要修改实际内容，只调整标题级别
- 保持所有格式，包括重要的空白
- 处理边缘情况，如包含 ## 符号的代码块章节
- 确保分片是可逆的（可以从分片重建原始文档）
