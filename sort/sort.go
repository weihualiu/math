package sort

// 排序

// 要排序内容的接口
// Compare function: 返回true表示正序排列(对象大于参数)，返回false表示倒序排列(对象小于参数)
// Equals function: 返回true表示
type Contexter interface {
	Compare(Contexter) bool
	Equals(Contexter) bool
}

// 排序算法接口
type Sorter interface {
	// 正序排序函数
	Sort() []Contexter
	// 组装原始未排序的值
	Append(context ...Contexter) Sorter
}

// 通用指定排序算法来进行排序的函数
func New(f func(context ...Contexter) Sorter, context ...Contexter) Sorter {
	return f(context...)
}

// 直接指定排序算法，传入待排序数据，返回已排序结果
func Sort(f func(context ...Contexter) Sorter, context ...Contexter) []Contexter {
	return f(context...).Sort()
}

// 指定排序算法并将结果进行反序排列
func SortReserve(f func(context ...Contexter) Sorter, context ...Contexter) []Contexter {
	var temp []Contexter
	result := Sort(f, context...)
	length := len(result)
	for i := length; i > 0; i-- {
		temp = append(temp, result[i])
	}
	return temp
}
