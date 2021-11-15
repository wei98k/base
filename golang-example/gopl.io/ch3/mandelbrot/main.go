package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

// 执行生成图片: go run main.go > out1.png

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	// 原文: NewRGBA returns a new RGBA image with the given bounds.
	// 译文: NewRGBA返回一个新的RGBA图像，并给定边界。
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// 设置坐标和颜色
			img.Set(px, py, mandelbrot(z))
		}
	}
	// 输出图片
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		// 取绝对值
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	// 原文: The real built-in function returns the real part of the complex number c. The return value will be floating point type corresponding to the type of c.
	// 译文: 实数内置函数返回复数c的实数部分，返回值将是与c的类型相对应的浮点类型。
	blue := uint8(real(v)*128) + 127
	// 原文: The imag built-in function returns the imaginary part of the complex number c. The return value will be floating point type corresponding to the type of c.
	// 译文: imag内置函数返回复数c的虚部，返回值将是与c的类型相对应的浮点类型。
	red := uint8(imag(v)*128) + 127
	// 原文: YCbCr represents a fully opaque 24-bit Y'CbCr color, having 8 bits each for one luma and two chroma components.
	// 译文: YCbCr表示完全不透明的24位Y'CbCr颜色，一个卢马和两个色度成分各占8位。
	// Tip: RGB和Y'CbCr之间的转换是有损失的
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

func newto(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
