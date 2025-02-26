package game

import (
	"fmt"
	"image/color"
	"main/calculations"
	"main/loops"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"gonum.org/v1/gonum/mat"
)

var points = []struct{ X, Y float32 }{}

const screenWidth, screenHeight = 1440, 720
const padding = 50

const (
	alpha = 45
	v0    = 750
	x     = 0
	y     = 0
	dt    = 1
	vM    = 999999999999
)

type Game struct {
	isRunning   bool
	v           *mat.VecDense
	a           *mat.VecDense
	r           *mat.VecDense
	leftScreen  *ebiten.Image
	rightScreen *ebiten.Image
	g1          *Graph
	g2          *Graph
	timer       float32
}

func NewGame() *Game {
	return &Game{
		g1: &Graph{},
		g2: &Graph{},
	}
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
			g.r = mat.NewVecDense(2, []float64{x, y})
			g.v = calculations.MakeVelocity(alpha, v0)
			g.a = calculations.G_Vec()
			// g.a = mat.NewVecDense(2, []float64{0, 0})
			// g.a = calculations.Nul_Acceleration(g.v, g.r.AtVec(1), vM)
			// g.a = calculations.Acceleration(g.v, vM)
		}
		time.Sleep(200 * time.Millisecond)
	}

	// Если процесс вычислений активен
	if g.isRunning {
		loops.Iter(g.r, g.v, g.a, dt, vM)
		if g.r.AtVec(1) < 0 {
			g.isRunning = false
			g.r = mat.NewVecDense(2, []float64{g.r.AtVec(0), 0})
		}
		g.g1.points = append(g.g1.points, struct {
			X float32
			Y float32
		}{float32(g.r.AtVec(0)), float32(g.r.AtVec(1))})
		g.g2.points = append(g.g2.points, struct {
			X float32
			Y float32
		}{g.timer * dt, float32(g.v.Norm(2))})
		fmt.Println("Полная энергия равна", g.timer, g.v.Norm(2)*g.v.Norm(2)/2+calculations.G_Scl()*g.r.AtVec(1))
		g.timer += 1
	}

	return nil
}
func (g *Game) Draw(screen *ebiten.Image) {
	// Левый экран
	g.leftScreen = ebiten.NewImage(screenWidth/2, screenHeight)
	g.leftScreen.Clear()
	g.g1.name = "Coordinate"
	g.g1.Draw(g.leftScreen)
	screen.DrawImage(g.leftScreen, nil)

	// Правый экран
	g.rightScreen = ebiten.NewImage(screenWidth/2, screenHeight)
	g.rightScreen.Clear()
	g.g2.name = "Velocity"
	g.g2.Draw(g.rightScreen)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(screenWidth/2), 0)
	screen.DrawImage(g.rightScreen, opts)

	// Рисуем кнопку "Добавить точку"
	buttonX, buttonY, buttonW, buttonH := screenWidth-120, 20, 100, 30
	ebitenutil.DrawRect(screen, float64(buttonX), float64(buttonY), float64(buttonW), float64(buttonH), color.RGBA{0, 0, 255, 255})
	ebitenutil.DebugPrintAt(screen, "Start", buttonX+20, buttonY+10)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
