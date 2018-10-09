package computer

type Spec struct {
	Maker string //大写声明的字段可以被导出
	model string //小写声明的字段不能导出
	Price int    //大写声明的字段可以被导出
}
