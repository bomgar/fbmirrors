package rank

// countWriter implements the io.Writer interface and counts the number of bytes written.
type countWriter struct {
	count int
}

// Write counts the bytes written to it and always succeeds without actually writing anywhere.
func (cw *countWriter) Write(p []byte) (int, error) {
	n := len(p)
	cw.count += n
	return n, nil
}
