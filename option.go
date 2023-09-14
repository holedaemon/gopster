package gopster

// Option configures a Chart.
type Option func(*Chart)

// Title sets a Chart's title. Defaults to an empty string.
func Title(t string) Option {
	return func(c *Chart) {
		c.title = t
	}
}

// Width sets a Chart's width. Defaults to 3.
func Width(w int) Option {
	return func(c *Chart) {
		c.width = w
	}
}

// Height sets a Chart's height. Defaults to 3.
func Height(h int) Option {
	return func(c *Chart) {
		c.height = h
	}
}

// ShowNumbers toggles number rendering on a Chart. Defaults to false.
// ShowTitles must also be true for this to work.
func ShowNumbers() Option {
	return func(c *Chart) {
		c.showNumbers = true
	}
}

// ShowTitles toggles title rendering on a Chart. Defaults to false.
func ShowTitles() Option {
	return func(c *Chart) {
		c.showTitles = true
	}
}

// Gap sets a Chart's gap. Defaults to 20.
func Gap(g float64) Option {
	return func(c *Chart) {
		c.gap = g
	}
}

// BackgroundColor sets a Chart's background color. Must be a valid hex color string.
// Defaults to black.
func BackgroundColor(bc string) Option {
	return func(c *Chart) {
		c.backgroundColor = bc
	}
}

// TextColor sets a Chart's text color. Must be a valid hex color string.
// Defaults to white.
func TextColor(tc string) Option {
	return func(c *Chart) {
		c.textColor = tc
	}
}
