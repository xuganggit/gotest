package main

import "fmt"

func main() {
	fmt.Println("hello word again")

	var b int  = 15
	var a int
	numbers := []int { 1, 2, 3 ,4 }					//初始化声明
	fmt.Printf("numbers的长度：%d\n", len(numbers))
	for a := 0; a < 10; a++ {						//for循环语法【a变量是局部变量】
		fmt.Printf("a的值为：%d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a的值为：%d\n", a)
	}

	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值：%d\n", i, x)
	}

	for i := 0; i < len(numbers); i++ {			//len函数可以获取集合长度
		fmt.Printf("numbers 第 %d 项的值为：%d\n", i, numbers[i])
	}

	fmt.Println(swap("徐刚", "你好，"))

	var book1 Book
	book1.Title = "数学课本"
	book1.Author = "徐刚"
	book1.Subject = "数学"
	book1.BookID = 1

	book2 := Book { "语文课本", "徐刚", "语文", 2 }
	book3 := Book { Title: "英语课本", Subject: "英语" }	//忽略的字段为0或空字符串
	printBook(book1)
	printBook(book2)
	printBook(book3)

	sliceTest()
	twoDimensionArray()
	testMap()

	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
	phone = new(IPhone)
	phone.call()
}
//函数的定义：注意
//			 1 形参声明方式
//			 2 函数返回值类型声明
func swap(x, y string) (string, string) {
	return y, x
}

func printBook(book Book) {
	//Printf方法可以求值，但不能换行
	fmt.Printf("BookID & Author & Title：%d -- %s -- %s\n", book.BookID, book.Author, book.Title)		
	fmt.Println("作者是空字符串：%s\n", book.Author == "")
}

type Book struct {
	Title string
	Author string
	Subject string
	BookID int
}

func sliceTest() {
	arr := []int {1,2,3,4,5,6}
	s1 := arr[:]
	for e := range s1 {
		fmt.Printf("element is: %d\n", s1[e])
	}

	s2 := arr[1:3]
	fmt.Printf("s2's length: %d\n", len(s2))

	s3 := make([]int, 3, 5)	//初始长度是3，容量是4
	s3 = append(s3, 666,777)	//长度变为5，容量变为5
	fmt.Printf("s3's len is: %d & s3's capacity is: %d\n", len(s3), cap(s3))
	for e := range s3 {
		fmt.Printf("element is: %d\n", s3[e])
	}
	s3 = append(s3, 888)	//长度变为6，容量变为
	fmt.Printf("s3's len is: %d & s3's capacity is: %d\n", len(s3), cap(s3))
}

func twoDimensionArray() {
	var arr = [][]int {{1,2},{3,4},{5},{6,7},{8,9}}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("arr[%d][%d] = %d\n", i, j, arr[i][j])
		}
	}
}

func testMap() {
	nums := []int {111,222,333}
	for i, num := range nums {
		fmt.Printf("num %d is: %d\n", i, num)
	}

	cmap := map[string]string{"a":"I'm a","b":"I'm b"}
	for key, value := range cmap {
		fmt.Printf("%s's value is: %s\n", key, value)
	}

	var countryMap map[string] string		//声明map实例
	countryMap = make(map[string]string)	//用make函数创建map实例
	countryMap["中国"] = "北京"
	countryMap["美国"] = "华盛顿"
	for key, capital := range countryMap {
		fmt.Printf("%s 首都 是：%s\n", key, capital)
	}
	delete(countryMap, "美国")
	fmt.Println("删除美国后map中的结果是：")
	for key, capital := range countryMap {
		fmt.Printf("%s 首都 是：%s\n", key, capital)
	}
}

type Phone interface {
	call()
}

type NokiaPhone struct {}

type IPhone struct {}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am nokia，I can call you")
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone，I can call you")
}
