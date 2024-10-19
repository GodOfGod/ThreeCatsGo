package tools

const reset = "\033[0m" 
const red = "\033[31m" 
const green = "\033[32m" 
const yellow = "\033[33m" 
const blue = "\033[34m" 
const magenta = "\033[35m" 
const cyan = "\033[36m" 
const gray = "\033[37m" 
const white = "\033[97m"

type colorFuncs struct {
	text string
}


func ColoredStr(str string) *colorFuncs {
	var retValues colorFuncs;

	retValues.text = str
	
	return &retValues
}

func (color *colorFuncs) Red() string {
	return red + color.text + reset;
}
func (color *colorFuncs) Green() string {
	return green + color.text + reset;
}
func (color *colorFuncs) Yellow() string {
	return yellow + color.text + reset;
}
func (color *colorFuncs) Blue() string {
	return blue + color.text + reset;
}
func (color *colorFuncs) Magenta() string {
	return magenta + color.text + reset;
}
func (color *colorFuncs) Cyan() string {
	return cyan + color.text + reset;
}
func (color *colorFuncs) White() string {
	return white + color.text + reset;
}
func (color *colorFuncs) Gray() string {
	return gray + color.text + reset;
}

