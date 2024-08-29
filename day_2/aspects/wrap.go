package santa

//looks for the wrapping paper length
func Wrap(l, w, h int) int {
	lw := l * w
	wh := w * h
	lh := l * h
	result := 0

	add := 0

	if lw <= wh && lw <= lh {
		add = lw
	} else if wh <= lw && wh <= lh {
		add = wh
	} else if lh <= lw && lh <= wh {
		add = lh
	}

	result = 2*lw + 2*wh + 2*lh + add
	return result
}