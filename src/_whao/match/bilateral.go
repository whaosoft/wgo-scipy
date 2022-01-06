package match

import (
	"fmt"
	"strings"
)

const (
	SLIDING_WINDOW int = 5 // 滑动窗口个数
	//LEMMA = "计算语言学课程有意思"
	LEMMA = "南京市长江大桥欢迎您"
	SEPARATION = "--"
	BLANK = " "
)
//
//var dict =[]string{"计算", "计算语言学", "课程", "有", "意思"}
var dict =[]string{"南京 ", "南京市 ", "长江大桥", " 大桥", " 市长", " 江大桥", " 南京市", " 欢迎您", " 欢迎", " 您"}

//
var result= ""
var tmpWinDow = SLIDING_WINDOW

//双向最大匹配-逆向

/*
//
func main() {

	var strArr []string

	//省的一个字占俩字节的处理,反正不会损失嘛性能
	for _,ss:= range LEMMA{
		strArr = append(strArr,string(ss) )
	}

	ll:=len(strArr)-1

	Reverse(strArr,SLIDING_WINDOW,ll)

	fmt.Println(result)
}
*/

//
func Reverse(str []string,len,from int){

	if from < 0 {
		return
	}

	tmpstr := ""

	currlen := from + 1 - len
	//fmt.Println("currlen",currlen)
	if currlen >= 0 {
		//
		tmparr:=str[from-len+1 : from+1]
		tmpstr = getstring(&tmparr)
	}
	//没什么用测半天也没得这个的时候
	/*
	else {
		tmparr:=str[0 : from+1]
		//fmt.Println("from",from)
		tmpstr = getstring(&tmparr)
	}
	*/

	//fmt.Println("==",tmpstr)
	if len ==SLIDING_WINDOW{
		fmt.Println("每一次滑动=",tmpstr)
	}

	if contain(tmpstr) {
		//fmt.Println("==",tmpstr,from-len)
		result=tmpstr+SEPARATION+result
		tmpWinDow=SLIDING_WINDOW
		Reverse(str,tmpWinDow,from-len)
	}else{

		if tmpWinDow>1 {
			tmpWinDow=tmpWinDow-1
			Reverse(str,tmpWinDow,from)
		}else{
			//==1
			//result=result+tmpstr+SEPARATION
			from=from-1
			tmpWinDow=SLIDING_WINDOW
			Reverse(str,tmpWinDow,from)
		}

	}

}


//
func getstring(arr *[]string)(result string){
	for _, i := range *arr {
		result += i
	}
	return result
}

//
func contain(ss string)bool{

	for _,dd:=range dict{
		//懒得删了在这加个
		if strings.Trim(dd,BLANK) == ss{
			return true
		}
	}
	return false
}