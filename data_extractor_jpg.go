package wkhtml

// DataExtractorJpg Jpg数据提取器。
type DataExtractorJpg struct{}

func (e *DataExtractorJpg) Extract(orgData []byte) []byte {
	head := []byte("\xff\xd8")
	tail := []byte("\xff\xd9")
	return extractData(orgData, head, tail)
}
