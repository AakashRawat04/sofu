package compressions

import (
	"bytes"
	"compress/gzip"
)

func gzipCompress(data string) string {
	var compressedData bytes.Buffer
	writer := gzip.NewWriter(&compressedData)
	defer writer.Close()
	_, err := writer.Write([]byte(data))
	if err != nil {
		return data // Return original data if compression fails
	}
	writer.Close()
	return compressedData.String()
}
