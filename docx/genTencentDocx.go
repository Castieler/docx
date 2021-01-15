package docx

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"

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
	fzip, err := os.Create(path)
	if err != nil {
		return errors.WithStack(err)
	}
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
	for _, ar := range allActivates {
		// 1. 拼接 "活动名称", "活动时间", "参与者资格", "活动规则" 样式及内容
		firstTitleContent := []string{
			ar.Name,                     // 活动名称
			ar.Time,                     //  活动时间
			ar.ParticipantQualification, // 参与资格
			"",                          // 活动规则内容为空
		}
		for index, ft := range template.ActivityFirstTitleList {
			activityStr += fmt.Sprintf(template.FirstTitleStyle,
				Num2ChineseStr(indexNum+index), ft)
			firstContentStr := firstTitleContent[index]
			if firstContentStr != "" {
				activityStr += handleCharAndGenSpanStyle(firstContentStr)
			}
		}
		// 2. 拼接活动规则下的数据
		var tabRowStyleStr, subRuleStr string
		for _, prize := range ar.Prizes {
			tabRowStyleStr += fmt.Sprintf(template.TableRow,
				handleCharAndGenSpanStyle(prize.Name),
				handleCharAndGenSpanStyle(prize.Description),
				handleCharAndGenSpanStyle(prize.Probability),
				handleCharAndGenSpanStyle(prize.Count),
			)
		}

		secondContentList := []string{
			handleCharAndGenSpanStyle(ar.ParticipateWay),
			handleCharAndGenSpanStyle(ar.PrizeContent) + fmt.Sprintf(template.TableStyle, tabRowStyleStr),
			handleCharAndGenSpanStyle(ar.Announcement),      // 开奖设置
			handleCharAndGenSpanStyle(ar.AwardDistribution), // 奖品发放
		}
		secondTitleList := template.SecondTitleList
		for _, name := range ar.SubRules {
			secondTitleList = append(secondTitleList, name.Name)
			secondContentList = append(secondContentList,
				handleCharAndGenSpanStyle(name.Content))
		}

		for index, name := range secondTitleList {
			subRuleStr += fmt.Sprintf(
				template.SecondTitleRuleStyle, Num2ChineseStr(index+1), name)
			subRuleStr += secondContentList[index]
		}
		activityStr += subRuleStr
		indexNum += len(template.ActivityFirstTitleList)
	}

	// 3. 拼接 "资格取消", "纳税义务", "免责条款", "管辖条款", "其他" 样式及内容
	firstTitleContent := []string{
		po.Disqualification,
		po.TaxDuty,
		po.ExemptionClause,
		po.JurisdictionClause,
		po.Other,
	}
	for index, ft := range template.EndFirstTitleList {
		activityStr += fmt.Sprintf(template.FirstTitleStyle,
			Num2ChineseStr(indexNum+index), ft)
		firstContentStr := firstTitleContent[index]
		if firstContentStr != "" {
			activityStr += handleCharAndGenSpanStyle(firstContentStr)
		}
	}

	docStr = fmt.Sprintf(template.DocxStyle,
		activityStr, // 所有活动样式字符串
	)

	return docStr
}

// 处理字符串中的换行符 & > < ' " 并生成一个段落样式
func handleCharAndGenSpanStyle(str string) string {
	str = strings.ReplaceAll(str, "&", "&amp;")
	str = strings.ReplaceAll(str, "<", "&lt;")
	str = strings.ReplaceAll(str, ">", "&gt;")
	str = strings.ReplaceAll(str, "\"", "&quot;")
	str = strings.ReplaceAll(str, "'", "&apos;")
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
