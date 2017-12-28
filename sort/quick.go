package sort

// 快速排序
// 通过一趟排序将待排序记录分割成独立的两部分，其中一部分记录的关键字均比另一部分关键字小，则分别继续对两部分进行排序，
//  直到整个序列有序
//

type Quick struct {
	context []Contexter
}

func NewQuick(context ...Contexter) Sorter {
	tmp := new(Quick)
	for _, v := range context {
		if v != nil {
			tmp.context = append(tmp.context, v)
		}
	}
	return tmp
}

func (this *Quick) Append(context ...Contexter) Sorter {
	for _, v := range context {
		if v != nil {
			this.context = append(this.context, v)
		}
	}
	return this
}

func (this *Quick) Sort() []Contexter {
	length := len(this.context)
	if length > 0 {
		this.quickSort(0, length-1)
	}
	return this.context
}

// 查找中轴
func (this *Quick) getMiddle(low, high int) int {
	// 定义中轴，一般将第一个元素作为中轴值，也可以通过math/rand方法在指定长度内随机指定一个中轴值
	middle := this.context[low]
	for low < high {
		for low < high && (this.context[high].Compare(middle) || this.context[high].Equals(middle)) {
			high--
		}
		this.context[low] = this.context[high]
		for low < high && !this.context[low].Compare(middle) {
			low++
		}
		this.context[high] = this.context[low]
	}
	this.context[low] = middle
	return low
}

// 递归形式的分治排序算法
func (this *Quick) quickSort(low, high int) {
	if low < high {
		middle := this.getMiddle(low, high)
		this.quickSort(low, middle-1)
		this.quickSort(middle+1, high)
	}
}
