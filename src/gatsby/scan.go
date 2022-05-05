package gatsby

import (
	"bufio"
	"bytes"
	"log"
	"strings"
)

// scanFile searches a markdown file for the first H1 line
func scanFile(data []byte) string {
	bytesReader := bytes.NewReader(data)
	scanner := bufio.NewScanner(bytesReader)

	for scanner.Scan() {
		firstH1 := strings.HasPrefix(scanner.Text(), "# ")
		if firstH1 {
			return strings.Replace(scanner.Text(), "# ", "", 1)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	log.Printf("no file data for %s", string(data))

	return ""
}
