# Tech Stack

## Technology Stack Table

| Category | Technology | Version | Purpose | Rationale |
|---|---|---|---|---|
| **Language** | Go | 1.17 | 后端开发语言 | 现有项目语言，高性能，并发支持 |
| **Web Framework** | Gin | v1.7.7 | 构建 RESTful API | 现有项目框架，轻量级，高性能 |
| **Configuration Management** | Viper | v1.11.0 | 读取和管理应用程序配置 | 现有项目配置方案，灵活 |
| **Database** | MySQL | (通过驱动推断) | 主要持久化存储 | 现有项目数据库，成熟稳定 |
| **MySQL Driver** | go-sql-driver/mysql | v1.6.0 | MySQL 数据库驱动 | GORM 底层使用 |
| **ORM** | GORM | v1.30.0 | 简化与 MySQL 数据库的交互 | 现有项目 ORM，功能强大 |
| **GORM MySQL Driver** | gorm.io/driver/mysql | v1.3.3 | GORM 的 MySQL 驱动 | 简化 GORM 与 MySQL 集成 |
| **Database Migration** | Migrate | v4.15.2 | 管理数据库版本和结构变更 | 现有项目迁移工具，自动化 |
| **Unique ID Generation** | xid | v1.4.0 | 生成全局唯一的 ID | 现有项目依赖，轻量级 |
| **Captcha Generation** | afocus/captcha | v0.0.0-20191010092841-4bd1f21c8868 | 生成图形验证码 | 现有项目依赖，特定功能 |
| **Operating System** | Ubuntu | 21 | 部署环境 | 现有项目操作系统 |
