# Gin + Gorm + PostgreSQL 实现关注/取消关注 REST 服务

## 项目简介
本项目使用 Go 的 Gin 框架、Gorm ORM 和 PostgreSQL 数据库实现了一个基础的用户关注和取消关注的 REST 服务。包括以下功能：

1. 用户注册
2. 用户登录
3. 关注其他用户
4. 取消关注

---

## 环境要求

- Go 1.19 或以上版本
- PostgreSQL 数据库

---

## 项目结构

```
.
├── main.go               // 主程序入口
├── controllers           // 控制器目录
│   ├── userController.go // 用户相关功能
│   ├── followController.go // 关注/取消关注功能
├── database              // 数据库配置
│   └── db.go
├── models                // 数据模型
│   ├── user.go           // 用户模型
│   ├── follow.go         // 关注模型
├── routes                // 路由配置
│   └── routes.go
├── middleware            // 中间件
│   └── auth.go           // 认证中间件
├── go.mod                // Go 模块依赖
├── go.sum                // 依赖锁文件
├── main_test.go          // 测试代码
```

---

## 数据库配置

在 `database/db.go` 中设置 PostgreSQL 的连接字符串：
```go
var dsn = "host=localhost user=your_user password=your_password dbname=your_db port=5432 sslmode=disable"
```

运行数据库迁移：
```go
database.DB.AutoMigrate(&models.User{}, &models.Follow{})
```

---

## 功能说明

### 1. 用户注册
**接口：**
```
POST /api/register
```
**请求体：**
```json
{
  "username": "test_user",
  "password": "password123"
}
```
**响应：**
```json
{
  "message": "User registered successfully!"
}
```

### 2. 用户登录
**接口：**
```
POST /api/login
```
**请求体：**
```json
{
  "username": "test_user",
  "password": "password123"
}
```
**响应：**
```json
{
  "token": "Bearer <your_token>"
}
```

### 3. 关注用户
**接口：**
```
POST /api/user/follow
```
**请求体：**
```json
{
  "followee_id": 2
}
```
**响应：**
```json
{
  "message": "Followed user successfully!"
}
```

### 4. 取消关注
**接口：**
```
DELETE /api/user/unfollow
```
**请求体：**
```json
{
  "followee_id": 2
}
```
**响应：**
```json
{
  "message": "Unfollowed user successfully!"
}
```

---

## 测试说明

项目使用 `httptest` 进行单元测试，测试代码位于 `main_test.go` 文件中。

运行测试命令：
```bash
go test ./...
```

测试覆盖以下功能：
1. 用户注册 (`TestRegister`)
2. 用户登录 (`TestLogin`)
3. 关注用户 (`TestFollowUser`)
4. 取消关注 (`TestUnfollowUser`)

---

## 项目启动

1. 启动 PostgreSQL 数据库。
2. 配置 `database/db.go` 文件中的数据库连接信息。
3. 运行以下命令启动服务：
```bash
go run main.go
```
4. 服务默认运行在 `http://localhost:8080`。

---

## 常见问题

### 1. 数据库连接失败
- 检查 `dsn` 配置是否正确。
- 确保 PostgreSQL 服务正在运行。

### 2. 注册或登录时报 `500` 错误
- 检查数据库是否已迁移。
- 确保用户名唯一性未被破坏。

---

## 作者
- 作者: 翁晓逸
- 日期: 2024 年 12 月

