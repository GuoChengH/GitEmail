package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
)

// TestGetForks 测试获取仓库的 Fork 用户列表
func TestGetForks(t *testing.T) {
	repoOwner := "zaphoyd"    // 替换为仓库所有者的用户名
	repoName := "websocketpp" // 替换为仓库的名称

	// 构造 API 请求 URL
	apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/forks", repoOwner, repoName)

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
		fmt.Println("Error fetching forks:", err)
		return
	}
	defer resp.Body.Close()

	// 解析返回的 JSON 数据
	var forks []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&forks)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println("Number of forks:", len(forks))
	// 提取所有 Fork 用户的用户名
	for _, fork := range forks {
		ownerInfo := fork["owner"].(map[string]interface{})
		username := ownerInfo["login"].(string)
		fmt.Println(username)
	}
	writeToFile(forks, "forks.txt")
}

func writeToFile(data []map[string]interface{}, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	bytes, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		return err
	}
	_, err = file.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}
