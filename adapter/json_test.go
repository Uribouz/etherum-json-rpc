package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestJsonDataToSlice(t *testing.T) {
	testCases := []struct{
		name string
		input string
		expected []interface{}
	}{
		{
			name: "Test JsonDataToSlice 1",
			input: `{
				"addresseses": ["0x28c6c06298d514db089934071355e5743bf21d60","0x12345678901","0x121231721231"]
			}`,
			expected: []interface{}{"0x28c6c06298d514db089934071355e5743bf21d60","0x12345678901","0x121231721231"},
		},
	}
	for _, each := range testCases {
		actual, err := JsonDataToAddresses(MockDataGetter{data:each.input})
		assert.NoError(t, err)
		assert.ElementsMatch(t, each.expected, actual)
	}
}

type MockDataGetter struct {data string}
func (m MockDataGetter) GetData() (string, error) {
	return m.data, nil
}