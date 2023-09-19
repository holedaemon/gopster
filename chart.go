package gopster

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"image/color"
	"io"
	"io/fs"

	"github.com/tdewolff/canvas"
)

//go:embed resources
var resources embed.FS

var resourcesDir fs.FS

func init() {
	var err error
	resourcesDir, err = fs.Sub(resources, "resources")
	if err != nil {
		panic(err)
	}
}

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

	ubuntuMono, err := resourcesDir.Open("ubuntu-mono.ttf")
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, ubuntuMono); err != nil {
		return nil, err
	}

	if err := family.LoadFont(buf.Bytes(), 0, canvas.FontRegular); err != nil {
		return nil, err
	}
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
