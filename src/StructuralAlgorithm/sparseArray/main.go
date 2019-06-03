package main

import "fmt"

/*
	稀疏数组:
		当一个数组中大部分元素为0，或者为同一个值得数组时，可以使用稀疏数组来保存该数组

	稀疏数组的处理方方是：
		1.记录数组一共有几行几列，有多少个不同的值
		2.把具有不同值得元素的行列及值记录在一个小规模的数组中，从而缩小程序的规模

	行(row) 列(col) 值(value)
	  11	  11		0    (第一行代表总的11*11 值0)
	  1        2	    1
 	  2	 	   3        2
*/

type ValNode struct {
	row int
	col int
	val int
}

func main(){
	//1.先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑棋
	chessMap[2][3] = 2 //蓝棋

	//2.输出看看原始数组
	for _,v := range chessMap{
		for _,v2 := range v{
			fmt.Printf("%d\t",v2)
		}
		fmt.Println()
	}


	//3.转成稀疏数组
	//(1).遍历chessMap,如果我们发现有一个元素的值不为0.创建一个node结构体
	//(2).将其放入到对应的切片中

	var sparseArr []ValNode
	//标准的一个稀疏数组应该还有一个表示记录原始二维数组的规模的行和列和默认值
	valNode := ValNode{
		row:11,
		col:11,
		val:0,
	}
	sparseArr = append(sparseArr,valNode)

	for i,v := range chessMap{
		for j,v2 := range v{
			if v2 != 0{
				//创建一个ValNode值节点
				valNode := ValNode{
					row:i,
					col:j,
					val:v2,
				}
				sparseArr = append(sparseArr,valNode)
			}
		}
	}

	//输出稀疏数组
	fmt.Println("当前的稀疏数组")
	for i,valNode := range sparseArr{
		fmt.Printf("%d: %d %d %d\n",i,valNode.row,valNode.col,valNode.val)
	}

	//将这个稀疏数组存盘

	//如何恢复原始的数组
	//1.打开存盘的文件=》恢复原始数组
	//2.这里使用稀疏数组恢复

	//先创建一个原始的数组
	var chessMap2 [11][11]int
	//遍历sparseArr[]遍历文件的每一行
	for i,valNode := range sparseArr{
		if i != 0{ //跳过第一行 第一行是规模
			chessMap2[valNode.row][valNode.col] = valNode.val
		}
	}
	//看看chessMap2是不是恢复了
	fmt.Println("恢复后的原始数据")
	for _,v := range chessMap2{
		for _,v2 := range v{
			fmt.Printf("%d\t",v2)
		}
		fmt.Println()
	}
}