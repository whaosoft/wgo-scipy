package optimize

import (
	"testing"
	"fmt"
)


/**

 */
func TestLeastSquares(t *testing.T) {

	x:=[]float64{50,40}
	y:=[]float64{110,140}
	a,b,e:=LeastSquares(x,y)

	fmt.Println("A:",a," B:",b," err:",e)

	a,b,x1,e:=LeastSquaresX(x,y,130)

	fmt.Println("A:",a," B:",b," X:",x1," err:",e)

}
