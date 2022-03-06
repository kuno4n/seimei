package seimei_test

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/glassmonkey/seimei"
)

func TestRun(t *testing.T) {
	t.Parallel()

	want := "田中 太郎"
	got := extractStdout(t, func() error {
		if err := seimei.Run("田中太郎", " "); err != nil {
			return fmt.Errorf("happen error: %w", err)
		}

		return nil
	})

	if got != want {
		t.Errorf("failed to test. got: %s, want: %s", got, want)
	}
}

// extractStdout
// refer: https://zenn.dev/glassonion1/articles/8ac939208bd455
func extractStdout(t *testing.T, runner func() error) string {
	t.Helper()

	// 既存のStdoutを退避する
	orgStdout := os.Stdout
	defer func() {
		// 出力先を元に戻す
		os.Stdout = orgStdout
	}()
	// パイプの作成(r: Reader, w: Writer)
	reader, writer, _ := os.Pipe()
	// Stdoutの出力先をパイプのwriterに変更する
	os.Stdout = writer
	// テスト対象の関数を実行する

	if err := runner(); err != nil {
		t.Fatalf("failed to run: %v", err)
	}
	// Writerをクローズする
	// Writerオブジェクトはクローズするまで処理をブロックするので注意
	writer.Close()
	// Bufferに書き込こまれた内容を読み出す
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(reader); err != nil {
		t.Fatalf("failed to read buf: %v", err)
	}
	// 文字列を取得する
	return strings.TrimRight(buf.String(), "\n")
}
