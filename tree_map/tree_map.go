package tree_map

type USign interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}
type Sign interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
type Integer interface {
	USign | Sign
}
type Float interface {
	~float32 | ~float64
}
type Ordered interface {
	Integer | Float | ~string
}

type TreeMap[K Ordered, V any] struct {
	root *node[K, V]
}

type node[K Ordered, V any] struct {
	parent *node[K, V]
	left   *node[K, V]
	right  *node[K, V]

	key   K
	value V
}

func (n *node[K, V]) height() int {
	if n == nil {
		return 0
	}

	if n.left == nil && n.right == nil {
		return 1
	} else if n.left == nil {
		return n.right.height() + 1
	} else if n.right == nil {
		return n.left.height() + 1
	} else {
		left := n.left.height()
		right := n.right.height()

		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}
}

func rightRotate[K Ordered, V any](n *node[K, V]) *node[K, V] {
	nn := n.left

	if n.parent != nil {
		if n.parent.left == n {
			n.parent.left = nn
		} else {
			n.parent.right = nn
		}
	}
	nn.parent = n.parent

	n.left = nn.right
	if nn.right != nil {
		nn.right.parent = n
	}

	nn.right = n
	n.parent = nn

	return nn
}

func leftRotate[K Ordered, V any](n *node[K, V]) *node[K, V] {
	nn := n.right

	if n.parent != nil {
		if n.parent.right == n {
			n.parent.right = nn
		} else {
			n.parent.left = nn
		}
	}
	nn.parent = n.parent

	n.right = nn.left
	if nn.left != nil {
		nn.left.parent = n
	}

	nn.left = n
	n.parent = nn

	return nn
}

func New[K Ordered, V any]() *TreeMap[K, V] {
	return &TreeMap[K, V]{}
}

func (t *TreeMap[K, V]) Get(key K) *V {
	tmp := t.root
	for tmp != nil {
		if tmp.key == key {
			return &tmp.value
		} else if tmp.key < key {
			tmp = tmp.right
		} else {
			tmp = tmp.left
		}
	}

	return nil
}

func (t *TreeMap[K, V]) Set(key K, value V) {
	if t.root == nil {
		t.root = new(node[K, V])
		t.root.key, t.root.value = key, value
		return
	}

	isRight := false
	tmp := t.root
	for {
		if key == tmp.key {
			tmp.value = value
			return
		}

		if tmp.key < key {
			if tmp.right == nil {
				isRight = true
				break
			}

			tmp = tmp.right
		} else {
			if tmp.left == nil {
				break
			}

			tmp = tmp.left
		}
	}

	if isRight {
		tmp.right = new(node[K, V])
		tmp.right.key, tmp.right.value = key, value
		tmp.right.parent = tmp
	} else {
		tmp.left = new(node[K, V])
		tmp.left.key, tmp.left.value = key, value
		tmp.left.parent = tmp
	}

	p := tmp
	for p != nil {
		lh := p.left.height()
		rh := p.right.height()

		if lh-rh > 1 {
			p = rightRotate(p)
		} else if rh-lh > 1 {
			p = leftRotate(p)
		}

		if p.parent == nil {
			t.root = p
		}

		p = p.parent
	}
}

func (t *TreeMap[K, V]) Range(f func(key K, value V)) {
	var dfs func(*node[K, V])
	dfs = func(root *node[K, V]) {
		if root == nil {
			return
		}

		if root.left != nil {
			dfs(root.left)
		}

		f(root.key, root.value)

		if root.right != nil {
			dfs(root.right)
		}
	}

	dfs(t.root)
}
