package main

import (
	"bufio"
	"fmt"
	"os"
)

// Console

func ConsoleOut(Out_ ...any) {
	for I := range Out_ {
		fmt.Print(Out_[I])
	}
}
func ConsoleIn() string {
	In := bufio.NewReader(os.Stdin)
	Line, _ := In.ReadString('\n')
	return Line
}

//

func GetLines(Program_ string) []string {
	var Lenght int = 0
	var Buffer int = 0

	for I := range Program_ {
		if Program_[I] == ';' {
			Lenght++
		}
	}

	Lines := make([]string, Lenght)

	for I := 0; I < Lenght; I++ {
		var BufferLenght int = 0

		for J := Buffer; Program_[J] != ';'; J++ {
			BufferLenght++
		}

		if I != 0 {
			BufferLenght--
		}

		BufferString := make([]byte, BufferLenght)
		BufferExtra := 0

		for J := Buffer; Program_[Buffer+BufferExtra] != ';'; {
			if Program_[Buffer+BufferExtra] == '\n' {
				BufferExtra++
				continue
			}
			if Program_[Buffer+BufferExtra] == '\t' {
				BufferExtra++
				continue
			}
			if Program_[Buffer+BufferExtra] == ' ' {
				BufferExtra++
				continue
			}
			BufferString[J-Buffer] = Program_[Buffer+BufferExtra]
			BufferExtra++
			J++
		}

		Buffer += BufferExtra

		Lines[I] = string(BufferString)

		Buffer++
	}

	return Lines
}

//

func GetFunction(Line_ string) string {
	var Lenght int = 0

	for I := 0; Line_[I] != '('; I++ {
		Lenght++
	}

	Return := make([]byte, Lenght)

	for I := 0; Line_[I] != '('; I++ {
		Return[I] = Line_[I]
	}

	return string(Return)
}

func RunFunctionOut(Line_ string) {
	if Line_[4] != '"' {
		for I := 4; Line_[I] != ')'; I++ {
			if I >= len(Line_)-1 {
				ConsoleOut("Ошибка OUT1: закрывающей скобки нет в функции out.\n")

				return
			}
			if []byte(Line_)[I] < '0' || []byte(Line_)[I] > '9' {
				ConsoleOut("Ошибка INT1: Неизвестный символ в переменной int.\n")

				return
			}
		}
		for I := 4; Line_[I] != ')'; I++ {
			fmt.Print(string(Line_[I]))
		}
	} else {
		for I := 5; Line_[I] != '"'; I++ {
			fmt.Print(string(Line_[I]))
		}
	}

	return
}

func main() {
	var Program string = "out(123H); out(124);"
	Code := GetLines(Program)

	for I := range Code {
		if GetFunction(Code[I]) == "out" {
			RunFunctionOut(Code[I])
		}
	}
}

/*

if GetFunction(Program) == "out" {
	RunFunctionOut(Program)
}

*/
