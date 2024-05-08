package note

import (
	"fmt"
	"gonote/util"
)

func SayHelloWorld() {
	fmt.Println("hello world!!")
}

// 轉義字符
func EscapedCharacters() {
	fmt.Println("\n1 雙引號")
	fmt.Println("\"hello world\"")
	fmt.Println("\n2 反斜線")
	fmt.Println("\\\\電子郵件說\\項目已取消\\清理文檔的時候\\我哭了")
	fmt.Println("\n3 警報聲")
	fmt.Println("\a即使一個程序只有3行常，也總有一天需要去維護他\a")
	fmt.Println("\n4 退格")
	fmt.Println("三天不編程, , \b生活變得毫無意義")
	fmt.Println("\n5 換頁")
	fmt.Println("一個程序猿正在寫軟件\f他的手指在鍵盤上飛舞\f程序編譯時沒有一條錯誤訓訊息\f運行起來就如同一陣微風")
	fmt.Println("\n6 回車")
	fmt.Println("我的感官很悠閒，我的精神自由地按照他自己的直覺前進\r我的程序如同自己在寫自己") // 一般與\n配合使用: \r\n
	fmt.Println("\n7 制表符")
	fmt.Println("一一\t坨坨\t方塊\n\r害羞的\t內向的\t靦腆的")
	fmt.Println("\n5 縱向制表符")
	fmt.Println("確實, 有時候我會遇到難題。 \v當我發現難題的時候, 我會慢下來, 安靜地觀察。\v然後我改變一行代碼, 困難就煙消雲散")
}

// 變量與常量
func VariablesAndConstant() {
	var v1 int
	var v2 int = 2
	var v3 = 3
	v1 = 1
	v4 := 4
	var (
		v5     = 5
		v6 int = 6
		v7 int
	)
	fmt.Printf("v1=%v,v2=%v,v3=%v,v4=%v,v5=%v,v6=%v,v7=%v\n", v1, v2, v3, v4, v5, v6, v7)
	fmt.Println("\n2 常量")

	const (
		c1 = 8
		c2 = iota
		c3 = iota
		c4
		c5 = 12
		c6
	)

	fmt.Printf("c1=%v,c2=%v,c3=%v,c4=%v,c5=%v,c6=%v\n", c1, c2, c3, c4, c5, c6)
}

func increase(n *int) {
	*n++ // n = n + 1
	fmt.Printf("\nincrease結束時n=%v\n的內存地址為%v\nn指向的值為%v\n", n, &n, *n)
}

// 指針
func Pointer() {
	var src = 2022
	var ptr = &src
	increase(ptr)
	fmt.Printf("\n調用increase(ptr)之後, src=%v\nsrc的內存地址為%v\n", src, &src)
	fmt.Printf("\n調用increase(ptr)之後, ptr=%v\nptr的內存地址為%v\nptr指向的值為%v\n", ptr, &ptr, *ptr)
}

func IfElse() {
	var age uint8
	fmt.Println("請輸入你的年齡:")
	fmt.Scanln(&age)
	if age < 13 {
		fmt.Println("小朋友不要學編程喔")
	} else if age < 25 {
		fmt.Println("年輕人學編程是正確的選擇")
	} else {
		fmt.Println("學編程的人都是大神")
	}
}

func SwitchCaseAge() {
	var age uint8
	fmt.Println("請輸入你的年齡:")
	fmt.Scanln(&age)
	switch {
	case age < 13:
		fmt.Println("小朋友不要學編程喔")
	case age < 25:
		fmt.Println("年輕人學編程是正確的選擇")
	default:
		fmt.Println("學編程的人都是大神")
	}
}

func SwitchCaseWeekday() {
	var weekday uint8
	fmt.Println("請輸入今天是星期幾:")
	fmt.Scanln(&weekday)
	switch weekday {
	case 1:
		fmt.Println("醬油炒飯")
	case 2:
		fmt.Println("醬油炒麵")
	default:
		fmt.Println("輸入有誤")
	}
}

