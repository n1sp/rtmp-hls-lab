package main

import (
	"os"
	"testing"
)

func TestIsAllowedStreamKey(t *testing.T) {
	// テスト用に一時的に環境変数を設定
	os.Setenv("AUTH_STREAM_KEY", "test_allowed")

	t.Cleanup(func() {
		os.Unsetenv("AUTH_STREAM_KEY")
	})

	allowedKey := "test_allowed"
	if IsAllowedStreamKey(allowedKey) == false {
		t.Error("trueを期待していましたが、falseが返されました。")
	}
	disallowedKey := "invalid_key"
	if IsAllowedStreamKey(disallowedKey) {
		t.Error("falseを期待していましたが、trueが返されました。")
	}

}
