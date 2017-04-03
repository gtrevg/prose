package tokenize

import (
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/jdkato/prose/internal/util"
	"github.com/stretchr/testify/assert"
)

var testdata = filepath.Join("..", "testdata")

func TestTreebankWordTokenizer(t *testing.T) {
	in := util.ReadDataFile(filepath.Join(testdata, "treebank_sents.json"))
	out := util.ReadDataFile(filepath.Join(testdata, "treebank_words.json"))

	input := []string{}
	output := [][]string{}

	util.CheckError(json.Unmarshal(in, &input))
	util.CheckError(json.Unmarshal(out, &output))

	word := NewTreebankWordTokenizer()
	for i, s := range input {
		assert.Equal(t, output[i], word.Tokenize(s))
	}
}
