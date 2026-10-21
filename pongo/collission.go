package pongo

import "github.com/devskar/pongo/pongo/entities"

func CheckCollission(b entities.Ball, p entities.Paddle) bool {
	return checkCollissionX(b, p) && checkCollissionY(b, p)
}

func checkCollissionY(b entities.Ball, p entities.Paddle) bool {
	var (
		ballUpperY float32 = b.Y - b.Radius
		ballLowerY float32 = b.Y + b.Radius

		paddleUpperY float32 = p.Y
		paddleLowerY float32 = p.Y + p.Height
	)

	return ballLowerY > paddleUpperY && ballUpperY < paddleLowerY
}

func checkCollissionX(b entities.Ball, p entities.Paddle) bool {
	return p.X > b.X-b.Radius && p.X < b.X+b.Radius
}
