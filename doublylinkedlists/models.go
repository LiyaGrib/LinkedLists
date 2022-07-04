package doublylinkedlists

type (
	Element struct {
		Body any
		Next *Element
		Prev *Element
	}

	Head struct {
		Len  uint
		Next *Element
	}
)
