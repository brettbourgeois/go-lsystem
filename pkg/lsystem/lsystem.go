package lsystem

import (
	"lsystems/pkg/models"
	"math"
	"strings"

	"github.com/fogleman/gg"
)

// lSystem struct defines an lsystem
type lSystem struct {
	variables     []string
	constants     []string
	axiom         string
	N             int
	c             int
	rules         map[string]string
	lSystemString string
	Angle         float64
}

// Draw will turtle graph an LSystem string and save a `canvas` sized image at the provided file path
func (ls *lSystem) Draw(canvas models.Canvas, fileName string) {
	// Setup drawing context with basic config
	dc := gg.NewContext(int(canvas.Size.X), int(canvas.Size.Y))
	dc.SetRGB(0, 125, 0)
	dc.SetLineWidth(2)
	// use drawing context to turtle graph the lsystem
	ls.turtleGraph(dc, canvas.Origin)
	dc.SavePNG(fileName)

}

// turtleGraph uses a drawing context to turtle graph an lsystem
func (ls *lSystem) turtleGraph(dc *gg.Context, origin models.Vector2) {
	drawVector := origin
	for i := 0; i < len(ls.lSystemString); i++ {
		switch string(ls.lSystemString[i]) {
		case "0":
			dc.Translate(0, -5)
			dc.DrawLine(drawVector.X, drawVector.Y, drawVector.X, drawVector.Y+5)
			dc.Stroke()
		case "1":
			dc.Translate(0, -5)
			dc.DrawLine(drawVector.X, drawVector.Y, drawVector.X, drawVector.Y+5)
			dc.Stroke()
		case "[":
			dc.Push()
			dc.RotateAbout(-45*math.Pi/180, drawVector.X, drawVector.Y)
		case "]":
			dc.Pop()
			dc.RotateAbout(45*math.Pi/180, drawVector.X, drawVector.Y)
		case "<":
			dc.Push()
		case ">":
			dc.Pop()
		case "A":
			dc.Translate(0, -10)
			dc.DrawLine(drawVector.X, drawVector.Y, drawVector.X, drawVector.Y+10)
			dc.Stroke()
		case "B":
			dc.Translate(0, -10)
			dc.DrawLine(drawVector.X, drawVector.Y, drawVector.X, drawVector.Y+10)
			dc.Stroke()
		case "F":
			dc.Translate(-10, 0)
			dc.DrawLine(drawVector.X, drawVector.Y, drawVector.X+10, drawVector.Y)
			dc.Stroke()
		case "G":
			dc.Translate(-10, 0)
			dc.DrawLine(drawVector.X, drawVector.Y, drawVector.X+10, drawVector.Y)
			dc.Stroke()
		case "+":
			dc.RotateAbout(ls.Angle*(math.Pi/180), drawVector.X, drawVector.Y)
		case "-":
			dc.RotateAbout(-ls.Angle*(math.Pi/180), drawVector.X, drawVector.Y)
		}
	}
}

func (S *lSystem) applyConstantRule() {
	for _, v := range S.constants {
		S.rules[v] = v
	}
}

// genlSystemString replaces strings based on rules
func (S *lSystem) genlSystemString() {
	if S.c <= S.N {
		if S.c == 0 {
			S.lSystemString = S.axiom
		} else {
			tmp := make([]string, 0)
			for r := range S.lSystemString {
				tmp = append(tmp, S.rules[string(S.lSystemString[r])])
			}
			S.lSystemString = strings.Join(tmp, "")
		}
	}
	if S.c != S.N {
		S.c++
		S.genlSystemString()
	}
}

// BinaryTree will return an BinaryTree lSystem
func BinaryTree(n int) lSystem {
	lSystem := lSystem{
		constants: []string{"[", "]"},
		variables: []string{"1,0"},
		axiom:     "0",
		N:         n,
		c:         0,
		rules:     map[string]string{"1": "11", "0": "1[0]0"},
	}
	lSystem.applyConstantRule()
	lSystem.genlSystemString()
	return lSystem
}

// SierpinskiTriangle will return an SierpinskiTriangle LSystem
func SierpinskiTriangle(n int) lSystem {
	lSystem := lSystem{
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
	lSystem.genlSystemString()
	return lSystem
}

// SierpinskiArrowHead will return an SierpinskiTriangle LSystem
func SierpinskiArrowHead(n int) lSystem {
	lSystem := lSystem{
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
	lSystem.genlSystemString()
	return lSystem
}

// KochCurve will return an KochCurve LSystem
func KochCurve(n int) lSystem {
	lSystem := lSystem{
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
	lSystem.genlSystemString()
	return lSystem
}

// DragonCurve will return an DragonCurve LSystem
func DragonCurve(n int) lSystem {
	lSystem := lSystem{
		variables: []string{"A", "B"},
		constants: []string{"F", "+", "-"},
		axiom:     "FA",
		rules: map[string]string{
			"A": "A+BF+",
			"B": "-FA-B",
		},
		N:     n,
		c:     0,
		Angle: 90,
	}
	lSystem.applyConstantRule()
	lSystem.genlSystemString()
	return lSystem
}

// FractalPlant will return an FractalPlant lSystem
func FractalPlant(n int) lSystem {
	lSystem := lSystem{
		variables: []string{"X", "A"},
		constants: []string{"+", "-", "<", ">"},
		axiom:     "X",
		rules: map[string]string{
			"X": "A+<<X>-X>-A<-AX>+X)",
			"A": "AA",
		},
		N:     n,
		c:     0,
		Angle: 25,
	}
	lSystem.applyConstantRule()
	lSystem.genlSystemString()
	return lSystem
}
