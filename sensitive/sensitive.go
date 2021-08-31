package main

import (
	"fmt"
	"strings"

	"github.com/importcjj/sensitive"
)

func main() {
	filter := sensitive.New()
	filter.AddWord("垃圾")
	filter.AddWord("尼玛")

	text := "@Harvey(黄炜) 除了 这些推送条件外，可否增加一个推送内容尼玛修改的功能@Harvey(黄炜) 除了 这些推送条件外，可否增加一个推送内容修改的功能垃圾"

	title := "中尤为垃圾重要尼玛。你或多或少都会有因为自己的提交，导致应用挂掉或服务宕机的经历。如果这个时候你的修改导致测试用例失败，你再重新审视自己的修改，发现之前的修改还有一些特殊场景没有包含，恭喜你减少了一次上库失误垃圾。也会有这样的情况，项目很大，启动环境很复杂，你优化了一个函数的性能，或是添加了某个新的特性，如果部署在正式环境上之后再进行测试，成本太高。对于这种场景，几个小小的测试用例或许就能够覆盖大部分的测试场景。而且在开发过程中，效率最高的莫过于所见即所得了，单元测试也能够帮助你做到这一点，试想一下，假如你一口气写完一千行代码，debug 的过程也不会轻之前的例子测试失败时使用 t.Error/t.Errorf，这个例子中使用 t.Fatal/t.Fatalf，区别在于前者遇错不停，还会继续执行其他的测试用例，后者遇错即停s。"

	var sb strings.Builder
	sb.WriteString(text)
	sb.WriteString(title)
	a, b := filter.FindIn(sb.String())
	fmt.Printf("a = %#v, b = %#v\n", a, b)

}
