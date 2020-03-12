package rubic

import (
	"fmt"
)

type Face struct {
	Cells    map[uint]string
	Size     uint
	Numbered bool
}

func (face *Face) SetFace(cells map[uint]string) {
	for cell, color := range cells {
		if cell < face.Size*face.Size {
			face.Cells[cell] = color
		}
	}
	return
}
func (face *Face) Init(size uint, color string, numbered bool) {
	face.Cells = make(map[uint]string)
	face.Size = size
	face.Numbered = numbered
	var cell uint
	for cell = 0; cell < face.Size*face.Size; cell++ {
		if numbered {
			face.Cells[cell] = color + fmt.Sprint(cell)
		} else {
			face.Cells[cell] = color
		}
	}
	return
}
func (face *Face) GetRow(index uint) (row []string) {
	row = make([]string, face.Size, face.Size)
	var cell uint
	for cell = 0; cell < face.Size; cell++ {
		row[cell] = face.Cells[(index*face.Size)+cell]
	}
	return
}
func (face *Face) SetRow(index uint, row []string) {
	var cell uint
	for cell = 0; cell < face.Size; cell++ {
		face.Cells[(index*face.Size)+cell] = row[cell]
	}
	return
}
func (face *Face) GetCol(index uint) (col []string) {
	col = make([]string, face.Size, face.Size)
	var cell uint
	for cell = 0; cell < face.Size; cell++ {
		col[cell] = face.Cells[index+(face.Size*cell)]
	}
	return
}
func (face *Face) SetCol(index uint, col []string) {
	var cell uint
	for cell = 0; cell < face.Size; cell++ {
		face.Cells[index+(face.Size*cell)] = col[cell]
	}
	return
}
func (face *Face) TurnClockwise() {
	newFace := Face{}
	newFace.Init(face.Size, "", false)
	var index uint
	for index = 0; index < face.Size; index++ {
		newRow := ReverseSlice(face.GetCol(index))
		newFace.SetRow(index, newRow)
	}
	face.Cells = newFace.Cells
}
func (face *Face) TurnCounterClockwise() {
	newFace := Face{}
	newFace.Init(face.Size, "", false)
	var newFaceIndex uint = 0
	var index uint
	for index = face.Size - 1; newFaceIndex < face.Size; index-- {
		newRow := face.GetCol(index)
		newFace.SetRow(newFaceIndex, newRow)
		newFaceIndex++
	}
	face.Cells = newFace.Cells
}

func (face *Face) Print() {
	var cell uint
	for cell = 0; cell < face.Size*face.Size; cell++ {
		if cell < face.Size*face.Size {
			if cell%face.Size == 0 {
				fmt.Println()
			}
			fmt.Print(face.Cells[cell])
		}
	}
}
