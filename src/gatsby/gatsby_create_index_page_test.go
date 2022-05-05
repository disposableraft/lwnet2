package gatsby

import "testing"

func TestCreateIndexPageLinks(t *testing.T) {
	var tests = []struct {
		name      string
		arguments []file
		expected  string
	}{
		{
			name: "When given one argument it returns a string",
			arguments: []file{
				{
					headline: "one argument test",
					path:     "/example/com/one-arg-test.html",
				},
			},
			expected: "<a href=\"/example/com/one-arg-test.html\">one argument test</a>",
		},
		{
			name: "When given two arguments it returns a string",
			arguments: []file{
				{
					headline: "a first argument",
					path:     "/example/com/first.html",
				},
				{
					headline: "a second argument",
					path:     "/example/com/second.html",
				},
			},
			expected: "<a href=\"/example/com/first.html\">a first argument</a><a href=\"/example/com/second.html\">a second argument</a>",
		},
	}

	for _, test := range tests {
		result := createIndexPageLinks(test.arguments)
		if result != test.expected {
			t.Errorf("%s expected %s got %s", test.name, test.expected, result)

		}
	}
}
