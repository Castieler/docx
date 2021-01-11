package main

import (
	docx2 "powerlaw.ai/platforn/docx/docx"
)

func main() {
	f := docx2.NewDocx()
	// add new paragraph
	para := f.AddParagraph()
	// add text
	para.AddText("test")

	para.AddText("test font size").Size(22)
	para.AddText("test color").Color("808080")
	para.AddText("test font size and color").Size(22).Color("121212")

	nextPara := f.AddParagraph()
	f.AddParagraphLink(nextPara, "google", `http://google.com`)

	f.Save("./test.docx")
}
