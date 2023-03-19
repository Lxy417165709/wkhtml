package wkhtml

import "bytes"

// DataExtractor 数据提取器。
type DataExtractor interface {
	Extract(orgData []byte) []byte
}

// ExtractData 从 orgData 中，提取出期望的数据。
func ExtractData(orgData []byte, extractor DataExtractor) []byte {
	return extractor.Extract(orgData)
}

// extractData 提取数据。
func extractData(orgData []byte, head, tail []byte) []byte {
	headIndex := bytes.Index(orgData, head)
	tailIndex := bytes.Index(orgData, tail)
	if headIndex == -1 || tailIndex == -1 {
		return nil
	}
	expectDataHeadIndex := headIndex
	expectDataTailIndex := tailIndex + len(tail) - 1
	if expectDataHeadIndex > expectDataTailIndex {
		return nil
	}
	if expectDataTailIndex+1 > len(orgData) {
		return nil
	}
	return orgData[expectDataHeadIndex : expectDataTailIndex+1]
}
