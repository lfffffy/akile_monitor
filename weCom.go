package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

// 解析CPU数量
func parseCPU(cpu string) int {
	re := regexp.MustCompile(`(\d+) (Virtual) Core`)

	// 查找匹配项
	matches := re.FindStringSubmatch(cpu)
	if len(matches) >= 2 {
		virtualCores := matches[1]

		vint, err := strconv.Atoi(virtualCores)
		if err != nil {
			return 0
		}
		return vint
	}
	return 0
}

// 格式化字节大小
func formatSize(size uint64) string {
	if size < 1024 {
		return fmt.Sprintf("%d B", size)
	} else if size < 1024*1024 {
		return fmt.Sprintf("%.2f KB", float64(size)/1024)
	} else if size < 1024*1024*1024 {
		return fmt.Sprintf("%.2f MB", float64(size)/1024/1024)
	} else if size < 1024*1024*1024*1024 {
		return fmt.Sprintf("%.2f GB", float64(size)/1024/1024/1024)
	} else if size < 1024*1024*1024*1024*1024 {
		return fmt.Sprintf("%.2f TB", float64(size)/1024/1024/1024/1024)
	} else if size < 1024*1024*1024*1024*1024*1024 {
		return fmt.Sprintf("%.2f PB", float64(size)/1024/1024/1024/1024/1024)
	} else {
		return fmt.Sprintf("%.2f EB", float64(size)/1024/1024/1024/1024/1024/1024)
	}
}

func sendWeComMessage(msg string) {
	// 格式化JSON数据
	jsonData := fmt.Sprintf(`{"msgtype":"text","text":{"content":"%s"}}`, msg)

	log.Println("jsonData:", jsonData)

	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + cfg.WeComKey

	// 发送POST请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	log.Println("Response status:", resp.Status)
	log.Println("Response body:", string(body))
}
