package rectangle

import (
	"math"
	"fmt"
)

func init() {
	fmt.Println("rectangle package initialized")
}

// 计算矩形面积
func Area(length float64, width float64) float64 {
	area := length * width
	return area
}

// 计算矩阵对角线长度
func Diagonal(length float64, width float64) float64 {
	diagonal := math.Sqrt((length * length) + (width * width))
	return diagonal
}
