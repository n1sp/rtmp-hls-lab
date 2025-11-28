package main

import (
	"fmt"
	"log"
	"net/http"
)

func publishHandler(w http.ResponseWriter, r *http.Request) {
	// nginx-rtmpから渡されるストリームキーは 'name' パラメータで取得
	streamKey := r.FormValue("name")
	// log
	log.Printf("[AUTH] stream key チェックを開始します。 Key: %s", streamKey)

	// *** 認証ロジックの仮実装 (Goバックエンド連携テスト用) ***
	// ストリームキーが 'test_allowed' なら許可、それ以外は拒否
	if IsAllowedStreamKey(streamKey) {
		// 許可（200）
		w.WriteHeader(http.StatusOK)
		// bodyに書き込み
		fmt.Fprint(w, "Stream を許可しました。")
		log.Printf("[AUTH] Stream Key '%s' が承認されました。", streamKey)
	} else {
		// 拒否（403）
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Stream は許可されていません。")
		log.Printf("[AUTH] Stream Key '%s' が拒否されました。", streamKey)

	}

}

func main() {
	// 認証エンドポイントを登録
	http.HandleFunc("/api/auth/publish", publishHandler)

	// Goバックエンドはポート8081でListen
	port := "8081"
	log.Printf("Go Backend listening on :%s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("バックエンドサーバーの起動に失敗しました。 %v", err)
	}

}

// StreamKeyの認証ロジック関数
func IsAllowedStreamKey(streamKey string) bool {
	AUTH_STREAM_KEY := "AUTH_STREAM_KEY"
	if streamKey == GetEnvString(AUTH_STREAM_KEY, "xxxxxxxxxx") {
		return true
	} else {
		return false
	}
}
