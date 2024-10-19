package tools

import (
	config "ThreeCatsGo/config"
	"fmt"
	"os"
)

func CreateFolder() {
	if !FileExist(config.AVATAR_IMG_FOLDER) {
		err := os.MkdirAll(config.AVATAR_IMG_FOLDER, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			panic(ColoredStr("create avatar folder failed").Red())
		}
	}
	if !FileExist(config.EVENT_FILE_FOLDER) {
		err := os.MkdirAll(config.EVENT_FILE_FOLDER, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			panic(ColoredStr("create avatar folder failed").Red())
		}
	}
	if !FileExist(config.EVENT_IMG_FOLDER) {
		err := os.MkdirAll(config.EVENT_IMG_FOLDER, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			panic(ColoredStr("create avatar folder failed").Red())
		}
	}
}
