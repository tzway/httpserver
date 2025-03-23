package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// 测试 replaceBadWord 函数
func TestReplaceBadWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "替换单个敏感词",
			input:    "What a kerfuffle!",
			expected: "What a ****!",
		},
		{
			name:     "替换多个敏感词",
			input:    "kerfuffle sharbert fornax",
			expected: "**** **** ****",
		},
		{
			name:     "大小写混合",
			input:    "KeRfUfFle ShArBeRt",
			expected: "**** ****",
		},
		{
			name:     "无敏感词",
			input:    "Hello world",
			expected: "Hello world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := replaceBadWord(tt.input)
			if result != tt.expected {
				t.Errorf("期望: %s, 实际: %s", tt.expected, result)
			}
		})
	}
}

// 测试 handleValidateChirp 处理器
func TestHandleValidateChirp(t *testing.T) {
	tests := []struct {
		name         string
		body         string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "内容过长",
			body:         `{"body": "` + strings.Repeat("a", 141) + `"}`,
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"Chirp is too long"}`,
		},
		{
			name:         "无效JSON",
			body:         `invalid_json`,
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"Invalid JSON"}`,
		},
		{
			name:         "正常请求",
			body:         `{"body": "Hello kerfuffle"}`,
			expectedCode: http.StatusOK,
			expectedBody: `{"cleaned_body":"Hello ****"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建模拟请求
			req := httptest.NewRequest("POST", "/validate", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")

			// 创建响应记录器
			rr := httptest.NewRecorder()

			// 执行处理器
			handleValidateChirp(rr, req)

			// 验证状态码
			if status := rr.Code; status != tt.expectedCode {
				t.Errorf("状态码错误: 期望 %d, 实际 %d", tt.expectedCode, status)
			}

			// 验证响应体
			actualBody := strings.TrimSpace(rr.Body.String())
			if actualBody != tt.expectedBody {
				t.Errorf("响应体错误:\n期望: %s\n实际: %s", tt.expectedBody, actualBody)
			}
		})
	}
}
