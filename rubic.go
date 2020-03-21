package rubic

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	Up    uint   = 0
	Down         = 1
	Right        = 2
	Left         = 3
	Front        = 4
	Back         = 5
	U     string = "U"
	D            = "D"
	R            = "R"
	L            = "L"
	F            = "F"
	B            = "B"
	M            = "M"
	E            = "E"
	S            = "S"
)

type Rubic struct {
	Faces []Face
	Size  uint
}

func (rubic *Rubic) Init(size uint, numbered bool) {
	rubic.Faces = make([]Face, 6, 6)
	rubic.Size = size
	rubic.Faces[Up].Init(size, U, numbered)
	rubic.Faces[Down].Init(size, D, numbered)
	rubic.Faces[Right].Init(size, R, numbered)
	rubic.Faces[Left].Init(size, L, numbered)
	rubic.Faces[Front].Init(size, F, numbered)
	rubic.Faces[Back].Init(size, B, numbered)

}
func (rubic *Rubic) Turn(notation string) {
	switch notation {
	case "F":
		rubic.TurnF()
	case "F'":

	}
}
func (rubic *Rubic) TurnF() {
	affectedData := make(map[uint][]string)
	rubic.Faces[Front].TurnClockwise()
	affectedData[Up] = rubic.Faces[Up].GetRow(rubic.Size - 1)
	affectedData[Right] = ReverseSlice(rubic.Faces[Right].GetCol(0))
	affectedData[Down] = rubic.Faces[Down].GetRow(0)
	affectedData[Left] = ReverseSlice(rubic.Faces[Left].GetCol(rubic.Size - 1))

	rubic.Faces[Right].SetCol(0, affectedData[Up])
	rubic.Faces[Down].SetRow(0, affectedData[Right])
	rubic.Faces[Left].SetCol(rubic.Size-1, affectedData[Down])
	rubic.Faces[Up].SetRow(rubic.Size-1, affectedData[Left])
}

func (rubic *Rubic) TurnFReverse() {
	affectedData := make(map[uint][]string)
	rubic.Faces[Front].TurnCounterClockwise()
	affectedData[Up] = ReverseSlice(rubic.Faces[Up].GetRow(rubic.Size - 1))
	affectedData[Right] = rubic.Faces[Right].GetCol(0)
	affectedData[Down] = ReverseSlice(rubic.Faces[Down].GetRow(0))
	affectedData[Left] = rubic.Faces[Left].GetCol(rubic.Size - 1)

	rubic.Faces[Up].SetRow(rubic.Size-1, affectedData[Right])
	rubic.Faces[Right].SetCol(0, affectedData[Down])
	rubic.Faces[Down].SetRow(0, affectedData[Left])
	rubic.Faces[Left].SetCol(rubic.Size-1, affectedData[Up])
}

func (rubic *Rubic) TurnR() {
	affectedData := make(map[uint][]string)
	rubic.Faces[Right].TurnClockwise()
	affectedData[Up] = ReverseSlice(rubic.Faces[Up].GetCol(rubic.Size - 1))
	affectedData[Back] = ReverseSlice(rubic.Faces[Back].GetCol(0))
	affectedData[Down] = rubic.Faces[Down].GetCol(rubic.Size - 1)
	affectedData[Front] = rubic.Faces[Front].GetCol(rubic.Size - 1)

	rubic.Faces[Back].SetCol(0, affectedData[Up])
	rubic.Faces[Down].SetCol(rubic.Size-1, affectedData[Back])
	rubic.Faces[Front].SetCol(rubic.Size-1, affectedData[Down])
	rubic.Faces[Up].SetCol(rubic.Size-1, affectedData[Front])
}
func (rubic *Rubic) TurnRReverse() {
	affectedData := make(map[uint][]string)
	rubic.Faces[Right].TurnCounterClockwise()
	affectedData[Up] = rubic.Faces[Up].GetCol(rubic.Size - 1)
	affectedData[Back] = ReverseSlice(rubic.Faces[Back].GetCol(0))
	affectedData[Down] = ReverseSlice(rubic.Faces[Down].GetCol(rubic.Size - 1))
	affectedData[Front] = rubic.Faces[Front].GetCol(rubic.Size - 1)

	rubic.Faces[Front].SetCol(rubic.Size-1, affectedData[Up])
	rubic.Faces[Up].SetCol(rubic.Size-1, affectedData[Back])
	rubic.Faces[Back].SetCol(0, affectedData[Down])
	rubic.Faces[Down].SetCol(rubic.Size-1, affectedData[Front])
}

