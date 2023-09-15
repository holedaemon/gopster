package gopster

// Option configures a chart.
type Option func(*Chart)

// Title sets a chart's title. Defaults to an empty string.
func Title(t string) Option {
	return func(c *Chart) {
		c.title = t
	}
}

// Width sets a chart's width. Defaults to 3.
func Width(w int) Option {
	return func(c *Chart) {
		c.width = w
	}
}

// Height sets a chart's height. Defaults to 3.
func Height(h int) Option {
	return func(c *Chart) {
		c.height = h
	}
}

// ShowNumbers toggles number rendering on a chart. Defaults to false.
// ShowTitles must also be true for this to work.
func ShowNumbers() Option {
	return func(c *Chart) {
		c.showNumbers = true
	}
}

// ShowTitles toggles title rendering on a chart. Defaults to false.
func ShowTitles() Option {
	return func(c *Chart) {
		c.showTitles = true
	}
}

// Gap sets a chart's gap. Defaults to 20.
func Gap(g float64) Option {
	return func(c *Chart) {
		c.gap = g
	}
}

// BackgroundColor sets a chart's background color. Must be a valid hex color string.
// Defaults to black.
func BackgroundColor(bc string) Option {
	return func(c *Chart) {
		c.backgroundColor = bc
	}
}

// TextColor sets a chart's text color. Must be a valid hex color string.
// Defaults to white.
func TextColor(tc string) Option {
	return func(c *Chart) {
		c.textColor = tc
	}
}
