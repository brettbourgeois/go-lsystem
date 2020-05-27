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
	turtlegraph.DrawLSystem(binaryTree, canvas.OriginBottomCenter(), "images/binaryTree.png")

	kochCurve := models.KochCurve(4)
	turtlegraph.DrawLSystem(kochCurve, canvas.OriginBottomRight(), "images/kochCurve.png")

	sierpinskiTriangle := models.SierpinskiTriangle(6)
	turtlegraph.DrawLSystem(sierpinskiTriangle, canvas.OriginTopRight(), "images/sTriangle.png")

	sierpinskiArrowHead := models.SierpinskiArrowHead(6)
	turtlegraph.DrawLSystem(sierpinskiArrowHead, canvas.OriginBottomRight(), "images/sArrowhead.png")

	dragonCurve := models.DragonCurve(10)
	turtlegraph.DrawLSystem(dragonCurve, canvas.OriginCenter(), "images/dragonCurve.png")

	barnsleyFern := models.BarnsleyFern(4)
	turtlegraph.DrawLSystem(barnsleyFern, canvas.OriginBottomCenter(), "images/fractalplant.png")
}
