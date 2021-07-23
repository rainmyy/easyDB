package common

/**
* this file const are golbal const for all data,import the package like "import . common"
 */
const (
	//data file type,0,file default type
	DataType = iota
	//ini file type, 1
	IniType
	//yaml file type, 2
	YamlType
	//json file type,3
	JsonType
)
const (
	//data file suffix
	DataSuffix = ".data"
	//ini file suffix
	IniSuffix = ".conf"
	//json file suffix
	JsonSuffix = ".json"
	//yaml file suffix
	YamlSuffix = ".yaml"
)
const (
	//left bracket, this data for turn tree data to string data
	LeftBracket rune = '['
	//right bracket,this data for turn tree data to string data
	RightBracket rune = ']'
	//right rance,this data for turn tree data to string data
	LeftRance  rune = '{'
	RightRance rune = '}'
	Colon      rune = ':'
	Comma      rune = ','
	None       rune = 'N'
	Slash      rune = '/'
	Hash       rune = '#'
	Asterisk   rune = '*'
	LineBreak  rune = '\n'
	Blank      rune = ' '
	Period     rune = '.'
)

const (
	BindTag = "bind"
	JsonTag = "json"
)
