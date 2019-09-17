package request

type Response struct {
	StatusCode int
	Data       string
	Err        error
	Speed      float64
}

type Responses []Response

func (rs Responses) Len() int { return len(rs) }
func (rs Responses) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}
func (rs Responses) Less(i, j int) bool {
	if rs[i].Speed < rs[j].Speed {
		return true
	}
	if rs[i].Speed > rs[j].Speed {
		return false
	}
	return rs[i].Speed < rs[j].Speed
}
