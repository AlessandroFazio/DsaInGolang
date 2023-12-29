package hashtables

const ArraySize = 10

type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key string
	next *bucketNode
}

func (t *HashTable) hasBucket(index int) bool {
	return t.array[index] != nil
}

func (t *HashTable) Insert(key string) {
	index := hash(key)
	if !t.hasBucket(index) {
		t.array[index] = &bucket{}
	}
	t.array[index].insert(key)
}

func (t *HashTable) Search(key string) bool {
	index := hash(key)
	if !t.hasBucket(index) {
		return false
	}
	return t.array[index].search(key)
}

func (t *HashTable) Delete(key string) bool {
	index := hash(key)
	if !t.hasBucket(index) {
		return false
	}
	return t.array[index].delete(key)
}

func (b *bucket) insert(key string) {
	if !b.searchInsert(key) {
		newNode := &bucketNode{key: key}
		newNode.next = b.head
		b.head = newNode
	}
}

func (b *bucket) search(key string) bool {
	current := b.head
	for current != nil {
		if current.key == key {
			return true
		} 
		current = current.next
	}
	return false
}

func (b *bucket) searchInsert(key string) bool {
	current := b.head
	for current != nil {
		if current.key == key {
			return true
		}
		current = current.next
	}
	return false
}

func (b *bucket) delete(key string) bool {
	if b.head.key == key {
		b.head = b.head.next
		return true
	}
	previous := b.head
	for previous.next != nil {
		if previous.next.key == key {
			previous.next = previous.next.next
			return true
		}
		previous = previous.next
	}
	return false
}


func hash(key string) int {
	sum := 0
	for _,v := range key{
		sum += int(v)
	}
	return sum % ArraySize
}
