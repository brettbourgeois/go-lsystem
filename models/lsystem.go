package models

import (
	"fmt"
	"strings"
)

// LSystem struct
type LSystem struct {
	variables     []string
	constants     []string
	axiom         string
	N             int
	c             int
	rules         map[string]string
	replacement   []string
	LSystemString string
	Angle         float64
}

// GenLSystemString replaces strings based on rules
func (S *LSystem) GenLSystemString() {
	if S.c <= S.N {
		if S.c == 0 {
			S.LSystemString = fmt.Sprintf("%s", S.axiom)
		} else {
			tmp := make([]string, 0)
			for r := range S.LSystemString {
				tmp = append(tmp, S.rules[string(S.LSystemString[r])])
			}
			S.LSystemString = strings.Join(tmp, "")
		}
	}
	if S.c != S.N {
		S.c++
		S.GenLSystemString()
	}
}

func (S *LSystem) applyConstantRule() {
	for _, v := range S.constants {
		S.rules[v] = v
	}
}

// BinaryTree will return an BinaryTree LSystem
func BinaryTree(n int) LSystem {
	lSystem := LSystem{
		constants: []string{"[", "]"},
		variables: []string{"1,0"},
		axiom:     "0",
		N:         n,
		c:         0,
		rules:     map[string]string{"1": "11", "0": "1[0]0"},
	}
	lSystem.applyConstantRule()
	lSystem.GenLSystemString()
	return lSystem
}

// SierpinskiTriangle will return an SierpinskiTriangle LSystem
func SierpinskiTriangle(n int) LSystem {
	lSystem := LSystem{
		variables: []string{"F", "G"},
		constants: []string{"+", "-"},
		axiom:     "F-G-G",
		rules: map[string]string{
			"F": "F-G+F+G-F",
			"G": "GG",
		},
		N:     n,
		c:     0,
		Angle: 120,
	}
	lSystem.applyConstantRule()
	lSystem.GenLSystemString()
	return lSystem
}

// SierpinskiArrowHead will return an SierpinskiTriangle LSystem
func SierpinskiArrowHead(n int) LSystem {
	lSystem := LSystem{
		variables: []string{"F", "G"},
		constants: []string{"+", "-"},
		axiom:     "F",
		rules: map[string]string{
			"F": "G-F-G",
			"G": "F+G+F",
		},
		N:     n,
		c:     0,
		Angle: 60,
	}
	lSystem.applyConstantRule()
	lSystem.GenLSystemString()
	return lSystem
}

// KochCurve will return an KochCurve LSystem
func KochCurve(n int) LSystem {
	lSystem := LSystem{
		variables: []string{"F"},
		constants: []string{"+", "-"},
		axiom:     "F",
		rules: map[string]string{
			"F": "F+F-F-F+F",
		},
		N:     n,
		c:     0,
		Angle: 90,
	}
	lSystem.applyConstantRule()
	lSystem.GenLSystemString()
	return lSystem
}

// DragonCurve
// variables : X Y
// constants : F + −
// start : FX
// rules : (X → X+YF+), (Y → −FX−Y)
// angle : 90°

// DragonCurve will return an DragonCurve LSystem
func DragonCurve(n int) LSystem {
	lSystem := LSystem{
		variables: []string{"X", "Y"},
		constants: []string{"F", "+", "-"},
		axiom:     "FX",
		rules: map[string]string{
			"X": "X+YF+",
			"Y": "-FX-Y",
		},
		N:     n,
		c:     0,
		Angle: 90,
	}
	lSystem.applyConstantRule()
	lSystem.GenLSystemString()
	return lSystem
}

// BarnsleyFern will return an BarnsleyFern LSystem
func BarnsleyFern(n int) LSystem {
	lSystem := LSystem{
		variables: []string{"X", "F"},
		constants: []string{"+", "-", "<", ">"},
		axiom:     "X",
		rules: map[string]string{
			"X": "F+<<X>-X>-F<-FX>+X)",
			"F": "FF",
		},
		N:     n,
		c:     0,
		Angle: 25,
	}
	lSystem.applyConstantRule()
	lSystem.GenLSystemString()
	return lSystem
}
