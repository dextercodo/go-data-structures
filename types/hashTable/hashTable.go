package hashtable

const ArraySize = 3

func Init() *HashTable {
	var ht HashTable
	for i := range ht.arr {
		ht.arr[i] = &bucket{}
	}
	return &ht
}

type HashTable struct {
	arr [ArraySize]*bucket
}

func (ht *HashTable) Insert(key string) {
	index := hash(key)
	ht.arr[index].insert(key)
}

func (ht *HashTable) Search(key string) bool {
	index := hash(key)
	return ht.arr[index].search(key)
}

func (ht *HashTable) Delete(key string) {
	index := hash(key)
	ht.arr[index].delete(key)
}
