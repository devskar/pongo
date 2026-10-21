package score

import (
	"bytes"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Scoreboard struct {
	score      Score
	fontSource *text.GoTextFaceSource
	fontFace   *text.GoTextFace
}

func NewScoreboard() Scoreboard {
	// Load variable-width font embedded in Ebitengine.
	var err error
	fontSource, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Panic(err)
	}

	// Create 16px font face from the above source.
	fontFace := &text.GoTextFace{
		Source: fontSource,
		Size:   42,
	}

	return Scoreboard{
		fontSource: fontSource,
		fontFace:   fontFace,
	}
}

func (s Scoreboard) Draw(screen *ebiten.Image) {

	size := screen.Bounds().Size()

	txtOp := &text.DrawOptions{}
	txtOp.GeoM.Translate(float64(size.X)/2, 0)

	// By default, the text is white. We can call ScaleWithColor to specify a different color.
	//colorGreen := color.RGBA{0, 255, 0, 255}
	//txtOp.ColorScale.ScaleWithColor(colorGreen)
	text.Draw(screen, fmt.Sprintf("%d : %d", s.score.P1, s.score.P2), s.fontFace, txtOp)
}

func (s *Scoreboard) Update(score Score) {
	s.score = score
}
