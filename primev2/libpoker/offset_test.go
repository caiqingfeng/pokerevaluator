package libpoker

import "testing"

func TestCountOfffset3(t *testing.T) {
	offset := countOffset3(13, 12, 11, 10, 9)
	if offset != 0 {
		t.Error("wrong offset", offset)
	}

	offset = countOffset3(1, 12, 11, 10, 9)
	if offset != 12*66+11 {
		t.Error("wrong offset", offset, 12*66+11)
	}

	offset = countOffset3(1, 2, 3, 4, 5)
	//22X6543, offset=12*66+(11+4)*8/2
	if offset != 12*66+(11+4)*8/2 {
		t.Error("wrong offset", offset, 12*66+(11+4)*8/2)
	}

	offset = countOffset3(11, 12, 9, 2, 5)
	//QQQKTX63, offset=2*66+11+1
	if offset != 2*66+11+1 {
		t.Error("wrong offset", offset, 2*66+11+1)
	}
}

func TestCountOffset4(t *testing.T) {
	offset := countOffset4(13, 12, 11, 2, 5)
	//AAKKQ , offset=0
	if offset != 0 {
		t.Error("wrong offset", offset)
	}

	offset = countOffset4(2, 1, 13, 6, 5)
	//3322A , offset=(12+...+2)*11
	if offset != 11*(12+2)*11/2 {
		t.Error("wrong offset", offset)
	}
}

func TestCountOffset5(t *testing.T) {
	offset := countOffset5(13, 12, 11, 10, 5, 6)
	//AAKQJ , offset=0
	if offset != 0 {
		t.Error("wrong offset", offset)
	}

	offset = countOffset5(1, 2, 3, 4, 5, 6)
	//22765 , offset=2859-3(22765,22764,22763)-2(22754,22753)-1(22743)
	// -2(2265x)-1(2264x)
	//最小的22543其offset是2859
	if offset != 2859-3-2-1-2-1 {
		t.Error("wrong offset", offset, 2859-3-2-1-2-1-1)
	}

	offset = countOffset5(1, 12, 8, 7, 3, 4)
	//22K98+45 offset=220*12+11*10/2(22Axx)+(22KQx)+(22KJx)+(22KTx)
	//
	if offset != 220*12+11*10/2+9+8+7 {
		t.Error("wrong offset", offset, 220*12+11*10/2+9+8+7)
	}

	offset = countOffset5(12, 11, 8, 5, 2, 4)
	//KKQ96+53 offset=220*1+11*10/2(KKAxx)+(KKQJx)+(KKQTx)+(KKQ9[8,7])
	//
	if offset != 220*1+11*10/2+9+8+2 {
		t.Error("wrong offset", offset, 220*1+11*10/2+9+8+2)
	}
}

func TestCountOffset(t *testing.T) {
	r := countOffset(12, 11, 2, 3)
	if r != 0 {
		t.Error("wrong result")
	}

	r = countOffset(0, 1, 1, 1)
	if r != 155 {
		t.Error("wrong result", r)
	}
}

