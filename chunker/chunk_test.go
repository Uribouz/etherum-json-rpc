package chunker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunk(t *testing.T) {
	testCases := []struct{
		name string
		input struct{
			workerNum int
			data []string
		}
		expected struct {
			data [][]string
		}
	}{
		{
			name: "Test Chunk 1",
			input: struct{
				workerNum int
				data []string
			}{
				workerNum: 3,
				data: []string{"ADDR1", "ADDR2", "ADDR3", "ADDR4", "ADDR5", "ADDR6", "ADDR7", "ADDR8", "ADDR9", "ADDR10", "ADDR11", "ADDR12", "ADDR13", "ADDR14"},
			},
			expected: struct {
				data [][]string
			}{
				data: [][]string{
					{"ADDR1", "ADDR2", "ADDR3", "ADDR4"},
					{"ADDR5", "ADDR6", "ADDR7", "ADDR8"},
					{"ADDR9", "ADDR10", "ADDR11", "ADDR12","ADDR13", "ADDR14"},
				},
			},
		},
	}
	for _, each := range testCases {
		actual := Chunk(each.input.workerNum, each.input.data)
		assert.ElementsMatch(t, each.expected.data, actual)
	}
}