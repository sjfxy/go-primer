package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//基础的函数库
//定义函数签名

type dieRollFunc func(int) int
//返回类型的给定的size 随机的数据 + 1 1自诶的数据
//
func dieRoll(size int) int  {
	rand.Seed(time.Now().UnixNano())
	return  rand.Intn(size) + 1
}
func fakeDieRoll(size int) int  {
	return  42
}
func sinDieRoll(size int) int  {
	rand.Seed(time.Now().UnixNano())
	return  rand.Intn(size) + 2;
}
func rollTwo(size1,size2 int)(int,int)  {
	return dieRoll(size1),sinDieRoll(size2)
}
//简单的tasks[] type Tasks func() int
//变成很多的类型funcs 签名一样的[]接口类型的方法去凤凰对应的接口
//func即可
func getDieRolls()[]dieRollFunc  {
	return  []dieRollFunc{
		sinDieRoll,
		fakeDieRoll,
		dieRoll,
	}
}
//theResut , err error
// thereust, err
// + ",
// + strconv .Itoa

func returnsNamed(input1 string,input2 int)(theResult string,err error)  {
   theResult = "modified" + input1 +","+strconv.Itoa(input2)
   return theResult,err
}
var rools []dieRollFunc
func init()  {
	rools = getDieRolls()
}
//可以了解一下对应的流程 init
//全局的属性 ->init ->main ->fmt->Qaunj ->init->main即可
func main()  {
	fmt.Printf("Rolled a die of size %d,result:%d\n",6,dieRoll(6))
   res1, res2 := rollTwo(6,10)
   fmt.Printf("Rolled a pair of dice(%d,%d),results:%d,%d\n",
   	6,10,res1,res2)
   named, err := returnsNamed("sin",22)
   // %d %v %s ''的使用 input string strconv.Itoa 转换的操作
   // 处理方式
   fmt.Printf("Named params returnd: '%s', %v\n",named,err)
   for index, rollFunc := range rools {
   	fmt.Printf("Die Roll Attempt #%d, result: %d\n",index,rollFunc(10))
   }


}


