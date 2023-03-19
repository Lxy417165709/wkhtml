package wkhtml

// DataExtractorPng Png数据提取器。
type DataExtractorPng struct{}

// Extract 提取数据。
func (e *DataExtractorPng) Extract(orgData []byte) []byte {
	head := []byte("\x89PNG\r\n\x1a\n")
	tail := []byte("\xae\x42\x60\x82")
	return extractData(orgData, head, tail)
}
