package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	baseQuad(x, y, '0', '-', '|')
}

func baseQuad(row, col int, cornerSym, rowSym, colSym rune) {
	if row < 0 || col < 0 {
		return
	}
	for c := 0; c < col; c++ {
		for r := 0; r < row; r++ {
			isLeft := r == 0
			isRight := r == row-1
			isTop := c == 0
			isBottom := c == col-1
			isCorner := (isTop && isLeft) || (isTop && isRight) || (isBottom && isLeft) || (isBottom && isRight)

			if isCorner {
				z01.PrintRune(cornerSym)
			} else if isTop || isBottom {
				z01.PrintRune(rowSym)
			} else if isLeft || isRight {
				z01.PrintRune(colSym)
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
