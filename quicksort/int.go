package quicksort

type intQuickSort struct {
}

func (this *intQuickSort) Sort(val interface{}) {
	low := 0
	svd := val.([]int)
	hight := len(svd) - 1
	this.intsort(low, hight, svd)
}

func (this *intQuickSort) intsort(low, hight int, svd []int) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint, svd)
		this.intsort(low, pivotpoint-1, svd)
		this.intsort(pivotpoint+1, hight, svd)
	}
}

func (this *intQuickSort) medianOfThree(low, hight int, svd []int) (int, int) {
	m := (hight + low) / 2
	if hight-low <= 1 {
		if svd[low] > svd[hight] {
			svd[low], svd[hight] = svd[hight], svd[low]
		}
		return svd[low], low
	}
	if svd[low] > svd[m] {
		svd[low], svd[m] = svd[m], svd[low]
	}
	if svd[low] > svd[hight] {
		svd[low], svd[hight] = svd[hight], svd[low]
	}
	if svd[m] > svd[hight] {
		svd[m], svd[hight] = svd[hight], svd[m]
	}
	svd[low+1], svd[m] = svd[m], svd[low+1]
	return svd[low+1], low + 1
}

func (this *intQuickSort) partition(low, hight int, pivotpoint *int, svd []int) {

	pivotitem, j := this.medianOfThree(low, hight, svd)
	max := hight - 1
	for i := low + 2; i <= max; i++ {
		if svd[i] < pivotitem {
			svd[i], svd[j] = svd[j], svd[i]
			j = i
		} else {
			svd[max], svd[i] = svd[i], svd[max]
			max--
			i--
		}
	}
	*pivotpoint = j

	/*
		var j = low
		//get first element as pivotitem.
		pivotitem := svd[low]

		for i := low + 1; i <= hight; i++ {
			if svd[i] < pivotitem {
				svd[i], svd[j] = svd[j], svd[i]
				j = i
			} else {
				for tj := hight; tj > i; tj-- {
					if svd[tj] < pivotitem {
						svd[tj], svd[i], svd[j] = svd[i], svd[j], svd[tj]
						hight--
						j = i + 1
						break
					}
					hight--
				}
			}
		}
		*pivotpoint = j
	*/
}

// int8
type int8QuickSort struct {
}

func (this *int8QuickSort) Sort(val interface{}) {
	low := 0
	svd := val.([]int8)
	hight := len(svd) - 1
	this.intsort(low, hight, svd)
}

func (this *int8QuickSort) intsort(low, hight int, svd []int8) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint, svd)
		this.intsort(low, pivotpoint-1, svd)
		this.intsort(pivotpoint+1, hight, svd)
	}
}

func (this *int8QuickSort) medianOfThree(low, hight int, svd []int8) (int8, int) {
	m := (hight + low) / 2
	if hight-low <= 1 {
		if svd[low] > svd[hight] {
			svd[low], svd[hight] = svd[hight], svd[low]
		}
		return svd[low], low
	}
	if svd[low] > svd[m] {
		svd[low], svd[m] = svd[m], svd[low]
	}
	if svd[low] > svd[hight] {
		svd[low], svd[hight] = svd[hight], svd[low]
	}
	if svd[m] > svd[hight] {
		svd[m], svd[hight] = svd[hight], svd[m]
	}
	svd[low+1], svd[m] = svd[m], svd[low+1]
	return svd[low+1], low + 1
}

func (this *int8QuickSort) partition(low, hight int, pivotpoint *int, svd []int8) {

	pivotitem, j := this.medianOfThree(low, hight, svd)
	max := hight - 1
	for i := low + 2; i <= max; i++ {
		if svd[i] < pivotitem {
			svd[i], svd[j] = svd[j], svd[i]
			j = i
		} else {
			svd[max], svd[i] = svd[i], svd[max]
			max--
			i--
		}
	}
	*pivotpoint = j
}

// int16
type int16QuickSort struct {
}

func (this *int16QuickSort) Sort(val interface{}) {
	low := 0
	svd := val.([]int16)
	hight := len(svd) - 1
	this.intsort(low, hight, svd)
}

func (this *int16QuickSort) intsort(low, hight int, svd []int16) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint, svd)
		this.intsort(low, pivotpoint-1, svd)
		this.intsort(pivotpoint+1, hight, svd)
	}
}

