package packr

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

var testBox = NewBox("./fixtures")

func Test_PackBytes(t *testing.T) {
	r := require.New(t)
	PackBytes(testBox.Path, "foo", []byte("bar"))
	s := testBox.String("foo")
	r.Equal("bar", s)
}

func Test_PackJSONBytes(t *testing.T) {
	r := require.New(t)
	b, err := json.Marshal([]byte("json bytes"))
	r.NoError(err)
	err = PackJSONBytes(testBox.Path, "the bytes", string(b))
	r.NoError(err)
	s, err := testBox.MustBytes("the bytes")
	r.NoError(err)
	r.Equal([]byte("json bytes"), s)
}