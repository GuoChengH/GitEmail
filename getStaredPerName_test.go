package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

// TestGetStargazers 测试获取仓库的 Star 用户列表
func TestGetStargazers(t *testing.T) {
	repoOwner := "viiftw" // 替换为仓库所有者的用户名
	repoName := "gfmail"  // 替换为仓库的名称

	// 构造 API 请求 URL
	apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/stargazers", repoOwner, repoName)

	// 创建 HTTP 请求客户端
	client := &http.Client{}

	// 创建 HTTP GET 请求
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 添加 Authorization 头
	req.Header.Set("Authorization", "token "+"token")

	// 发送 HTTP 请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching stargazers:", err)
		return
	}
	defer resp.Body.Close()

	// 解析返回的 JSON 数据
	var stargazers []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&stargazers)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println("Stargazers:", len(stargazers))
	// 提取所有 Star 用户的用户名
	for _, stargazer := range stargazers {
		username := stargazer["login"].(string)
		fmt.Println(username)
	}
}
