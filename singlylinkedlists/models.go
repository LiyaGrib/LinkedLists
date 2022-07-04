package singlylinkedlists

type (
	Element struct {
		Body any
		Next *Element
	}

	List struct {
		Len  uint
		Next *Element
	}
)
