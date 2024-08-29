package santa

//looks for the length of the ribbon length
func Ribbon(l, w, h int) int {
	result := 0
	
	p1 := 2 * (l + w)
	p2 := 2 * (w + h)
	p3 := 2 * (l + h)
	minPeri := 0
	if p1 <= p2 && p1<=p3 {
		minPeri = p1
	} else if p2 <= p1 && p2 <= p3 {
		minPeri = p2
	} else if p3 <= p1 && p3<=p2 {
		minPeri = p3
	} 
// fmt.Println(minPeri)
	result = minPeri + (l * w * h)
	return result
}