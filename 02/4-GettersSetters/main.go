package main

import (
	"errors"
)

type Circle struct {
	radius float64
}

func (c *Circle) Radius() float64 {
	return c.radius
}

func (c *Circle) SetRadius(radius float64) {
	c.radius = radius
}

type Square struct {
	side float64
}

func (s *Square) Side() float64 {
	return s.side
}

func (s *Square) SetSide(side float64) error {
	if side <= 0 {
		return errors.New("invalid side")
	}
	s.side = side
	return nil
}

func main() {

}
