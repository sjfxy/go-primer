package npcs

import (
	"fmt"
	"math"
)

//对应的类型的提供的方法
func (loc Location) String() string  {
	return  fmt.Sprintf("(%f,%f,%f)",loc.X,loc.Y,loc.Z)
}
//使用对Location 提供计算距离的操作
// float64 00.00

func (power Power) String() string  {
	return  fmt.Sprintf("(这个是我们内部提供的String方法 %d,%d)",power.Attack,power.Defense)
}
func (loc Location) EuclideanDistance(target Location) float64  {
	return  math.Sqrt(
		(loc.X - target.X) *(loc.X-target.X)+
			(loc.Y -target.Y) * (loc.Y-target.Y) +
			(loc.Z -target.Z) * (loc.Z -target.Z))

}
//我们可以随便的时候
// Nppalayer Power Poweer Loc Location
// func(loca Location)
// npc.loc.ElcDistanef()
//tar.loc
//即可

func (npc NonPlayerCharacter)DistanceTo(target NonPlayerCharacter) float64 {
return  npc.Loc.EuclideanDistance(target.Loc)
}
//没有任何的侵入的代码和设计模型
// String()
//在打印的对象的时候会主动的带有 Loc.String()
//我们对应的都提供的对应的打印一个对象的String
// 我们Spring %v Nodplayefc
//会自动的调用String()
//然后对应的对象 则会地质队 地道又Loccation.String()
// Power.String()
// go test
// go get
// golide
// go buled
// go env
// go run
// push
// git/Gitlab/第三方的仓库和
//dockerf.yaml
//werkef.yaml
//daoclu.yaml
//对应的yaml
//push->git->wehdcker->自动执行 wejceklya.
//测试 检测 ci部署
//推送到rwsgiert
//仓库
//即可
//设置对应的dockerf_hub的地址密码
//这里可以设置对应的regsier一切的用户名和密码
//然后我们进行发布不说
//rune/pop
// mogjnd
// mysql
//myod
//doke
//

func (npc NonPlayerCharacter) String() string {
return  fmt.Sprintf("%s %s %s",npc.Name,npc.Loc,npc.Power)
}

