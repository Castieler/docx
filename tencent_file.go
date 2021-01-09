package docx

import (
	"archive/zip"
	"io"
	"os"

	"github.com/gingfrederik/docx/template"
)

type PrizePO struct {
	Name        string
	Description string
	Probability string
	Count       string
}

type TencentPO struct {
	ActivityName             string
	ActivityTime             string
	ParticipantQualification string
	ParticipateWay           string
	WinningRules             string
	Prizes                   []PrizePO
}

// Save save file to path
func SaveTencentDocx(po TencentPO, path string) (err error) {
	f := NewFile()
	fzip, _ := os.Create(path)
	defer fzip.Close()

	zipWriter := zip.NewWriter(fzip)
	defer zipWriter.Close()
	docStr := genDocumentStr(po)
	return f.packWithDocStr(zipWriter, docStr)
}

func WriteTencentDocx(po TencentPO, writer io.Writer) (err error) {
	f := NewFile()
	zipWriter := zip.NewWriter(writer)
	defer zipWriter.Close()
	docStr := genDocumentStr(po)
	return f.packWithDocStr(zipWriter, docStr)
}

func genDocumentStr(po TencentPO) string {
	return template.WordDocumentDemo + po.ActivityName + template.WordDocDemo2 +
		po.ActivityTime + template.WordDocDemo3 +
		po.ParticipateWay + template.WordDocDemo5 + po.WinningRules + template.WordDocDemo6
}
