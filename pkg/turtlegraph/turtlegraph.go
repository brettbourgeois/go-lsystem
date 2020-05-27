package turtlegraph

import (
	"lsystems/models"
	"math"

	"github.com/fogleman/gg"
)

// DrawLSystem will turtle graph an LSystem string onto a 'canvas' sized image 
func DrawLSystem(lsystem models.LSystem, canvas models.Canvas, fileName string) {
	dc := gg.NewContext(int(canvas.Size.X), int(canvas.Size.Y))
	turtleGraph(dc, canvas.Origin, lsystem)
	dc.SavePNG(fileName)
}

func turtleGraph(dc *gg.Context, origin models.Vector2, lsystem models.LSystem) {
	drawVector := origin
	dc.SetRGB(0, 125, 0)
	dc.SetLineWidth(2)
	for i := 0; i < len(lsystem.LSystemString); i++ {
		switch string(lsystem.LSystemString[i]) {
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
			dc.RotateAbout(lsystem.Angle*(math.Pi/180), drawVector.X, drawVector.Y)
		case "-":
			dc.RotateAbout(-lsystem.Angle*(math.Pi/180), drawVector.X, drawVector.Y)
		}
	}
}
