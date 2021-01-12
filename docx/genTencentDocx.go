package docx

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"strings"

	"powerlaw.ai/platform/docx/docx/template"
)

type PrizePO struct {
	Name        string
	Description string
	Probability string
	Count       string
}

type EventsVO struct {
	ActivityRules      []ActivityRule
	Disqualification   string // 资格取消
	TaxDuty            string // 纳税义务
	ExemptionClause    string // 免责条款
	JurisdictionClause string // 管辖条款
	Other              string
}

type ActivityRule struct {
	Name                     string
	Time                     string
	ParticipantQualification string
	ParticipateWay           string
	WinningRules             string
	Prizes                   []PrizePO
	PrizeContent             string
	Announcement             string
	AwardDistribution        string
	SubRules                 []NameVO
}

type NameVO struct {
	Name    string
	Content string
}

// Save save file to path
func SaveTencentDocx(po EventsVO, path string) (err error) {
	f := NewDocx()
	fzip, _ := os.Create(path)
	defer fzip.Close()

	zipWriter := zip.NewWriter(fzip)
	defer zipWriter.Close()
	docStr := genDocumentStr(po)
	return f.packWithDocStr(zipWriter, docStr)
}

func WriteTencentDocx(po EventsVO, writer io.Writer) (err error) {
	f := NewDocx()
	zipWriter := zip.NewWriter(writer)
	defer zipWriter.Close()
	docStr := genDocumentStr(po)
	return f.packWithDocStr(zipWriter, docStr)
}

func genDocumentStr(po EventsVO) string {
	allActivates := po.ActivityRules
	var docStr, activityStr string
	var indexNum = 1
	// 一级标题数量：活动名称、活动时间、参与者资格、活动规则
	firstTitleNum := 4
	// 二级标题数量：参与方式、奖项设置、开奖设置、奖品发放
	secondTitleNum := 4
	for _, ar := range allActivates {
		var tabRowStyleStr, subRuleStr string
		// 构建奖品表中的数据样式字符串
		for _, prize := range ar.Prizes {
			tabRowStyleStr += fmt.Sprintf(template.TableRow,
				prize.Name, prize.Description, prize.Probability, prize.Count,
			)
		}

		// 构建活动中子标题样式字符串
		for index, name := range ar.SubRules {
			subRuleStr += fmt.Sprintf(
				template.SubRuleStyle, Num2ChineseStr(secondTitleNum+index+1),
				name.Name, handleStrNewLine(name.Content))
		}
		activityStr += fmt.Sprintf(template.ActivateRuleStyle,
			Num2ChineseStr(indexNum), handleStrNewLine(ar.Name), // 汉字序号，活动名称
			Num2ChineseStr(indexNum+1), handleStrNewLine(ar.Time), // 汉字序号，活动时间
			Num2ChineseStr(indexNum+2), handleStrNewLine(ar.ParticipantQualification), // 汉字序号，参与资格
			Num2ChineseStr(indexNum+3),             // 活动规则汉字序号
			handleStrNewLine(ar.ParticipateWay),    // 参与方式
			handleStrNewLine(ar.PrizeContent),      // 奖项设置
			tabRowStyleStr,                         // 奖品内容
			handleStrNewLine(ar.Announcement),      // 开奖设置
			handleStrNewLine(ar.AwardDistribution), // 奖品发放
			subRuleStr,                             // 自定义子标题
		)

		indexNum += firstTitleNum
	}
	docStr = fmt.Sprintf(template.DocxStyle,
		activityStr,                                                     // 所有活动样式字符串
		Num2ChineseStr(indexNum), handleStrNewLine(po.Disqualification), // 汉字序号,资格取消
		Num2ChineseStr(indexNum+1), handleStrNewLine(po.TaxDuty), // 汉字序号，纳税义务
		Num2ChineseStr(indexNum+2), handleStrNewLine(po.ExemptionClause), // 汉字序号，免责条款
		Num2ChineseStr(indexNum+3), handleStrNewLine(po.JurisdictionClause), // 汉字序号，管辖条款
		Num2ChineseStr(indexNum+4), handleStrNewLine(po.Other), // 汉字序号，其他
	)

	return docStr
}

// 处理字符串中的换行符
func handleStrNewLine(str string) string {
	strList := strings.Split(str, "\n")
	var styleStr string
	if len(strList) == 0 {
		return str
	}
	for _, s := range strList {
		if s == "" {
			continue
		}
		styleStr += fmt.Sprintf(template.NewLineStyle, s)
	}
	return styleStr
}