func (rubic *Rubic) TurnU() {
	affectedData := make(map[uint][]string)
	rubic.Faces[Up].TurnClockwise()
	affectedData[Left] = rubic.Faces[Left].GetRow(0)
	affectedData[Back] = rubic.Faces[Back].GetRow(0)
	affectedData[Right] = rubic.Faces[Right].GetRow(0)
	affectedData[Front] = rubic.Faces[Front].GetRow(0)

	rubic.Faces[Back].SetRow(0, affectedData[Left])
	rubic.Faces[Right].SetRow(0, affectedData[Back])
	rubic.Faces[Front].SetRow(0, affectedData[Right])
	rubic.Faces[Left].SetRow(0, affectedData[Front])
}
func (rubic *Rubic) TurnUReverse() {
	affectedData := make(map[uint][]string)
	rubic.Faces[Up].TurnCounterClockwise()
	affectedData[Left] = rubic.Faces[Left].GetRow(0)
	affectedData[Back] = rubic.Faces[Back].GetRow(0)
	affectedData[Right] = rubic.Faces[Right].GetRow(0)
	affectedData[Front] = rubic.Faces[Front].GetRow(0)

	rubic.Faces[Back].SetRow(0, affectedData[Right])
	rubic.Faces[Right].SetRow(0, affectedData[Front])
	rubic.Faces[Front].SetRow(0, affectedData[Left])
	rubic.Faces[Left].SetRow(0, affectedData[Back])
}

func (rubic *Rubic) TurnL() {
	affectedData := make(map[uint][]string)
	rubic.Faces[Left].TurnClockwise()
	affectedData[Up] = rubic.Faces[Up].GetCol(0)
	affectedData[Front] = rubic.Faces[Front].GetCol(0)
	affectedData[Down] = ReverseSlice(rubic.Faces[Down].GetCol(0))
	affectedData[Back] = ReverseSlice(rubic.Faces[Back].GetCol(rubic.Size - 1))

	rubic.Faces[Up].SetCol(0, affectedData[Back])
	rubic.Faces[Front].SetCol(0, affectedData[Up])
	rubic.Faces[Down].SetCol(0, affectedData[Front])
	rubic.Faces[Back].SetCol(rubic.Size-1, affectedData[Down])
}
func (rubic *Rubic) TurnLReverse() {
	affectedData := make(map[uint][]string)
	rubic.Faces[Left].TurnCounterClockwise()
	affectedData[Up] = ReverseSlice(rubic.Faces[Up].GetCol(0))
	affectedData[Front] = rubic.Faces[Front].GetCol(0)
	affectedData[Down] = rubic.Faces[Down].GetCol(0)
	affectedData[Back] = ReverseSlice(rubic.Faces[Back].GetCol(rubic.Size - 1))

	rubic.Faces[Up].SetCol(0, affectedData[Front])
	rubic.Faces[Front].SetCol(0, affectedData[Down])
	rubic.Faces[Down].SetCol(0, affectedData[Back])
	rubic.Faces[Back].SetCol(rubic.Size-1, affectedData[Up])
}

