package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"awesomeProject/database"
	"awesomeProject/models"
	"awesomeProject/routes"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	routes.SetupRoutes(r)
	return r
}

func TestRegister(t *testing.T) {
	database.Connect() // 连接数据库

	// 清理可能的测试残留数据
	database.DB.Where("username = ?", "test_user").Unscoped().Delete(&models.User{})

	r := setupRouter()

	// 构造请求
	payload := map[string]string{
		"username": "test_user",
		"password": "password123",
	}
	jsonPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected HTTP status 200, got %d", w.Code)
	}
}

func TestLogin(t *testing.T) {
	database.Connect()

	r := setupRouter()

	// 构造请求
	payload := map[string]string{
		"username": "test_user",
		"password": "password123",
	}
	jsonPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected HTTP status 200, got %d", w.Code)
	}

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)

	if _, exists := response["token"]; !exists {
		t.Errorf("Expected token in response")
	}
}

func TestFollowUser(t *testing.T) {
	database.Connect()
	r := setupRouter()

	// 清理可能的测试残留数据
	database.DB.Where("username IN ?", []string{"follower", "followee"}).Unscoped().Delete(&models.User{})

	// 创建测试用户
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user1 := models.User{Username: "follower", Password: string(hashedPassword)}
	database.DB.Create(&user1)
	user2 := models.User{Username: "followee", Password: string(hashedPassword)}
	database.DB.Create(&user2)

	// 模拟登录获取 token
	loginPayload := map[string]string{
		"username": "follower",
		"password": "password123",
	}
	loginJSON, _ := json.Marshal(loginPayload)

	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(loginJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Login failed with status: %d", w.Code)
	}

	var loginResponse map[string]string
	json.Unmarshal(w.Body.Bytes(), &loginResponse)
	token := loginResponse["token"]

	// 测试关注用户
	followPayload := map[string]uint{
		"followee_id": user2.ID,
	}
	followJSON, _ := json.Marshal(followPayload)

	req, _ = http.NewRequest("POST", "/api/user/follow", bytes.NewBuffer(followJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected HTTP status 200, got %d", w.Code)
	}

	var followResponse map[string]string
	json.Unmarshal(w.Body.Bytes(), &followResponse)
	if followResponse["message"] != "Followed user successfully!" {
		t.Errorf("Unexpected response: %s", followResponse["message"])
	}

	// 清理测试数据
	database.DB.Unscoped().Delete(&user1)
	database.DB.Unscoped().Delete(&user2)
}

func TestUnfollowUser(t *testing.T) {
	database.Connect()
	r := setupRouter()

	// 清理可能的测试残留数据
	database.DB.Where("username IN ?", []string{"follower", "followee"}).Unscoped().Delete(&models.User{})

	// 创建测试用户
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user1 := models.User{Username: "follower", Password: string(hashedPassword)}
	database.DB.Create(&user1)
	user2 := models.User{Username: "followee", Password: string(hashedPassword)}
	database.DB.Create(&user2)

	// 模拟登录获取 token
	loginPayload := map[string]string{
		"username": "follower",
		"password": "password123",
	}
	loginJSON, _ := json.Marshal(loginPayload)

	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(loginJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Login failed with status: %d", w.Code)
	}

	var loginResponse map[string]string
	json.Unmarshal(w.Body.Bytes(), &loginResponse)
	token := loginResponse["token"]

	// 模拟关注用户
	followPayload := map[string]uint{
		"followee_id": user2.ID,
	}
	followJSON, _ := json.Marshal(followPayload)

	req, _ = http.NewRequest("POST", "/api/user/follow", bytes.NewBuffer(followJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Follow failed with status: %d", w.Code)
	}

	// 测试取消关注用户
	unfollowPayload := map[string]uint{
		"followee_id": user2.ID,
	}
	unfollowJSON, _ := json.Marshal(unfollowPayload)

	req, _ = http.NewRequest("DELETE", "/api/user/unfollow", bytes.NewBuffer(unfollowJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected HTTP status 200, got %d", w.Code)
	}

	var unfollowResponse map[string]string
	json.Unmarshal(w.Body.Bytes(), &unfollowResponse)
	if unfollowResponse["message"] != "Unfollowed user successfully!" {
		t.Errorf("Unexpected response: %s", unfollowResponse["message"])
	}

	// 清理测试数据
	database.DB.Unscoped().Delete(&user1)
	database.DB.Unscoped().Delete(&user2)
}
