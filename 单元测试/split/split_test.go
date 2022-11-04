package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) { //单元测试
	got := Split("我爱你", "爱")           //got : 得到的
	want := []string{"我", "你"}         // want ： 想得到的
	if !reflect.DeepEqual(got, want) { // got 和 want 比较
		t.Errorf("want:%v got:%v", want, got)
	}
}

func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected: %v, got:%v", want, got)
	}
}

// 子测试
func TestSplit2(t *testing.T) {
	type test struct { //定义结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ //测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}
	for name, tc := range tests { //对tests map序列 遍历
		t.Run(name, func(t *testing.T) { //使用t.Run()执行子测试
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected: %#v, got: %#v", tc.want, got) //将测试用例的name格式化输出
			}
		})
	}
}

//测试覆盖率 : go test -cover  -coverprofile=c.out; Go还提供了一个额外的-coverprofile参数，
//用来将覆盖率相关的记录信息输出到一个文件
/*
基准测试:
基准测试就是在一定的工作负载之下检测程序性能的一种方法。基准测试的基本格式如下：
func BenchmarkName(b *testing.B){
	// ...
}
基准测试以Benchmark为前缀，需要一个*testing.B类型的参数b，基准测试必须要执行b.N次，这样的测试才有对照性，
b.N的值是系统根据实际情况去调整的，从而保证测试的稳定性。
*/
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}

/*基准测试并不会默认执行，需要增加-bench参数，所以我们通过执行go test -bench=Split命令执行基准测试，
/go test -bench=Split
split $ go test -bench=Split
goos: darwin
goarch: amd64
pkg: github.com/Q1mi/studygo/code_demo/test_demo/split
BenchmarkSplit-8        10000000               203 ns/op
PASS
ok      github.com/Q1mi/studygo/code_demo/test_demo/split       2.255s
---
其中BenchmarkSplit-8表示对Split函数进行基准测试，数字8表示GOMAXPROCS的值，这个对于并发基准测试很重要。
10000000和203ns/op表示每次调用Split函数耗时203ns，这个结果是10000000次调用的平均值。

我们还可以为基准测试添加-benchmem参数，来获得内存分配的统计数据。
split $ go test -bench=Split -benchmem
goos: darwin
goarch: amd64
pkg: github.com/Q1mi/studygo/code_demo/test_demo/split
BenchmarkSplit-8        10000000               215 ns/op             112 B/op          3 allocs/op
PASS
ok      github.com/Q1mi/studygo/code_demo/test_demo/split       2.394s
*/

/*性能比较函数
性能比较函数通常是一个带有参数的函数，被多个不同的Benchmark函数传入不同的值来调用。举个例子如下：

func benchmark(b *testing.B, size int){ ... }
func Benchmark10(b *testing.B){ benchmark(b, 10) }
func Benchmark100(b *testing.B){ benchmark(b, 100) }
func Benchmark1000(b *testing.B){ benchmark(b, 1000) }
*/

/*重置时间
b.ResetTimer之前的处理不会放到执行时间里，也不会输出到报告中，所以可以在之前做一些不计划作为测试报告的操作。例如：

func BenchmarkSplit(b *testing.B) {
	time.Sleep(5 * time.Second) // 假设需要做一些耗时的无关操作
	b.ResetTimer()              // 重置计时器
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}
*/
