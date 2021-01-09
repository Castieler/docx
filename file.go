package docx

import (
	"archive/zip"
	"encoding/xml"
	"io"
	"os"
	"strconv"

	"github.com/gingfrederik/docx/template"
)

type File struct {
	Document    Document
	DocRelation DocRelation

	rId int
}

func NewFile() *File {
	defaultRel := []*RelationShip{
		{
			ID:     "rId1",
			Type:   `http://schemas.openxmlformats.org/officeDocument/2006/relationships/customXml`,
			Target: "../customXml/item1.xml",
		},
		{
			ID:     "rId2",
			Type:   `http://schemas.openxmlformats.org/officeDocument/2006/relationships/numbering`,
			Target: "numbering.xml",
		},
		{
			ID:     "rId3",
			Type:   `http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles`,
			Target: "styles.xml",
		},
		{
			ID:         "rId4",
			Type:       "http://schemas.openxmlformats.org/officeDocument/2006/relationships/settings",
			Target:     "settings.xml",
		},
		{
			ID:         "rId5",
			Type:       "http://schemas.openxmlformats.org/officeDocument/2006/relationships/webSettings",
			Target:     "webSettings.xml",
		},

		{
			ID:         "rId6",
			Type:       "http://schemas.openxmlformats.org/officeDocument/2006/relationships/comments",
			Target:     "comments.xml",
		},
		{
			ID:         "rId7",
			Type:       "http://schemas.microsoft.com/office/2011/relationships/commentsExtended",
			Target:     "commentsExtended.xml",
		},
		{
			ID:         "rId8",
			Type:       "http://schemas.microsoft.com/office/2016/09/relationships/commentsIds",
			Target:     "commentsIds.xml",
		},
		{
			ID:         "rId9",
			Type:       "http://schemas.microsoft.com/office/2018/08/relationships/commentsExtensible",
			Target:     "commentsExtensible.xml",
		},
		{
			ID:         "rId10",
			Type:       "http://schemas.openxmlformats.org/officeDocument/2006/relationships/fontTable",
			Target:     "fontTable.xml",
		},
		{
			ID:         "rId11",
			Type:       "http://schemas.microsoft.com/office/2011/relationships/people",
			Target:     "people.xml",
		},
		{
			ID:         "rId12",
			Type:       "http://schemas.openxmlformats.org/officeDocument/2006/relationships/theme",
			Target:     "theme/theme1.xml",
		},
	}

	f := &File{
		Document: Document{
			XMLName: xml.Name{
				Space: "w",
			},
			XMLW: XMLNS_W,
			XMLR: XMLNS_R,
			Body: &Body{
				XMLName: xml.Name{
					Space: "w",
				},
				Paragraph: make([]*Paragraph, 0),
			},
		},
		DocRelation: DocRelation{
			Xmlns:        XMLNS,
			Relationship: defaultRel,
		},
		rId: 4,
	}

	return f
}

// Save save file to path
func (f *File) Save(path string) (err error) {
	fzip, _ := os.Create(path)
	defer fzip.Close()

	zipWriter := zip.NewWriter(fzip)
	defer zipWriter.Close()

	return f.pack(zipWriter)
}

func (f *File) Write(writer io.Writer) (err error) {
	zipWriter := zip.NewWriter(writer)
	defer zipWriter.Close()

	return f.pack(zipWriter)
}

// AddParagraph add new paragraph
func (f *File) AddParagraph() *Paragraph {
	p := &Paragraph{
		Data: make([]interface{}, 0),
		file: f,
	}

	f.Document.Body.Paragraph = append(f.Document.Body.Paragraph, p)
	return p
}

func (f *File) addLinkRelation(link string) string {
	rel := &RelationShip{
		ID:         "rId" + strconv.Itoa(f.rId),
		Type:       REL_HYPERLINK,
		Target:     link,
		TargetMode: REL_TARGETMODE,
	}

	f.rId += 1

	f.DocRelation.Relationship = append(f.DocRelation.Relationship, rel)

	return rel.ID
}

func (f *File) pack(zipWriter *zip.Writer) (err error) {
	files := map[string]string{}

	files["_rels/.rels"] = template.TEMP_REL_NEW
	files["docProps/app.xml"] = template.TEMP_DOCPROPS_APP_NEW
	files["docProps/core.xml"] = template.TEMP_DOCPROPS_CORE_NEW
	files["word/theme/theme1.xml"] = template.TEMP_WORD_THEME_THEME_NEW
	files["word/styles.xml"] = template.TEMP_WORD_STYLE_NEW
	files["[Content_Types].xml"] = template.TEMP_CONTENT_NEW
	files["word/_rels/document.xml.rels"], err = marshal(f.DocRelation)
	if err != nil {
		return err
	}
	//files["word/_rels/document.xml.rels"] = template.DOC_Relation_NEW
	
	//files["word/document.xml"], err = marshal(f.Document)
	//if err != nil {
	//	return err
	//}
	files["word/document.xml"] = template.WOED_DOCUMENT_NEW

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

func marshal(data interface{}) (out string, err error) {
	body, err := xml.Marshal(data)
	if err != nil {
		return
	}

	out = xml.Header + string(body)
	return
}
