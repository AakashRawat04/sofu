package compressions

import "fmt"

func HandleCompression(scheme string, data string) (string, error) {
	switch scheme {
	case "gzip":
		return gzipCompress(data), nil
	default:
		return data, fmt.Errorf("unsupported compression scheme: %s", scheme)
	}
}
