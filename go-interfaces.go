package main

import "fmt"

//这个是接口的测试
type attacker struct {
	attackerpower int
	dmggonus int
}
type sword struct {
	attacker
	twohanded bool
}
//这个基础的组合的方式不好的地方是 基础损坏则导致组合的收到升或者降级幸运
type gun struct {
	attacker
	bullestreamining int
}
type chair struct {
	legcount int
	leather bool
}
type weapon interface {
	Wield() bool
}
type Fun func() bool
//这个最后变成了 Wield
//我们封装了过去fun() bool
//即可
// 但是我们Fun(qqlogin)->Interface 即可
func (t Fun) Wield() bool  {
	return  t()
}
//因为只需要给包含的bool即可
//我们Fun()对应的func()类型的但是不是
//Fun func()
//func wikld
// type SinFunc func(w http.Requestwwo,r *http.Request)()
// func(f SinFunc) WideNew()
//即可
func (s sword) Wield() bool  {
 fmt.Println("You've wielded a sword!")
 return  true
}
func (g gun) Wield() bool  {
	fmt.Println("You've wielded a gun!")
	return  true
}
func (c chair) Wield() bool  {
	fmt.Println("You've wielded a chair! You tm 是封了吗?")
	return  true
}
func  wielder(w weapon) bool  {
	fmt.Println("Wielding....")
	return w.Wield()
}
func wiedld() bool  {
	fmt.Println("我是随便的函数 我需要自己伪造吗？")
	return  true
}
func main()  {
	sword1 := sword{
		attacker{attackerpower:1,dmggonus:5},
		true,
	}
	gun1 := gun{
		attacker{attackerpower:10,dmggonus:20},
		11,
	}
	chair1 := chair{legcount:3,leather:true}

	fmt.Printf("Weapons:sword:%v, gun:%v\n ",sword1,gun1)

	wielder(sword1)
	wielder(gun1)
	wielder(chair1)

	ce := Fun(wiedld)
	wielder(ce)
	//这个就是dunckeftupe functier
	// type unc  wikefd
	// func f ufn c
	// 这个会有一个就是 Nodo
	// func
	// func
	// int
	// bool
	// fun
	// weiler
	// bool
	//去完成了我需要吗


}