func (this *int16QuickSort) medianOfThree(low, hight int, svd []int16) (int16, int) {
	m := (hight + low) / 2
	if hight-low <= 1 {
		if svd[low] > svd[hight] {
			svd[low], svd[hight] = svd[hight], svd[low]
		}
		return svd[low], low
	}
	if svd[low] > svd[m] {
		svd[low], svd[m] = svd[m], svd[low]
	}
	if svd[low] > svd[hight] {
		svd[low], svd[hight] = svd[hight], svd[low]
	}
	if svd[m] > svd[hight] {
		svd[m], svd[hight] = svd[hight], svd[m]
	}
	svd[low+1], svd[m] = svd[m], svd[low+1]
	return svd[low+1], low + 1
}

func (this *int16QuickSort) partition(low, hight int, pivotpoint *int, svd []int16) {

	pivotitem, j := this.medianOfThree(low, hight, svd)
	max := hight - 1
	for i := low + 2; i <= max; i++ {
		if svd[i] < pivotitem {
			svd[i], svd[j] = svd[j], svd[i]
			j = i
		} else {
			svd[max], svd[i] = svd[i], svd[max]
			max--
			i--
		}
	}
	*pivotpoint = j
}

// int32
type int32QuickSort struct {
}

func (this *int32QuickSort) Sort(val interface{}) {
	low := 0
	svd := val.([]int32)
	hight := len(svd) - 1
	this.intsort(low, hight, svd)
}

func (this *int32QuickSort) intsort(low, hight int, svd []int32) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint, svd)
		this.intsort(low, pivotpoint-1, svd)
		this.intsort(pivotpoint+1, hight, svd)
	}
}

func (this *int32QuickSort) medianOfThree(low, hight int, svd []int32) (int32, int) {
	m := (hight + low) / 2
	if hight-low <= 1 {
		if svd[low] > svd[hight] {
			svd[low], svd[hight] = svd[hight], svd[low]
		}
		return svd[low], low
	}
	if svd[low] > svd[m] {
		svd[low], svd[m] = svd[m], svd[low]
	}
	if svd[low] > svd[hight] {
		svd[low], svd[hight] = svd[hight], svd[low]
	}
	if svd[m] > svd[hight] {
		svd[m], svd[hight] = svd[hight], svd[m]
	}
	svd[low+1], svd[m] = svd[m], svd[low+1]
	return svd[low+1], low + 1
}

func (this *int32QuickSort) partition(low, hight int, pivotpoint *int, svd []int32) {

	pivotitem, j := this.medianOfThree(low, hight, svd)
	max := hight - 1
	for i := low + 2; i <= max; i++ {
		if svd[i] < pivotitem {
			svd[i], svd[j] = svd[j], svd[i]
			j = i
		} else {
			svd[max], svd[i] = svd[i], svd[max]
			max--
			i--
		}
	}
	*pivotpoint = j
}

// int64
type int64QuickSort struct {
}

func (this *int64QuickSort) Sort(val interface{}) {
	low := 0
	svd := val.([]int64)
	hight := len(svd) - 1
	this.intsort(low, hight, svd)
}

func (this *int64QuickSort) intsort(low, hight int, svd []int64) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint, svd)
		this.intsort(low, pivotpoint-1, svd)
		this.intsort(pivotpoint+1, hight, svd)
	}
}

func (this *int64QuickSort) medianOfThree(low, hight int, svd []int64) (int64, int) {
	m := (hight + low) / 2
	if hight-low <= 1 {
		if svd[low] > svd[hight] {
			svd[low], svd[hight] = svd[hight], svd[low]
		}
		return svd[low], low
	}
	if svd[low] > svd[m] {
		svd[low], svd[m] = svd[m], svd[low]
	}
	if svd[low] > svd[hight] {
		svd[low], svd[hight] = svd[hight], svd[low]
	}
	if svd[m] > svd[hight] {
		svd[m], svd[hight] = svd[hight], svd[m]
	}
	svd[low+1], svd[m] = svd[m], svd[low+1]
	return svd[low+1], low + 1
}

func (this *int64QuickSort) partition(low, hight int, pivotpoint *int, svd []int64) {

	pivotitem, j := this.medianOfThree(low, hight, svd)
	max := hight - 1
	for i := low + 2; i <= max; i++ {
		if svd[i] < pivotitem {
			svd[i], svd[j] = svd[j], svd[i]
			j = i
		} else {
			svd[max], svd[i] = svd[i], svd[max]
			max--
			i--
		}
	}
	*pivotpoint = j
}

func init() {
	register("int", new(intQuickSort))
	register("int8", new(int8QuickSort))
	register("int16", new(int16QuickSort))
	register("int32", new(int32QuickSort))
	register("int64", new(int64QuickSort))
}
