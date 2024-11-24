// V0.1.1
// Author: wunderbarb
// © Nov 2024

package toolbox

import (
	"bufio"
	"strings"
	"testing"

	"github.com/wunderbarb/test"
)

func TestReadTheLine(t *testing.T) {
	require, assert := test.Describe(t)

	tests := []struct {
		line        string
		strip       bool
		found       string
		expectedErr bool
	}{
		{"qwerty\n", false, "qwerty\n", false},
		{"asd\njkk", false, "asd\n", false},
		{"zcv", false, "", true},
		{"qwerty\n", true, "qwerty", false},
		{"asd\njkk", true, "asd", false},
		{"zcv", true, "", true},
	}
	for _, tt := range tests {
		res, err := ReadTheLine(bufio.NewReader(strings.NewReader(tt.line)), tt.strip)
		if tt.expectedErr {
			require.Error(err)
		} else {
			require.NoError(err)
			assert.Equal(tt.found, res)
		}
	}
}
func Test_ReadFieldsFromLine(t *testing.T) {
	require, assert := test.Describe(t)

	const line = "scan 3m 56\n"
	r := strings.NewReader(line)
	b, err := ReadFieldsInLine(r)
	require.NoError(err, "could not parse")
	require.Equal(3, len(b), "did not returned the proper number of fields")
	assert.True((b[0] == "scan") && (b[1] == "3m") && (b[2] == "56"), "wrong fields %v", b)

	// redo the test with added \r
	const line2 = "scan 3m 56\r\n"
	r = strings.NewReader(line2)
	b, err = ReadFieldsInLine(r)
	require.NoError(err, "could not parse")
	require.Equal(3, len(b), "did not return the proper number of fields")
	if (b[0] != "scan") || (b[1] != "3m") || (b[2] != "56") {
		t.Errorf("wrong fields %v", b)
	}
}

func Test_ParseFieldsFromLine(t *testing.T) {
	require, assert := test.Describe(t)

	const line = "scan 3m 56\n"
	r := bufio.NewReader(strings.NewReader(line))
	ans, err := ParseFieldsInLine(r, "scan")
	require.NoError(err)
	require.Len(ans, 2)
	assert.Equal("3m", ans[0], "wrong fields %v", ans)
	assert.Equal("56", ans[1], "wrong fields %v", ans)
	r1 := bufio.NewReader(strings.NewReader(line))
	_, err = ParseFieldsInLine(r1, "bad")
	assert.ErrorIs(err, ErrParsingFailed)

	r2 := test.FaultyReader{}
	_, err = ParseFieldsInLine(r2, "scan")
	assert.Error(err)
}