func (rubic *Rubic) TurnB() {
	affectedData := make(map[uint][]string)
	rubic.Faces[Back].TurnClockwise()
	affectedData[Up] = ReverseSlice(rubic.Faces[Up].GetRow(0))
	affectedData[Left] = rubic.Faces[Left].GetCol(0)
	affectedData[Down] = ReverseSlice(rubic.Faces[Down].GetRow(rubic.Size - 1))
	affectedData[Right] = rubic.Faces[Right].GetCol(rubic.Size - 1)

	rubic.Faces[Up].SetRow(0, affectedData[Right])
	rubic.Faces[Left].SetCol(0, affectedData[Up])
	rubic.Faces[Down].SetRow(rubic.Size-1, affectedData[Left])
	rubic.Faces[Right].SetCol(rubic.Size-1, affectedData[Down])
}
func (rubic *Rubic) TurnBReverse() {
	affectedData := make(map[uint][]string)
	rubic.Faces[Back].TurnCounterClockwise()
	affectedData[Up] = rubic.Faces[Up].GetRow(0)
	affectedData[Left] = ReverseSlice(rubic.Faces[Left].GetCol(0))
	affectedData[Down] = rubic.Faces[Down].GetRow(rubic.Size - 1)
	affectedData[Right] = ReverseSlice(rubic.Faces[Right].GetCol(rubic.Size - 1))

	rubic.Faces[Up].SetRow(0, affectedData[Left])
	rubic.Faces[Left].SetCol(0, affectedData[Down])
	rubic.Faces[Down].SetRow(rubic.Size-1, affectedData[Right])
	rubic.Faces[Right].SetCol(rubic.Size-1, affectedData[Up])
}

func (rubic *Rubic) TurnD() {
	affectedData := make(map[uint][]string)
	rubic.Faces[Down].TurnClockwise()
	affectedData[Front] = rubic.Faces[Front].GetRow(rubic.Size - 1)
	affectedData[Left] = rubic.Faces[Left].GetRow(rubic.Size - 1)
	affectedData[Back] = rubic.Faces[Back].GetRow(rubic.Size - 1)
	affectedData[Right] = rubic.Faces[Right].GetRow(rubic.Size - 1)

	rubic.Faces[Front].SetRow(rubic.Size-1, affectedData[Left])
	rubic.Faces[Left].SetRow(rubic.Size-1, affectedData[Back])
	rubic.Faces[Back].SetRow(rubic.Size-1, affectedData[Right])
	rubic.Faces[Right].SetRow(rubic.Size-1, affectedData[Front])
}
func (rubic *Rubic) TurnDReverse() {
	affectedData := make(map[uint][]string)
	rubic.Faces[Down].TurnCounterClockwise()
	affectedData[Front] = rubic.Faces[Front].GetRow(rubic.Size - 1)
	affectedData[Left] = rubic.Faces[Left].GetRow(rubic.Size - 1)
	affectedData[Back] = rubic.Faces[Back].GetRow(rubic.Size - 1)
	affectedData[Right] = rubic.Faces[Right].GetRow(rubic.Size - 1)

	rubic.Faces[Front].SetRow(rubic.Size-1, affectedData[Right])
	rubic.Faces[Left].SetRow(rubic.Size-1, affectedData[Front])
	rubic.Faces[Back].SetRow(rubic.Size-1, affectedData[Left])
	rubic.Faces[Right].SetRow(rubic.Size-1, affectedData[Back])
}

