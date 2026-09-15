package main

import "fmt"

func main() {
	// 1. 创建一个包含10个整数的切片，初始值为1到10
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("初始切片:", slice)

	// 2. 使用切片操作获取第3到第7个元素（包含第7个）
	// 注意：Go语言中切片的索引从0开始，所以第3个元素的索引是2，第7个元素的索引是6
	slice = slice[2:7] // 包含索引2，不包含索引7，所以是2-6的元素
	fmt.Println("获取第3到第7个元素:", slice)

	// 3. 在切片末尾添加三个新元素：11, 12, 13
	slice = append(slice, 11, 12, 13)
	fmt.Println("添加新元素后:", slice)

	// 4. 删除切片中的第5个元素
	// 注意：此时切片的长度已经变化，第5个元素的索引是4
	slice = append(slice[:4], slice[5:]...)
	fmt.Println("删除第5个元素后:", slice)

	// 5. 将切片中的所有元素乘以2
	for i := range slice {
		slice[i] *= 2
	}
	fmt.Println("所有元素乘以2后:", slice)

	// 6. 打印最终切片的内容和容量
	fmt.Println("最终切片内容:", slice)
	fmt.Println("最终切片容量:", cap(slice))
}
