package model

import (
	"bytes"
	"errors"
	"io"
	"math"
)

// Arrow defines the graphical output for an event in a state machine.
// To help visualize, here is an example arrow *going right*:

//		Manual Inventory Removal
// ----------------+---------------->

type Direction int64

var ErrDirectionNotDefined = errors.New("direction provided not defined")

const (
	Right Direction = iota
	Left
	Up
	Down
)

const (
	DownArrow  = "V"
	LabelMark  = "+"
	LeftArrow  = "<"
	RightArrow = ">"
	UpArrow    = "^"
)

// TODO: start with the base 4 (check!) and then expand

// DrawArrow takes a direction and a size and draws an arrow with an arrow in the appropriate direction at then end.
func DrawArrow(w io.Writer, dir Direction, size int, label string) error {
	switch dir {
	case Right:
		return drawRightArrow(w, size, label)
	case Left:
		return drawLeftArrow(w, size, label)
	case Up:
		return drawUpArrow(w, size, label)
	case Down:
		return drawDownArrow(w, size, label)
	default:
		return ErrDirectionNotDefined
	}
}

func drawRightArrow(w io.Writer, size int, label string) error {
	var byteBuff bytes.Buffer

	// first the label
	offset := (size - len(label)) / 2
	for i := 0; i < offset; i++ {
		byteBuff.Write([]byte(Space))
	}

	byteBuff.Write([]byte(label + NewLine))

	drawHorizontalBase(&byteBuff, size)
	// at the end, add an arrow!
	byteBuff.Write([]byte(RightArrow))

	if _, err := w.Write(byteBuff.Bytes()); err != nil {
		return err
	}

	return nil
}

func drawLeftArrow(w io.Writer, size int, label string) error {
	var byteBuff bytes.Buffer

	// first the label
	offset := ((size - len(label)) / 2) + 1
	for i := 0; i < offset; i++ {
		byteBuff.Write([]byte(Space))
	}

	byteBuff.Write([]byte(label + NewLine))

	// at the beginning, add an arrow!
	byteBuff.Write([]byte(LeftArrow))
	drawHorizontalBase(&byteBuff, size)

	if _, err := w.Write(byteBuff.Bytes()); err != nil {
		return err
	}

	return nil
}

func drawUpArrow(w io.Writer, size int, label string) error {
	var byteBuff bytes.Buffer
	// add up arrow to start
	byteBuff.Write([]byte(UpArrow + NewLine))
	drawVerticalBase(&byteBuff, size, label)

	if _, err := w.Write(byteBuff.Bytes()); err != nil {
		return err
	}

	return nil
}

func drawDownArrow(w io.Writer, size int, label string) error {
	var byteBuff bytes.Buffer
	drawVerticalBase(&byteBuff, size, label)
	// finish with a down arrow
	byteBuff.Write([]byte(DownArrow))

	if _, err := w.Write(byteBuff.Bytes()); err != nil {
		return err
	}

	return nil
}

// arrowCenter finds the middle of the arrow by taking the floor of the size divided by 2.
func arrowCenter(size int) int {
	return int(math.Floor(float64(size / 2)))
}

// drawHorizontalBase draws the base of a horizontal arrow
func drawHorizontalBase(buff *bytes.Buffer, size int) {
	for i := 1; i < size; i++ { // 1 smaller than the size to account for the arrow
		if i == arrowCenter(size) {
			buff.Write([]byte(LabelMark))
		} else {
			buff.Write([]byte(HorizontalLine))
		}
	}
}

func drawVerticalBase(buff *bytes.Buffer, size int, label string) {
	for i := 1; i < size; i++ {
		if i == arrowCenter(size) {
			buff.Write([]byte(LabelMark + Space + label + NewLine))
		} else {
			buff.Write([]byte(VerticalLine + NewLine))
		}
	}
}
