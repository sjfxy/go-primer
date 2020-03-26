package main

import (
	"fmt"
	"www.cloudnative.com/go-primer/npcs"
)

func main()  {
	mob := npcs.NonPlayerCharacter{Name:"Sin"}
	fmt.Println(mob)
	//String()一次的调用即可
	//每个都会默认的调用 对象
	// types 定义对应的对象 struct
	// npcs 对应的每个人提供的接口实现类你好呢的hi哦的
	// 仓库的去寻找然后对象池的使用接口的规范
	//封装的过程
	// Power LCoation locayon Nkep Power locaton nagke
	// String()
	//location
	//自动调用 mob.String()
	//npc.Name npc.Loc String()
	hortense := npcs.NonPlayerCharacter{Name: "fxy",
		Power: npcs.Power{Attack: 12, Defense: 13},
		Loc:   npcs.Location{X: 10.0, Y: 23.0, Z: 12.0},
	}
	fmt.Println(hortense)
	fmt.Printf("Sin is %f units from fxy \n",mob.DistanceTo(hortense))
}
//对象比较的时候 直接的loc.x loc.当前的实例对象 this.x -tar.loca
//npc.loca=thisnp.clocatoar.x /y
//我们只能说这个Go语言的很舒服的

