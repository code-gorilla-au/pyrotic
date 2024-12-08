package engine

import (
	"testing"

	"github.com/code-gorilla-au/odize"
)

var emptyMap = map[string]string{}

func Test_generateMeta(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]string
	}{
		{
			name:  "simple key-value pair",
			input: "key=value",
			expected: map[string]string{
				"key": "value",
			},
		},
		{
			name:  "multiple key-value pairs",
			input: "key1=value1,key2=value2",
			expected: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name:  "quoted values with commas",
			input: `field="name->string,age->int",enabled=true`,
			expected: map[string]string{
				"field":   "name->string,age->int",
				"enabled": "true",
			},
		},
		{
			name:     "empty input",
			input:    "",
			expected: map[string]string{},
		},
		{
			name:  "spaces around equals",
			input: "key = value, space = needed",
			expected: map[string]string{
				"key":   "value",
				"space": "needed",
			},
		},
		{
			name:  "quoted strings with spaces",
			input: `name="John Doe",age=30`,
			expected: map[string]string{
				"name": "John Doe",
				"age":  "30",
			},
		},
		{
			name:  "complex nested structure",
			input: `fields="name->string,age->int",config="host=localhost,port=8080"`,
			expected: map[string]string{
				"fields": "name->string,age->int",
				"config": "host=localhost,port=8080",
			},
		},
		{
			name:  "complex nested structure, different separators",
			input: `fields="name:string;age:int",config="host=localhost,port=8080"`,
			expected: map[string]string{
				"fields": "name:string;age:int",
				"config": "host=localhost,port=8080",
			},
		},
		{
			name:     "invalid format missing value",
			input:    "key=",
			expected: emptyMap,
		},
		{
			name:     "invalid format missing equals",
			input:    "keyvalue",
			expected: emptyMap,
		},
		{
			name:  "single quoted value",
			input: `single='value'`,
			expected: map[string]string{
				"single": "value",
			},
		},
		{
			name:  "single double quoted value",
			input: `single="value"`,
			expected: map[string]string{
				"single": "value",
			},
		},
		{
			name:  "unmatched quotes",
			input: `key="value,other=test`,
			expected: map[string]string{
				"key": "value,other=test",
			},
		},
		{
			name:  "multiple single quoted sections",
			input: `first='quoted value',second='another value'`,
			expected: map[string]string{
				"first":  "quoted value",
				"second": "another value",
			},
		},
		{
			name:  "multiple double quoted sections",
			input: `first="quoted value",second="another value"`,
			expected: map[string]string{
				"first":  "quoted value",
				"second": "another value",
			},
		},
		{
			name:  "empty quoted string",
			input: `empty=""`,
			expected: map[string]string{
				"empty": "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateMeta(tt.input)
			odize.AssertEqual(t, tt.expected, result)
		})
	}
}

// TestGenerateMetaFuzz adds some fuzz testing for the parser
func TestGenerateMetaFuzz(t *testing.T) {
	problematicInputs := []string{
		`key=="value"`,         // Double equals
		`key=value,`,           // Trailing comma
		`,key=value`,           // Leading comma
		`key=value,,key2=val2`, // Double comma
		`"key"="value"`,        // Quoted key
		`key=value=extra`,      // Multiple equals
		`=value`,               // Missing key
		`key=`,                 // Missing value
		`"`,                    // Single quote
		`""`,                   // Empty quotes
		`key="`,                // Unclosed quote
		`key=",`,               // Quote followed by comma
	}

	for _, input := range problematicInputs {
		t.Run("fuzz_"+input, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("parser panicked on input %q: %v", input, r)
				}
			}()

			generateMeta(input)
		})
	}
}
