package avl

type node struct {
	key   string
	value string

	height int

	left  *node
	right *node
}

type Tree struct {
	root *node
}

func (a *Tree) Add(key string, value string) {
	a.root = addNode(a.root, key, value)
}

func (a *Tree) Get(key string) string {
	n := getNode(a.root, key)
	if n == nil {
		return ""
	}

	return n.value
}

func (a *Tree) Remove(key string) {
	n := getNode(a.root, key)
	if n == nil {
		return
	}
	a.root = removeNode(a.root, key)
}

func addNode(root *node, key, value string) *node {
	if root == nil {
		return &node{
			key:    key,
			value:  value,
			height: 0,
		}
	}

	if root.key == key {
		root.value = value
		return root
	}

	if key < root.key {
		root.left = addNode(root.left, key, value)
	} else {
		root.right = addNode(root.right, key, value)
	}

	root.height = max(getHeight(root.left), getHeight(root.right)) + 1

	r := root

	if getFactor(r) > 1 {
		if getFactor(r.left) > 0 {
			// LL
			r = rightRotate(r)
		} else if getFactor(r.left) < 0 {
			// LR
			r.left = leftRotate(r.left)
			r = rightRotate(r)
		}
	} else if getFactor(r) < -1 {
		if getFactor(r.right) < 0 {
			// RR
			r = leftRotate(r)
		} else if getFactor(r.right) > 0 {
			// RL
			r.right = rightRotate(r.right)
			r = leftRotate(r)
		}
	}

	return r
}

func removeNode(root *node, key string) *node {
	if root == nil {
		return nil
	}

	if key < root.key {
		root.left = removeNode(root.left, key)
		return root
	}

	if key > root.key {
		root.right = removeNode(root.right, key)
		return root
	}

	if root.left == nil {
		r := root.right
		root.right = nil
		return r
	}

	if root.right == nil {
		r := root.left
		root.left = nil
		return r
	}

	r := getSuccessor(root.right)
	r.right = removeNode(root.right, r.key)
	r.left = root.left

	root.left, root.right = nil, nil

	if r == nil {
		return nil
	}

	r.height = max(getHeight(r.left), getHeight(r.right)) + 1

	if getFactor(r) > 1 {
		if getFactor(r.left) > 0 {
			// LL
			r = rightRotate(r)
		} else if getFactor(r.left) < 0 {
			// LR
			r.left = leftRotate(r.left)
			r = rightRotate(r)
		}
	} else if getFactor(r) < -1 {
		if getFactor(r.right) < 0 {
			// RR
			r = leftRotate(r)
		} else if getFactor(r.right) > 0 {
			// RL
			r.right = rightRotate(r.right)
			r = leftRotate(r)
		}
	}

	return r
}

func getNode(root *node, key string) *node {
	if root == nil {
		return nil
	}

	if key == root.key {
		return root
	} else if key < root.key {
		return getNode(root.left, key)
	} else {
		return getNode(root.right, key)
	}
}

func getFactor(root *node) int {
	return getHeight(root.left) - getHeight(root.right)
}

func getHeight(root *node) int {
	if root == nil {
		return 0
	}
	return root.height
}

func getSuccessor(root *node) *node {
	if root == nil {
		return nil
	}

	if root.left == nil {
		return root
	}

	return getSuccessor(root.left)
}

func rightRotate(root *node) *node {
	r := root.left
	root.left = r.right
	r.right = root

	r.height = max(getHeight(r.left), getHeight(r.right)) + 1
	root.height = max(getHeight(root.left), getHeight(root.right)) + 1

	return r
}

func leftRotate(root *node) *node {
	r := root.right
	root.right = r.left
	r.left = root

	r.height = max(getHeight(r.left), getHeight(r.right)) + 1
	root.height = max(getHeight(root.left), getHeight(root.right)) + 1

	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
