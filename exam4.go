package exam

import (
	"errors"
	"fmt"
)

func main() {

	filename := ""
	fmt.Scan(&filename)

	if checkfilee(filename + ".txt") {

		var n = 0

		for i:= 0;i<n;i++{
			a:=fmt.Sprint(i)
        var a = filename
		a=a+strconv.Itoa(rand.Intn(100))+".txt"
	
	}

	}

}

func checkfilee(filename string) bool {

	_, err := os.Stat(filename)

	return !errors.Is(err, os.ErrNoExist)

}
