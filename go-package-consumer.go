package main

import (
	"github.com/olekukonko/tablewriter"
	"os"
)
//这个可以作为需要生成表格的所有的代码的处理方式我们可以去啦 无想去将变成微服务变成高可用的可以有自动 的cpu 内初
//bianc weh.yaml docke.yaml doa.eynod
// goiangle
//ci-cd
//可以配置
//面向接口变成
//
func main()  {
	data := [][]string{
		[]string{"Alfred", "15", "10/20", "(10.32, 56.21, 30.25)"},
		[]string{"Beelzebub", "30", "30/50", "(1,1,1)"},
		[]string{"Hortense", "21", "80/80", "(1,1,1)"},
		[]string{"Pokey", "8", "30/40", "(1,1,1)"},
	}
	table := tablewriter.NewWriter(os.Stdout)//设置输出的dout 地方hsez dta rneder
	//
	table.SetHeader([]string{"NPC", "Speed", "Power", "Location"})
	table.AppendBulk(data)
	table.Render()
	//表格的内容打印出来 设置 标题 数据处理渲染 到的滴定法 漏

}