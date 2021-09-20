package util

type Range struct {
	Start []byte
	Limit []byte
}


// 不是很理解这个处理是要干什
//逻辑是找到最后一个0xff以内的字符, 截位, 然后最后一位+1 TODO
func BytesPrefix(prefix []byte) *Range {
	var limit []byte
	for i := len(prefix) - 1; i >= 0; i-- {
		c := prefix[i]
		if c < 0xff {
			limit = make([]byte, i+1)
			copy(limit, prefix)
			limit[i] = c + 1
			break
		}
	}
	return &Range{prefix,limit}
}
