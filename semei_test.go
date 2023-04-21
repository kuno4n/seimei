package seimei_test

import (
	"github.com/glassmonkey/seimei/v2/parser"
	"testing"

	"github.com/kuno4n/seimei"
	"github.com/stretchr/testify/assert"
)

func TestDivideSeiMei(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name      string
		inputName string
		want      seimei.DividedName
		wantErr   error
	}

	tests := []testdata{
		{
			name:      "サンプル",
			inputName: "田中太郎",
			want: seimei.DividedName{
				LastName:  "田中",
				FirstName: "太郎",
			},
		},
		{
			name:      "ルールベースで動作する",
			inputName: "乙一",
			want: seimei.DividedName{
				LastName:  "乙",
				FirstName: "一",
			},
		},
		{
			name:      "統計量ベースで動作する",
			inputName: "竈門炭治郎",
			want: seimei.DividedName{
				LastName:  "竈門",
				FirstName: "炭治郎",
			},
		},
		{
			name:      "統計量ベースで分割できる",
			inputName: "中曽根康弘",
			want: seimei.DividedName{
				LastName:  "中曽根",
				FirstName: "康弘",
			},
		},
		{
			name:      "空白があったらそれを優先",
			inputName: "田 中太郎",
			want: seimei.DividedName{
				LastName:  "田",
				FirstName: "中太郎",
			},
		},
		{
			name:      "空白が1つだけあったらそれを分割場所として優先",
			inputName: "田 中太郎",
			want: seimei.DividedName{
				LastName:  "田",
				FirstName: "中太郎",
			},
		},
		{
			name:      "全角空白も対応、左右はトリム",
			inputName: "  　　  田　中太郎  　　  ",
			want: seimei.DividedName{
				LastName:  "田",
				FirstName: "中太郎",
			},
		},
		{
			name:      "空白が2つ以上あったら空白を除去して通常の姓名分割",
			inputName: "田 中太　郎",
			want: seimei.DividedName{
				LastName:  "田中",
				FirstName: "太郎",
			},
		},
		{
			name:      "1文字は分割できない",
			inputName: "あ",
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
				assert.Equal(t, tt.want, *dividedName)
			}
		})
	}
}
