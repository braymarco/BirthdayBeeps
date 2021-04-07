package main

import (
	"fmt"
	"syscall"
)

var (
	C  = 523
	_C = 554
	D  = 587
	_D = 622
	E  = 659
	F  = 698
	_F = 740
	G  = 784
	_G = 831
	A  = 880
	_A = 932
	B  = 988
)

func main() {
	fmt.Println(`                     /^\
           /         (/^\)     /
      \   ( \         \ /     ( \     /^\
     / )   \ |        _|_      \ |   |/^\|
    | /    _|_        | |      _|_    \ /
    _|_    | |        | |      | |    _|_
    | |    | |        | |      | |    | |
    | |    | |    ****| |******| |    | |
    | |****| |****    | |      | |****| |
   *| |    | |                 | |    | |*****
 *  | |   F  E  L  I  Z               | |      *
*                                               *
| *          C  U  M  P  L  E  A  Ñ  O  S  !  * |
|  *****                                 *****  |
|@      **********             **********      @|
| @   @           *************           @   @ |
|  @@@ @    @                       @    @ @@@  |
|       @@@@ @      @       @      @ @@@@       |
 *            @@@@@@ @     @ @@@@@@            *
  *                   @@@@@                   *
   *****                                 *****
        **********             **********
                  *************

------------------------------------------------`)
	play()
	play()
}

func play() {


	Beep(C, 282)
	Beep(C, 282)
	Beep(D, 400)
	Beep(C, 382)
	Beep(F, 500)
	Beep(E, 400)

	Beep(C, 282)
	Beep(C, 282)
	Beep(D, 400)
	Beep(C, 382)
	Beep(G, 500)
	Beep(F, 600)

	Beep(C, 282)
	Beep(C, 282)
	Beep(_C, 350)
	Beep(A, 302)
	Beep(F, 400)
	Beep(E, 400)
	Beep(D, 500)

	Beep(_A, 340)
	Beep(_A, 220)
	Beep(A, 220)
	Beep(F, 228)
	Beep(G, 286)
	Beep(F, 270)
}

func Beep(freq int, duration int) {
	kernel32, _ := syscall.LoadLibrary("kernel32.dll")
	beep32, _ := syscall.GetProcAddress(kernel32, "Beep")

	defer syscall.FreeLibrary(kernel32)

	 syscall.Syscall(beep32, uintptr(2), uintptr(freq), uintptr(duration), 0)


}
