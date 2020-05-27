// Objective of this main.go is to generate an image using lsystems
package main

import (
	"lsystems/models"
	"lsystems/pkg/turtlegraph"
)

func main() {
	canvasSize := models.Vector2{X: 850, Y: 600}
	canvas := models.Canvas{
		Size:    canvasSize,
		Padding: 20,
	}

	binaryTree := models.BinaryTree(6)
	turtlegraph.DrawLSystem(binaryTree, canvas.OriginBottomCenter(), "examples/binaryTree.png")

	kochCurve := models.KochCurve(4)
	turtlegraph.DrawLSystem(kochCurve, canvas.OriginBottomRight(), "examples/kochCurve.png")

	sierpinskiTriangle := models.SierpinskiTriangle(4)
	turtlegraph.DrawLSystem(sierpinskiTriangle, canvas.OriginCenter(), "examples/sTriangle.png")

	sierpinskiArrowHead := models.SierpinskiArrowHead(4)
	turtlegraph.DrawLSystem(sierpinskiArrowHead, canvas.OriginCenter(), "examples/sArrowhead.png")

	dragonCurve := models.DragonCurve(10)
	turtlegraph.DrawLSystem(dragonCurve, canvas.OriginCenter(), "examples/dragonCurve.png")

	barnsleyFern := models.BarnsleyFern(4)
	turtlegraph.DrawLSystem(barnsleyFern, canvas.OriginCenter(), "examples/fractalplant.png")
}