// for 循環
func For() {
	fmt.Println("\n 無限循環")
	i := 1
	for {
		fmt.Print(i, "\t")
		i++
		if i == 11 {
			fmt.Println()
			break
		}
	}

	fmt.Println("\n 條件循環")
	i = 1
	for i < 11 {
		fmt.Print(i, "\t")
		i++
	}
	fmt.Println()

	fmt.Println("\n 標準for循環")
	for j := 1; j < 11; j++ {
		fmt.Print(j, "\t")
	}
	fmt.Println()
}

// 函數

func Function() {
	res1, res2 := func(n1, n2 int) (sum, difference int) {
		sum = n1 + n2
		difference = n1 - n2
		return
	}(2, 3)
	fmt.Println("res1=", res1, ", res2=", res2)

	// fmt.Printf("getRes=%v,Type of getRes=%T\n", getRes, getRes)
}

// defer
func deferUtil() func(int) int {
	i := 0
	return func(n int) int {
		fmt.Printf("本次調用接收到n=%v\n", n)
		i++
		fmt.Printf("本次調用是第%v次調用\n", i)
		return i
	}
}

func Defer() int {
	f := deferUtil()
	defer f(1)
	defer f(2)
	defer f(3)
	return f(4)
}

func DeferRecover() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	n := 0
	fmt.Println(3 / n)
}

// 數組
func Array() {
	// 聲明
	var a = [...]int{
		1,
		456,
		789,
	}
	a[0] = 123
	fmt.Println("for 遍歷")
	for i := 0; i < len(a); i++ {
		fmt.Println("a[%v]=%v", i, a[i])
	}
	fmt.Println("\n for...range遍歷")
	for i, v := range a {
		fmt.Println("a[%v]=%v", i, v)
	}
	fmt.Println(("\n 多維數組"))
	var twoDimensionalArray [3][4]int = [3][4]int{
		{1, 2, 3, 4},
		{2, 3, 4, 5},
		{3, 4, 5, 6},
	}
	for i, v := range twoDimensionalArray {
		for i2, v2 := range v {
			fmt.Printf("a[%v][%v]=%v\n", i, i2, v2)
		}
		fmt.Println()
	}

}

// 4.2 切片
func Slice() {
	array := [5]int{1, 2, 3, 4, 5}
	var s1 []int = array[1:4] //[開始引用的index:結束引用的index+1]//[0:len(array)]等效于[:]
	s1[0] = 0
	fmt.Println("array=", array)
	s2 := s1[1:]
	s2[0] = 0
	fmt.Println("array=", array)
	var s3 []int
	fmt.Println("does s3==nil?", s3 == nil)
	s3 = make([]int, 3) //make([]Type, len, cap)//如果不寫cap, 默認與len相同
	fmt.Printf("len(s3)=%v,cap(s3)=%v\n", len(s3), cap(s3))
	s4 := []int{1, 2, 3} //由系統自動創建底層數組
	fmt.Println("s4=", s4)

	fmt.Println("\n4.2.3 追加元素")
	s1 = append(s1, 6, 7, 8) //底層創建了新的數組,不在引用原數組
	s1[4] = 0
	fmt.Println("array=", array)
	fmt.Println("s1=", s1)
	s5 := append(s1, s2...)
	fmt.Println("s5=", s5)

	fmt.Println("\n4.2.4 複製數組")
	s6 := []int{1, 1}
	copy(s5, s6) //容量能接收多少，就接收多少
	fmt.Println("s5=", s5)

	fmt.Println("\n4.2.5 string與[]byte")
	str := "hello 世界"
	fmt.Printf("[]byte(str)=%v\n[]byte(str)=%s\n", []byte(str), []byte(str))
	for i, v := range str {
		fmt.Printf("str[%v]=%c\n", i, v)
	}
	key := util.SelectByKey("註冊", "登錄", "退出")
	fmt.Println("接收到Key=", key)

}
