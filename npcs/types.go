package npcs
//定义类型 type 结构声明 定义 再操作
//Power 描述这个攻击的NPC的力量指标
type Power struct {
	Attack int
	Defense int
}
//wherfe inthe vie NPC exites
type Location struct {
	X float64
	Y float64
	Z float64
}
type NonPlayerCharacter struct {
	Name string
	Speed int
	HP int
	Power Power
	Loc Location
}
