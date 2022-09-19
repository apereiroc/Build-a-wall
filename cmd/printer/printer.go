package printer

const fullBrick = "■■"
const halfBrick = "■"
const mortar = "|"

type Printer struct {
	rows   int
	bricks int
}

func NewPrinter(rows, bricks int) *Printer {
	return &Printer{rows: rows, bricks: bricks}
}

func (p Printer) GetOutput() string {
	var rowBase string = ""

	for i := 0; i < p.bricks; i++ {
		rowBase += fullBrick + mortar
	}
	rowBase = rowBase[:len(rowBase)-1]

	var rowTop string = halfBrick + mortar
	for i := 0; i < p.bricks-1; i++ {
		rowTop += fullBrick + mortar
	}
	rowTop += halfBrick

	var output string = ""
	for i := 0; i < p.rows; i++ {
		if i%2 == 0 {
			output += rowBase + "\n"
		} else {
			output += rowTop + "\n"
		}
	}

	return output
}
