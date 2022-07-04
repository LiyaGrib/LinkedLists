package methods

func ListInit() (l *List) {
	l = new(List)
	l.Len = 0
	l.Next = nil
	return l
}
