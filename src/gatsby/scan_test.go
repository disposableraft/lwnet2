package gatsby

import "testing"

func TestScanFile(t *testing.T) {
	var tests = []struct {
		name     string
		expected string
		file     file
	}{
		{
			name: "returns a title from the first line",
			file: file{
				name:     "test-file.md",
				markdown: templateData{},
				data:     []byte("# title foo\n### 2021-11-11 {#date}\n### tag1, tag2, tag3 {#tags}"),
			},
			expected: "# title foo",
		},
		{
			name: "returns a title from two consecuitive titles",
			file: file{
				name:     "",
				markdown: templateData{},
				data:     []byte("# First H1 tag\n# Second H1 tag\n### tag1, tag2, tag3 {#tags}"),
			},
			expected: "# First H1 tag",
		},
		{
			name: "returns a title that occurs after other tags",
			file: file{
				data: []byte("## tag, tags, taggers\n## 000--00--\n# H1 in third\n"),
			},
			expected: "# H1 in third",
		},
	}

	for _, test := range tests {
		result := scanFile(test.file.data)
		if result != test.expected {
			t.Errorf("%s expected %s got %s", test.name, test.expected, test.file.headline)
		}
	}
}
