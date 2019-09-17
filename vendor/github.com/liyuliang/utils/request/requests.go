package request

import "sort"

type html struct {
	Index   int
	Content string
	Err     error
}

type htmls []html

func (hs htmls) Len() int { return len(hs) }
func (hs htmls) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}
func (hs htmls) Less(i, j int) bool {
	if hs[i].Index < hs[j].Index {
		return true
	}
	if hs[i].Index > hs[j].Index {
		return false
	}
	return hs[i].Index < hs[j].Index
}

func HttpGets(urls []string) (resp []html) {

	chs := make(chan html, len(urls))

	for i, url := range urls {

		go func(uri string, num int) {
			resp := DoReq(uri, "")
			chs <- html{
				Index:   num,
				Content: resp.Data,
				Err:     resp.Err,
			}
		}(url, i)
	}
	for i := 0; i < len(urls); i++ {
		h := <-chs
		resp = append(resp, h)
	}

	sort.Sort(htmls(resp))
	return resp
}
