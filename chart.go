package gopster

import (
	"errors"
	"fmt"
	"image/color"

	"github.com/tdewolff/canvas"
)

// ErrorChart is returned when a chart is misconfigured.
var ErrorChart = errors.New("gopster: error creating chart")

const (
	maxSize     = 3
	defaultSize = 3

	maxGap     = 150
	defaultGap = 20

	titleMargin = 60
	titlePt     = 120.0

	chartItemTitleMargin = 20
	chartItemSize        = 260.0
	chartItemTitlePt     = 60.0

	mmToPixel = 3.7795275591
)

// Chart is a Topster chart.
type Chart struct {
	title           string
	items           []*chartItem
	width           int
	height          int
	backgroundColor string
	textColor       string
	showNumbers     bool
	showTitles      bool
	gap             float64

	background color.Color
	color      color.Color
	family     *canvas.FontFamily
}

// NewChart creates a new chart with the given settings.
func NewChart(opts ...Option) (*Chart, error) {
	c := &Chart{
		items:  make([]*chartItem, 0),
		width:  defaultSize,
		height: defaultSize,
		gap:    defaultGap,
	}

	family := canvas.NewFontFamily("ubuntu-mono")
	family.LoadFontFile("resources/ubuntu-mono.ttf", canvas.FontRegular)
	c.family = family

	for _, o := range opts {
		o(c)
	}

	if c.width <= 0 || c.width > maxSize {
		return nil, fmt.Errorf("%w: width must be a number between 0 and %d", ErrorChart, maxSize)
	}

	if c.height <= 0 || c.height > maxSize {
		return nil, fmt.Errorf("%w: height must be a number between 0 and %d", ErrorChart, maxSize)
	}

	if c.gap < 0 || c.gap > maxGap {
		return nil, fmt.Errorf("%w: gap must be a number between 0 and %d", ErrorChart, maxGap)
	}

	if c.backgroundColor != "" {
		bc, err := parseHexColor(c.backgroundColor)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrorChart, err)
		}

		c.background = bc
	} else {
		c.background = color.Black
	}

	if c.textColor != "" {
		tc, err := parseHexColor(c.textColor)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrorChart, err)
		}

		c.color = tc
	} else {
		c.color = color.White
	}

	return c, nil
}
