package quicksort

//float32

type floatQuickSort32 struct {
}

func (this *floatQuickSort32) Sort(val interface{}) {
	low := 0
	svd := val.([]float32)
	hight := len(svd) - 1
	this.intsort(low, hight, svd)
}

func (this *floatQuickSort32) intsort(low, hight int, svd []float32) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint, svd)
		this.intsort(low, pivotpoint-1, svd)
		this.intsort(pivotpoint+1, hight, svd)
	}
}

func (this *floatQuickSort32) medianOfThree(low, hight int, svd []float32) (float32, int) {
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

func (this *floatQuickSort32) partition(low, hight int, pivotpoint *int, svd []float32) {

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

// float64
type floatQuickSort64 struct {
	val []float64
}

func (this *floatQuickSort64) Sort(val interface{}) {
	low := 0
	svd := val.([]float64)
	hight := len(svd) - 1
	this.intsort(low, hight, svd)
}

func (this *floatQuickSort64) intsort(low, hight int, svd []float64) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint, svd)
		this.intsort(low, pivotpoint-1, svd)
		this.intsort(pivotpoint+1, hight, svd)
	}
}

func (this *floatQuickSort64) medianOfThree(low, hight int, svd []float64) (float64, int) {
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

func (this *floatQuickSort64) partition(low, hight int, pivotpoint *int, svd []float64) {

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
	register("float32", new(floatQuickSort32))
	register("float64", new(floatQuickSort64))
}
