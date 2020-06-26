package main

var (
	outDir = "./image/"
)

func main() {
	// 生成直方图
	histogram("./data/iris.csv")

	// 生成箱型图
	boxPlot("./data/iris.csv")
}
