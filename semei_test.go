package seimei_test

import (
	"github.com/glassmonkey/seimei/v2/parser"
	"testing"

	"github.com/kuno4n/seimei"
	"github.com/stretchr/testify/assert"
)

func TestDivideName(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name      string
		inputName string
		want      string
		wantErr   error
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
			name:      "1文字は分割できない",
			inputName: "あ",
			want:      "",
			wantErr:   parser.ErrNameLength,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			dividedName, err := seimei.DivideSeiMei(tt.inputName)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.Equal(t, tt.want, dividedName.LastName+" "+dividedName.FirstName)
			}
		})
	}
}
