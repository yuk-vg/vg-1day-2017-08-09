package bot

import (
	"context"

	"github.com/yuk-vg/vg-1day-2017-08-09/kaji/model"
)

type (
	// Poster はInに渡されたmessageをPOSTするための構造体です
	Poster struct {
		In chan *model.Message
	}
)

// Run はPosterを起動します
func (p *Poster) Run(ctx context.Context, url string) {
	for {
		select {
		case <-ctx.Done():
			close(p.In)
			return
		case m := <-p.In:
			postJSON(url+"/api/messages", m, nil)
		}
	}
}

// NewPoster は新しいPoster構造体のポインタを返します
func NewPoster(bufferSize int) *Poster {
	in := make(chan *model.Message, bufferSize)
	return &Poster{
		In: in,
	}
}
