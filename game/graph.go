package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var stepSizes = []float32{10, 20, 50, 100, 200, 500, 1000, 2000}

func chooseStep(rangeVal float32) float32 {
	for _, step := range stepSizes {
		if rangeVal/step <= 10 {
			return step
		}
	}
	return stepSizes[len(stepSizes)-1]
}

func findBounds(points []struct{ X, Y float32 }) (maxX, maxY float32) {
	maxX, maxY = 0, 0
	for _, p := range points {
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	return maxX, maxY
}

type Graph struct {
	points []struct {
		X float32
		Y float32
	}
	maxX, maxY     float32
	scaleX, scaleY float32
	stepX, stepY   float32
	name           string
}

func (g *Graph) Draw(screen *ebiten.Image) {
	g.maxX, g.maxY = findBounds(g.points)
	g.stepX = chooseStep(g.maxX)
	g.stepY = chooseStep(g.maxY)
	g.scaleX = (screenWidth/2 - 2*padding) / g.maxX
	g.scaleY = (screenHeight - 2*padding) / g.maxY

	// Рисуем подпись графика
	ebitenutil.DebugPrintAt(screen, g.name, screenWidth/4-80, 5)

	// Рисуем оси
	vector.StrokeLine(screen, padding, screenHeight-padding, screenWidth/2-padding, screenHeight-padding, 1, WhiteColor{}, false) // X
	vector.StrokeLine(screen, padding, screenHeight-padding, padding, padding, 1, WhiteColor{}, false)                            // Y

	// Подписи осей
	ebitenutil.DebugPrintAt(screen, "X", screenWidth/2-padding-30, screenHeight-padding+10)
	ebitenutil.DebugPrintAt(screen, "Y", padding-30, padding-30)

	// Рисуем сетку с подписями
	for x := float32(0.0); x <= g.maxX; x += g.stepX {
		px := padding + x*g.scaleX
		vector.StrokeLine(screen, px, padding, px, screenHeight-padding, 1, WhiteColor{}, false)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%.0f", x), int(px)-15, screenHeight-padding+5)
	}

	for y := float32(0.0); y <= g.maxY; y += g.stepY {
		py := screenHeight - padding - y*g.scaleY
		vector.StrokeLine(screen, padding, py, screenWidth/2-padding, py, 1, WhiteColor{}, false)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%.0f", y), padding-35, int(py)-5)
	}

	// Рисуем точки и линии
	for i := 0; i < len(g.points)-1; i++ {
		x1 := padding + g.points[i].X*g.scaleX
		y1 := screenHeight - padding - g.points[i].Y*g.scaleY
		x2 := padding + g.points[i+1].X*g.scaleX
		y2 := screenHeight - padding - g.points[i+1].Y*g.scaleY

		vector.StrokeLine(screen, x1, y1, x2, y2, 1, Red{}, true)      // Красная линия
		vector.DrawFilledRect(screen, x1-2, y1-2, 4, 4, Green{}, true) // Зеленая точка
	}
}

func (g *Graph) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth / 2, screenHeight
}
