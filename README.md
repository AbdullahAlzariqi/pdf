# PDF Text Extraction Library

[![Built with WeBuild](https://raw.githubusercontent.com/webuild-community/badge/master/svg/WeBuild.svg)](https://webuild.community)

A robust Go library for extracting text from PDF files with dynamic text spacing handling. This library handles complex PDFs including those with LaTeX content, small fonts, and multi-language text.

## Features

- **Dynamic Text Extraction**: Intelligently handles text spacing regardless of PDF structure
- **Multi-Language Support**: Full Unicode support including UTF-16 encoded text and various character encodings
- **Small Font Handling**: Accurate extraction of text regardless of font size
- **LaTeX Support**: Handles mathematical notation and special characters
- **Multiple Output Formats**:
  - Plain text extraction
  - Styled text with formatting information (font, size, position, color)
  - Text organized by rows or columns
  - Hierarchical content structure (blocks, lines, spans)
- **Encoding Support**: WinAnsiEncoding, MacRomanEncoding, PDFDocEncoding, and custom character maps
- **Flexible Merging**: Configurable text span merging with baseline and spacing tolerance

## Installation

```bash
go get -u github.com/AbdullahAlzariqi/pdf
```

## Quick Start

### Basic Text Extraction

```go
package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/AbdullahAlzariqi/pdf"
)

func main() {
	f, r, err := pdf.Open("document.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		log.Fatal(err)
	}
	buf.ReadFrom(b)
	fmt.Println(buf.String())
}
```

### Advanced: Styled Text Extraction

```go
package main

import (
	"fmt"
	"log"

	"github.com/AbdullahAlzariqi/pdf"
)

func main() {
	f, r, err := pdf.Open("document.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	spans, err := r.GetStyledTexts()
	if err != nil {
		log.Fatal(err)
	}

	for _, span := range spans {
		fmt.Printf("Text: %s | Font: %s | Size: %.2f | Position: (%.2f, %.2f)\n",
			span.S, span.Font, span.FontSize, span.X, span.Y)
	}
}
```

### Text Organized by Rows

```go
package main

import (
	"fmt"
	"log"

	"github.com/AbdullahAlzariqi/pdf"
)

func main() {
	f, r, err := pdf.Open("document.pdf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for pageIndex := 1; pageIndex <= r.NumPage(); pageIndex++ {
		page := r.Page(pageIndex)
		rows, err := page.GetTextByRow()
		if err != nil {
			log.Printf("Error reading page %d: %v", pageIndex, err)
			continue
		}

		fmt.Printf("=== Page %d ===\n", pageIndex)
		for _, row := range rows {
			fmt.Printf("Row Y-position: %d\n", row.Position)
			for _, text := range row.Content {
				fmt.Printf("  %s", text.S)
			}
			fmt.Println()
		}
	}
}
```

### Custom Text Merging

```go
package main

import (
	"fmt"
	"log"

	"github.com/AbdullahAlzariqi/pdf"
)

func main() {
	// Configure custom merging options
	customMergeOptions := pdf.MergeOptions{
		BaselineTolerance: 1.0,  // Stricter baseline alignment
		SpaceMultiplier:   2.0,  // Allow wider gaps between words
	}

	// Use with font analysis
	font := pdf.Font{} // Initialize from actual PDF font
	span1 := pdf.Span{Font: "Arial", FontSize: 12, X: 0, Y: 100, W: 30}
	span2 := pdf.Span{Font: "Arial", FontSize: 12, X: 35, Y: 100, W: 25}

	canMerge := pdf.CanBeMerged(span1, span2, font, 0, 0, customMergeOptions)
	fmt.Printf("Spans can be merged: %t\n", canMerge)
}
```

## API Reference

### Main Types

- **`Reader`**: Main PDF reader instance
- **`Page`**: Represents a PDF page
- **`Text`**: Individual text element with position and formatting
- **`Span`**: Contiguous text run with consistent styling
- **`Line`**: Collection of spans on the same baseline
- **`Block`**: Group of lines forming a paragraph-like structure

### Key Methods

#### Reader Methods
- `GetPlainText() (io.Reader, error)`: Extract all text as plain string
- `GetStyledTexts() ([]Span, error)`: Get all text with styling information
- `NumPage() int`: Get total number of pages
- `Page(num int) Page`: Get specific page

#### Page Methods
- `GetPlainText(fonts map[string]*Font) (string, error)`: Extract page text
- `GetTextByRow() (Rows, error)`: Get text organized by rows
- `GetTextByColumn() (Columns, error)`: Get text organized by columns
- `Content() Content`: Get raw page content with positioning

### Character Encoding Support

The library supports multiple character encodings:
- **UTF-16**: Automatic detection and decoding
- **PDFDocEncoding**: Standard PDF character encoding
- **WinAnsiEncoding**: Windows ANSI character set
- **MacRomanEncoding**: Mac Roman character set
- **Custom Encodings**: Through ToUnicode character maps

### Error Handling

The library uses Go's standard error handling patterns. Most functions return an error as the last return value. For debugging difficult PDFs, enable debug mode:

```go
pdf.DebugOn = true
```

## Examples

See the `examples/` directory for complete working examples:
- `examples/read_plain_text/`: Basic text extraction
- `examples/read_text_with_styles/`: Styled text extraction

## Advanced Features

### Dynamic Spacing Handling
The library automatically handles PDFs with inconsistent spacing by analyzing font metrics, character spacing (Tc), and word spacing (Tw) operators.

### Multi-Language Text
Full Unicode support ensures proper extraction of text in any language, including right-to-left scripts and complex character compositions.

### Font Fallback
When font information is unavailable or corrupted, the library gracefully falls back to ensure text extraction continues.

## Testing

Run the test suite:

```bash
go test -v
```

The library includes comprehensive tests for text merging algorithms, character encoding, and edge cases.

## Contributing

Contributions are welcome! Please ensure:
1. All tests pass
2. New features include tests
3. Code follows Go conventions
4. Documentation is updated

## Publishing on pkg.go.dev

To publish the module documentation on pkg.go.dev:

1. Ensure `go.mod` declares `module github.com/AbdullahAlzariqi/pdf`.
2. Push the repository to a public Git host.
3. Create and push a semantic version tag, e.g., `git tag v0.1.0` and `git push origin v0.1.0`.
4. Visit https://pkg.go.dev/github.com/AbdullahAlzariqi/pdf to view the documentation.

## License

This project is licensed under the same terms as the original Go PDF library.
