package quicksort_test

import (
	"fmt"
	"sort"
	"testing"
	"time"

	"github.com/hzwy23/quicksort"
)

func TestSort(t *testing.T) {

	var t2 []int = []int{11, 31, 43, 3, 48, 23}
	quicksort.QuickSort(t2)
	fmt.Println("int", t2)

	var t28 []int8 = []int8{11, 31, 43, 3, 48, 23}
	quicksort.QuickSort(t28)
	fmt.Println("int8", t28)

	var t216 []int16 = []int16{11, 31, 43, 3, 48, 23}
	quicksort.QuickSort(t216)
	fmt.Println("int16", t216)

	var t232 []int32 = []int32{11, 31, 43, 3, 48, 23}
	quicksort.QuickSort(t232)
	fmt.Println("int32", t232)

	var t264 []int64 = []int64{11, 31, 43, 3, 48, 23}
	quicksort.QuickSort(t264)
	fmt.Println("int64", t264)

	var t3 []string = []string{"123", "211", "abc", "def", "1sd", "6dw", "235"}
	quicksort.QuickSort(t3)
	fmt.Println("string", t3)

	var tf32 []float32 = []float32{3.2345, 1.242345, 3.11234, 6.41234, 4.11235}
	quicksort.QuickSort(tf32)
	fmt.Println("float32", tf32)

	var t4 []float64 = []float64{3.2345, 1.242345, 3.11234, 6.41234, 4.11235}
	quicksort.QuickSort(t4)
	fmt.Println("float64", t4)
	var t8 []int8 = []int8{19, 43, 63, 27, 41, 24, 64, 24, 64, 34, 65, 24, 6, 1, 23, 5, 43, 6, -1, -5, -2, 23, 4}
	quicksort.QuickSort(t8)
	fmt.Println("t8", t8)
	var t16 []int16 = []int16{19, 43, 63, 27, 41, 24, 64, 24, 64, 234, 645, 234, 6, 1, 23, 5, 43, 6, -1, -5, -2, 4123, 4234}
	quicksort.QuickSort(t16)
	fmt.Println("t16", t16)
	var t32 []int32 = []int32{19, 43, 63, 27, 41, 24, 64, 24, 64, 234, 64345, 234, 6, 1, 23, 5, 43, 6, -1, -5, -2, 4123, 4234}
	quicksort.QuickSort(t32)
	fmt.Println("t32", t32)
	var t64 []int64 = []int64{19, 43, 63, 27, 41, 24, 64, 24, 64, 234, 64345, 234, 6, 1, 23, 5, 43, 6, -1, -5, -2, 4123, 4234}
	quicksort.QuickSort(t64)
	fmt.Println("t64", t64)

	var utint []uint = []uint{12, 3, 4, 5}
	quicksort.QuickSort(utint)
	fmt.Println("utint", utint)

	var utint8 []uint8 = []uint8{12, 3, 4, 5}
	quicksort.QuickSort(utint8)
	fmt.Println("utint", utint8)

	var utint16 []uint16 = []uint16{12, 3, 4, 5}
	quicksort.QuickSort(utint16)
	fmt.Println("utint", utint16)

	var utint32 []uint32 = []uint32{12, 3, 4, 5}
	quicksort.QuickSort(utint32)
	fmt.Println("utint", utint32)

	var utint64 []uint64 = []uint64{12, 3, 4, 5}
	quicksort.QuickSort(utint64)
	fmt.Println("utint", utint64)

}

func TestSysSort(t *testing.T) {
	fmt.Println("start system sort")
	curTime := time.Now().Unix()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	for i := 0; i < 10000000; i++ {
		var t []int = []int{11, -1, 3, 4, 5, 12, 545, 67456, 34, 6, 3, 45, 7, 3, 5, 5, 6, 3, 3, 4, 7, 764, 465, 5423}
		a := sort.IntSlice(t[0:])
		sort.Sort(a)
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	endTime := time.Now().Unix()
	fmt.Println("end system sort,costs time:", (endTime - curTime))
}

func TestMySort(t *testing.T) {
	fmt.Println("start my sort")
	curTime := time.Now().Unix()
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	for i := 0; i < 10000000; i++ {
		var t []int = []int{11, -1, 3, 4, 5, 12, 545, 67456, 34, 6, 3, 45, 7, 3, 5, 5, 6, 3, 3, 4, 7, 764, 465, 5423}
		quicksort.QuickSort(t)
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	endTime := time.Now().Unix()
	fmt.Println("end my sort,costs time:", (endTime - curTime))
}
