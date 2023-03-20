package order_map

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Complex interface {
	~complex64 | ~complex128
}

type Ordered interface {
	Integer | Float | ~string
}

type OrderMap[K Ordered, V any] struct {
	data      map[K]V
	orderList []K
}

func New[K Ordered, V any]() *OrderMap[K, V] {
	return &OrderMap[K, V]{
		data:      make(map[K]V),
		orderList: make([]K, 0),
	}
}

func (o *OrderMap[K, V]) Get(key K) V {
	return o.data[key]
}

func (o *OrderMap[K, V]) Set(key K, value V) {
	if _, ok := o.data[key]; ok {
		return
	}

	o.data[key] = value

	left, right := 0, len(o.orderList)
	for left < right {
		mid := (left + right) / 2
		if o.orderList[mid] < key {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left == len(o.orderList) {
		o.orderList = append(o.orderList, key)
	} else {
		nList := make([]K, len(o.orderList)+1)
		copy(nList, o.orderList[:left])
		copy(nList[left+1:], o.orderList[left:])
		nList[left] = key
		o.orderList = nList
	}
}
