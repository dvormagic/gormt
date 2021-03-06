package gormt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type simpleStruct struct {
	StringField string `json:"stringField"`
	IntField    int64  `json:"intField"`
}

func TestSetValueSimpleStruct(t *testing.T) {
	var testJSON JSON[simpleStruct]

	foo := simpleStruct{
		StringField: "foo",
		IntField:    5,
	}
	testJSON.Set(foo)

	value, err := testJSON.Value()
	require.NoError(t, err)

	b := value.([]byte)
	require.Equal(t, `{"stringField":"foo","intField":5}`, string(b))
}

func TestScanGetSimpleStruct(t *testing.T) {
	var testJSON JSON[simpleStruct]

	err := testJSON.Scan([]byte(`{"stringField":"foo","intField":5}`))
	require.NoError(t, err)

	foo := testJSON.Get()
	require.Equal(t, "foo", foo.StringField)
	require.Equal(t, int64(5), foo.IntField)
}

func TestGetDataType(t *testing.T) {
	var testJSON JSON[simpleStruct]

	require.Equal(t, "json", testJSON.GormDataType())
}

func TestErrByteAssertion(t *testing.T) {
	var testJSON JSON[simpleStruct]

	err := testJSON.Scan(`{"stringField":"foo","intField":5}`)
	require.Error(t, err)
	require.Equal(t, "[]byte assertion failed", err.Error())
}

func TestErrUnmarshal(t *testing.T) {
	var testJSON JSON[simpleStruct]

	err := testJSON.Scan([]byte(`{"stringField":foo,"intField":"5"}`))
	require.Error(t, err)
}
