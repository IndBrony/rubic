package rubic

import (
	"reflect"
	"testing"
)

func basicTurnTesting(t *testing.T, funcName string) {
	rubic := Rubic{}
	rubic.Init(3, true)
	reflect.ValueOf(&rubic).MethodByName(funcName).Call([]reflect.Value{})
	reflect.ValueOf(&rubic).MethodByName(funcName + "Reverse").Call([]reflect.Value{})
	rubic2 := Rubic{}
	rubic2.Init(3, true)
	if !reflect.DeepEqual(rubic, rubic2) {
		t.Fail()
	}
}
func TestRubicTurnS(t *testing.T) {
	notations := []string{F, R, U, L, B, D, M, E, S}
	for _, notation := range notations {
		basicTurnTesting(t, "Turn"+notation)
	}
}

func TestDoTurns(t *testing.T) {
	rubic := Rubic{}
	rubic.Init(3, true)
	rubic.DoTurns([]string{"F"})
	rubic2 := Rubic{}
	rubic2.Init(3, true)
	rubic2.TurnF()
	if !reflect.DeepEqual(rubic, rubic2) {
		t.Fail()
	}

	rubic = Rubic{}
	rubic.Init(3, true)
	rubic.DoTurns([]string{"F'"})
	rubic2 = Rubic{}
	rubic2.Init(3, true)
	rubic2.TurnFReverse()
	if !reflect.DeepEqual(rubic, rubic2) {
		t.Fail()
	}

	rubic = Rubic{}
	rubic.Init(3, true)
	rubic.DoTurns([]string{"2F'"})
	rubic2 = Rubic{}
	rubic2.Init(3, true)
	rubic2.TurnFReverse()
	rubic2.TurnFReverse()
	if !reflect.DeepEqual(rubic, rubic2) {
		t.Fail()
	}

	rubic = Rubic{}
	rubic.Init(3, true)
	rubic.DoTurns([]string{"2F"})
	rubic2 = Rubic{}
	rubic2.Init(3, true)
	rubic2.TurnF()
	rubic2.TurnF()
	if !reflect.DeepEqual(rubic, rubic2) {
		t.Fail()
	}
}
