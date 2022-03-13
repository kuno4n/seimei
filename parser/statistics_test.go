package parser_test

import (
	"testing"

	"github.com/glassmonkey/seimei"
	"github.com/glassmonkey/seimei/feature"
	"github.com/glassmonkey/seimei/parser"
	"github.com/google/go-cmp/cmp"
)

func TestStatisticsParser_Parse(t *testing.T) {
	t.Parallel()

	type testdata struct {
		name  string
		input parser.FullName
		want  parser.DividedName
		skip  bool
	}

	separator := parser.Separator("/")
	tests := []testdata{
		{
			name:  "3文字",
			input: "菅義偉",
			want: parser.DividedName{
				LastName:  "菅",
				FirstName: "義偉",
				Separator: separator,
				Score:     0.3703703703703704,
				Algorithm: parser.Statistics,
			},
			skip: false,
		},
		{
			name:  "4文字",
			input: "阿部晋三",
			want: parser.DividedName{
				LastName:  "阿部",
				FirstName: "晋三",
				Separator: separator,
				Score:     0.995819397993311,
				Algorithm: parser.Statistics,
			},
			skip: false,
		},
		{
			name:  "5文字",
			input: "中曽根康弘",
			want: parser.DividedName{
				LastName:  "中曽根",
				FirstName: "康弘",
				Separator: separator,
				Score:     0.1111111111111111, // patch work score, todo fix.
				Algorithm: parser.Statistics,
			},
			skip: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.skip {
				t.Skip()
			}
			sut := parser.StatisticsParser{
				OrderCalculator: feature.KanjiFeatureOrderCalculator{
					Manager: seimei.InitKanjiFeatureManager(),
				},
			}
			got, err := sut.Parse(tt.input, separator)
			if err != nil {
				t.Errorf("error is not nil, err=%v", err)
			}

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("divided name mismatch (-got +want):\n%s", diff)
			}
		})
	}
}
