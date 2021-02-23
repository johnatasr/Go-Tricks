package main

type fileParserInterface interface {
        loadDataSet(path string, filename string)
        extractText(full_file_path string)
}

type pdfParser struct {
        // Data goes here ...
}

type emlParser struct {
        // Data goes here ...
}

func (p pdfParser) loadDataSet() {
        // Method definition ...
}

func (p pdfParser) extractText() {
        // Method definition ...
}

func (e emlParser) loadDataSet() {
        // Method definition ...
}

func (e emlParser) extractText() {
        // Method definition ...
}

func main() {
        // Main entrypoint
}