package gopster

type Option func(*Chart)

func Title(t string) Option {
	return func(c *Chart) {
		c.title = t
	}
}

func Width(w int) Option {
	return func(c *Chart) {
		c.width = w
	}
}

func Height(h int) Option {
	return func(c *Chart) {
		c.height = h
	}
}

func ShowNumbers() Option {
	return func(c *Chart) {
		c.showNumbers = true
	}
}

func ShowTitles() Option {
	return func(c *Chart) {
		c.showTitles = true
	}
}

func Gap(g float64) Option {
	return func(c *Chart) {
		c.gap = g
	}
}

func BackgroundColor(bc string) Option {
	return func(c *Chart) {
		c.backgroundColor = bc
	}
}

func TextColor(tc string) Option {
	return func(c *Chart) {
		c.textColor = tc
	}
}
