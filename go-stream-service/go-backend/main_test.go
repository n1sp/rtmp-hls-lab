package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestPublishHandler_Allowed(t *testing.T) {
	// テスト用に一時的に環境変数を設定
	os.Setenv("AUTH_STREAM_KEYS", "valid_key")

	req := httptest.NewRequest("POST", "/api/auth/publish", strings.NewReader("name=valid_key"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// responseを受け取る箱
	rr := httptest.NewRecorder()

	publishHandler(rr, req)

	// status codeの確認
	if rr.Code != http.StatusOK {
		t.Fatalf("期待していたステータスコード %d ではなく、%d が返されました。", http.StatusOK, rr.Code)
	}

	fmt.Println(rr.Body.String())

}

func TestPublishHandler_Rejected(t *testing.T) {
	// テスト用に一時的に環境変数を設定
	os.Setenv("AUTH_STREAM_KEYS", "invalid_key")

	req := httptest.NewRequest("POST", "/api/auth/publish", strings.NewReader("name=valid_key"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// responseを受け取る箱
	rr := httptest.NewRecorder()

	publishHandler(rr, req)

	// status codeの確認
	if rr.Code != http.StatusForbidden {
		t.Fatalf("期待していたステータスコード %d ではなく、%d が返されました。", http.StatusForbidden, rr.Code)
	}

	fmt.Println(rr.Body.String())

}

func TestIsAllowedStreamKey(t *testing.T) {
	// テストケースの定義
	testCases := []struct {
		name     string // テスト名
		envValue string // 環境変数 AUTH_STREAM_KEYS
		inputKey string // 認証対象のストリームキー
		want     bool   // 期待する結果
	}{
		{
			name:     "正常系：完全一致するキーが含まれる",
			envValue: "key1,key2,key3",
			inputKey: "key3",
			want:     true,
		},
		{
			name:     "正常系：キーが含まれていない",
			envValue: "key1,key2,key3",
			inputKey: "key99",
			want:     false,
		},
		{
			name:     "準正常系：環境変数の値に空白あり（TrimSpaceの確認）",
			envValue: "key1,key2 ,key3",
			inputKey: "key2",
			want:     true,
		},
		{
			name:     "準正常系：引数が空文字",
			envValue: "key1,key2",
			inputKey: "",
			want:     false,
		},
		{
			name:     "準正常系：環境変数が空文字",
			envValue: "",
			inputKey: "key1",
			want:     false,
		},
		{
			name:     "エッジケース：１つだけ設定されている場合",
			envValue: "singleKey",
			inputKey: "singleKey",
			want:     true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// テスト実行用の環境変数を設定
			t.Setenv("AUTH_STREAM_KEYS", tt.envValue)

			got := IsAllowedStreamKey(tt.inputKey)

			if got != tt.want {
				t.Errorf("IsAllowedStreamKey(%q) = %v, want %v (env: %q)", tt.inputKey, got, tt.want, tt.envValue)
			}
		})
	}
}
