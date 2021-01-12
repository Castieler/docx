package template

const (
	TableStyle = `        <w:tbl>
            <w:tblPr>
                <w:tblStyle w:val="a3"/>
                <w:tblW w:w="5000" w:type="pct"/>
                <w:tblLook w:val="04A0" w:firstRow="1" w:lastRow="0" w:firstColumn="1" w:lastColumn="0" w:noHBand="0" w:noVBand="1"/>
            </w:tblPr>
            <w:tblGrid>
                <w:gridCol w:w="1830"/>
                <w:gridCol w:w="3288"/>
                <w:gridCol w:w="2069"/>
                <w:gridCol w:w="1103"/>
            </w:tblGrid>
            <w:tr w:rsidR="0055086D" w:rsidRPr="00743371" w14:paraId="0E6E14AA" w14:textId="77777777" w:rsidTr="0055086D">
                <w:tc>
                    <w:tcPr>
                        <w:tcW w:w="1104" w:type="pct"/>
                        <w:shd w:val="clear" w:color="auto" w:fill="D9D9D9" w:themeFill="background1" w:themeFillShade="D9"/>
                    </w:tcPr>
                    <w:p w14:paraId="6A75AB5D" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00BF7A43" w:rsidP="00353DA9">
                        <w:pPr>
                            <w:spacing w:after="0" w:line="240" w:lineRule="auto"/>
                        </w:pPr>
                        <w:r w:rsidRPr="0029777C">
                            <w:rPr>
                                <w:rFonts w:hint="eastAsia"/>
                            </w:rPr>
                            <w:t>奖项名称</w:t>
                        </w:r>
                    </w:p>
                </w:tc>
                <w:tc>
                    <w:tcPr>
                        <w:tcW w:w="1983" w:type="pct"/>
                        <w:shd w:val="clear" w:color="auto" w:fill="D9D9D9" w:themeFill="background1" w:themeFillShade="D9"/>
                    </w:tcPr>
                    <w:p w14:paraId="3F3C0566" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00BF7A43" w:rsidP="00353DA9">
                        <w:pPr>
                            <w:spacing w:after="0" w:line="240" w:lineRule="auto"/>
                        </w:pPr>
                        <w:r w:rsidRPr="0029777C">
                            <w:rPr>
                                <w:rFonts w:hint="eastAsia"/>
                            </w:rPr>
                            <w:t>奖品描述</w:t>
                        </w:r>
                    </w:p>
                </w:tc>
                <w:tc>
                    <w:tcPr>
                        <w:tcW w:w="1248" w:type="pct"/>
                        <w:shd w:val="clear" w:color="auto" w:fill="D9D9D9" w:themeFill="background1" w:themeFillShade="D9"/>
                    </w:tcPr>
                    <w:p w14:paraId="7CA7E222" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00BF7A43" w:rsidP="00353DA9">
                        <w:pPr>
                            <w:spacing w:after="0" w:line="240" w:lineRule="auto"/>
                        </w:pPr>
                        <w:r w:rsidRPr="0029777C">
                            <w:rPr>
                                <w:rFonts w:hint="eastAsia"/>
                            </w:rPr>
                            <w:t>中奖概率</w:t>
                        </w:r>
                    </w:p>
                </w:tc>
                <w:tc>
                    <w:tcPr>
                        <w:tcW w:w="665" w:type="pct"/>
                        <w:shd w:val="clear" w:color="auto" w:fill="D9D9D9" w:themeFill="background1" w:themeFillShade="D9"/>
                    </w:tcPr>
                    <w:p w14:paraId="4F9B2883" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00BF7A43" w:rsidP="00353DA9">
                        <w:pPr>
                            <w:spacing w:after="0" w:line="240" w:lineRule="auto"/>
                        </w:pPr>
                        <w:r w:rsidRPr="0029777C">
                            <w:rPr>
                                <w:rFonts w:hint="eastAsia"/>
                            </w:rPr>
                            <w:t>数量</w:t>
                        </w:r>
                    </w:p>
                </w:tc>
            </w:tr>
			%s
        </w:tbl>`

	TableRow = `<w:tr w:rsidR="0029777C" w:rsidRPr="00743371" w14:paraId="525E0653" w14:textId="77777777" w:rsidTr="0055086D">
                <w:tc>
                    <w:tcPr>
                        <w:tcW w:w="1104" w:type="pct"/>
                    </w:tcPr>
                    <w:p w14:paraId="301B86E4" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00BF7A43" w:rsidP="00353DA9">
                        <w:pPr>
                            <w:spacing w:after="0" w:line="240" w:lineRule="auto"/>
                        </w:pPr>
                        <w:r w:rsidRPr="0029777C">
                            <w:rPr>
                                <w:rFonts w:hint="eastAsia"/>
                            </w:rPr>
                            <w:t>%s</w:t>
                        </w:r>
                    </w:p>
                </w:tc>
                <w:tc>
                    <w:tcPr>
                        <w:tcW w:w="1983" w:type="pct"/>
                    </w:tcPr>
                    <w:p w14:paraId="2B129718" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00BF7A43" w:rsidP="00353DA9">
                        <w:pPr>
                            <w:spacing w:after="0" w:line="240" w:lineRule="auto"/>
                        </w:pPr>
                        <w:r w:rsidRPr="0029777C">
                            <w:rPr>
                                <w:rFonts w:hint="eastAsia"/>
                            </w:rPr>
                            <w:t>%s</w:t>
                        </w:r>
                    </w:p>
                </w:tc>
                <w:tc>
                    <w:tcPr>
                        <w:tcW w:w="1248" w:type="pct"/>
                    </w:tcPr>
                    <w:p w14:paraId="033B4B64" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00BF7A43" w:rsidP="00353DA9">
                        <w:pPr>
                            <w:spacing w:after="0" w:line="240" w:lineRule="auto"/>
                        </w:pPr>
                        <w:r w:rsidRPr="0029777C">
                            <w:t>%s</w:t>
                        </w:r>
                    </w:p>
                </w:tc>
                <w:tc>
                    <w:tcPr>
                        <w:tcW w:w="665" w:type="pct"/>
                    </w:tcPr>
                    <w:p w14:paraId="14C7DC60" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00BF7A43" w:rsidP="00353DA9">
                        <w:pPr>
                            <w:spacing w:after="0" w:line="240" w:lineRule="auto"/>
                        </w:pPr>
                        <w:r w:rsidRPr="0029777C">
                            <w:t>%s</w:t>
                        </w:r>
                    </w:p>
                </w:tc>
            </w:tr>`

	DocxStyle = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:document
        xmlns:wpc="http://schemas.microsoft.com/office/word/2010/wordprocessingCanvas"
        xmlns:cx="http://schemas.microsoft.com/office/drawing/2014/chartex"
        xmlns:cx1="http://schemas.microsoft.com/office/drawing/2015/9/8/chartex"
        xmlns:cx2="http://schemas.microsoft.com/office/drawing/2015/10/21/chartex"
        xmlns:cx3="http://schemas.microsoft.com/office/drawing/2016/5/9/chartex"
        xmlns:cx4="http://schemas.microsoft.com/office/drawing/2016/5/10/chartex"
        xmlns:cx5="http://schemas.microsoft.com/office/drawing/2016/5/11/chartex"
        xmlns:cx6="http://schemas.microsoft.com/office/drawing/2016/5/12/chartex"
        xmlns:cx7="http://schemas.microsoft.com/office/drawing/2016/5/13/chartex"
        xmlns:cx8="http://schemas.microsoft.com/office/drawing/2016/5/14/chartex"
        xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006"
        xmlns:aink="http://schemas.microsoft.com/office/drawing/2016/ink"
        xmlns:am3d="http://schemas.microsoft.com/office/drawing/2017/model3d"
        xmlns:o="urn:schemas-microsoft-com:office:office"
        xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships"
        xmlns:m="http://schemas.openxmlformats.org/officeDocument/2006/math"
        xmlns:v="urn:schemas-microsoft-com:vml"
        xmlns:wp14="http://schemas.microsoft.com/office/word/2010/wordprocessingDrawing"
        xmlns:wp="http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"
        xmlns:w10="urn:schemas-microsoft-com:office:word"
        xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"
        xmlns:w14="http://schemas.microsoft.com/office/word/2010/wordml"
        xmlns:w15="http://schemas.microsoft.com/office/word/2012/wordml"
        xmlns:w16cex="http://schemas.microsoft.com/office/word/2018/wordml/cex"
        xmlns:w16cid="http://schemas.microsoft.com/office/word/2016/wordml/cid"
        xmlns:w16="http://schemas.microsoft.com/office/word/2018/wordml"
        xmlns:w16se="http://schemas.microsoft.com/office/word/2015/wordml/symex"
        xmlns:wpg="http://schemas.microsoft.com/office/word/2010/wordprocessingGroup"
        xmlns:wpi="http://schemas.microsoft.com/office/word/2010/wordprocessingInk"
        xmlns:wne="http://schemas.microsoft.com/office/word/2006/wordml"
        xmlns:wps="http://schemas.microsoft.com/office/word/2010/wordprocessingShape" mc:Ignorable="w14 w15 w16se w16cid w16 w16cex wp14">
    <w:body>
		%s
    </w:body>
</w:document>`

	FirstTitleStyle = `<w:p w14:paraId="5ED56D03" w14:textId="7981F797" w:rsidR="00425149" w:rsidRPr="00743371" w:rsidRDefault="00C736C8" w:rsidP="00C736C8">
            <w:pPr>
                <w:pStyle w:val="1"/>
                <w:keepNext w:val="0"/>
                <w:keepLines w:val="0"/>
                <w:spacing w:line="240" w:lineRule="auto"/>
            </w:pPr>
            <w:r>
                <w:rPr>
                    <w:rFonts w:hint="eastAsia"/>
                </w:rPr>
                <w:t>%s、</w:t>
            </w:r>
            <w:r w:rsidRPr="00743371">
                <w:rPr>
                    <w:rFonts w:hint="eastAsia"/>
                </w:rPr>
                <w:t>%s</w:t>
            </w:r>
        </w:p>`

	SecondTitleRuleStyle = `<w:p w14:paraId="411045CE" w14:textId="17711249" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00C736C8" w:rsidP="00C736C8">
            <w:pPr>
                <w:pStyle w:val="2"/>
                <w:keepNext w:val="0"/>
                <w:keepLines w:val="0"/>
                <w:spacing w:line="240" w:lineRule="auto"/>
            </w:pPr>
            <w:r>
                <w:rPr>
                    <w:rFonts w:hint="eastAsia"/>
                </w:rPr>
                <w:t>（%s）</w:t>
            </w:r>
            <w:r w:rsidRPr="0029777C">
                <w:t>%s</w:t>
            </w:r>
        </w:p>`

	NewLineStyle = `<w:p w14:paraId="365F210D" w14:textId="77777777" w:rsidR="00E3674D" w:rsidRDefault="00C736C8" w:rsidP="00C736C8">
            <w:pPr>
                <w:spacing w:line="240" w:lineRule="auto"/>
            </w:pPr>
            <w:r>
                <w:rPr>
                    <w:rFonts w:hint="eastAsia"/>
                </w:rPr>
                <w:t>%s</w:t>
            </w:r>
        </w:p>`
)

var (
	ActivityFirstTitleList = []string{"活动名称", "活动时间", "参与者资格", "活动规则"}

	SecondTitleList = []string{"参与方式", "奖项设置", "开奖设置", "奖品发放"}

	EndFirstTitleList = []string{"资格取消", "纳税义务", "免责条款", "管辖条款", "其他"}
)
