package examples

import (
	hashtable "data-structures/types/hashTable"
	"fmt"
)

func HashTableExample() {
	ht := hashtable.Init()
	// fmt.Println(ht)

	ht.Insert("James")
	ht.Insert("Farrah")
	ht.Insert("Kamran")
	ht.Insert("Marvel")
	ht.Insert("Dumbass")
	ht.Insert("YelloMan")
	ht.Insert("Eric")
	ht.Insert("Randy")

	x := ht.Search("xxx")
	fmt.Println(x)
	x = ht.Search("Kamran")
	fmt.Println(x)
	x = ht.Search("Marvel")
	fmt.Println(x)

	ht.Delete("Kamran")
	x = ht.Search("Kamran")
	fmt.Println(x)

}
