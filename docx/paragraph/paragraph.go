package paragraph

import (
	"encoding/xml"

	run2 "powerlaw.ai/platforn/docx/docx/run"
)

type Paragraph struct {
	XMLName xml.Name `xml:"w:p"`
	Data    []interface{}
}

// AddText add text to paragraph
func (p *Paragraph) AddText(text string) *run2.Run {
	t := &run2.Text{
		Text: text,
	}

	run := &run2.Run{
		Text:          t,
		RunProperties: &run2.RunProperties{},
	}

	p.Data = append(p.Data, run)

	return run
}
