package main

import "testing"

func TestWordCount(t *testing.T) {
	
	// Arrange
	
	testTable := []struct {
		s        string
		expected map[string]int
	}{
		{
			s: " as sa as fe we e e e",
			expected: map[string]int{
				"as": 2,
				"sa": 1,
				"fe": 1,
				"we": 1,
				"e":  3,
			},
		},
		{
			s: "aaaaa",
			expected: map[string]int{
				"aaaaa": 1,
			},
		},
		{
			s: "a1 a2 a3",
			expected: map[string]int{
				"a1": 1,
				"a2": 1,
				"a3": 1,
			},
		},
		{
			s: "1 1 asda w w w w ewqeqwe e e e e      d d d ",
			expected: map[string]int{
				"1": 2,
				"asda": 1,
				"w": 4,
				"ewqeqwe": 1,
				"e": 4,
				"d": 3,
			},
		},
	}

	// Act
	
	for _, testCase := range testTable{
		result := wordCount(testCase.s)
		t.Logf("Calling func(%#v)\n\t\t result %#v\n",
			testCase.s, result)

		// Assert
		if len(result) != len(testCase.expected) {
			t.Errorf("Error!\n Expected : %#v;\nResult: %#v\n", testCase.expected, result)
		} else {
			for i, _ := range result {
				if result[i] != testCase.expected[i] {
					t.Errorf("Error!\n Expected : %#v;\nResult: %#v\n", testCase.expected, result)
				}
			}
		}
	}
}

