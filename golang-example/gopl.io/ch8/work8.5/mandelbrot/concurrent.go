package mandelbrot

import (
	"image"
	"image/color"
	"math/cmplx"
	"sync"
)

func ConcurrentRender(workers int) *image.RGBA {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	// 原文: NewRGBA returns a new RGBA image with the given bounds.
	// 译文: NewRGBA返回一个新的RGBA图像，并给定边界。
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// update-channel
	rows := make(chan int, height)
	for row := 0; row < height; row++ {
		rows <- row
	}
	close(rows)

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			for py := range rows {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					// 设置坐标和颜色
					img.Set(px, py, mandelbrot(z))
				}
			}
			wg.Done()
		}()

	}
	wg.Wait()
	return img
	// 输出图片
	// png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
