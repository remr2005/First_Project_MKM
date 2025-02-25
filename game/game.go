package game

import (
	"fmt"
	"main/calculations"
	"main/loops"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"gonum.org/v1/gonum/mat"
)

var points = []struct{ X, Y float32 }{}

const screenWidth, screenHeight = 640, 480
const padding = 50

const (
	alpha = 45
	v0    = 750
	x     = 0
	y     = 0
	dt    = 0.2
	vM    = 150
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

type Game struct {
	maxX, maxY     float32
	scaleX, scaleY float32
	stepX, stepY   float32
	isRunning      bool
	v              *mat.VecDense
	a              *mat.VecDense
	r              *mat.VecDense
}

func (g *Game) Update() error {
	// Проверка нажатия кнопки "Старт"
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		// Проверка, находится ли клик в области кнопки "Старт"
		if mx >= screenWidth-120 && mx <= screenWidth-20 && my >= 20 && my <= 50 {
			points = append(points, struct {
				X float32
				Y float32
			}{x, y})
			g.isRunning = true
			g.v = calculations.MakeVelocity(alpha, v0)
			g.a = mat.NewVecDense(2, nil) // nil → нулевые значения
			g.r = mat.NewVecDense(2, []float64{x, y})
		}
	}

	// Если процесс вычислений активен
	if g.isRunning {
		loops.Iter(g.r, g.v, g.a, dt, vM)
		points = append(points, struct {
			X float32
			Y float32
		}{float32(g.r.AtVec(0)), float32(g.r.AtVec(1))})
		if g.r.AtVec(1) < 0 {
			g.isRunning = false
		}
		if g.r.AtVec(0) > float64(g.maxX) {
			g.maxX = float32(g.r.AtVec(0))
		}
		if g.r.AtVec(1) > float64(g.maxY) {
			g.maxY = float32(g.r.AtVec(1))
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.stepX = chooseStep(g.maxX)
	g.stepY = chooseStep(g.maxY)
	g.scaleX = (screenWidth - 2*padding) / g.maxX
	g.scaleY = (screenHeight - 2*padding) / g.maxY

	// Рисуем подпись графика
	ebitenutil.DebugPrintAt(screen, "Coordinate", screenWidth/2-80, 5)

	// Рисуем оси
	vector.StrokeLine(screen, padding, screenHeight-padding, screenWidth-padding, screenHeight-padding, 1, WhiteColor{}, false) // X
	vector.StrokeLine(screen, padding, screenHeight-padding, padding, padding, 1, WhiteColor{}, false)                          // Y

	// Подписи осей
	ebitenutil.DebugPrintAt(screen, "X", screenWidth-padding-30, screenHeight-padding+10)
	ebitenutil.DebugPrintAt(screen, "Y", padding-30, padding-30)

	// Рисуем сетку с подписями
	for x := float32(0.0); x <= g.maxX; x += g.stepX {
		px := padding + x*g.scaleX
		vector.StrokeLine(screen, px, padding, px, screenHeight-padding, 1, WhiteColor{}, false)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%.0f", x), int(px)-15, screenHeight-padding+5)
	}

	for y := float32(0.0); y <= g.maxY; y += g.stepY {
		py := screenHeight - padding - y*g.scaleY
		vector.StrokeLine(screen, padding, py, screenWidth-padding, py, 1, WhiteColor{}, false)
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%.0f", y), padding-35, int(py)-5)
	}

	// Рисуем точки и линии
	for i := 0; i < len(points)-1; i++ {
		x1 := padding + points[i].X*g.scaleX
		y1 := screenHeight - padding - points[i].Y*g.scaleY
		x2 := padding + points[i+1].X*g.scaleX
		y2 := screenHeight - padding - points[i+1].Y*g.scaleY

		vector.StrokeLine(screen, x1, y1, x2, y2, 1, Red{}, true)      // Красная линия
		vector.DrawFilledRect(screen, x1-2, y1-2, 4, 4, Green{}, true) // Зеленая точка
	}

	// Рисуем кнопку "Старт"
	buttonX, buttonY, buttonW, buttonH := screenWidth-120, 20, 100, 30
	vector.DrawFilledRect(screen, float32(buttonX), float32(buttonY), float32(buttonW), float32(buttonH), Blue{}, false)
	ebitenutil.DebugPrintAt(screen, "Start", buttonX+25, buttonY+10)

	// Показываем состояние процесса
	if g.isRunning {
		ebitenutil.DebugPrintAt(screen, "Calculating", 20, 40)
	} else {
		ebitenutil.DebugPrintAt(screen, "It s over", 20, 40)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
