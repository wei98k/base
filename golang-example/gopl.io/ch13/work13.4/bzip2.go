package bzip

import (
	"io"
	"os/exec"
)

// 练习 13.4： 因为C库依赖的限制。
// 使用os/exec包启动/bin/bzip2命令作为一个子进程，
// 提供一个纯Go的bzip.NewWriter的替代实现
// （译注：虽然是纯Go实现，但是运行时将依赖/bin/bzip2命令，
// 其他操作系统可能无法运行）。

type writer struct {
	in  io.WriteCloser
	cmd *exec.Cmd
}

func NewWriter(w io.Writer) io.WriteCloser {
	cmd := exec.Command("bzip2")

	cmd.Stdout = w

	in, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	wc := &writer{in, cmd}
	return wc
}

func (w *writer) Write(data []byte) (int, error) {
	return w.in.Write(data)
}

func (w *writer) Close() error {
	err := w.in.Close()
	cmdErr := w.cmd.Wait()
	if err != nil {
		return err
	}
	if cmdErr != nil {
		return cmdErr
	}
	return nil
}
