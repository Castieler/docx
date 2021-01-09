package template

const (
	TEMP_REL_NEW = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
	<Relationships
		xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
		<Relationship Id="rId3" Type="http://schemas.openxmlformats.org/package/2006/relationships/metadata/core-properties" Target="docProps/core.xml"/>
		<Relationship Id="rId2" Type="http://schemas.microsoft.com/office/2011/relationships/webextensiontaskpanes" Target="word/webextensions/taskpanes.xml"/>
		<Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument" Target="word/document.xml"/>
		<Relationship Id="rId4" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/extended-properties" Target="docProps/app.xml"/>
	</Relationships>`
	TEMP_DOCPROPS_APP_NEW = `
		<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
		<Properties
				xmlns="http://schemas.openxmlformats.org/officeDocument/2006/extended-properties"
				xmlns:vt="http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes">
			<Template>Normal.dotm</Template>
			<TotalTime>27</TotalTime>
			<Pages>2</Pages>
			<Words>134</Words>
			<Characters>765</Characters>
			<Application>Microsoft Office Word</Application>
			<DocSecurity>0</DocSecurity>
			<Lines>6</Lines>
			<Paragraphs>1</Paragraphs>
			<ScaleCrop>false</ScaleCrop>
			<Company></Company>
			<LinksUpToDate>false</LinksUpToDate>
			<CharactersWithSpaces>898</CharactersWithSpaces>
			<SharedDoc>false</SharedDoc>
			<HyperlinksChanged>false</HyperlinksChanged>
			<AppVersion>16.0000</AppVersion>
		</Properties>
`
	TEMP_DOCPROPS_CORE_NEW = `
		<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
		<cp:coreProperties
			xmlns:cp="http://schemas.openxmlformats.org/package/2006/metadata/core-properties"
			xmlns:dc="http://purl.org/dc/elements/1.1/"
			xmlns:dcterms="http://purl.org/dc/terms/"
			xmlns:dcmitype="http://purl.org/dc/dcmitype/"
			xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
			<dc:title></dc:title>
			<dc:subject></dc:subject>
			<dc:creator>唐 穆</dc:creator>
			<cp:keywords></cp:keywords>
			<dc:description></dc:description>
			<cp:lastModifiedBy>唐 穆</cp:lastModifiedBy>
			<cp:revision>7</cp:revision>
			<dcterms:created xsi:type="dcterms:W3CDTF">2020-10-27T10:11:00Z</dcterms:created>
			<dcterms:modified xsi:type="dcterms:W3CDTF">2020-10-27T10:42:00Z</dcterms:modified>
		</cp:coreProperties>
`


	TEMP_CONTENT_NEW = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Types
        xmlns="http://schemas.openxmlformats.org/package/2006/content-types">
    <Default Extension="rels" ContentType="application/vnd.openxmlformats-package.relationships+xml"/>
    <Default Extension="xml" ContentType="application/xml"/>
    <Override PartName="/word/document.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml"/>
    <Override PartName="/customXml/itemProps1.xml" ContentType="application/vnd.openxmlformats-officedocument.customXmlProperties+xml"/>
    <Override PartName="/word/numbering.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.numbering+xml"/>
    <Override PartName="/word/styles.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.styles+xml"/>
    <Override PartName="/word/settings.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.settings+xml"/>
    <Override PartName="/word/webSettings.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.webSettings+xml"/>
    <Override PartName="/word/comments.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.comments+xml"/>
    <Override PartName="/word/commentsExtended.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.commentsExtended+xml"/>
    <Override PartName="/word/commentsIds.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.commentsIds+xml"/>
    <Override PartName="/word/commentsExtensible.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.commentsExtensible+xml"/>
    <Override PartName="/word/fontTable.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.fontTable+xml"/>
    <Override PartName="/word/people.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.people+xml"/>
    <Override PartName="/word/webextensions/taskpanes.xml" ContentType="application/vnd.ms-office.webextensiontaskpanes+xml"/>
    <Override PartName="/word/webextensions/webextension1.xml" ContentType="application/vnd.ms-office.webextension+xml"/>
    <Override PartName="/word/theme/theme1.xml" ContentType="application/vnd.openxmlformats-officedocument.theme+xml"/>
    <Override PartName="/docProps/core.xml" ContentType="application/vnd.openxmlformats-package.core-properties+xml"/>
    <Override PartName="/docProps/app.xml" ContentType="application/vnd.openxmlformats-officedocument.extended-properties+xml"/>
</Types>`

	TEMP_WORD_THEME_THEME_NEW = `
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<a:theme
        xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" name="Office 主题​​">
    <a:themeElements>
        <a:clrScheme name="Office">
            <a:dk1>
                <a:sysClr val="windowText" lastClr="000000"/>
            </a:dk1>
            <a:lt1>
                <a:sysClr val="window" lastClr="FFFFFF"/>
            </a:lt1>
            <a:dk2>
                <a:srgbClr val="44546A"/>
            </a:dk2>
            <a:lt2>
                <a:srgbClr val="E7E6E6"/>
            </a:lt2>
            <a:accent1>
                <a:srgbClr val="4472C4"/>
            </a:accent1>
            <a:accent2>
                <a:srgbClr val="ED7D31"/>
            </a:accent2>
            <a:accent3>
                <a:srgbClr val="A5A5A5"/>
            </a:accent3>
            <a:accent4>
                <a:srgbClr val="FFC000"/>
            </a:accent4>
            <a:accent5>
                <a:srgbClr val="5B9BD5"/>
            </a:accent5>
            <a:accent6>
                <a:srgbClr val="70AD47"/>
            </a:accent6>
            <a:hlink>
                <a:srgbClr val="0563C1"/>
            </a:hlink>
            <a:folHlink>
                <a:srgbClr val="954F72"/>
            </a:folHlink>
        </a:clrScheme>
        <a:fontScheme name="Office">
            <a:majorFont>
                <a:latin typeface="等线 Light" panose="020F0302020204030204"/>
                <a:ea typeface=""/>
                <a:cs typeface=""/>
                <a:font script="Jpan" typeface="游ゴシック Light"/>
                <a:font script="Hang" typeface="맑은 고딕"/>
                <a:font script="Hans" typeface="等线 Light"/>
                <a:font script="Hant" typeface="新細明體"/>
                <a:font script="Arab" typeface="Times New Roman"/>
                <a:font script="Hebr" typeface="Times New Roman"/>
                <a:font script="Thai" typeface="Angsana New"/>
                <a:font script="Ethi" typeface="Nyala"/>
                <a:font script="Beng" typeface="Vrinda"/>
                <a:font script="Gujr" typeface="Shruti"/>
                <a:font script="Khmr" typeface="MoolBoran"/>
                <a:font script="Knda" typeface="Tunga"/>
                <a:font script="Guru" typeface="Raavi"/>
                <a:font script="Cans" typeface="Euphemia"/>
                <a:font script="Cher" typeface="Plantagenet Cherokee"/>
                <a:font script="Yiii" typeface="Microsoft Yi Baiti"/>
                <a:font script="Tibt" typeface="Microsoft Himalaya"/>
                <a:font script="Thaa" typeface="MV Boli"/>
                <a:font script="Deva" typeface="Mangal"/>
                <a:font script="Telu" typeface="Gautami"/>
                <a:font script="Taml" typeface="Latha"/>
                <a:font script="Syrc" typeface="Estrangelo Edessa"/>
                <a:font script="Orya" typeface="Kalinga"/>
                <a:font script="Mlym" typeface="Kartika"/>
                <a:font script="Laoo" typeface="DokChampa"/>
                <a:font script="Sinh" typeface="Iskoola Pota"/>
                <a:font script="Mong" typeface="Mongolian Baiti"/>
                <a:font script="Viet" typeface="Times New Roman"/>
                <a:font script="Uigh" typeface="Microsoft Uighur"/>
                <a:font script="Geor" typeface="Sylfaen"/>
                <a:font script="Armn" typeface="Arial"/>
                <a:font script="Bugi" typeface="Leelawadee UI"/>
                <a:font script="Bopo" typeface="Microsoft JhengHei"/>
                <a:font script="Java" typeface="Javanese Text"/>
                <a:font script="Lisu" typeface="Segoe UI"/>
                <a:font script="Mymr" typeface="Myanmar Text"/>
                <a:font script="Nkoo" typeface="Ebrima"/>
                <a:font script="Olck" typeface="Nirmala UI"/>
                <a:font script="Osma" typeface="Ebrima"/>
                <a:font script="Phag" typeface="Phagspa"/>
                <a:font script="Syrn" typeface="Estrangelo Edessa"/>
                <a:font script="Syrj" typeface="Estrangelo Edessa"/>
                <a:font script="Syre" typeface="Estrangelo Edessa"/>
                <a:font script="Sora" typeface="Nirmala UI"/>
                <a:font script="Tale" typeface="Microsoft Tai Le"/>
                <a:font script="Talu" typeface="Microsoft New Tai Lue"/>
                <a:font script="Tfng" typeface="Ebrima"/>
            </a:majorFont>
            <a:minorFont>
                <a:latin typeface="等线" panose="020F0502020204030204"/>
                <a:ea typeface=""/>
                <a:cs typeface=""/>
                <a:font script="Jpan" typeface="游明朝"/>
                <a:font script="Hang" typeface="맑은 고딕"/>
                <a:font script="Hans" typeface="等线"/>
                <a:font script="Hant" typeface="新細明體"/>
                <a:font script="Arab" typeface="Arial"/>
                <a:font script="Hebr" typeface="Arial"/>
                <a:font script="Thai" typeface="Cordia New"/>
                <a:font script="Ethi" typeface="Nyala"/>
                <a:font script="Beng" typeface="Vrinda"/>
                <a:font script="Gujr" typeface="Shruti"/>
                <a:font script="Khmr" typeface="DaunPenh"/>
                <a:font script="Knda" typeface="Tunga"/>
                <a:font script="Guru" typeface="Raavi"/>
                <a:font script="Cans" typeface="Euphemia"/>
                <a:font script="Cher" typeface="Plantagenet Cherokee"/>
                <a:font script="Yiii" typeface="Microsoft Yi Baiti"/>
                <a:font script="Tibt" typeface="Microsoft Himalaya"/>
                <a:font script="Thaa" typeface="MV Boli"/>
                <a:font script="Deva" typeface="Mangal"/>
                <a:font script="Telu" typeface="Gautami"/>
                <a:font script="Taml" typeface="Latha"/>
                <a:font script="Syrc" typeface="Estrangelo Edessa"/>
                <a:font script="Orya" typeface="Kalinga"/>
                <a:font script="Mlym" typeface="Kartika"/>
                <a:font script="Laoo" typeface="DokChampa"/>
                <a:font script="Sinh" typeface="Iskoola Pota"/>
                <a:font script="Mong" typeface="Mongolian Baiti"/>
                <a:font script="Viet" typeface="Arial"/>
                <a:font script="Uigh" typeface="Microsoft Uighur"/>
                <a:font script="Geor" typeface="Sylfaen"/>
                <a:font script="Armn" typeface="Arial"/>
                <a:font script="Bugi" typeface="Leelawadee UI"/>
                <a:font script="Bopo" typeface="Microsoft JhengHei"/>
                <a:font script="Java" typeface="Javanese Text"/>
                <a:font script="Lisu" typeface="Segoe UI"/>
                <a:font script="Mymr" typeface="Myanmar Text"/>
                <a:font script="Nkoo" typeface="Ebrima"/>
                <a:font script="Olck" typeface="Nirmala UI"/>
                <a:font script="Osma" typeface="Ebrima"/>
                <a:font script="Phag" typeface="Phagspa"/>
                <a:font script="Syrn" typeface="Estrangelo Edessa"/>
                <a:font script="Syrj" typeface="Estrangelo Edessa"/>
                <a:font script="Syre" typeface="Estrangelo Edessa"/>
                <a:font script="Sora" typeface="Nirmala UI"/>
                <a:font script="Tale" typeface="Microsoft Tai Le"/>
                <a:font script="Talu" typeface="Microsoft New Tai Lue"/>
                <a:font script="Tfng" typeface="Ebrima"/>
            </a:minorFont>
        </a:fontScheme>
        <a:fmtScheme name="Office">
            <a:fillStyleLst>
                <a:solidFill>
                    <a:schemeClr val="phClr"/>
                </a:solidFill>
                <a:gradFill rotWithShape="1">
                    <a:gsLst>
                        <a:gs pos="0">
                            <a:schemeClr val="phClr">
                                <a:lumMod val="110000"/>
                                <a:satMod val="105000"/>
                                <a:tint val="67000"/>
                            </a:schemeClr>
                        </a:gs>
                        <a:gs pos="50000">
                            <a:schemeClr val="phClr">
                                <a:lumMod val="105000"/>
                                <a:satMod val="103000"/>
                                <a:tint val="73000"/>
                            </a:schemeClr>
                        </a:gs>
                        <a:gs pos="100000">
                            <a:schemeClr val="phClr">
                                <a:lumMod val="105000"/>
                                <a:satMod val="109000"/>
                                <a:tint val="81000"/>
                            </a:schemeClr>
                        </a:gs>
                    </a:gsLst>
                    <a:lin ang="5400000" scaled="0"/>
                </a:gradFill>
                <a:gradFill rotWithShape="1">
                    <a:gsLst>
                        <a:gs pos="0">
                            <a:schemeClr val="phClr">
                                <a:satMod val="103000"/>
                                <a:lumMod val="102000"/>
                                <a:tint val="94000"/>
                            </a:schemeClr>
                        </a:gs>
                        <a:gs pos="50000">
                            <a:schemeClr val="phClr">
                                <a:satMod val="110000"/>
                                <a:lumMod val="100000"/>
                                <a:shade val="100000"/>
                            </a:schemeClr>
                        </a:gs>
                        <a:gs pos="100000">
                            <a:schemeClr val="phClr">
                                <a:lumMod val="99000"/>
                                <a:satMod val="120000"/>
                                <a:shade val="78000"/>
                            </a:schemeClr>
                        </a:gs>
                    </a:gsLst>
                    <a:lin ang="5400000" scaled="0"/>
                </a:gradFill>
            </a:fillStyleLst>
            <a:lnStyleLst>
                <a:ln w="6350" cap="flat" cmpd="sng" algn="ctr">
                    <a:solidFill>
                        <a:schemeClr val="phClr"/>
                    </a:solidFill>
                    <a:prstDash val="solid"/>
                    <a:miter lim="800000"/>
                </a:ln>
                <a:ln w="12700" cap="flat" cmpd="sng" algn="ctr">
                    <a:solidFill>
                        <a:schemeClr val="phClr"/>
                    </a:solidFill>
                    <a:prstDash val="solid"/>
                    <a:miter lim="800000"/>
                </a:ln>
                <a:ln w="19050" cap="flat" cmpd="sng" algn="ctr">
                    <a:solidFill>
                        <a:schemeClr val="phClr"/>
                    </a:solidFill>
                    <a:prstDash val="solid"/>
                    <a:miter lim="800000"/>
                </a:ln>
            </a:lnStyleLst>
            <a:effectStyleLst>
                <a:effectStyle>
                    <a:effectLst/>
                </a:effectStyle>
                <a:effectStyle>
                    <a:effectLst/>
                </a:effectStyle>
                <a:effectStyle>
                    <a:effectLst>
                        <a:outerShdw blurRad="57150" dist="19050" dir="5400000" algn="ctr" rotWithShape="0">
                            <a:srgbClr val="000000">
                                <a:alpha val="63000"/>
                            </a:srgbClr>
                        </a:outerShdw>
                    </a:effectLst>
                </a:effectStyle>
            </a:effectStyleLst>
            <a:bgFillStyleLst>
                <a:solidFill>
                    <a:schemeClr val="phClr"/>
                </a:solidFill>
                <a:solidFill>
                    <a:schemeClr val="phClr">
                        <a:tint val="95000"/>
                        <a:satMod val="170000"/>
                    </a:schemeClr>
                </a:solidFill>
                <a:gradFill rotWithShape="1">
                    <a:gsLst>
                        <a:gs pos="0">
                            <a:schemeClr val="phClr">
                                <a:tint val="93000"/>
                                <a:satMod val="150000"/>
                                <a:shade val="98000"/>
                                <a:lumMod val="102000"/>
                            </a:schemeClr>
                        </a:gs>
                        <a:gs pos="50000">
                            <a:schemeClr val="phClr">
                                <a:tint val="98000"/>
                                <a:satMod val="130000"/>
                                <a:shade val="90000"/>
                                <a:lumMod val="103000"/>
                            </a:schemeClr>
                        </a:gs>
                        <a:gs pos="100000">
                            <a:schemeClr val="phClr">
                                <a:shade val="63000"/>
                                <a:satMod val="120000"/>
                            </a:schemeClr>
                        </a:gs>
                    </a:gsLst>
                    <a:lin ang="5400000" scaled="0"/>
                </a:gradFill>
            </a:bgFillStyleLst>
        </a:fmtScheme>
    </a:themeElements>
    <a:objectDefaults/>
    <a:extraClrSchemeLst/>
    <a:extLst>
        <a:ext uri="{05A4C25C-085E-4340-85A3-A5531E510DB2}">
            <thm15:themeFamily
                    xmlns:thm15="http://schemas.microsoft.com/office/thememl/2012/main" name="Office Theme" id="{62F939B6-93AF-4DB8-9C6B-D6C7DFDC589F}" vid="{4A3C46E8-61CC-4603-A589-7422A47A8E4A}"/>
        </a:ext>
    </a:extLst>
</a:theme>`
	TEMP_WORD_STYLE_NEW = `
	<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
	<w:styles
			xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006"
			xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships"
			xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"
			xmlns:w14="http://schemas.microsoft.com/office/word/2010/wordml"
			xmlns:w15="http://schemas.microsoft.com/office/word/2012/wordml"
			xmlns:w16cex="http://schemas.microsoft.com/office/word/2018/wordml/cex"
			xmlns:w16cid="http://schemas.microsoft.com/office/word/2016/wordml/cid"
			xmlns:w16="http://schemas.microsoft.com/office/word/2018/wordml"
			xmlns:w16se="http://schemas.microsoft.com/office/word/2015/wordml/symex" mc:Ignorable="w14 w15 w16se w16cid w16 w16cex">
		<w:docDefaults>
			<w:rPrDefault>
				<w:rPr>
					<w:rFonts w:asciiTheme="minorHAnsi" w:eastAsiaTheme="minorEastAsia" w:hAnsiTheme="minorHAnsi" w:cstheme="minorBidi"/>
					<w:kern w:val="2"/>
					<w:sz w:val="21"/>
					<w:szCs w:val="24"/>
					<w:lang w:val="en-US" w:eastAsia="zh-CN" w:bidi="ar-SA"/>
				</w:rPr>
			</w:rPrDefault>
			<w:pPrDefault/>
		</w:docDefaults>
		<w:latentStyles w:defLockedState="0" w:defUIPriority="99" w:defSemiHidden="0" w:defUnhideWhenUsed="0" w:defQFormat="0" w:count="376">
			<w:lsdException w:name="Normal" w:uiPriority="0"/>
			<w:lsdException w:name="heading 1" w:uiPriority="9" w:qFormat="1"/>
			<w:lsdException w:name="heading 2" w:semiHidden="1" w:uiPriority="9" w:unhideWhenUsed="1" w:qFormat="1"/>
			<w:lsdException w:name="heading 3" w:semiHidden="1" w:uiPriority="9" w:unhideWhenUsed="1" w:qFormat="1"/>
			<w:lsdException w:name="heading 4" w:semiHidden="1" w:uiPriority="9" w:unhideWhenUsed="1" w:qFormat="1"/>
			<w:lsdException w:name="heading 5" w:semiHidden="1" w:uiPriority="9" w:unhideWhenUsed="1" w:qFormat="1"/>
			<w:lsdException w:name="heading 6" w:semiHidden="1" w:uiPriority="9" w:unhideWhenUsed="1" w:qFormat="1"/>
			<w:lsdException w:name="heading 7" w:semiHidden="1" w:uiPriority="9" w:unhideWhenUsed="1" w:qFormat="1"/>
			<w:lsdException w:name="heading 8" w:semiHidden="1" w:uiPriority="9" w:unhideWhenUsed="1" w:qFormat="1"/>
			<w:lsdException w:name="heading 9" w:semiHidden="1" w:uiPriority="9" w:unhideWhenUsed="1" w:qFormat="1"/>
			<w:lsdException w:name="index 1" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="index 2" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="index 3" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="index 4" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="index 5" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="index 6" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="index 7" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="index 8" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="index 9" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="toc 1" w:semiHidden="1" w:uiPriority="39" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="toc 2" w:semiHidden="1" w:uiPriority="39" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="toc 3" w:semiHidden="1" w:uiPriority="39" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="toc 4" w:semiHidden="1" w:uiPriority="39" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="toc 5" w:semiHidden="1" w:uiPriority="39" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="toc 6" w:semiHidden="1" w:uiPriority="39" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="toc 7" w:semiHidden="1" w:uiPriority="39" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="toc 8" w:semiHidden="1" w:uiPriority="39" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="toc 9" w:semiHidden="1" w:uiPriority="39" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Normal Indent" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="footnote text" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="annotation text" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="header" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="footer" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="index heading" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="caption" w:semiHidden="1" w:uiPriority="35" w:unhideWhenUsed="1" w:qFormat="1"/>
			<w:lsdException w:name="table of figures" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="envelope address" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="envelope return" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="footnote reference" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="annotation reference" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="line number" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="page number" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="endnote reference" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="endnote text" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="table of authorities" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="macro" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="toa heading" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="List" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="List Bullet" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="List Number" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="List 2" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="List 3" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="List 4" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="List 5" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="List Bullet 2" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="List Bullet 3" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="List Bullet 4" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="List Bullet 5" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="List Number 2" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="List Number 3" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="List Number 4" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="List Number 5" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Title" w:uiPriority="10" w:qFormat="1"/>
			<w:lsdException w:name="Closing" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Signature" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Default Paragraph Font" w:semiHidden="1" w:uiPriority="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Body Text" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Body Text Indent" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="List Continue" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="List Continue 2" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="List Continue 3" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="List Continue 4" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="List Continue 5" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Message Header" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Subtitle" w:uiPriority="11" w:qFormat="1"/>
			<w:lsdException w:name="Salutation" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Date" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Body Text First Indent" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Body Text First Indent 2" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Note Heading" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Body Text 2" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Body Text 3" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Body Text Indent 2" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Body Text Indent 3" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Block Text" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Hyperlink" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="FollowedHyperlink" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Strong" w:uiPriority="22" w:qFormat="1"/>
			<w:lsdException w:name="Emphasis" w:uiPriority="20" w:qFormat="1"/>
			<w:lsdException w:name="Document Map" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Plain Text" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="E-mail Signature" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="HTML Top of Form" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="HTML Bottom of Form" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Normal (Web)" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="HTML Acronym" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="HTML Address" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="HTML Cite" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="HTML Code" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="HTML Definition" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="HTML Keyboard" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="HTML Preformatted" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="HTML Sample" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="HTML Typewriter" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="HTML Variable" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Normal Table" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="annotation subject" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="No List" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Outline List 1" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Outline List 2" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Outline List 3" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Simple 1" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Simple 2" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Simple 3" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Classic 1" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Classic 2" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Classic 3" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Classic 4" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Colorful 1" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Colorful 2" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Colorful 3" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Columns 1" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Columns 2" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Columns 3" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Columns 4" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Columns 5" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Grid 1" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Grid 2" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Grid 3" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Grid 4" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Grid 5" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Grid 6" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Grid 7" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Grid 8" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table List 1" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table List 2" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table List 3" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table List 4" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table List 5" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table List 6" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table List 7" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table List 8" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table 3D effects 1" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table 3D effects 2" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table 3D effects 3" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Contemporary" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Elegant" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Professional" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Subtle 1" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Subtle 2" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Web 1" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Web 2" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Web 3" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Balloon Text" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Table Grid" w:uiPriority="39"/>
			<w:lsdException w:name="Table Theme" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Placeholder Text" w:semiHidden="1"/>
			<w:lsdException w:name="No Spacing" w:uiPriority="1" w:qFormat="1"/>
			<w:lsdException w:name="Light Shading" w:uiPriority="60"/>
			<w:lsdException w:name="Light List" w:uiPriority="61"/>
			<w:lsdException w:name="Light Grid" w:uiPriority="62"/>
			<w:lsdException w:name="Medium Shading 1" w:uiPriority="63"/>
			<w:lsdException w:name="Medium Shading 2" w:uiPriority="64"/>
			<w:lsdException w:name="Medium List 1" w:uiPriority="65"/>
			<w:lsdException w:name="Medium List 2" w:uiPriority="66"/>
			<w:lsdException w:name="Medium Grid 1" w:uiPriority="67"/>
			<w:lsdException w:name="Medium Grid 2" w:uiPriority="68"/>
			<w:lsdException w:name="Medium Grid 3" w:uiPriority="69"/>
			<w:lsdException w:name="Dark List" w:uiPriority="70"/>
			<w:lsdException w:name="Colorful Shading" w:uiPriority="71"/>
			<w:lsdException w:name="Colorful List" w:uiPriority="72"/>
			<w:lsdException w:name="Colorful Grid" w:uiPriority="73"/>
			<w:lsdException w:name="Light Shading Accent 1" w:uiPriority="60"/>
			<w:lsdException w:name="Light List Accent 1" w:uiPriority="61"/>
			<w:lsdException w:name="Light Grid Accent 1" w:uiPriority="62"/>
			<w:lsdException w:name="Medium Shading 1 Accent 1" w:uiPriority="63"/>
			<w:lsdException w:name="Medium Shading 2 Accent 1" w:uiPriority="64"/>
			<w:lsdException w:name="Medium List 1 Accent 1" w:uiPriority="65"/>
			<w:lsdException w:name="Revision" w:semiHidden="1"/>
			<w:lsdException w:name="List Paragraph" w:uiPriority="34" w:qFormat="1"/>
			<w:lsdException w:name="Quote" w:uiPriority="29" w:qFormat="1"/>
			<w:lsdException w:name="Intense Quote" w:uiPriority="30" w:qFormat="1"/>
			<w:lsdException w:name="Medium List 2 Accent 1" w:uiPriority="66"/>
			<w:lsdException w:name="Medium Grid 1 Accent 1" w:uiPriority="67"/>
			<w:lsdException w:name="Medium Grid 2 Accent 1" w:uiPriority="68"/>
			<w:lsdException w:name="Medium Grid 3 Accent 1" w:uiPriority="69"/>
			<w:lsdException w:name="Dark List Accent 1" w:uiPriority="70"/>
			<w:lsdException w:name="Colorful Shading Accent 1" w:uiPriority="71"/>
			<w:lsdException w:name="Colorful List Accent 1" w:uiPriority="72"/>
			<w:lsdException w:name="Colorful Grid Accent 1" w:uiPriority="73"/>
			<w:lsdException w:name="Light Shading Accent 2" w:uiPriority="60"/>
			<w:lsdException w:name="Light List Accent 2" w:uiPriority="61"/>
			<w:lsdException w:name="Light Grid Accent 2" w:uiPriority="62"/>
			<w:lsdException w:name="Medium Shading 1 Accent 2" w:uiPriority="63"/>
			<w:lsdException w:name="Medium Shading 2 Accent 2" w:uiPriority="64"/>
			<w:lsdException w:name="Medium List 1 Accent 2" w:uiPriority="65"/>
			<w:lsdException w:name="Medium List 2 Accent 2" w:uiPriority="66"/>
			<w:lsdException w:name="Medium Grid 1 Accent 2" w:uiPriority="67"/>
			<w:lsdException w:name="Medium Grid 2 Accent 2" w:uiPriority="68"/>
			<w:lsdException w:name="Medium Grid 3 Accent 2" w:uiPriority="69"/>
			<w:lsdException w:name="Dark List Accent 2" w:uiPriority="70"/>
			<w:lsdException w:name="Colorful Shading Accent 2" w:uiPriority="71"/>
			<w:lsdException w:name="Colorful List Accent 2" w:uiPriority="72"/>
			<w:lsdException w:name="Colorful Grid Accent 2" w:uiPriority="73"/>
			<w:lsdException w:name="Light Shading Accent 3" w:uiPriority="60"/>
			<w:lsdException w:name="Light List Accent 3" w:uiPriority="61"/>
			<w:lsdException w:name="Light Grid Accent 3" w:uiPriority="62"/>
			<w:lsdException w:name="Medium Shading 1 Accent 3" w:uiPriority="63"/>
			<w:lsdException w:name="Medium Shading 2 Accent 3" w:uiPriority="64"/>
			<w:lsdException w:name="Medium List 1 Accent 3" w:uiPriority="65"/>
			<w:lsdException w:name="Medium List 2 Accent 3" w:uiPriority="66"/>
			<w:lsdException w:name="Medium Grid 1 Accent 3" w:uiPriority="67"/>
			<w:lsdException w:name="Medium Grid 2 Accent 3" w:uiPriority="68"/>
			<w:lsdException w:name="Medium Grid 3 Accent 3" w:uiPriority="69"/>
			<w:lsdException w:name="Dark List Accent 3" w:uiPriority="70"/>
			<w:lsdException w:name="Colorful Shading Accent 3" w:uiPriority="71"/>
			<w:lsdException w:name="Colorful List Accent 3" w:uiPriority="72"/>
			<w:lsdException w:name="Colorful Grid Accent 3" w:uiPriority="73"/>
			<w:lsdException w:name="Light Shading Accent 4" w:uiPriority="60"/>
			<w:lsdException w:name="Light List Accent 4" w:uiPriority="61"/>
			<w:lsdException w:name="Light Grid Accent 4" w:uiPriority="62"/>
			<w:lsdException w:name="Medium Shading 1 Accent 4" w:uiPriority="63"/>
			<w:lsdException w:name="Medium Shading 2 Accent 4" w:uiPriority="64"/>
			<w:lsdException w:name="Medium List 1 Accent 4" w:uiPriority="65"/>
			<w:lsdException w:name="Medium List 2 Accent 4" w:uiPriority="66"/>
			<w:lsdException w:name="Medium Grid 1 Accent 4" w:uiPriority="67"/>
			<w:lsdException w:name="Medium Grid 2 Accent 4" w:uiPriority="68"/>
			<w:lsdException w:name="Medium Grid 3 Accent 4" w:uiPriority="69"/>
			<w:lsdException w:name="Dark List Accent 4" w:uiPriority="70"/>
			<w:lsdException w:name="Colorful Shading Accent 4" w:uiPriority="71"/>
			<w:lsdException w:name="Colorful List Accent 4" w:uiPriority="72"/>
			<w:lsdException w:name="Colorful Grid Accent 4" w:uiPriority="73"/>
			<w:lsdException w:name="Light Shading Accent 5" w:uiPriority="60"/>
			<w:lsdException w:name="Light List Accent 5" w:uiPriority="61"/>
			<w:lsdException w:name="Light Grid Accent 5" w:uiPriority="62"/>
			<w:lsdException w:name="Medium Shading 1 Accent 5" w:uiPriority="63"/>
			<w:lsdException w:name="Medium Shading 2 Accent 5" w:uiPriority="64"/>
			<w:lsdException w:name="Medium List 1 Accent 5" w:uiPriority="65"/>
			<w:lsdException w:name="Medium List 2 Accent 5" w:uiPriority="66"/>
			<w:lsdException w:name="Medium Grid 1 Accent 5" w:uiPriority="67"/>
			<w:lsdException w:name="Medium Grid 2 Accent 5" w:uiPriority="68"/>
			<w:lsdException w:name="Medium Grid 3 Accent 5" w:uiPriority="69"/>
			<w:lsdException w:name="Dark List Accent 5" w:uiPriority="70"/>
			<w:lsdException w:name="Colorful Shading Accent 5" w:uiPriority="71"/>
			<w:lsdException w:name="Colorful List Accent 5" w:uiPriority="72"/>
			<w:lsdException w:name="Colorful Grid Accent 5" w:uiPriority="73"/>
			<w:lsdException w:name="Light Shading Accent 6" w:uiPriority="60"/>
			<w:lsdException w:name="Light List Accent 6" w:uiPriority="61"/>
			<w:lsdException w:name="Light Grid Accent 6" w:uiPriority="62"/>
			<w:lsdException w:name="Medium Shading 1 Accent 6" w:uiPriority="63"/>
			<w:lsdException w:name="Medium Shading 2 Accent 6" w:uiPriority="64"/>
			<w:lsdException w:name="Medium List 1 Accent 6" w:uiPriority="65"/>
			<w:lsdException w:name="Medium List 2 Accent 6" w:uiPriority="66"/>
			<w:lsdException w:name="Medium Grid 1 Accent 6" w:uiPriority="67"/>
			<w:lsdException w:name="Medium Grid 2 Accent 6" w:uiPriority="68"/>
			<w:lsdException w:name="Medium Grid 3 Accent 6" w:uiPriority="69"/>
			<w:lsdException w:name="Dark List Accent 6" w:uiPriority="70"/>
			<w:lsdException w:name="Colorful Shading Accent 6" w:uiPriority="71"/>
			<w:lsdException w:name="Colorful List Accent 6" w:uiPriority="72"/>
			<w:lsdException w:name="Colorful Grid Accent 6" w:uiPriority="73"/>
			<w:lsdException w:name="Subtle Emphasis" w:uiPriority="19" w:qFormat="1"/>
			<w:lsdException w:name="Intense Emphasis" w:uiPriority="21" w:qFormat="1"/>
			<w:lsdException w:name="Subtle Reference" w:uiPriority="31" w:qFormat="1"/>
			<w:lsdException w:name="Intense Reference" w:uiPriority="32" w:qFormat="1"/>
			<w:lsdException w:name="Book Title" w:uiPriority="33" w:qFormat="1"/>
			<w:lsdException w:name="Bibliography" w:semiHidden="1" w:uiPriority="37" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="TOC Heading" w:semiHidden="1" w:uiPriority="39" w:unhideWhenUsed="1" w:qFormat="1"/>
			<w:lsdException w:name="Plain Table 1" w:uiPriority="41"/>
			<w:lsdException w:name="Plain Table 2" w:uiPriority="42"/>
			<w:lsdException w:name="Plain Table 3" w:uiPriority="43"/>
			<w:lsdException w:name="Plain Table 4" w:uiPriority="44"/>
			<w:lsdException w:name="Plain Table 5" w:uiPriority="45"/>
			<w:lsdException w:name="Grid Table Light" w:uiPriority="40"/>
			<w:lsdException w:name="Grid Table 1 Light" w:uiPriority="46"/>
			<w:lsdException w:name="Grid Table 2" w:uiPriority="47"/>
			<w:lsdException w:name="Grid Table 3" w:uiPriority="48"/>
			<w:lsdException w:name="Grid Table 4" w:uiPriority="49"/>
			<w:lsdException w:name="Grid Table 5 Dark" w:uiPriority="50"/>
			<w:lsdException w:name="Grid Table 6 Colorful" w:uiPriority="51"/>
			<w:lsdException w:name="Grid Table 7 Colorful" w:uiPriority="52"/>
			<w:lsdException w:name="Grid Table 1 Light Accent 1" w:uiPriority="46"/>
			<w:lsdException w:name="Grid Table 2 Accent 1" w:uiPriority="47"/>
			<w:lsdException w:name="Grid Table 3 Accent 1" w:uiPriority="48"/>
			<w:lsdException w:name="Grid Table 4 Accent 1" w:uiPriority="49"/>
			<w:lsdException w:name="Grid Table 5 Dark Accent 1" w:uiPriority="50"/>
			<w:lsdException w:name="Grid Table 6 Colorful Accent 1" w:uiPriority="51"/>
			<w:lsdException w:name="Grid Table 7 Colorful Accent 1" w:uiPriority="52"/>
			<w:lsdException w:name="Grid Table 1 Light Accent 2" w:uiPriority="46"/>
			<w:lsdException w:name="Grid Table 2 Accent 2" w:uiPriority="47"/>
			<w:lsdException w:name="Grid Table 3 Accent 2" w:uiPriority="48"/>
			<w:lsdException w:name="Grid Table 4 Accent 2" w:uiPriority="49"/>
			<w:lsdException w:name="Grid Table 5 Dark Accent 2" w:uiPriority="50"/>
			<w:lsdException w:name="Grid Table 6 Colorful Accent 2" w:uiPriority="51"/>
			<w:lsdException w:name="Grid Table 7 Colorful Accent 2" w:uiPriority="52"/>
			<w:lsdException w:name="Grid Table 1 Light Accent 3" w:uiPriority="46"/>
			<w:lsdException w:name="Grid Table 2 Accent 3" w:uiPriority="47"/>
			<w:lsdException w:name="Grid Table 3 Accent 3" w:uiPriority="48"/>
			<w:lsdException w:name="Grid Table 4 Accent 3" w:uiPriority="49"/>
			<w:lsdException w:name="Grid Table 5 Dark Accent 3" w:uiPriority="50"/>
			<w:lsdException w:name="Grid Table 6 Colorful Accent 3" w:uiPriority="51"/>
			<w:lsdException w:name="Grid Table 7 Colorful Accent 3" w:uiPriority="52"/>
			<w:lsdException w:name="Grid Table 1 Light Accent 4" w:uiPriority="46"/>
			<w:lsdException w:name="Grid Table 2 Accent 4" w:uiPriority="47"/>
			<w:lsdException w:name="Grid Table 3 Accent 4" w:uiPriority="48"/>
			<w:lsdException w:name="Grid Table 4 Accent 4" w:uiPriority="49"/>
			<w:lsdException w:name="Grid Table 5 Dark Accent 4" w:uiPriority="50"/>
			<w:lsdException w:name="Grid Table 6 Colorful Accent 4" w:uiPriority="51"/>
			<w:lsdException w:name="Grid Table 7 Colorful Accent 4" w:uiPriority="52"/>
			<w:lsdException w:name="Grid Table 1 Light Accent 5" w:uiPriority="46"/>
			<w:lsdException w:name="Grid Table 2 Accent 5" w:uiPriority="47"/>
			<w:lsdException w:name="Grid Table 3 Accent 5" w:uiPriority="48"/>
			<w:lsdException w:name="Grid Table 4 Accent 5" w:uiPriority="49"/>
			<w:lsdException w:name="Grid Table 5 Dark Accent 5" w:uiPriority="50"/>
			<w:lsdException w:name="Grid Table 6 Colorful Accent 5" w:uiPriority="51"/>
			<w:lsdException w:name="Grid Table 7 Colorful Accent 5" w:uiPriority="52"/>
			<w:lsdException w:name="Grid Table 1 Light Accent 6" w:uiPriority="46"/>
			<w:lsdException w:name="Grid Table 2 Accent 6" w:uiPriority="47"/>
			<w:lsdException w:name="Grid Table 3 Accent 6" w:uiPriority="48"/>
			<w:lsdException w:name="Grid Table 4 Accent 6" w:uiPriority="49"/>
			<w:lsdException w:name="Grid Table 5 Dark Accent 6" w:uiPriority="50"/>
			<w:lsdException w:name="Grid Table 6 Colorful Accent 6" w:uiPriority="51"/>
			<w:lsdException w:name="Grid Table 7 Colorful Accent 6" w:uiPriority="52"/>
			<w:lsdException w:name="List Table 1 Light" w:uiPriority="46"/>
			<w:lsdException w:name="List Table 2" w:uiPriority="47"/>
			<w:lsdException w:name="List Table 3" w:uiPriority="48"/>
			<w:lsdException w:name="List Table 4" w:uiPriority="49"/>
			<w:lsdException w:name="List Table 5 Dark" w:uiPriority="50"/>
			<w:lsdException w:name="List Table 6 Colorful" w:uiPriority="51"/>
			<w:lsdException w:name="List Table 7 Colorful" w:uiPriority="52"/>
			<w:lsdException w:name="List Table 1 Light Accent 1" w:uiPriority="46"/>
			<w:lsdException w:name="List Table 2 Accent 1" w:uiPriority="47"/>
			<w:lsdException w:name="List Table 3 Accent 1" w:uiPriority="48"/>
			<w:lsdException w:name="List Table 4 Accent 1" w:uiPriority="49"/>
			<w:lsdException w:name="List Table 5 Dark Accent 1" w:uiPriority="50"/>
			<w:lsdException w:name="List Table 6 Colorful Accent 1" w:uiPriority="51"/>
			<w:lsdException w:name="List Table 7 Colorful Accent 1" w:uiPriority="52"/>
			<w:lsdException w:name="List Table 1 Light Accent 2" w:uiPriority="46"/>
			<w:lsdException w:name="List Table 2 Accent 2" w:uiPriority="47"/>
			<w:lsdException w:name="List Table 3 Accent 2" w:uiPriority="48"/>
			<w:lsdException w:name="List Table 4 Accent 2" w:uiPriority="49"/>
			<w:lsdException w:name="List Table 5 Dark Accent 2" w:uiPriority="50"/>
			<w:lsdException w:name="List Table 6 Colorful Accent 2" w:uiPriority="51"/>
			<w:lsdException w:name="List Table 7 Colorful Accent 2" w:uiPriority="52"/>
			<w:lsdException w:name="List Table 1 Light Accent 3" w:uiPriority="46"/>
			<w:lsdException w:name="List Table 2 Accent 3" w:uiPriority="47"/>
			<w:lsdException w:name="List Table 3 Accent 3" w:uiPriority="48"/>
			<w:lsdException w:name="List Table 4 Accent 3" w:uiPriority="49"/>
			<w:lsdException w:name="List Table 5 Dark Accent 3" w:uiPriority="50"/>
			<w:lsdException w:name="List Table 6 Colorful Accent 3" w:uiPriority="51"/>
			<w:lsdException w:name="List Table 7 Colorful Accent 3" w:uiPriority="52"/>
			<w:lsdException w:name="List Table 1 Light Accent 4" w:uiPriority="46"/>
			<w:lsdException w:name="List Table 2 Accent 4" w:uiPriority="47"/>
			<w:lsdException w:name="List Table 3 Accent 4" w:uiPriority="48"/>
			<w:lsdException w:name="List Table 4 Accent 4" w:uiPriority="49"/>
			<w:lsdException w:name="List Table 5 Dark Accent 4" w:uiPriority="50"/>
			<w:lsdException w:name="List Table 6 Colorful Accent 4" w:uiPriority="51"/>
			<w:lsdException w:name="List Table 7 Colorful Accent 4" w:uiPriority="52"/>
			<w:lsdException w:name="List Table 1 Light Accent 5" w:uiPriority="46"/>
			<w:lsdException w:name="List Table 2 Accent 5" w:uiPriority="47"/>
			<w:lsdException w:name="List Table 3 Accent 5" w:uiPriority="48"/>
			<w:lsdException w:name="List Table 4 Accent 5" w:uiPriority="49"/>
			<w:lsdException w:name="List Table 5 Dark Accent 5" w:uiPriority="50"/>
			<w:lsdException w:name="List Table 6 Colorful Accent 5" w:uiPriority="51"/>
			<w:lsdException w:name="List Table 7 Colorful Accent 5" w:uiPriority="52"/>
			<w:lsdException w:name="List Table 1 Light Accent 6" w:uiPriority="46"/>
			<w:lsdException w:name="List Table 2 Accent 6" w:uiPriority="47"/>
			<w:lsdException w:name="List Table 3 Accent 6" w:uiPriority="48"/>
			<w:lsdException w:name="List Table 4 Accent 6" w:uiPriority="49"/>
			<w:lsdException w:name="List Table 5 Dark Accent 6" w:uiPriority="50"/>
			<w:lsdException w:name="List Table 6 Colorful Accent 6" w:uiPriority="51"/>
			<w:lsdException w:name="List Table 7 Colorful Accent 6" w:uiPriority="52"/>
			<w:lsdException w:name="Mention" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Smart Hyperlink" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Hashtag" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Unresolved Mention" w:semiHidden="1" w:unhideWhenUsed="1"/>
			<w:lsdException w:name="Smart Link" w:semiHidden="1" w:unhideWhenUsed="1"/>
		</w:latentStyles>
		<w:style w:type="paragraph" w:default="1" w:styleId="a">
			<w:name w:val="Normal"/>
			<w:rsid w:val="0055086D"/>
			<w:pPr>
				<w:spacing w:after="360" w:line="360" w:lineRule="exact"/>
			</w:pPr>
			<w:rPr>
				<w:rFonts w:ascii="微软雅黑" w:eastAsia="微软雅黑" w:hAnsi="微软雅黑" w:cs="宋体"/>
				<w:kern w:val="0"/>
			</w:rPr>
		</w:style>
		<w:style w:type="paragraph" w:styleId="1">
			<w:name w:val="heading 1"/>
			<w:basedOn w:val="a"/>
			<w:next w:val="a"/>
			<w:link w:val="10"/>
			<w:uiPriority w:val="9"/>
			<w:qFormat/>
			<w:rsid w:val="000E39C6"/>
			<w:pPr>
				<w:keepNext/>
				<w:keepLines/>
				<w:spacing w:before="340" w:after="330" w:line="578" w:lineRule="auto"/>
				<w:outlineLvl w:val="0"/>
			</w:pPr>
			<w:rPr>
				<w:b/>
				<w:bCs/>
				<w:kern w:val="44"/>
				<w:sz w:val="28"/>
				<w:szCs w:val="28"/>
			</w:rPr>
		</w:style>
		<w:style w:type="paragraph" w:styleId="2">
			<w:name w:val="heading 2"/>
			<w:basedOn w:val="a"/>
			<w:next w:val="a"/>
			<w:link w:val="20"/>
			<w:uiPriority w:val="9"/>
			<w:unhideWhenUsed/>
			<w:qFormat/>
			<w:rsid w:val="000E39C6"/>
			<w:pPr>
				<w:keepNext/>
				<w:keepLines/>
				<w:spacing w:before="260" w:after="260" w:line="416" w:lineRule="auto"/>
				<w:outlineLvl w:val="1"/>
			</w:pPr>
			<w:rPr>
				<w:rFonts w:asciiTheme="majorHAnsi" w:eastAsiaTheme="majorEastAsia" w:hAnsiTheme="majorHAnsi" w:cstheme="majorBidi"/>
				<w:b/>
				<w:bCs/>
				<w:sz w:val="24"/>
			</w:rPr>
		</w:style>
		<w:style w:type="character" w:default="1" w:styleId="a0">
			<w:name w:val="Default Paragraph Font"/>
			<w:uiPriority w:val="1"/>
			<w:semiHidden/>
			<w:unhideWhenUsed/>
		</w:style>
		<w:style w:type="table" w:default="1" w:styleId="a1">
			<w:name w:val="Normal Table"/>
			<w:uiPriority w:val="99"/>
			<w:semiHidden/>
			<w:unhideWhenUsed/>
			<w:tblPr>
				<w:tblInd w:w="0" w:type="dxa"/>
				<w:tblCellMar>
					<w:top w:w="0" w:type="dxa"/>
					<w:left w:w="108" w:type="dxa"/>
					<w:bottom w:w="0" w:type="dxa"/>
					<w:right w:w="108" w:type="dxa"/>
				</w:tblCellMar>
			</w:tblPr>
		</w:style>
		<w:style w:type="numbering" w:default="1" w:styleId="a2">
			<w:name w:val="No List"/>
			<w:uiPriority w:val="99"/>
			<w:semiHidden/>
			<w:unhideWhenUsed/>
		</w:style>
		<w:style w:type="character" w:customStyle="1" w:styleId="10">
			<w:name w:val="标题 1 字符"/>
			<w:basedOn w:val="a0"/>
			<w:link w:val="1"/>
			<w:uiPriority w:val="9"/>
			<w:rsid w:val="0029777C"/>
			<w:rPr>
				<w:rFonts w:ascii="微软雅黑" w:eastAsia="微软雅黑" w:hAnsi="微软雅黑" w:cs="宋体"/>
				<w:b/>
				<w:bCs/>
				<w:kern w:val="44"/>
				<w:sz w:val="28"/>
				<w:szCs w:val="28"/>
			</w:rPr>
		</w:style>
		<w:style w:type="character" w:customStyle="1" w:styleId="20">
			<w:name w:val="标题 2 字符"/>
			<w:basedOn w:val="a0"/>
			<w:link w:val="2"/>
			<w:uiPriority w:val="9"/>
			<w:rsid w:val="0029777C"/>
			<w:rPr>
				<w:rFonts w:asciiTheme="majorHAnsi" w:eastAsiaTheme="majorEastAsia" w:hAnsiTheme="majorHAnsi" w:cstheme="majorBidi"/>
				<w:b/>
				<w:bCs/>
				<w:kern w:val="0"/>
				<w:sz w:val="24"/>
			</w:rPr>
		</w:style>
		<w:style w:type="table" w:styleId="a3">
			<w:name w:val="Table Grid"/>
			<w:basedOn w:val="a1"/>
			<w:uiPriority w:val="39"/>
			<w:rsid w:val="00743371"/>
			<w:tblPr>
				<w:tblBorders>
					<w:top w:val="single" w:sz="4" w:space="0" w:color="auto"/>
					<w:left w:val="single" w:sz="4" w:space="0" w:color="auto"/>
					<w:bottom w:val="single" w:sz="4" w:space="0" w:color="auto"/>
					<w:right w:val="single" w:sz="4" w:space="0" w:color="auto"/>
					<w:insideH w:val="single" w:sz="4" w:space="0" w:color="auto"/>
					<w:insideV w:val="single" w:sz="4" w:space="0" w:color="auto"/>
				</w:tblBorders>
			</w:tblPr>
		</w:style>
		<w:style w:type="character" w:styleId="a4">
			<w:name w:val="annotation reference"/>
			<w:basedOn w:val="a0"/>
			<w:uiPriority w:val="99"/>
			<w:semiHidden/>
			<w:unhideWhenUsed/>
			<w:rsid w:val="0029777C"/>
			<w:rPr>
				<w:sz w:val="21"/>
				<w:szCs w:val="21"/>
			</w:rPr>
		</w:style>
		<w:style w:type="paragraph" w:styleId="a5">
			<w:name w:val="annotation text"/>
			<w:basedOn w:val="a"/>
			<w:link w:val="a6"/>
			<w:uiPriority w:val="99"/>
			<w:unhideWhenUsed/>
			<w:rsid w:val="0029777C"/>
		</w:style>
		<w:style w:type="character" w:customStyle="1" w:styleId="a6">
			<w:name w:val="批注文字 字符"/>
			<w:basedOn w:val="a0"/>
			<w:link w:val="a5"/>
			<w:uiPriority w:val="99"/>
			<w:rsid w:val="0029777C"/>
			<w:rPr>
				<w:rFonts w:ascii="微软雅黑" w:eastAsia="微软雅黑" w:hAnsi="微软雅黑" w:cs="宋体"/>
				<w:kern w:val="0"/>
			</w:rPr>
		</w:style>
		<w:style w:type="paragraph" w:styleId="a7">
			<w:name w:val="annotation subject"/>
			<w:basedOn w:val="a5"/>
			<w:next w:val="a5"/>
			<w:link w:val="a8"/>
			<w:uiPriority w:val="99"/>
			<w:semiHidden/>
			<w:unhideWhenUsed/>
			<w:rsid w:val="0029777C"/>
			<w:rPr>
				<w:b/>
				<w:bCs/>
			</w:rPr>
		</w:style>
		<w:style w:type="character" w:customStyle="1" w:styleId="a8">
			<w:name w:val="批注主题 字符"/>
			<w:basedOn w:val="a6"/>
			<w:link w:val="a7"/>
			<w:uiPriority w:val="99"/>
			<w:semiHidden/>
			<w:rsid w:val="0029777C"/>
			<w:rPr>
				<w:rFonts w:ascii="微软雅黑" w:eastAsia="微软雅黑" w:hAnsi="微软雅黑" w:cs="宋体"/>
				<w:b/>
				<w:bCs/>
				<w:kern w:val="0"/>
			</w:rPr>
		</w:style>
		<w:style w:type="paragraph" w:styleId="a9">
			<w:name w:val="Balloon Text"/>
			<w:basedOn w:val="a"/>
			<w:link w:val="aa"/>
			<w:uiPriority w:val="99"/>
			<w:semiHidden/>
			<w:unhideWhenUsed/>
			<w:rsid w:val="0029777C"/>
			<w:pPr>
				<w:spacing w:after="0" w:line="240" w:lineRule="auto"/>
			</w:pPr>
			<w:rPr>
				<w:rFonts w:ascii="宋体" w:eastAsia="宋体"/>
				<w:sz w:val="18"/>
				<w:szCs w:val="18"/>
			</w:rPr>
		</w:style>
		<w:style w:type="character" w:customStyle="1" w:styleId="aa">
			<w:name w:val="批注框文本 字符"/>
			<w:basedOn w:val="a0"/>
			<w:link w:val="a9"/>
			<w:uiPriority w:val="99"/>
			<w:semiHidden/>
			<w:rsid w:val="0029777C"/>
			<w:rPr>
				<w:rFonts w:ascii="宋体" w:eastAsia="宋体" w:hAnsi="微软雅黑" w:cs="宋体"/>
				<w:kern w:val="0"/>
				<w:sz w:val="18"/>
				<w:szCs w:val="18"/>
			</w:rPr>
		</w:style>
	</w:styles>
`
	DOC_Relation_NEW=`<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
		<Relationships
			xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
			<Relationship Id="rId8" Type="http://schemas.microsoft.com/office/2016/09/relationships/commentsIds" Target="commentsIds.xml"/>
			<Relationship Id="rId3" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles" Target="styles.xml"/>
			<Relationship Id="rId7" Type="http://schemas.microsoft.com/office/2011/relationships/commentsExtended" Target="commentsExtended.xml"/>
			<Relationship Id="rId12" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/theme" Target="theme/theme1.xml"/>
			<Relationship Id="rId2" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/numbering" Target="numbering.xml"/>
			<Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/customXml" Target="../customXml/item1.xml"/>
			<Relationship Id="rId6" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/comments" Target="comments.xml"/>
			<Relationship Id="rId11" Type="http://schemas.microsoft.com/office/2011/relationships/people" Target="people.xml"/>
			<Relationship Id="rId5" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/webSettings" Target="webSettings.xml"/>
			<Relationship Id="rId10" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/fontTable" Target="fontTable.xml"/>
			<Relationship Id="rId4" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/settings" Target="settings.xml"/>
			<Relationship Id="rId9" Type="http://schemas.microsoft.com/office/2018/08/relationships/commentsExtensible" Target="commentsExtensible.xml"/>
		</Relationships>`

	WOED_DOCUMENT_NEW=`
<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
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
        <w:p w14:paraId="5ED56D03" w14:textId="77777777" w:rsidR="00425149" w:rsidRPr="00743371" w:rsidRDefault="00743371" w:rsidP="0055086D">
            <w:pPr>
                <w:pStyle w:val="1"/>
            </w:pPr>
            <w:r w:rsidRPr="00743371">
                <w:rPr>
                    <w:rFonts w:hint="eastAsia"/>
                </w:rPr>
                <w:t>活动名称</w:t>
            </w:r>
        </w:p>
        <w:p w14:paraId="3CC3F9EF" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00743371" w:rsidP="0055086D">
            <w:r w:rsidRPr="0029777C">
                <w:rPr>
                    <w:rFonts w:hint="eastAsia"/>
                </w:rPr>
                <w:t>Spark幸运鹅</w:t>
            </w:r>
        </w:p>
        <w:p w14:paraId="45582A9A" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="00743371" w:rsidRDefault="00743371" w:rsidP="0055086D">
            <w:pPr>
                <w:pStyle w:val="1"/>
            </w:pPr>
            <w:r w:rsidRPr="00743371">
                <w:t>活动时间</w:t>
            </w:r>
        </w:p>
        <w:p w14:paraId="62673B8C" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="00743371" w:rsidRDefault="00743371" w:rsidP="0055086D">
            <w:r w:rsidRPr="00743371">
                <w:t>2019/11/02 00:00:00</w:t>
            </w:r>
        </w:p>
        <w:p w14:paraId="2D644F63" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00743371" w:rsidP="0055086D">
            <w:pPr>
                <w:pStyle w:val="1"/>
            </w:pPr>
            <w:r w:rsidRPr="0029777C">
                <w:t>参与者资格</w:t>
            </w:r>
        </w:p>
        <w:p w14:paraId="7C466785" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="00743371" w:rsidRDefault="00743371" w:rsidP="0055086D">
            <w:commentRangeStart w:id="0"/>
            <w:r w:rsidRPr="00743371">
                <w:t>参与者应具备以下全部资格：每个自然人仅可使用一个自己合法有权的QQ号码(微信账号、手机号码等)参与一次活动。参与者应具备以下全部资格：每个自然人仅可使用一个自己合法有权的QQ号码(微信账号、手机号码等)参与一次活动。参与者应具备以下全部资格</w:t>
            </w:r>
            <w:ins w:id="1" w:author="唐 穆" w:date="2020-10-27T18:32:00Z">
                <w:r w:rsidR="0029777C">
                    <w:rPr>
                        <w:rFonts w:hint="eastAsia"/>
                    </w:rPr>
                    <w:t>阿萨德阿萨德</w:t>
                </w:r>
            </w:ins>
            <w:r w:rsidRPr="00743371">
                <w:t>：每个自然人仅可使用一个自己合法有权的QQ号码(微信账号、手机号码等)参与一次活动。参与者应具备以下参与者应具备以下全部资格：每个自然人仅可使用一个自己合法有权的QQ号码(微信账号、可使用一个自己合法有权的QQ号码(微信账号、手机号码等)参与一次活动。参与者应具备以下参与者应具备以下全部资格：每个自然人仅可使用一个自己合法有权的QQ号码(微信账号</w:t>
            </w:r>
            <w:commentRangeEnd w:id="0"/>
            <w:r w:rsidR="0029777C">
                <w:rPr>
                    <w:rStyle w:val="a4"/>
                </w:rPr>
                <w:commentReference w:id="0"/>
            </w:r>
        </w:p>
        <w:p w14:paraId="33FE625C" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="00743371" w:rsidRDefault="00743371" w:rsidP="0055086D">
            <w:pPr>
                <w:pStyle w:val="1"/>
                <w:pPrChange w:id="2" w:author="唐 穆" w:date="2020-10-27T18:40:00Z">
                    <w:pPr>
                        <w:pStyle w:val="1"/>
                        <w:ind w:leftChars="-4" w:left="-8" w:firstLineChars="1"/>
                    </w:pPr>
                </w:pPrChange>
            </w:pPr>
            <w:r w:rsidRPr="00743371">
                <w:t>活动规则</w:t>
            </w:r>
        </w:p>
        <w:p w14:paraId="411045CE" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00743371" w:rsidP="0055086D">
            <w:pPr>
                <w:pStyle w:val="2"/>
                <w:pPrChange w:id="3" w:author="唐 穆" w:date="2020-10-27T18:40:00Z">
                    <w:pPr>
                        <w:pStyle w:val="1"/>
                    </w:pPr>
                </w:pPrChange>
            </w:pPr>
            <w:r w:rsidRPr="0029777C">
                <w:t>参与方式</w:t>
            </w:r>
        </w:p>
        <w:p w14:paraId="2852B645" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00743371" w:rsidP="0055086D">
            <w:pPr>
                <w:pPrChange w:id="4" w:author="唐 穆" w:date="2020-10-27T18:40:00Z">
                    <w:pPr>
                        <w:pStyle w:val="2"/>
                    </w:pPr>
                </w:pPrChange>
            </w:pPr>
            <w:r w:rsidRPr="0029777C">
                <w:t>请扫描“腾讯游戏年度发布会”官方小程序码，或打开微信搜一搜搜索“腾讯游戏”，进入“腾讯游戏年度发布会”小程序;玩家预约观看发布会后参与预约环节组队活动并选择想要参与的抽奖，小队满员后即可自动获得抽奖参与资格。请扫描“腾讯游戏年度发布会”官方小程序码，或打开微信搜一搜搜索“腾讯游戏”，进入“腾讯游戏年度发布会”小程序;玩家预约观看发布会后参与预约环节组队活动并选择想要参与的抽奖，小队满员后即可自动获得抽奖参与资格。抽奖，小队满员后即可自动获得抽奖参与资格。</w:t>
            </w:r>
        </w:p>
        <w:p w14:paraId="08274505" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="00743371" w:rsidRDefault="00743371" w:rsidP="0055086D">
            <w:pPr>
                <w:pStyle w:val="2"/>
            </w:pPr>
            <w:r w:rsidRPr="00743371">
                <w:lastRenderedPageBreak/>
                <w:t>中奖规则</w:t>
            </w:r>
        </w:p>
        <w:p w14:paraId="658F7D06" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="00743371" w:rsidRDefault="00743371" w:rsidP="0055086D">
            <w:r w:rsidRPr="00743371">
                <w:t>请扫描“腾讯游戏年度发布会”官方小程序码，或打开微信搜一搜搜索“腾讯游戏”，进入“腾讯游戏年度发布会”小程序;玩家预约观看发布会后参与预约环节组队活动并选择想要参与的抽奖，小队满员后即可自动获得抽奖参与资格。请扫描“腾讯游戏年度发布会”官方小程序码，或打开微信搜一搜搜索“腾讯游戏”，进入“腾讯游戏年度发布会”小程序;玩家预约观看发布会后参与预约环节组队活动并选择想要参与的抽奖，小队满员后即可自动获得抽奖参与资格。抽奖，小队满员后即可自动获得抽奖参与资格。</w:t>
            </w:r>
        </w:p>
        <w:p w14:paraId="4CBC2F76" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="00743371" w:rsidRDefault="00743371" w:rsidP="0055086D">
            <w:pPr>
                <w:pStyle w:val="2"/>
            </w:pPr>
            <w:r w:rsidRPr="00743371">
                <w:t>奖品内容</w:t>
            </w:r>
        </w:p>
        <w:tbl>
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
                <w:tblGridChange w:id="5">
                    <w:tblGrid>
                        <w:gridCol w:w="1830"/>
                        <w:gridCol w:w="3288"/>
                        <w:gridCol w:w="2069"/>
                        <w:gridCol w:w="1103"/>
                    </w:tblGrid>
                </w:tblGridChange>
            </w:tblGrid>
            <w:tr w:rsidR="0055086D" w:rsidRPr="00743371" w14:paraId="0E6E14AA" w14:textId="77777777" w:rsidTr="0055086D">
                <w:tc>
                    <w:tcPr>
                        <w:tcW w:w="1104" w:type="pct"/>
                        <w:shd w:val="clear" w:color="auto" w:fill="D9D9D9" w:themeFill="background1" w:themeFillShade="D9"/>
                    </w:tcPr>
                    <w:p w14:paraId="6A75AB5D" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00743371" w:rsidP="0055086D">
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
                    <w:p w14:paraId="3F3C0566" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00743371" w:rsidP="0055086D">
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
                    <w:p w14:paraId="7CA7E222" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00743371" w:rsidP="0055086D">
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
                    <w:p w14:paraId="4F9B2883" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00743371" w:rsidP="0055086D">
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
            <w:tr w:rsidR="0029777C" w:rsidRPr="00743371" w14:paraId="525E0653" w14:textId="77777777" w:rsidTr="0055086D">
                <w:tc>
                    <w:tcPr>
                        <w:tcW w:w="1104" w:type="pct"/>
                    </w:tcPr>
                    <w:p w14:paraId="301B86E4" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00743371" w:rsidP="0055086D">
                        <w:pPr>
                            <w:spacing w:after="0" w:line="240" w:lineRule="auto"/>
                        </w:pPr>
                        <w:r w:rsidRPr="0029777C">
                            <w:rPr>
                                <w:rFonts w:hint="eastAsia"/>
                            </w:rPr>
                            <w:t>一等奖</w:t>
                        </w:r>
                    </w:p>
                </w:tc>
                <w:tc>
                    <w:tcPr>
                        <w:tcW w:w="1983" w:type="pct"/>
                    </w:tcPr>
                    <w:p w14:paraId="2B129718" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00743371" w:rsidP="0055086D">
                        <w:pPr>
                            <w:spacing w:after="0" w:line="240" w:lineRule="auto"/>
                        </w:pPr>
                        <w:r w:rsidRPr="0029777C">
                            <w:rPr>
                                <w:rFonts w:hint="eastAsia"/>
                            </w:rPr>
                            <w:t>奖品名称奖品名称</w:t>
                        </w:r>
                    </w:p>
                </w:tc>
                <w:tc>
                    <w:tcPr>
                        <w:tcW w:w="1248" w:type="pct"/>
                    </w:tcPr>
                    <w:p w14:paraId="033B4B64" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00743371" w:rsidP="0055086D">
                        <w:pPr>
                            <w:spacing w:after="0" w:line="240" w:lineRule="auto"/>
                        </w:pPr>
                        <w:r w:rsidRPr="0029777C">
                            <w:t>0.00052%</w:t>
                        </w:r>
                    </w:p>
                </w:tc>
                <w:tc>
                    <w:tcPr>
                        <w:tcW w:w="665" w:type="pct"/>
                    </w:tcPr>
                    <w:p w14:paraId="14C7DC60" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00743371" w:rsidP="0055086D">
                        <w:pPr>
                            <w:spacing w:after="0" w:line="240" w:lineRule="auto"/>
                        </w:pPr>
                        <w:r w:rsidRPr="0029777C">
                            <w:t>10</w:t>
                        </w:r>
                    </w:p>
                </w:tc>
            </w:tr>
            <w:tr w:rsidR="0029777C" w:rsidRPr="00743371" w14:paraId="799EC489" w14:textId="77777777" w:rsidTr="0055086D">
                <w:tc>
                    <w:tcPr>
                        <w:tcW w:w="1104" w:type="pct"/>
                    </w:tcPr>
                    <w:p w14:paraId="7C257883" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00743371" w:rsidP="0055086D">
                        <w:pPr>
                            <w:spacing w:after="0" w:line="240" w:lineRule="auto"/>
                        </w:pPr>
                        <w:r w:rsidRPr="0029777C">
                            <w:rPr>
                                <w:rFonts w:hint="eastAsia"/>
                            </w:rPr>
                            <w:t>二等奖</w:t>
                        </w:r>
                    </w:p>
                </w:tc>
                <w:tc>
                    <w:tcPr>
                        <w:tcW w:w="1983" w:type="pct"/>
                    </w:tcPr>
                    <w:p w14:paraId="448D3C0A" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00743371" w:rsidP="0055086D">
                        <w:pPr>
                            <w:spacing w:after="0" w:line="240" w:lineRule="auto"/>
                        </w:pPr>
                        <w:r w:rsidRPr="0029777C">
                            <w:rPr>
                                <w:rFonts w:hint="eastAsia"/>
                            </w:rPr>
                            <w:t>奖品名称奖品名称</w:t>
                        </w:r>
                    </w:p>
                </w:tc>
                <w:tc>
                    <w:tcPr>
                        <w:tcW w:w="1248" w:type="pct"/>
                    </w:tcPr>
                    <w:p w14:paraId="79452E9C" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00743371" w:rsidP="0055086D">
                        <w:pPr>
                            <w:spacing w:after="0" w:line="240" w:lineRule="auto"/>
                        </w:pPr>
                        <w:r w:rsidRPr="0029777C">
                            <w:t>0.00052%</w:t>
                        </w:r>
                    </w:p>
                </w:tc>
                <w:tc>
                    <w:tcPr>
                        <w:tcW w:w="665" w:type="pct"/>
                    </w:tcPr>
                    <w:p w14:paraId="6E1A14D9" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="0029777C" w:rsidRDefault="00743371" w:rsidP="0055086D">
                        <w:pPr>
                            <w:spacing w:after="0" w:line="240" w:lineRule="auto"/>
                        </w:pPr>
                        <w:r w:rsidRPr="0029777C">
                            <w:t>10</w:t>
                        </w:r>
                    </w:p>
                </w:tc>
            </w:tr>
        </w:tbl>
        <w:p w14:paraId="5AE6B7C6" w14:textId="77777777" w:rsidR="00743371" w:rsidRPr="00743371" w:rsidRDefault="00743371" w:rsidP="0055086D"/>
        <w:sectPr w:rsidR="00743371" w:rsidRPr="00743371" w:rsidSect="00294703">
            <w:pgSz w:w="11900" w:h="16840"/>
            <w:pgMar w:top="1440" w:right="1800" w:bottom="1440" w:left="1800" w:header="851" w:footer="992" w:gutter="0"/>
            <w:cols w:space="425"/>
            <w:docGrid w:type="lines" w:linePitch="312"/>
        </w:sectPr>
    </w:body>
</w:document>
`
)
