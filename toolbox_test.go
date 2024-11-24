// V0.2.0
// Author: wunderbarb
// © Nov 2024

package toolbox

import (
	"testing"

	"github.com/wunderbarb/test"
)

func Test_IsDirectory(t *testing.T) {
	_, assert := test.Describe(t)

	assert.True(IsDirectory("testdata"), "did not detect testdata as a directory")
	assert.False(IsDirectory("toolbox.go"), "did detect file as a directory")
	assert.False(IsDirectory("tool"), "did detect a non existing directory")
	assert.False(IsDirectory("testdata/faux"))
}

func Test_List_WithExtension(t *testing.T) {
	require, assert := test.Describe(t)

	l, err := List("../toolbox", WithExtension(test.SwapCase(".md")))
	require.NoError(err)
	assert.Len(l, 2)
	l, err = List("../toolbox", WithExtension(".go"), WithExtension(".md"))
	require.NoError(err)
	assert.Greater(len(l), 2)
	l, _ = List("../toolbox", WithExtension(".gold"))
	assert.Zero(len(l))
	l, _ = List("testdata2", WithExtension(".go"))
	assert.Zero(len(l))
}

func TestList(t *testing.T) {
	require, assert := test.Describe(t)

	l, err := List("../toolbox", WithSubDir())
	require.NoError(err)
	assert.NotEqual(0, len(l))
	assert.Contains(l, "testdata")

	l, err = List("../toolbox")
	require.NoError(err)
	assert.NotEqual(0, len(l))
	_, err = List("testdata2", WithSubDir())
	assert.Error(err)
	_, err = List("toolbox_test.go", WithSubDir())
	assert.Error(err)

	_, err = List("bad")
	assert.Error(err)
}

func TestList_WithOrdered(t *testing.T) {
	require, assert := test.Describe(t)
	dirT := t.TempDir()

	f1, _ := test.RandomFileWithDir(100, "tst", dirT)
	f2, _ := test.RandomFileWithDir(1, "tst", dirT)
	f3, _ := test.RandomFileWithDir(10, "tst", dirT)
	f4, _ := test.RandomFileWithDir(50, "tst", dirT)

	l, err := List(dirT, WithOrderedSize())
	require.Len(l, 4)
	require.NoError(err)
	assert.Equal(f2, l[0])
	assert.Equal(f3, l[1])
	assert.Equal(f4, l[2])
	assert.Equal(f1, l[3])

}

// Tests the Strip method.
func Test_Strip(t *testing.T) {
	_, assert := test.Describe(t)
	var data = []struct {
		f   string
		ext string
		s   string
	}{
		{"try.aes", ".aes", "try"},
		{"try.aes", "aes", "try"},
		{"try.go.aes", ".aes", "try.go"},
		{"try.go.aes", "aes", "try.go"},
		{"try.go.aes", ".go.aes", "try"},
		{"try.go.aes", "go.aes", "try"},
		{"try.aes", ".go", "try.aes"},
		{"try.aes", "go", "try.aes"},
		{"try.aes", "", "try.aes"},
		{"try.pcd.1.spe", ".1.spe", "try.pcd"},
		{"try.pcd.1.spe", "1.spe", "try.pcd"},
	}
	for _, tt := range data {
		assert.Equal(tt.s, Strip(tt.f, test.SwapCase(tt.ext)))
	}
}
