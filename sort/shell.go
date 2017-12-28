package sort

// 希尔排序
// 先将整个待排序的记录序列分割为若干子序列，分别进行直接插入排序，待整个序列中的记录“基本有序”时，再对整体记录进行依次直接插入排序。

type Shell struct {
	context []Contexter
}

func NewShell(context ...Contexter) Sorter {
	tmp := new(Shell)
	for _, v := range context {
		if v == nil {
			continue
		}
		tmp.context = append(tmp.context, v)
	}
	return tmp
}

func (this *Shell) Append(context ...Contexter) Sorter {
	for _, v := range context {
		if v == nil {
			continue
		}
		this.context = append(this.context, v)
	}
	return this
}

func (this *Shell) Sort() []Contexter {
	this.increment(len(this.context)/2 + len(this.context)%2)
	return this.context
}

func (this *Shell) increment(gap int) int {
	for i := 0; i < gap; i++ {
		for j := i; j < len(this.context)-gap; j += gap {
			if this.context[j].Compare(this.context[j+gap]) {
				this.context[j], this.context[j+gap] = this.context[j+gap], this.context[j]
			}
		}
	}
	if gap == 1 {
		return 0
	}
	return this.increment(gap/2 + gap%2)
}
