package shapes

type Cuboid struct {
	Length, Width, Height float64
}

func NewCuboid(length, width, height float64) *Cuboid {
	return &Cuboid{Length: length, Width: width, Height: height}
}

func (c *Cuboid) Area() float64 {
	return 2 * (c.Length*c.Width + c.Width*c.Height + c.Height*c.Length)
}

func (c *Cuboid) Volume() float64 {
	return c.Length * c.Width * c.Height
}

func (c *Cuboid) Calculate() float64 {
	return c.Volume()
}
