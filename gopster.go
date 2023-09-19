package gopster

import (
	"fmt"
	"image"
	"math"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/rasterizer"
)

// Generate renders a chart into an image.
func (c *Chart) Generate() image.Image {
	width := (float64(c.width) * (chartItemSize + c.gap)) + c.gap
	height := (float64(c.height) * (chartItemSize + c.gap)) + c.gap

	var chartTitleMargin float64
	if c.title != "" {
		chartTitleMargin = titleMargin
	}
	height += chartTitleMargin

	c.buildTitles()

	maxTitleWidth := c.maxTitleWidth()
	if c.showTitles {
		width += maxTitleWidth

		minTitleHeight := c.minTitleHeight(chartTitleMargin)
		if height < minTitleHeight {
			height = minTitleHeight
		}
	}

	cvs := canvas.New(width, height)
	ctx := canvas.NewContext(cvs)

	ctx.SetCoordSystem(canvas.CartesianIV)
	ctx.SetFillColor(c.background)
	ctx.DrawPath(0.0, 0.0, canvas.Rectangle(width, height))

	if c.title != "" {
		font := c.family.Face(titlePt, c.color, canvas.FontBlack)
		ctx.DrawText(
			width/2,
			((c.gap + 90) / 2),
			canvas.NewTextLine(font, c.title, canvas.Center),
		)
	}

	// Draw chart items
	for i, item := range c.items {
		x := float64((i % c.width))
		y := float64(i / c.width)

		imgWidth, imgHeight := c.scaledDimensions(item.coverImage)

		rect := canvas.Rect{
			X: (x * (chartItemSize + c.gap)) + c.gap + c.findCenteringOffset(imgWidth),
			Y: (y * (chartItemSize + c.gap)) + c.gap + c.findCenteringOffset(imgHeight) + chartTitleMargin,
			W: imgWidth,
			H: imgHeight,
		}
		ctx.FitImage(
			item.coverImage,
			rect,
			canvas.ImageCover,
		)
	}

	if c.showTitles {
		font := c.family.Face(chartItemTitlePt, c.color, canvas.FontBlack, canvas.FontNormal)

		currentHeight := chartTitleMargin + c.gap
		for i, title := range c.titles {
			if i == c.width*c.height {
				break
			}

			if i%c.width == 0 && i != 0 {
				currentHeight += 25
			}

			currentHeight += 25

			ctx.DrawText(
				width-maxTitleWidth+10,
				currentHeight,
				canvas.NewTextLine(font, title, canvas.Left),
			)
		}
	}

	return rasterizer.Draw(cvs, canvas.DPI(72.0), canvas.DefaultColorSpace)
}

func (c *Chart) maxTitleWidth() float64 {
	var maxTitleWidth float64

	font := c.family.Face(chartItemTitlePt, c.color, canvas.FontBlack, canvas.FontNormal)

	for _, i := range c.titles {
		width := font.TextWidth(i)
		if width > maxTitleWidth {
			maxTitleWidth = width
		}
	}

	return maxTitleWidth + chartItemTitleMargin + c.gap
}

func (c *Chart) minTitleHeight(chartTitleMargin float64) float64 {
	minTitleHeight := (c.gap * 2) + chartTitleMargin

	for i := range c.items {
		if i == c.width*c.height {
			break
		}

		minTitleHeight += 25
		if i%c.width == 0 && i != 0 {
			minTitleHeight += 25
		}
	}

	return minTitleHeight
}

func (c *Chart) scaledDimensions(img image.Image) (float64, float64) {
	diffPercent := 1.0

	width := float64(img.Bounds().Dx())
	height := float64(img.Bounds().Dy())

	if width > chartItemSize && height > chartItemSize {
		diffPercent = math.Min((chartItemSize / width), (chartItemSize / height))
	} else if width > chartItemSize {
		diffPercent = chartItemSize / width
	} else if height > chartItemSize {
		diffPercent = chartItemSize / height
	} else if width < chartItemSize && height < chartItemSize {
		diffPercent = math.Min((chartItemSize / width), (chartItemSize / height))
	}

	return math.Floor(width * diffPercent), math.Floor(height * diffPercent)
}

func (c *Chart) findCenteringOffset(dim float64) float64 {
	if dim < chartItemSize {
		return math.Floor((chartItemSize - dim) / 2)
	}
	return 0
}

func (c *Chart) buildTitles() {
	count := 1
	for i, item := range c.items {
		if i == c.width*c.height {
			break
		}

		titleString := item.title
		if item.creator != "" {
			titleString = fmt.Sprintf("%s - %s", item.creator, titleString)
		}

		if c.showNumbers {
			titleString = fmt.Sprintf("%d. %s", count, titleString)
		}

		c.titles = append(c.titles, titleString)
		count++
	}
}
