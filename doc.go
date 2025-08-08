// Package pdf implements reading of PDF files with enhanced text extraction capabilities.
//
// This package provides robust PDF text extraction with support for:
//   - Dynamic text spacing handling
//   - Multi-language support with full Unicode
//   - Small font handling
//   - LaTeX and mathematical notation support
//   - Multiple output formats (plain text, styled text, hierarchical content)
//   - Various character encodings (WinAnsi, MacRoman, PDFDoc, UTF-16)
//   - Configurable text span merging
//
// Basic usage:
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
// For more advanced usage including styled text extraction and custom merging options,
// see the examples in the examples/ directory.
package pdf
