package doublylinkedlists

type (
	Node struct {
		Body any
		next *Node
		prev *Node
	}

	List struct {
		len  uint
		head *Node
		last *Node
	}
)
