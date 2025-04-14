package compressions

import (
	"slices"
	"strings"
)

// IsSupported checks if the given compression scheme is supported
func IsSupported(scheme string) bool {
	return slices.Contains(SupportedCompressionSchemes, scheme)
}

// HandleCompression processes Accept-Encoding header and returns compressed data and used scheme
func HandleCompression(acceptEncoding string, data string) (string, string) {
	schemes := parseAcceptEncoding(acceptEncoding)

	for _, scheme := range schemes {
		if IsSupported(scheme) {
			// Apply the first supported compression scheme
			switch scheme {
			case "gzip":
				return gzipCompress(data), "gzip"
			}
		}
	}

	// No supported compression found
	return data, ""
}

// parseAcceptEncoding splits the Accept-Encoding header into individual schemes
func parseAcceptEncoding(acceptEncoding string) []string {
	schemes := strings.Split(acceptEncoding, ",")
	result := make([]string, 0, len(schemes))

	for _, scheme := range schemes {
		result = append(result, strings.TrimSpace(scheme))
	}

	return result
}
