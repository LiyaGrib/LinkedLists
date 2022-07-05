package singlylinkedlists

type (
	Node struct {
		Body any
		next *Node
	}

	List struct {
		len  uint
		head *Node
	}
)
