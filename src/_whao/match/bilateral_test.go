package match

import (
	"testing"
	"fmt"
)

/**

 */
func TestBilateralReverse(t *testing.T) {

	var strArr []string

	//省的一个字占俩字节的处理,反正不会损失嘛性能
	for _,ss:= range LEMMA{
		strArr = append(strArr,string(ss) )
	}

	ll:=len(strArr)-1

	Reverse(strArr,SLIDING_WINDOW,ll)

	fmt.Println(result)

}