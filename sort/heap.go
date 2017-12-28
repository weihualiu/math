package sort

// 堆排序
// 将待排序序列构造成一个大顶堆，此时，整个序列的最大值就是堆顶的根节点。将其与末尾元素进行交换，此时末尾就为最大值。
//  然后将剩余n-1个元素重新构造成一个堆，这样会得到n个元素的次小值。如此反复执行，就得到一个有序序列了。

// 大顶堆/小顶堆

// 大顶堆/小顶堆

type Heap struct {
	context []Contexter
}

func NewHeap(context ...Contexter) Sorter {
	tmp := new(Heap)
	for _, v := range context {
		if v == nil {
			continue
		}
		tmp.context = append(tmp.context, v)
	}
	return tmp
}

func (this *Heap) Append(context ...Contexter) Sorter {
	for _, v := range context {
		if v == nil {
			continue
		}
		this.context = append(this.context, v)
	}
	return this
}

func (this *Heap) Sort() []Contexter {
	length := len(this.context)
	for i := length/2 - 1; i >= 0; i-- {
		this.adjustHeap(i, length)
	}
	for i := length - 1; i > 0; i-- {
		this.context[0], this.context[i] = this.context[i], this.context[0]
		this.adjustHeap(0, i)
	}
	return this.context
}

func (this *Heap) adjustHeap(i, length int) {
	temp := this.context[i]
	for k := i*2 + 1; k < length; k = k*2 + 1 {
		if k+1 < length && this.context[k+1].Compare(this.context[k]) {
			k++
		}
		if this.context[k].Compare(temp) {
			this.context[i] = this.context[k]
			i = k
		} else {
			break
		}
	}
	this.context[i] = temp
}
