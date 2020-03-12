package rubic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	face := Face{}
	face.Init(3, "U", true)
	expected := []string{"U0", "U1", "U2"}
	if row := face.GetRow(0); !reflect.DeepEqual(row, expected) {
		t.Errorf("Test Failed to get row index 0 expecting %v but got %v", expected, row)
	}
	expected = []string{"U3", "U4", "U5"}
	if row := face.GetRow(1); !reflect.DeepEqual(row, expected) {
		t.Errorf("Test Failed to get row index 1 expecting %v but got %v", expected, row)
	}
	expected = []string{"U6", "U7", "U8"}
	if row := face.GetRow(2); !reflect.DeepEqual(row, expected) {
		t.Errorf("Test Failed to get row index 2 expecting %v but got %v", expected, row)
	}
}
func TestSetRow(t *testing.T) {
	face := Face{}
	face.Init(3, "U", true)
	newRow := []string{"U0", "U1", "U2"}
	face.SetRow(0, newRow)
	if row := face.GetRow(0); !reflect.DeepEqual(row, newRow) {
		t.Errorf("Test Failed to set row index 0 expecting %v but got %v", newRow, row)
	}
	newRow = []string{"U3", "U4", "U5"}
	face.SetRow(1, newRow)
	if row := face.GetRow(1); !reflect.DeepEqual(row, newRow) {
		t.Errorf("Test Failed to set row index 1 expecting %v but got %v", newRow, row)
	}
	newRow = []string{"U6", "U7", "U8"}
	face.SetRow(2, newRow)
	if row := face.GetRow(2); !reflect.DeepEqual(row, newRow) {
		t.Errorf("Test Failed to set row index 2 expecting %v but got %v", newRow, row)
	}
}

func TestGetCol(t *testing.T) {
	face := Face{}
	face.Init(3, "U", true)
	expected := []string{"U0", "U3", "U6"}
	if col := face.GetCol(0); !reflect.DeepEqual(col, expected) {
		t.Errorf("Test Failed to get col index 0 expecting %v but got %v", expected, col)
	}
	expected = []string{"U1", "U4", "U7"}
	if col := face.GetCol(1); !reflect.DeepEqual(col, expected) {
		t.Errorf("Test Failed to get col index 1 expecting %v but got %v", expected, col)
	}
	expected = []string{"U2", "U5", "U8"}
	if col := face.GetCol(2); !reflect.DeepEqual(col, expected) {
		t.Errorf("Test Failed to get col index 2 expecting %v but got %v", expected, col)
	}
}

func TestSetCol(t *testing.T) {
	face := Face{}
	face.Init(3, "U", true)
	newCol := []string{"U0", "U3", "U6"}
	face.SetCol(0, newCol)
	if col := face.GetCol(0); !reflect.DeepEqual(col, newCol) {
		t.Errorf("Test Failed to set col index 0 expecting %v but got %v", newCol, col)
	}
	newCol = []string{"U1", "U4", "U7"}
	face.SetCol(1, newCol)
	if col := face.GetCol(1); !reflect.DeepEqual(col, newCol) {
		t.Errorf("Test Failed to set col index 1 expecting %v but got %v", newCol, col)
	}
	newCol = []string{"U2", "U5", "U8"}
	face.SetCol(2, newCol)
	if col := face.GetCol(2); !reflect.DeepEqual(col, newCol) {
		t.Errorf("Test Failed to set col index 2 expecting %v but got %v", newCol, col)
	}
}

func TestTurn(t *testing.T) {
	face := Face{}
	face.Init(3, "U", true)
	face2 := Face{}
	face2.Init(3, "U", true)
	face.TurnClockwise()
	face.TurnCounterClockwise()
	if !reflect.DeepEqual(face, face2) {
		t.Errorf("Test Failed TurnClockwise and TurnCounterClockwise should cancel eachother out but instead got :\n")
		face.Print()
	}

	face.TurnClockwise()
	expected := map[uint]string{0: "U6", 1: "U3", 2: "U0", 3: "U7", 4: "U4", 5: "U1", 6: "U8", 7: "U5", 8: "U2"}
	for index, cell := range face.Cells {
		if expected[index] != cell {
			t.Errorf("Test Failed expecting \n%s\n but instead got \n%s\n", fmt.Sprint(expected), fmt.Sprint(face.Cells))
		}
	}
}
