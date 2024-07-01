// Objective of this main.go is to generate an image using lsystems
package main

import (
	"lsystems/pkg/lsystem"
	"lsystems/pkg/models"
)

func main() {
	canvasSize := models.Vector2{X: 850, Y: 600}
	canvas := models.Canvas{
		Size:    canvasSize,
		Padding: 20,
	}

	binaryTree := lsystem.BinaryTree(6)
	binaryTree.Draw(canvas.OriginBottomCenter(), "images/binaryTree.png")

	kochCurve := lsystem.KochCurve(4)
	kochCurve.Draw(canvas.OriginBottomRight(), "images/kochCurve.png")

	sierpinskiTriangle := lsystem.SierpinskiTriangle(6)
	sierpinskiTriangle.Draw(canvas.OriginTopRight(), "images/sTriangle.png")

	sierpinskiArrowHead := lsystem.SierpinskiArrowHead(6)
	sierpinskiArrowHead.Draw(canvas.OriginBottomRight(), "images/sArrowhead.png")

	canvas.Padding = 20
	fractalPlant := lsystem.FractalPlant(4)
	fractalPlant.Draw(canvas.OriginBottomCenter(), "images/fractalplant.png")

	canvas.Padding = 150
	dragonCurve := lsystem.DragonCurve(10)
	dragonCurve.Draw(canvas.OriginBottomCenter(), "images/dragonCurve.png")

	// Generate an image for the README.md
	canvas = models.Canvas{
		Size:    models.Vector2{X: 325, Y: 425},
		Padding: 20,
	}

	exampleImage := lsystem.FractalPlant(4)
	exampleImage.Draw(canvas.OriginBottomCenter(), "images/readme-example.png")

}
