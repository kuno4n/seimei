package seimei

import (
	// Using embed.
	_ "embed"
	"encoding/csv"
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/glassmonkey/seimei/v2/feature"
	"github.com/glassmonkey/seimei/v2/parser"
)

const separator = " "

type DividedName struct {
	LastName  string
	FirstName string
}

var (
	//go:embed namedivider-python/assets/kanji.csv
	assets string

	nameParser          parser.NameParser
	kanjiFeatureManager feature.KanjiFeatureManager
)

func init() {
	kanjiFeatureManager = initKanjiFeatureManager()
	nameParser = initNameParser(separator, kanjiFeatureManager)
}

func initNameParser(parseString string, manager feature.KanjiFeatureManager) parser.NameParser {
	return parser.NewNameParser(parser.Separator(parseString), manager)
}

func initKanjiFeatureManager() feature.KanjiFeatureManager {
	r := csv.NewReader(strings.NewReader(assets))
	m := make(map[feature.Character]feature.KanjiFeature)

	for i := 0; ; i++ {
		record, err := r.Read()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			panic(err)
		}

		if i == 0 {
			continue
		}

		c := feature.Character(record[0])

		var os, ls []float64

		for _, o := range record[feature.CharacterFeatureSize : feature.CharacterFeatureSize+feature.OrderFeatureSize] {
			s, err := strconv.ParseFloat(o, 64)
			if err != nil {
				panic(err)
			}

			os = append(os, s)
		}

		for _, l := range record[feature.CharacterFeatureSize+feature.OrderFeatureSize : feature.CharacterFeatureSize+feature.OrderFeatureSize+feature.LengthFeatureSize] {
			s, err := strconv.ParseFloat(l, 64)
			if err != nil {
				panic(err)
			}

			ls = append(ls, s)
		}

		kf, err := feature.NewKanjiFeature(c, os, ls)
		if err != nil {
			panic(err)
		}

		m[c] = kf
	}

	return feature.KanjiFeatureManager{
		KanjiFeatureMap: m,
	}
}

func DivideName(fullName string) (*DividedName, error) {
	dividedName, err := nameParser.Parse(parser.FullName(fullName))
	if err != nil {
		return nil, err
	}
	return &DividedName{
		LastName:  string(dividedName.LastName),
		FirstName: string(dividedName.FirstName),
	}, nil
}
