package other

type Status int

type Token struct {
	tok string
}

const (
	ACCEPT Status = 1
	ERROR  Status = 0
)

var END = Token{tok: "END"}

func gettoken() Token {
	return Token{}
}

func shift(t Token) bool {
	return false
}

func reduce(t Token) bool {
	return false
}

func forParse() Status {
	var tok Token
	for {
		tok = gettoken()
		if tok == END {
			return ACCEPT
		}
		for !shift(tok) {
			if !reduce(tok) {
				return ERROR
			}
		}
	}
}

func gotoParse() Status {

	var tok Token
reading:
	tok = gettoken()
	if tok == END {
		return ACCEPT
	}
shifting:
	if shift(tok) {
		goto reading
	}
	//reducing:
	if reduce(tok) {
		goto shifting
	}
	return ERROR
}
