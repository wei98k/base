package ch07

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
    *c  += ByteCounter(len(p))
    return len(p), nil
}