func (rubic *Rubic) TurnM() {
	affectedData := make(map[uint][]string)
	affectedData[Front] = rubic.Faces[Front].GetCol(1)
	affectedData[Down] = ReverseSlice(rubic.Faces[Down].GetCol(1))
	affectedData[Back] = ReverseSlice(rubic.Faces[Back].GetCol(1))
	affectedData[Up] = rubic.Faces[Up].GetCol(1)

	rubic.Faces[Front].SetCol(1, affectedData[Up])
	rubic.Faces[Down].SetCol(1, affectedData[Front])
	rubic.Faces[Back].SetCol(1, affectedData[Down])
	rubic.Faces[Up].SetCol(1, affectedData[Back])
}
func (rubic *Rubic) TurnMReverse() {
	affectedData := make(map[uint][]string)
	affectedData[Front] = rubic.Faces[Front].GetCol(1)
	affectedData[Down] = rubic.Faces[Down].GetCol(1)
	affectedData[Back] = ReverseSlice(rubic.Faces[Back].GetCol(1))
	affectedData[Up] = ReverseSlice(rubic.Faces[Up].GetCol(1))

	rubic.Faces[Front].SetCol(1, affectedData[Down])
	rubic.Faces[Down].SetCol(1, affectedData[Back])
	rubic.Faces[Back].SetCol(1, affectedData[Up])
	rubic.Faces[Up].SetCol(1, affectedData[Front])
}
func (rubic *Rubic) TurnE() {
	affectedData := make(map[uint][]string)
	affectedData[Front] = rubic.Faces[Front].GetRow(1)
	affectedData[Right] = rubic.Faces[Right].GetRow(1)
	affectedData[Back] = rubic.Faces[Back].GetRow(1)
	affectedData[Left] = rubic.Faces[Left].GetRow(1)

	rubic.Faces[Front].SetRow(1, affectedData[Left])
	rubic.Faces[Right].SetRow(1, affectedData[Front])
	rubic.Faces[Back].SetRow(1, affectedData[Right])
	rubic.Faces[Left].SetRow(1, affectedData[Back])
}
func (rubic *Rubic) TurnEReverse() {
	affectedData := make(map[uint][]string)
	affectedData[Front] = rubic.Faces[Front].GetRow(1)
	affectedData[Right] = rubic.Faces[Right].GetRow(1)
	affectedData[Back] = rubic.Faces[Back].GetRow(1)
	affectedData[Left] = rubic.Faces[Left].GetRow(1)

	rubic.Faces[Front].SetRow(1, affectedData[Right])
	rubic.Faces[Right].SetRow(1, affectedData[Back])
	rubic.Faces[Back].SetRow(1, affectedData[Left])
	rubic.Faces[Left].SetRow(1, affectedData[Front])
}

func (rubic *Rubic) TurnS() {
	affectedData := make(map[uint][]string)
	affectedData[Up] = rubic.Faces[Up].GetRow(1)
	affectedData[Right] = ReverseSlice(rubic.Faces[Right].GetCol(1))
	affectedData[Down] = rubic.Faces[Down].GetRow(1)
	affectedData[Left] = ReverseSlice(rubic.Faces[Left].GetCol(1))

	rubic.Faces[Up].SetRow(1, affectedData[Left])
	rubic.Faces[Right].SetCol(1, affectedData[Up])
	rubic.Faces[Down].SetRow(1, affectedData[Right])
	rubic.Faces[Left].SetCol(1, affectedData[Down])
}
func (rubic *Rubic) TurnSReverse() {
	affectedData := make(map[uint][]string)
	affectedData[Up] = ReverseSlice(rubic.Faces[Up].GetRow(1))
	affectedData[Right] = rubic.Faces[Right].GetCol(1)
	affectedData[Down] = ReverseSlice(rubic.Faces[Down].GetRow(1))
	affectedData[Left] = rubic.Faces[Left].GetCol(1)

	rubic.Faces[Up].SetRow(1, affectedData[Right])
	rubic.Faces[Right].SetCol(1, affectedData[Down])
	rubic.Faces[Down].SetRow(1, affectedData[Left])
	rubic.Faces[Left].SetCol(1, affectedData[Up])
}

func (rubic *Rubic) Print() {
	for index, face := range rubic.Faces {
		fmt.Println()
		faceDir := ""
		switch uint(index) {
		case Up:
			faceDir = "Up"
		case Down:
			faceDir = "Down"
		case Right:
			faceDir = "Right"
		case Left:
			faceDir = "Left"
		case Front:
			faceDir = "Front"
		case Back:
			faceDir = "Back"
		}
		fmt.Printf("%s Face: \n", faceDir)
		face.Print()
		fmt.Println()
	}
}
func (rubic *Rubic) DoTurns(notations []string) {
	for _, notation := range notations {
		subNotation := strings.Split(notation, "")
		funcName := "Turn"
		turns := 1
		switch len(subNotation) {
		case 1:
			funcName += notation
		case 2:
			if subNotation[0] == "2" {
				turns++
				funcName += subNotation[1]
			} else {
				funcName += subNotation[0] + "Reverse"
			}
		case 3:
			turns++
			funcName += subNotation[1]
		}
		for i := 0; i < turns; i++ {
			reflect.ValueOf(rubic).MethodByName(funcName).Call([]reflect.Value{})
		}
	}
	rubic.Print()
}
