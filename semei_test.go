package seimei_test

import (
	"testing"

	"github.com/kuno4n/seimei"
)

func TestDivideName(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name       string
		inputName  string
		want       string
		wantErrMsg string
	}

	tests := []testdata{
		{
			name:      "サンプル",
			inputName: "田中太郎",
			want:      "田中 太郎",
		},
		{
			name:      "ルールベースで動作する",
			inputName: "乙一",
			want:      "乙 一",
		},
		{
			name:      "統計量ベースで動作する",
			inputName: "竈門炭治郎",
			want:      "竈門 炭治郎",
		},
		{
			name:      "統計量ベースで分割できる",
			inputName: "中曽根康弘",
			want:      "中曽根 康弘",
		},
		{
			name:       "1文字は分割できない",
			inputName:  "あ",
			want:       "",
			wantErrMsg: "parse error: name length needs at least 2 chars\n",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if err := seimei.ParseName(stdout, stderr, tt.inputName, tt.inputParser); err != nil {
				t.Fatalf("happen error: %v", err)
			}

			if stdout.String() != tt.want {
				t.Errorf("failed to test. got: %s, want: %s", stdout, tt.want)
			}
			if stderr.String() != tt.wantErrMsg {
				t.Errorf("failed to test. got: %s, want: %s", stderr, tt.wantErrMsg)
			}
		})
	}
}
