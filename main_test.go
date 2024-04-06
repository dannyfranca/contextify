package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestContextify(t *testing.T) {
	// Build the binary
	buildCmd := exec.Command("go", "build", "-o", "bin/ctxfy", ".")
	err := buildCmd.Run()
	if err != nil {
		t.Fatalf("Failed to build the binary: %v", err)
	}

	// Create a temporary output file
	tempFile, err := os.CreateTemp("", "codebase*.md")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	// Set the command-line arguments
	os.Args = []string{"bin/ctxfy", "gen", "-o", tempFile.Name(), "-e", "go,md,json,html", "assets"}

	// Run the main function
	main()

	// Read the generated output file
	outputBytes, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	outputString := string(outputBytes)

	// Define the expected output
	expectedOutput := `# assets/dir1/file1.go

` + "```go" + `
package main

func main() {
	println("Hello, World!")
}
` + "```" + `

# assets/dir1/file2.md

` + "```md" + `
# Heading

This is a markdown file.
` + "```" + `

# assets/dir2/file3.json

` + "```json" + `
{
  "name": "John Doe",
  "age": 30
}
` + "```" + `

# assets/file5.html

` + "```html" + `
<!DOCTYPE html>
<html>

<head>
    <title>Example HTML</title>
</head>

<body>
    <h1>Welcome</h1>
    <p>This is an HTML file.</p>
</body>

</html>
` + "```" + `

`

	// Compare the generated output with the expected output
	println(strings.TrimSpace(outputString))
	println("")
	println("")
	println(strings.TrimSpace(expectedOutput))
	if strings.TrimSpace(outputString) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Generated output does not match the expected output")
	}
}
