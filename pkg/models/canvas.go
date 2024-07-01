package models

// Vector3 is vector3
type Vector3 struct {
	x int
	y int
	z int
}

// Vector2 is vector2
type Vector2 struct {
	X float64
	Y float64
}

// Canvas is for output image
type Canvas struct {
	Size    Vector2
	Origin  Vector2
	Padding float64
}

// OriginCenter centers origin in the middle of image
func (c *Canvas) OriginCenter() Canvas {
	c.Origin = Vector2{
		X: c.Size.X / 2,
		Y: c.Size.Y / 2,
	}
	return *c
}

// OriginTopRight centers origin in the 'top right'
func (c *Canvas) OriginTopRight() Canvas {
	c.Origin = Vector2{
		X: c.Size.X - c.Padding,
		Y: c.Padding,
	}
	return *c
}

// OriginBottomRight centers origin in the 'bottom right'
func (c *Canvas) OriginBottomRight() Canvas {
	c.Origin = Vector2{
		X: c.Size.X - c.Padding,
		Y: c.Size.Y - c.Padding,
	}
	return *c
}

// OriginBottomCenter centers origin in the 'bottom center'
func (c *Canvas) OriginBottomCenter() Canvas {
	c.Origin = Vector2{
		X: c.Size.X / 2,
		Y: c.Size.Y - c.Padding,
	}
	return *c
}
