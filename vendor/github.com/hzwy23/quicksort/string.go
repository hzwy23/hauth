package quicksort

type stringQuickSort struct {
}

func (this *stringQuickSort) Sort(val interface{}) {
	low := 0
	svd := val.([]string)
	hight := len(svd) - 1
	this.sort(low, hight, svd)
}

func (this *stringQuickSort) sort(low, hight int, svd []string) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint, svd)
		this.sort(low, pivotpoint-1, svd)
		this.sort(pivotpoint+1, hight, svd)
	}
}

func (this *stringQuickSort) medianOfThree(low, hight int, svd []string) (string, int) {
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

func (this *stringQuickSort) partition(low, hight int, pivotpoint *int, svd []string) {

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
	register("string", new(stringQuickSort))
}
