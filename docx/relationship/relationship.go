package relationship

import "encoding/xml"

const (
	XMLNS         = `http://schemas.openxmlformats.org/package/2006/relationships`
	REL_HYPERLINK = `http://schemas.openxmlformats.org/officeDocument/2006/relationships/hyperlink`

	REL_TARGETMODE = "External"
)

var DefaultRel = []*RelationShip{
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
		ID:     "rId4",
		Type:   "http://schemas.openxmlformats.org/officeDocument/2006/relationships/settings",
		Target: "settings.xml",
	},
	{
		ID:     "rId5",
		Type:   "http://schemas.openxmlformats.org/officeDocument/2006/relationships/webSettings",
		Target: "webSettings.xml",
	},

	{
		ID:     "rId6",
		Type:   "http://schemas.openxmlformats.org/officeDocument/2006/relationships/comments",
		Target: "comments.xml",
	},
	{
		ID:     "rId7",
		Type:   "http://schemas.microsoft.com/office/2011/relationships/commentsExtended",
		Target: "commentsExtended.xml",
	},
	{
		ID:     "rId8",
		Type:   "http://schemas.microsoft.com/office/2016/09/relationships/commentsIds",
		Target: "commentsIds.xml",
	},
	{
		ID:     "rId9",
		Type:   "http://schemas.microsoft.com/office/2018/08/relationships/commentsExtensible",
		Target: "commentsExtensible.xml",
	},
	{
		ID:     "rId10",
		Type:   "http://schemas.openxmlformats.org/officeDocument/2006/relationships/fontTable",
		Target: "fontTable.xml",
	},
	{
		ID:     "rId11",
		Type:   "http://schemas.microsoft.com/office/2011/relationships/people",
		Target: "people.xml",
	},
	{
		ID:     "rId12",
		Type:   "http://schemas.openxmlformats.org/officeDocument/2006/relationships/theme",
		Target: "theme/theme1.xml",
	},
}

type DocRelation struct {
	XMLName      xml.Name        `xml:"Relationships"`
	Xmlns        string          `xml:"xmlns,attr"`
	Relationship []*RelationShip `xml:"Relationship"`
}

type RelationShip struct {
	XMLName    xml.Name `xml:"Relationship"`
	ID         string   `xml:"Id,attr"`
	Type       string   `xml:"Type,attr"`
	Target     string   `xml:"Target,attr"`
	TargetMode string   `xml:"TargetMode,attr,omitempty"`
}
