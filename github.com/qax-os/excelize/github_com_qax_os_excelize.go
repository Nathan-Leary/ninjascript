package excelize
import (
github_com_qax_os_excelize "github.com/qax-os/excelize"
)
func Load() {
if _, ok := DefaultGojaInterfaces["github.com/qax-os/excelize"]; !ok {
   DefaultGojaInterfaces["github.com/qax-os/excelize"] = map[string]interface{}{}
}
DefaultGojaInterfaces["github.com/qax-os/excelize"]["STCellFormulaTypeArray"] = github_com_qax_os_excelize.STCellFormulaTypeArray
DefaultGojaInterfaces["github.com/qax-os/excelize"]["STCellFormulaTypeDataTable"] = github_com_qax_os_excelize.STCellFormulaTypeDataTable
DefaultGojaInterfaces["github.com/qax-os/excelize"]["STCellFormulaTypeNormal"] = github_com_qax_os_excelize.STCellFormulaTypeNormal
DefaultGojaInterfaces["github.com/qax-os/excelize"]["STCellFormulaTypeShared"] = github_com_qax_os_excelize.STCellFormulaTypeShared
DefaultGojaInterfaces["github.com/qax-os/excelize"]["Bar"] = github_com_qax_os_excelize.Bar
DefaultGojaInterfaces["github.com/qax-os/excelize"]["BarStacked"] = github_com_qax_os_excelize.BarStacked
DefaultGojaInterfaces["github.com/qax-os/excelize"]["BarPercentStacked"] = github_com_qax_os_excelize.BarPercentStacked
DefaultGojaInterfaces["github.com/qax-os/excelize"]["Bar3DClustered"] = github_com_qax_os_excelize.Bar3DClustered
DefaultGojaInterfaces["github.com/qax-os/excelize"]["Bar3DStacked"] = github_com_qax_os_excelize.Bar3DStacked
DefaultGojaInterfaces["github.com/qax-os/excelize"]["Bar3DPercentStacked"] = github_com_qax_os_excelize.Bar3DPercentStacked
DefaultGojaInterfaces["github.com/qax-os/excelize"]["Col"] = github_com_qax_os_excelize.Col
DefaultGojaInterfaces["github.com/qax-os/excelize"]["ColStacked"] = github_com_qax_os_excelize.ColStacked
DefaultGojaInterfaces["github.com/qax-os/excelize"]["ColPercentStacked"] = github_com_qax_os_excelize.ColPercentStacked
DefaultGojaInterfaces["github.com/qax-os/excelize"]["Col3DClustered"] = github_com_qax_os_excelize.Col3DClustered
DefaultGojaInterfaces["github.com/qax-os/excelize"]["Col3D"] = github_com_qax_os_excelize.Col3D
DefaultGojaInterfaces["github.com/qax-os/excelize"]["Col3DStacked"] = github_com_qax_os_excelize.Col3DStacked
DefaultGojaInterfaces["github.com/qax-os/excelize"]["Col3DPercentStacked"] = github_com_qax_os_excelize.Col3DPercentStacked
DefaultGojaInterfaces["github.com/qax-os/excelize"]["Doughnut"] = github_com_qax_os_excelize.Doughnut
DefaultGojaInterfaces["github.com/qax-os/excelize"]["Line"] = github_com_qax_os_excelize.Line
DefaultGojaInterfaces["github.com/qax-os/excelize"]["Pie"] = github_com_qax_os_excelize.Pie
DefaultGojaInterfaces["github.com/qax-os/excelize"]["Pie3D"] = github_com_qax_os_excelize.Pie3D
DefaultGojaInterfaces["github.com/qax-os/excelize"]["Radar"] = github_com_qax_os_excelize.Radar
DefaultGojaInterfaces["github.com/qax-os/excelize"]["Scatter"] = github_com_qax_os_excelize.Scatter
DefaultGojaInterfaces["github.com/qax-os/excelize"]["DataValidationTypeCustom"] = github_com_qax_os_excelize.DataValidationTypeCustom
DefaultGojaInterfaces["github.com/qax-os/excelize"]["DataValidationTypeDate"] = github_com_qax_os_excelize.DataValidationTypeDate
DefaultGojaInterfaces["github.com/qax-os/excelize"]["DataValidationTypeDecimal"] = github_com_qax_os_excelize.DataValidationTypeDecimal
DefaultGojaInterfaces["github.com/qax-os/excelize"]["DataValidationTypeTextLeng"] = github_com_qax_os_excelize.DataValidationTypeTextLeng
DefaultGojaInterfaces["github.com/qax-os/excelize"]["DataValidationTypeTime"] = github_com_qax_os_excelize.DataValidationTypeTime
DefaultGojaInterfaces["github.com/qax-os/excelize"]["DataValidationTypeWhole"] = github_com_qax_os_excelize.DataValidationTypeWhole
DefaultGojaInterfaces["github.com/qax-os/excelize"]["DataValidationOperatorBetween"] = github_com_qax_os_excelize.DataValidationOperatorBetween
DefaultGojaInterfaces["github.com/qax-os/excelize"]["DataValidationOperatorEqual"] = github_com_qax_os_excelize.DataValidationOperatorEqual
DefaultGojaInterfaces["github.com/qax-os/excelize"]["DataValidationOperatorGreaterThan"] = github_com_qax_os_excelize.DataValidationOperatorGreaterThan
DefaultGojaInterfaces["github.com/qax-os/excelize"]["DataValidationOperatorGreaterThanOrEqual"] = github_com_qax_os_excelize.DataValidationOperatorGreaterThanOrEqual
DefaultGojaInterfaces["github.com/qax-os/excelize"]["DataValidationOperatorLessThan"] = github_com_qax_os_excelize.DataValidationOperatorLessThan
DefaultGojaInterfaces["github.com/qax-os/excelize"]["DataValidationOperatorLessThanOrEqual"] = github_com_qax_os_excelize.DataValidationOperatorLessThanOrEqual
DefaultGojaInterfaces["github.com/qax-os/excelize"]["DataValidationOperatorNotBetween"] = github_com_qax_os_excelize.DataValidationOperatorNotBetween
DefaultGojaInterfaces["github.com/qax-os/excelize"]["DataValidationOperatorNotEqual"] = github_com_qax_os_excelize.DataValidationOperatorNotEqual
DefaultGojaInterfaces["github.com/qax-os/excelize"]["SourceRelationship"] = github_com_qax_os_excelize.SourceRelationship
DefaultGojaInterfaces["github.com/qax-os/excelize"]["SourceRelationshipChart"] = github_com_qax_os_excelize.SourceRelationshipChart
DefaultGojaInterfaces["github.com/qax-os/excelize"]["SourceRelationshipComments"] = github_com_qax_os_excelize.SourceRelationshipComments
DefaultGojaInterfaces["github.com/qax-os/excelize"]["SourceRelationshipImage"] = github_com_qax_os_excelize.SourceRelationshipImage
DefaultGojaInterfaces["github.com/qax-os/excelize"]["SourceRelationshipTable"] = github_com_qax_os_excelize.SourceRelationshipTable
DefaultGojaInterfaces["github.com/qax-os/excelize"]["SourceRelationshipDrawingML"] = github_com_qax_os_excelize.SourceRelationshipDrawingML
DefaultGojaInterfaces["github.com/qax-os/excelize"]["SourceRelationshipDrawingVML"] = github_com_qax_os_excelize.SourceRelationshipDrawingVML
DefaultGojaInterfaces["github.com/qax-os/excelize"]["SourceRelationshipHyperLink"] = github_com_qax_os_excelize.SourceRelationshipHyperLink
DefaultGojaInterfaces["github.com/qax-os/excelize"]["SourceRelationshipWorkSheet"] = github_com_qax_os_excelize.SourceRelationshipWorkSheet
DefaultGojaInterfaces["github.com/qax-os/excelize"]["SourceRelationshipChart201506"] = github_com_qax_os_excelize.SourceRelationshipChart201506
DefaultGojaInterfaces["github.com/qax-os/excelize"]["SourceRelationshipChart20070802"] = github_com_qax_os_excelize.SourceRelationshipChart20070802
DefaultGojaInterfaces["github.com/qax-os/excelize"]["SourceRelationshipChart2014"] = github_com_qax_os_excelize.SourceRelationshipChart2014
DefaultGojaInterfaces["github.com/qax-os/excelize"]["SourceRelationshipCompatibility"] = github_com_qax_os_excelize.SourceRelationshipCompatibility
DefaultGojaInterfaces["github.com/qax-os/excelize"]["NameSpaceDrawingML"] = github_com_qax_os_excelize.NameSpaceDrawingML
DefaultGojaInterfaces["github.com/qax-os/excelize"]["NameSpaceDrawingMLChart"] = github_com_qax_os_excelize.NameSpaceDrawingMLChart
DefaultGojaInterfaces["github.com/qax-os/excelize"]["NameSpaceDrawingMLSpreadSheet"] = github_com_qax_os_excelize.NameSpaceDrawingMLSpreadSheet
DefaultGojaInterfaces["github.com/qax-os/excelize"]["NameSpaceSpreadSheet"] = github_com_qax_os_excelize.NameSpaceSpreadSheet
DefaultGojaInterfaces["github.com/qax-os/excelize"]["NameSpaceXML"] = github_com_qax_os_excelize.NameSpaceXML
DefaultGojaInterfaces["github.com/qax-os/excelize"]["EMU"] = github_com_qax_os_excelize.EMU
DefaultGojaInterfaces["github.com/qax-os/excelize"]["HSLModel"] = github_com_qax_os_excelize.HSLModel
DefaultGojaInterfaces["github.com/qax-os/excelize"]["HSLToRGB"] = github_com_qax_os_excelize.HSLToRGB
DefaultGojaInterfaces["github.com/qax-os/excelize"]["RGBToHSL"] = github_com_qax_os_excelize.RGBToHSL
DefaultGojaInterfaces["github.com/qax-os/excelize"]["ReadZipReader"] = github_com_qax_os_excelize.ReadZipReader
DefaultGojaInterfaces["github.com/qax-os/excelize"]["ThemeColor"] = github_com_qax_os_excelize.ThemeColor
DefaultGojaInterfaces["github.com/qax-os/excelize"]["TitleToNumber"] = github_com_qax_os_excelize.TitleToNumber
DefaultGojaInterfaces["github.com/qax-os/excelize"]["ToAlphaString"] = github_com_qax_os_excelize.ToAlphaString
DefaultGojaInterfaces["github.com/qax-os/excelize"]["Comment"] = github_com_qax_os_excelize.Comment{}
DefaultGojaInterfaces["github.com/qax-os/excelize"]["DataValidation"] = github_com_qax_os_excelize.DataValidation{}
DefaultGojaInterfaces["github.com/qax-os/excelize"]["NewDataValidation"] = github_com_qax_os_excelize.NewDataValidation
DefaultGojaInterfaces["github.com/qax-os/excelize"]["DataValidationErrorStyleStop"] = github_com_qax_os_excelize.DataValidationErrorStyleStop
DefaultGojaInterfaces["github.com/qax-os/excelize"]["DataValidationErrorStyleWarning"] = github_com_qax_os_excelize.DataValidationErrorStyleWarning
DefaultGojaInterfaces["github.com/qax-os/excelize"]["DataValidationErrorStyleInformation"] = github_com_qax_os_excelize.DataValidationErrorStyleInformation
DefaultGojaInterfaces["github.com/qax-os/excelize"]["ErrSheetNotExist"] = github_com_qax_os_excelize.ErrSheetNotExist{}
DefaultGojaInterfaces["github.com/qax-os/excelize"]["File"] = github_com_qax_os_excelize.File{}
DefaultGojaInterfaces["github.com/qax-os/excelize"]["NewFile"] = github_com_qax_os_excelize.NewFile
DefaultGojaInterfaces["github.com/qax-os/excelize"]["OpenFile"] = github_com_qax_os_excelize.OpenFile
DefaultGojaInterfaces["github.com/qax-os/excelize"]["OpenReader"] = github_com_qax_os_excelize.OpenReader
DefaultGojaInterfaces["github.com/qax-os/excelize"]["HSL"] = github_com_qax_os_excelize.HSL{}
DefaultGojaInterfaces["github.com/qax-os/excelize"]["Rows"] = github_com_qax_os_excelize.Rows{}
DefaultGojaInterfaces["github.com/qax-os/excelize"]["HSLToRGB"] = github_com_qax_os_excelize.HSLToRGB
DefaultGojaInterfaces["github.com/qax-os/excelize"]["RGBToHSL"] = github_com_qax_os_excelize.RGBToHSL
DefaultGojaInterfaces["github.com/qax-os/excelize"]["ReadZipReader"] = github_com_qax_os_excelize.ReadZipReader
DefaultGojaInterfaces["github.com/qax-os/excelize"]["ThemeColor"] = github_com_qax_os_excelize.ThemeColor
DefaultGojaInterfaces["github.com/qax-os/excelize"]["TitleToNumber"] = github_com_qax_os_excelize.TitleToNumber
DefaultGojaInterfaces["github.com/qax-os/excelize"]["ToAlphaString"] = github_com_qax_os_excelize.ToAlphaString
DefaultGojaInterfaces["github.com/qax-os/excelize"]["NewDataValidation"] = github_com_qax_os_excelize.NewDataValidation
DefaultGojaInterfaces["github.com/qax-os/excelize"]["NewFile"] = github_com_qax_os_excelize.NewFile
DefaultGojaInterfaces["github.com/qax-os/excelize"]["OpenFile"] = github_com_qax_os_excelize.OpenFile
DefaultGojaInterfaces["github.com/qax-os/excelize"]["OpenReader"] = github_com_qax_os_excelize.OpenReader
}
