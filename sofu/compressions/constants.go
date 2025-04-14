package compressions

// Constants for common compression schemes
const (
	CompressionGzip    = "gzip"
	CompressionDeflate = "deflate"
)

// array of supported compression schemes
var SupportedCompressionSchemes = []string{
	CompressionGzip,
}
