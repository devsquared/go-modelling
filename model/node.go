package model

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"strings"
)

// Node defines the graphical output for a single node of a state machine.
// To help visualize, here is an example node:

//          +----------+-----------+
//			|     ASN Expected     |
//			+----------------------+
//			| "Status": "expected" |
//			+----------+-----------+

const maxNodeWidth = 50
const maxLabelWidth = 46
const padding = maxNodeWidth - maxLabelWidth

func DrawStateNode(w io.Writer, label, content string) error {
	if len(label) > maxLabelWidth {
		label = label[:maxLabelWidth]
	}

	trimmedContent := strings.TrimSpace(content)

	if len(trimmedContent) > maxLabelWidth {
		trimmedContent = trimmedContent[:maxLabelWidth-3]
		trimmedContent += "..."
	}

	return drawNodeBox(w, label, trimmedContent)
}

func drawNodeBox(w io.Writer, label, content string) error {
	var nodeWidth int
	if len(label) > len(content) {
		nodeWidth = len(label) + padding
		content = padString(label, content)
	} else {
		nodeWidth = len(content) + padding
		label = padString(content, label)
	}

	var buff bytes.Buffer

	if err := drawNodeHeader(&buff, nodeWidth, label); err != nil {
		return fmt.Errorf("trouble drawing node header: %w", err)
	}

	if err := drawContentBox(&buff, nodeWidth, content); err != nil {
		return fmt.Errorf("trouble drawing node content box: %w", err)
	}

	if _, err := w.Write(buff.Bytes()); err != nil {
		return fmt.Errorf("error drawing state node: %w", err)
	}

	return nil
}

func padString(biggerString string, smallerString string) string {
	labelNeededPadding := len(biggerString) - len(smallerString)
	var (
		prePadding  int
		postPadding int
	)

	if labelNeededPadding%2 == 0 {
		prePadding = labelNeededPadding / 2
		postPadding = labelNeededPadding / 2
	} else {
		prePadding = labelNeededPadding / 2
		postPadding = (labelNeededPadding / 2) + 1
	}

	var newLabel string
	for i := 0; i < prePadding; i++ {
		newLabel += " "
	}

	newLabel += smallerString

	for i := 0; i < postPadding; i++ {
		newLabel += " "
	}

	return newLabel
}

func drawNodeHeader(buff *bytes.Buffer, nodeWidth int, label string) error {
	if err := drawFullHorizontalLine(buff, nodeWidth, true); err != nil {
		return errors.New("cannot draw horizontal line")
	}

	// draw label
	buff.Write([]byte(VerticalLine + Space + label + Space + VerticalLine + NewLine))

	if err := drawFullHorizontalLine(buff, nodeWidth, false); err != nil {
		return errors.New("cannot draw horizontal line")
	}

	return nil
}

func drawContentBox(buff *bytes.Buffer, nodeWidth int, content string) error {
	// draw content
	buff.Write([]byte(VerticalLine + Space + content + Space + VerticalLine + NewLine))

	if err := drawFullHorizontalLine(buff, nodeWidth, true); err != nil {
		return errors.New("cannot draw horizontal line")
	}

	return nil
}

func drawFullHorizontalLine(buff *bytes.Buffer, width int, isOuter bool) error {
	buff.Write([]byte(Corner))
	for i := 2; i < width; i++ {
		if isOuter && i == findCenter(width) {
			buff.Write([]byte(LabelMark))
		} else {
			buff.Write([]byte(HorizontalLine))
		}
	}

	buff.Write([]byte(Corner))

	buff.Write([]byte(NewLine))

	return nil
}

func findCenter(size int) int {
	return int(math.Floor(float64(size / 2)))
}
