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
	os.Setenv("AUTH_STREAM_KEY", "valid_key")

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
	os.Setenv("AUTH_STREAM_KEY", "invalid_key")

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
	// テスト用に一時的に環境変数を設定
	os.Setenv("AUTH_STREAM_KEY", "test_allowed")

	t.Cleanup(func() {
		os.Unsetenv("AUTH_STREAM_KEY")
	})

	const allowedKey = "test_allowed"
	if IsAllowedStreamKey(allowedKey) == false {
		t.Error("trueを期待していましたが、falseが返されました。")
	}
	disallowedKey := "invalid_key"
	if IsAllowedStreamKey(disallowedKey) {
		t.Error("falseを期待していましたが、trueが返されました。")
	}

}
