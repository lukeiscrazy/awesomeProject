# 用户关注/取消关注 REST 服务

## 项目简介
本项目是一个基于 Gin、Gorm 和 PostgreSQL 实现的后端 REST 服务，提供用户注册、登录、关注其他用户、取消关注以及查看关注列表的功能。

---

## 功能列表
1. **用户注册**
   - 路由：`POST /api/register`
   - 描述：用户可以通过用户名和密码注册账号。

2. **用户登录**
   - 路由：`POST /api/login`
   - 描述：用户可以通过用户名和密码登录系统，成功后返回 JWT Token。

3. **关注用户**
   - 路由：`POST /api/user/follow`
   - 描述：已登录用户可以关注其他用户。

4. **取消关注用户**
   - 路由：`DELETE /api/user/unfollow`
   - 描述：已登录用户可以取消关注其他用户。

5. **查看关注列表**
   - 路由：`GET /api/user/following`
   - 描述：已登录用户可以查看自己关注的所有用户列表。

---

## 环境配置

### 必备条件
- Go 1.19 或更高版本
- PostgreSQL 数据库
- 配置以下依赖包：
  ```bash
  go get github.com/gin-gonic/gin
  go get gorm.io/gorm
  go get gorm.io/driver/postgres
  go get github.com/dgrijalva/jwt-go
  go get golang.org/x/crypto
  ```

---

## 项目结构
```
.
├── controllers
│   ├── userController.go     # 用户注册、登录相关功能
│   ├── followController.go   # 用户关注、取消关注功能
│   ├── getfollowingController.go   # 查看关注列表功能
├── database
│   └── db.go                 # 数据库连接和初始化
├── models
│   ├── user.go               # 用户模型
│   ├── follow.go             # 关注模型
├── routes
│   └── routes.go             # 路由定义
├── middleware
│   └── auth.go               # JWT 中间件
├── main.go                   # 主程序入口
├── main_test.go              # 测试用例
├── go.mod                    # Go 模块配置
├── go.sum                    # Go 依赖锁定
```

---

## 使用方法

### 数据库配置
1. 创建 PostgreSQL 数据库。
   ```sql
   CREATE DATABASE your_db;
   ```
2. 更新 `database/db.go` 文件中的数据库连接配置。
   ```go
   dsn := "host=localhost user=your_user password=your_password dbname=your_db port=5432 sslmode=disable"
   ```

### 运行项目
1. 初始化数据库并启动服务：
   ```bash
   go run main.go
   ```
2. 服务将运行在 `http://localhost:8080`。

### API 测试
推荐使用 Postman 或 curl 工具测试以下 API：

1. 注册用户：
   ```bash
   curl -X POST http://localhost:8080/api/register -H "Content-Type: application/json" -d '{"username": "test_user", "password": "password123"}'
   ```

2. 用户登录：
   ```bash
   curl -X POST http://localhost:8080/api/login -H "Content-Type: application/json" -d '{"username": "test_user", "password": "password123"}'
   ```

3. 关注用户：
   ```bash
   curl -X POST http://localhost:8080/api/user/follow -H "Authorization: Bearer <your_token>" -H "Content-Type: application/json" -d '{"followee_id": 2}'
   ```

4. 取消关注：
   ```bash
   curl -X DELETE http://localhost:8080/api/user/unfollow -H "Authorization: Bearer <your_token>" -H "Content-Type: application/json" -d '{"followee_id": 2}'
   ```

5. 查看关注列表：
   ```bash
   curl -X GET http://localhost:8080/api/user/following -H "Authorization: Bearer <your_token>"
   ```

---

## 测试
1. 运行测试用例：
   ```bash
   go test ./...
   ```
2. 测试覆盖以下功能：
   - 用户注册
   - 用户登录
   - 关注用户
   - 取消关注
   - 查看关注列表

---

## 注意事项
1. 确保 `username` 字段唯一。
2. 在生产环境中，将 JWT 密钥替换为更加安全的密钥，并运行服务时设置为 `release` 模式。
   ```bash
   export GIN_MODE=release
   ```

