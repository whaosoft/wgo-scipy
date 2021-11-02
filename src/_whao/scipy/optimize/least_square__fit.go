package optimize

import (
	"errors"
)

//功能 : 最小二乘法直线拟合 y = a·x + b， 计算系数a 和 b
//参数 : x -- 辐照度的数组
//       y --  功率的数组
//       num 是数组包含的元素个数，x[]和y[]的元素个数必须相等
//       a,b 都是返回值
//返回 : 拟合计算成功返回true, 拟合计算失败返回false

/**
    求a,b值
 */
func LeastSquares(x[]float64,y[]float64)(a float64,b float64,e error){

	// x是横坐标数据,y是纵坐标数据
	// a是斜率，b是截距
	xi := float64(0)
	x2 := float64(0)
	yi := float64(0)
	xy := float64(0)

	if len(x)!= len(y) {
		return 0,0, errors.New("最小二乘时，两数组长度不一致!")
	}else {
		length := float64(len(x))
		for i := 0; i < len(x); i++{
			xi += x[i]
			x2 += x[i] * x[i]
			yi += y[i]
			xy += x[i] * y[i]
		}
		a = (yi * xi - xy * length) / (xi * xi - x2 * length)
		b = (yi * x2 - xy * xi) / (x2 * length - xi * xi)
	}
	return a,b,nil
}


/**
    求x值
 */
func LeastSquaresX(x_arr[]float64,y_arr[]float64,y float64)(a float64,b float64,x float64,e error){

	a,b,e=LeastSquares(x_arr,y_arr)
	if nil!=e{
		return 0,0,0,e
	}

	return a,b,( y-b ) /a ,nil
}