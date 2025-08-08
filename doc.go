// Package pdf implements reading of PDF files with enhanced text extraction capabilities.
//
// This package provides robust PDF text extraction with support for dynamic text spacing
// handling, multi-language content, and various output formats. It handles complex PDFs
// including those with LaTeX content, small fonts, and multi-language text.
//
// # Features
//
//   - Dynamic text extraction with intelligent spacing handling
//   - Multi-language support with full Unicode including UTF-16
//   - Small font handling and mathematical notation support
//   - Multiple output formats: plain text, styled text, hierarchical content
//   - Character encoding support: WinAnsi, MacRoman, PDFDoc, UTF-16, custom encodings
//   - Configurable text span merging with baseline and spacing tolerance
//   - Text organization by rows, columns, or hierarchical structure
//
// # Basic Usage
//
// Extract plain text from a PDF:
//
//	f, r, err := pdf.Open("document.pdf")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer f.Close()
//
//	// Extract plain text
//	var buf bytes.Buffer
//	b, err := r.GetPlainText()
//	if err != nil {
//		log.Fatal(err)
//	}
//	buf.ReadFrom(b)
//	fmt.Println(buf.String())
//
// # Advanced Usage
//
// Extract styled text with formatting information:
//
//	spans, err := r.GetStyledTexts()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, span := range spans {
//		fmt.Printf("Text: %s | Font: %s | Size: %.2f | Position: (%.2f, %.2f)\n",
//			span.S, span.Font, span.FontSize, span.X, span.Y)
//	}
//
// Extract text organized by rows:
//
//	for pageIndex := 1; pageIndex <= r.NumPage(); pageIndex++ {
//		page := r.Page(pageIndex)
//		rows, err := page.GetTextByRow()
//		if err != nil {
//			continue
//		}
//
//		for _, row := range rows {
//			for _, text := range row.Content {
//				fmt.Printf("%s", text.S)
//			}
//		}
//	}
//
// # Custom Text Merging
//
// Configure custom merging options for specialized text extraction:
//
//	customMergeOptions := pdf.MergeOptions{
//		BaselineTolerance: 1.0,  // Stricter baseline alignment
//		SpaceMultiplier:   2.0,  // Allow wider gaps between words
//	}
//
//	span1 := pdf.Span{Font: "Arial", FontSize: 12, X: 0, Y: 100, W: 30}
//	span2 := pdf.Span{Font: "Arial", FontSize: 12, X: 35, Y: 100, W: 25}
//	canMerge := pdf.CanBeMerged(span1, span2, font, 0, 0, customMergeOptions)
//
// # Main Types
//
// Reader: Main PDF reader instance providing access to pages and text extraction methods.
//
// Page: Represents a PDF page with methods for extracting text in various formats.
//
// Text: Individual text element with position and formatting information.
//
// Span: Contiguous text run with consistent styling including font, size, and color.
//
// Line: Collection of spans on the same baseline.
//
// Block: Group of lines forming a paragraph-like structure.
//
// # Error Handling
//
// The library follows Go's standard error handling patterns. Most functions return
// an error as the last return value. For debugging difficult PDFs, enable debug mode:
//
//	pdf.DebugOn = true
//
// # Character Encoding
//
// The library automatically handles various character encodings including UTF-16,
// PDFDocEncoding, WinAnsiEncoding, MacRomanEncoding, and custom character maps
// through ToUnicode mappings.
//
// For complete examples and advanced usage patterns, see the examples/ directory.
package pdf
