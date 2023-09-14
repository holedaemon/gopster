package gopster

import (
	"errors"
	"fmt"
	"image"
)

// ErrorItem is returned when a chart item is misconfigured.
var ErrorItem = errors.New("gopster: chart item is misconfigured")

type chartItem struct {
	title      string
	creator    string
	coverImage image.Image
}

// AddItem adds an item to a Chart.
func (c *Chart) AddItem(title string, creator string, img image.Image) error {
	if c.items == nil {
		c.items = make([]*chartItem, 0)
	}

	if len(c.items) == c.width*c.height {
		return fmt.Errorf("%w: maximum number of items have been added", ErrorItem)
	}

	if title == "" || creator == "" || img == nil {
		return fmt.Errorf("%w: missing title/creator/image", ErrorItem)
	}

	item := &chartItem{
		title:      title,
		creator:    creator,
		coverImage: img,
	}

	c.items = append(c.items, item)

	return nil
}

// MustAddItem is the same as AddItem, except it will panic on error.
func (c *Chart) MustAddItem(title, creator string, img image.Image) {
	err := c.AddItem(title, creator, img)
	if err != nil {
		panic(err)
	}
}
