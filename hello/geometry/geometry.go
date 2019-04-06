package geometry

type Vertex struct {
	X int
	Y int
}

type Square struct {
	A Vertex
	B Vertex
}

func (s Square) Area() int {
	a, b := s.getSides()
	return a * b
}

func (s Square) Perimeter() int {
	a, b := s.getSides()
	return 2*a + 2*b
}

func (s Square) getSides() (int, int) {
	a := s.B.X - s.A.X
	b := s.B.Y - s.A.Y
	return a, b
}
