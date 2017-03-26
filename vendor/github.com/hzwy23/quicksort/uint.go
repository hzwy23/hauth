package quicksort

type uintQuickSort struct {
}

func (this *uintQuickSort) Sort(val interface{}) {
	low := 0
	svd := val.([]uint)
	hight := len(svd) - 1
	this.intsort(low, hight, svd)
}

func (this *uintQuickSort) intsort(low, hight int, svd []uint) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint, svd)
		this.intsort(low, pivotpoint-1, svd)
		this.intsort(pivotpoint+1, hight, svd)
	}
}

func (this *uintQuickSort) medianOfThree(low, hight int, svd []uint) (uint, int) {
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

func (this *uintQuickSort) partition(low, hight int, pivotpoint *int, svd []uint) {

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
type uint8QuickSort struct {
}

func (this *uint8QuickSort) Sort(val interface{}) {
	low := 0
	svd := val.([]uint8)
	hight := len(svd) - 1
	this.intsort(low, hight, svd)
}

func (this *uint8QuickSort) intsort(low, hight int, svd []uint8) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint, svd)
		this.intsort(low, pivotpoint-1, svd)
		this.intsort(pivotpoint+1, hight, svd)
	}
}

func (this *uint8QuickSort) medianOfThree(low, hight int, svd []uint8) (uint8, int) {
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

func (this *uint8QuickSort) partition(low, hight int, pivotpoint *int, svd []uint8) {

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
type uint16QuickSort struct {
}

func (this *uint16QuickSort) Sort(val interface{}) {
	low := 0
	svd := val.([]uint16)
	hight := len(svd) - 1
	this.intsort(low, hight, svd)
}

func (this *uint16QuickSort) intsort(low, hight int, svd []uint16) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint, svd)
		this.intsort(low, pivotpoint-1, svd)
		this.intsort(pivotpoint+1, hight, svd)
	}
}

func (this *uint16QuickSort) medianOfThree(low, hight int, svd []uint16) (uint16, int) {
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

func (this *uint16QuickSort) partition(low, hight int, pivotpoint *int, svd []uint16) {

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
type uint32QuickSort struct {
}

func (this *uint32QuickSort) Sort(val interface{}) {
	low := 0
	svd := val.([]uint32)
	hight := len(svd) - 1
	this.intsort(low, hight, svd)
}

func (this *uint32QuickSort) intsort(low, hight int, svd []uint32) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint, svd)
		this.intsort(low, pivotpoint-1, svd)
		this.intsort(pivotpoint+1, hight, svd)
	}
}

func (this *uint32QuickSort) medianOfThree(low, hight int, svd []uint32) (uint32, int) {
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

func (this *uint32QuickSort) partition(low, hight int, pivotpoint *int, svd []uint32) {

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
type uint64QuickSort struct {
}

func (this *uint64QuickSort) Sort(val interface{}) {
	low := 0
	svd := val.([]uint64)
	hight := len(svd) - 1
	this.intsort(low, hight, svd)
}

func (this *uint64QuickSort) intsort(low, hight int, svd []uint64) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint, svd)
		this.intsort(low, pivotpoint-1, svd)
		this.intsort(pivotpoint+1, hight, svd)
	}
}

func (this *uint64QuickSort) medianOfThree(low, hight int, svd []uint64) (uint64, int) {
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

func (this *uint64QuickSort) partition(low, hight int, pivotpoint *int, svd []uint64) {

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
	register("uint", new(uintQuickSort))
	register("uint8", new(uint8QuickSort))
	register("uint16", new(uint16QuickSort))
	register("uint32", new(uint32QuickSort))
	register("uint64", new(uint64QuickSort))
}
