package common

const (
	DataType = iota
	IniType
	YamlType
	JsonType
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
	Period       rune = '.'
)

const (
	BindTag = "bind"
	JsonTag = "json"
)
