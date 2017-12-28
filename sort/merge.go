package sort

// 归并排序
//  利用归并的思想实现的排序方法，该算法采用经典的分治策略（分治法将问题分成一些小的问题然后递归求解，
//  而治的阶段则将分的阶段得到的答案“修补”在一起，即分而治之）

type Merge struct {
	context []Contexter
}

func NewMerge(context ...Contexter) Sorter {
	tmp := new(Merge)
	for _, v := range context {
		if v == nil {
			continue
		}
		tmp.context = append(tmp.context, v)
	}
	return tmp
}

func (this *Merge) Append(context ...Contexter) Sorter {
	for _, v := range context {
		if v == nil {
			continue
		}
		this.context = append(this.context, v)
	}
	return this
}

func (this *Merge) Sort() []Contexter {
	this.sort(0, len(this.context)-1)
	return this.context
}

func (this *Merge) sort(low, high int) {
	if low < high {
		mid := (low + high) / 2
		this.sort(low, mid)
		this.sort(mid+1, high)
		this.merge(low, mid, high)
	}
}

func (this *Merge) merge(low, mid, high int) {
	i := low
	j := mid + 1
	t := 0
	temp := make([]Contexter, high-low+1)
	for i <= mid && j <= high {
		if this.context[j].Compare(this.context[i]) {
			temp[t] = this.context[i]
			i++
		} else {
			temp[t] = this.context[j]
			j++
		}
		t++
	}
	for i <= mid {
		temp[t] = this.context[i]
		t++
		i++
	}
	for j <= high {
		temp[t] = this.context[j]
		t++
		j++
	}
	i = low
	t = 0
	for i <= high {
		this.context[i] = temp[t]
		t++
		i++
	}
}
