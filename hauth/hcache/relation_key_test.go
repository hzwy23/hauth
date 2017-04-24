package hcache

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	relKeyObj := NewRelationKey()

	relKeyObj.SetRelatKey("A", "ABC")
	relKeyObj.SetRelatKey("B", "ABC")
	relKeyObj.SetRelatKey("C", "ABC")
	relKeyObj.SetRelatKey("E", "B")
	relKeyObj.SetRelatKey("F", "B")
	relKeyObj.SetRelatKey("G", "B")
	relKeyObj.SetRelatKey("H", "F")
	relKeyObj.SetRelatKey("I", "E")
	relKeyObj.SetRelatKey("I", "A")
	relKeyObj.SetRelatKey("I", "C")
	fmt.Println(relKeyObj.GetRelatKey("ABC"))
	fmt.Println(relKeyObj.GetRelatKey("H"))
	fmt.Println(relKeyObj.GetRelatKey("DEL"))
	fmt.Println(relKeyObj.GetRelatKey("Enter"))
}
