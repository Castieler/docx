package docx

import "encoding/xml"

const (
	XMLNS_W = `http://schemas.openxmlformats.org/wordprocessingml/2006/main`
	XMLNS_R = `http://schemas.openxmlformats.org/officeDocument/2006/relationships`
	XMLNS_W14 = `http://schemas.microsoft.com/office/word/2010/wordml`
)

type Document struct {
	XMLName xml.Name `xml:"w:document"`
	XMLW    string   `xml:"xmlns:w,attr"`
	XMLR    string   `xml:"xmlns:r,attr"`
	XMLW14  string   `xml:"xmlns:w14,attr"`
	Body    *Body
}

type Body struct {
	XMLName   xml.Name `xml:"w:body"`
	Paragraph []*Paragraph
}
