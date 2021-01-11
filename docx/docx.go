package docx

import (
	"archive/zip"
	"encoding/xml"
	"io"
	"os"
	"strconv"

	"powerlaw.ai/platform/docx/docx/paragraph"
	"powerlaw.ai/platform/docx/docx/relationship"
	run2 "powerlaw.ai/platform/docx/docx/run"
	"powerlaw.ai/platform/docx/docx/template"
)

const (
	XMLNS_W   = `http://schemas.openxmlformats.org/wordprocessingml/2006/main`
	XMLNS_R   = `http://schemas.openxmlformats.org/officeDocument/2006/relationships`
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
	Paragraph []*paragraph.Paragraph
}

type Docx struct {
	Document    Document
	DocRelation relationship.DocRelation

	rId int
}

func NewDocx() *Docx {
	f := &Docx{
		Document: Document{
			XMLName: xml.Name{
				Space: "w",
			},
			XMLW:   XMLNS_W,
			XMLR:   XMLNS_R,
			XMLW14: XMLNS_W14,
			Body: &Body{
				XMLName: xml.Name{
					Space: "w",
				},
				Paragraph: make([]*paragraph.Paragraph, 0),
			},
		},
		DocRelation: relationship.DocRelation{
			Xmlns:        relationship.XMLNS,
			Relationship: relationship.DefaultRel,
		},
		rId: 4,
	}

	return f
}

// Save save file to path
func (f *Docx) Save(path string) (err error) {
	fzip, _ := os.Create(path)
	defer fzip.Close()

	zipWriter := zip.NewWriter(fzip)
	defer zipWriter.Close()

	return f.pack(zipWriter)
}

func (f *Docx) Write(writer io.Writer) (err error) {
	zipWriter := zip.NewWriter(writer)
	defer zipWriter.Close()

	return f.pack(zipWriter)
}

// AddParagraph add new paragraph
func (f *Docx) AddParagraph() *paragraph.Paragraph {
	p := &paragraph.Paragraph{
		Data: make([]interface{}, 0),
	}

	f.Document.Body.Paragraph = append(f.Document.Body.Paragraph, p)
	return p
}

func (f *Docx) pack(zipWriter *zip.Writer) (err error) {
	docStr, err := marshal(f.Document)
	if err != nil {
		return err
	}

	return f.packWithDocStr(zipWriter, docStr)
}

func marshal(data interface{}) (out string, err error) {
	body, err := xml.Marshal(data)
	if err != nil {
		return
	}

	out = xml.Header + string(body)
	return
}

// AddParagraphLink add hyperlink to paragraph
func (f *Docx) AddParagraphLink(p *paragraph.Paragraph, text string, link string) *run2.Hyperlink {
	rId := f.addLinkRelation(link)
	hyperlink := &run2.Hyperlink{
		ID: rId,
		Run: run2.Run{
			RunProperties: &run2.RunProperties{
				RunStyle: &run2.RunStyle{
					Val: run2.HYPERLINK_STYLE,
				},
			},
			InstrText: text,
		},
	}

	p.Data = append(p.Data, hyperlink)
	return hyperlink
}

func (f *Docx) addLinkRelation(link string) string {
	rel := &relationship.RelationShip{
		ID:         "rId" + strconv.Itoa(f.rId),
		Type:       relationship.REL_HYPERLINK,
		Target:     link,
		TargetMode: relationship.REL_TARGETMODE,
	}

	f.rId += 1

	f.DocRelation.Relationship = append(f.DocRelation.Relationship, rel)

	return rel.ID
}

func (f *Docx) packWithDocStr(zipWriter *zip.Writer,
	documentStr string) (err error) {
	files := map[string]string{}
	files["_rels/.rels"] = template.Rel

	files["customXml/_rels/item1.xml.rels"] =
		template.CustomXmlRelsItem1XmlRels
	files["customXml/item1.xml"] =
		template.CustomXmlItem1
	files["customXml/itemProps1.xml"] =
		template.CustomXmlItemProps1

	files["docProps/app.xml"] = template.DocPropsApp
	files["docProps/core.xml"] = template.DocPropsCore

	files["word/_rels/document.xml.rels"], err = marshal(f.DocRelation)
	if err != nil {
		return err
	}
	files["word/theme/theme1.xml"] = template.WordThemeTheme
	files["word/webextensions/taskpanes.xml"] = template.WordWebextensionsTaskpanes
	files["word/webextensions/webextension1.xml"] = template.WordWebextensionsWebextension1
	files["word/webextensions/_rels/taskpanes.xml.rels"] = template.WordWebextensionsRelsTaskpanesXmlRels

	files["word/commentsExtended.xml"] = template.WordCommentsExtended
	files["word/commentsExtensible.xml"] = template.WordCommentsExtensible
	files["word/commentsIds.xml"] = template.WordCommentIds
	files["word/styles.xml"] = template.WordStyle
	files["word/numbering.xml"] = template.WordNumbering
	files["word/fontTable.xml"] = template.WordFontTableNew
	files["word/webSettings.xml"] = template.WordWebSettingsNew

	files["[Content_Types].xml"] = template.ContentTypes

	files["word/document.xml"] = documentStr

	for path, data := range files {
		w, err := zipWriter.Create(path)
		if err != nil {
			return err
		}

		_, err = w.Write([]byte(data))
		if err != nil {
			return err
		}
	}

	return
}
