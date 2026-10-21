package pongo

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/devskar/pongo/pongo/entities"
	"github.com/devskar/pongo/pongo/input"
	"github.com/devskar/pongo/pongo/score"
	"github.com/devskar/pongo/vectors"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

var (
	ballRadius float32 = ScreenWidth / 100
	ballX      float32 = ScreenWidth/2 - ballRadius
	ballY      float32 = ScreenHeight/2 - ballRadius

	paddleHeight float32 = ScreenHeight / 10
	paddleWidht  float32 = ScreenWidth / 100
	paddleY      float32 = ScreenHeight/2 - paddleHeight/2
	paddleLeftX  float32 = 5 * paddleWidht
	paddleRightX float32 = ScreenWidth - paddleLeftX - paddleWidht
)

type Game struct {
	p1Input input.PlayerInput
	p2Input input.PlayerInput

	score      score.Score
	scoreboard score.Scoreboard

	leftPaddle  entities.Paddle
	rightPaddle entities.Paddle

	ball entities.Ball

	board entities.Board
}

func NewGame() Game {

	return Game{
		p1Input: input.NewPlayerInput(ebiten.KeyW, ebiten.KeyS),
		p2Input: input.NewPlayerInput(ebiten.KeyUp, ebiten.KeyDown),

		leftPaddle:  entities.NewPaddle(paddleLeftX, paddleY, paddleWidht, paddleHeight),
		rightPaddle: entities.NewPaddle(paddleRightX, paddleY, paddleWidht, paddleHeight),
		ball:        entities.NewBall(ballX, ballY, ballRadius, vectors.Vector2{X: 3, Y: 3}),

		board: entities.Board{},

		score:      score.Score{P1: 0, P2: 0},
		scoreboard: score.NewScoreboard(),
	}
}

func (g *Game) Update() error {
	g.p1Input.Update()
	g.p2Input.Update()

	g.leftPaddle.Update(&g.p1Input)
	g.rightPaddle.Update(&g.p2Input)

	// Ball wall collission
	if 0 > g.ball.X-g.ball.Radius {
		g.score.P2++
		g.ball.Speed.X = g.ball.Speed.X * -1

		g.resetPos()
	}

	if g.ball.X+g.ball.Radius > ScreenWidth {
		g.score.P1++
		g.ball.Speed.X = g.ball.Speed.X * -1

		g.resetPos()
	}

	if 0 > g.ball.Y-g.ball.Radius || g.ball.Y+g.ball.Radius > ScreenHeight {
		g.ball.Speed.Y *= -1.1
		g.ball.Speed.X *= 1.1
	}

	if CheckCollission(g.ball, g.leftPaddle) || CheckCollission(g.ball, g.rightPaddle) {
		g.ball.Speed.X *= -1.1
		g.ball.Speed.Y *= 1.1
	}

	g.ball.Update()
	g.scoreboard.Update(g.score)

	return nil
}

func (g *Game) resetPos() {

	g.leftPaddle.Y = paddleY
	g.rightPaddle.Y = paddleY
	g.ball.X = ballX
	g.ball.Y = ballY

	g.ball.Speed.X = 3
	g.ball.Speed.Y = 3
}

func (g *Game) Draw(screen *ebiten.Image) {
	es := []entities.Entity{g.board, g.leftPaddle, g.ball, g.rightPaddle, g.scoreboard}

	for _, entity := range es {
		entity.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
