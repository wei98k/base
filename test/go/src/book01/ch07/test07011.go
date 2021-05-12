package ch07

type Writer interface {
    Write(p []byte) (n int, err error)
}
