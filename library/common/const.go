package common

const (
	IniType = iota
	YamlType
	JsonType
	DataType
)

const (
	LeftBracket  rune = '['
	RightBracket rune = ']'
	LeftRrance   rune = '{'
	RightRrance  rune = '}'
	Colon        rune = ':'
	Comma        rune = ','
	None         rune = 'N'
	Slash        rune = '/'
	Hash         rune = '#'
	Asterisk     rune = '*'
	LineBreak    rune = '\n'
	Blank        rune = ' '
)

const (
	BindTag = "bind"
	Json    = "json"
)
