package util

import "hash/crc32"

var table = crc32.MakeTable(crc32.Castagnoli)

type CRC uint32

func NewCRC(b []byte) CRC {
	return CRC(0).Update(b)
}

func (c CRC) Update(b []byte) CRC {
	return CRC(crc32.Update(uint32(c), table,b))
}

func (c CRC) Value() uint32 {
	// c = 3861378113
	// c >> 15 = 117839
	// c << 17 = 506118552027136
	// | = 506118552144975
	// 0xa282ead8 = 2726488792
	// 2332473127
	return uint32(c >> 15 | c << 17) + 0xa282ead8
}

