// Objective of this main.go is to generate an image using lsystems
package main

import (
	"lsystems/models"
	"lsystems/models/lsystem"
	"math"

	"github.com/fogleman/gg"
)

func main() {
	binaryTree := lsystem.BinaryTree(6)
	drawFractal(binaryTree, "binaryTree.png")

	kochCurve := lsystem.KochCurve(3)
	drawFractal(kochCurve, "kochCurve.png")

	sierpinskiTriangle := lsystem.SierpinskiTriangle(4)
	drawFractal(sierpinskiTriangle, "sTriangle.png")

	sierpinskiArrowHead := lsystem.SierpinskiArrowHead(4)
	drawFractal(sierpinskiArrowHead, "sArrowhead.png")
}

func drawFractal(lsystem lsystem.LSystem, fileName string) {
	canvas := models.Canvas{
		Height: 1024,
		Width:  1024,
	}
	dc := gg.NewContext(int(canvas.Width), int(canvas.Height))
	turtleGraph(dc, canvas, lsystem)
	dc.SavePNG("examples/" + fileName)
}
func turtleGraph(dc *gg.Context, canvas models.Canvas, lsystem lsystem.LSystem) {
	vector := models.Vector2{X: canvas.Width / 2, Y: canvas.Height / 2}
	dc.SetRGB(0, 125, 0)
	dc.SetLineWidth(2)
	for i := 0; i < len(lsystem.LSystemString); i++ {
		switch string(lsystem.LSystemString[i]) {
		case "0":
			dc.Translate(0, -5)
			dc.DrawLine(vector.X, vector.Y, vector.X, vector.Y+5)
			dc.Stroke()
		case "1":
			dc.Translate(0, -5)
			dc.DrawLine(vector.X, vector.Y, vector.X, vector.Y+5)
			dc.Stroke()
		case "[":
			dc.Push()
			dc.RotateAbout(-45*math.Pi/180, vector.X, vector.Y)
		case "]":
			dc.Pop()
			dc.RotateAbout(45*math.Pi/180, vector.X, vector.Y)
		case "F":
			dc.Translate(-10, 0)
			dc.DrawLine(vector.X, vector.Y, vector.X+10, vector.Y)
			dc.Stroke()
		case "G":
			dc.Translate(-10, 0)
			dc.DrawLine(vector.X, vector.Y, vector.X+10, vector.Y)
			dc.Stroke()
		case "+":
			dc.RotateAbout(lsystem.Angle*(math.Pi/180), vector.X, vector.Y)
		case "-":
			dc.RotateAbout(-lsystem.Angle*(math.Pi/180), vector.X, vector.Y)
		}
	}
}
