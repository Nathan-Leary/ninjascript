package excelize
import (
github_com_qax_os_excelize "github.com/qax-os/excelize"
)
func Load(Interfaces ...interface{}) {
	if len(Interfaces) > 5 {
		if i, ok := Interfaces[0].(map[string]map[string]interface{}) {
			 DefaultGojaInterfaces: = i
			 if i, ok := Interfaces[1].(*goja.Runtime) {
				vm := i
				if i, ok := Interfaces[2].(*quickjs.Context) {
					ctx := i	
					if i, ok := Interfaces[3].(func(*quickjs.Context, interface{}) (quickjs.Value, string)) {
						ConvertInterfaceToQuickJS := i
						if i, ok := Interfaces[4].(func(interface{}, ...interface{}) interface{}) {
							ExecuteFunction := i
if _, ok := DefaultGojaInterfaces["excelize"]; !ok {
   DefaultGojaInterfaces["excelize"] = map[string]interface{}{}
}
DefaultGojaInterfaces["excelize"]["STCellFormulaTypeArray"] = github_com_qax_os_excelize.STCellFormulaTypeArray
DefaultGojaInterfaces["excelize"]["STCellFormulaTypeDataTable"] = github_com_qax_os_excelize.STCellFormulaTypeDataTable
DefaultGojaInterfaces["excelize"]["STCellFormulaTypeNormal"] = github_com_qax_os_excelize.STCellFormulaTypeNormal
DefaultGojaInterfaces["excelize"]["STCellFormulaTypeShared"] = github_com_qax_os_excelize.STCellFormulaTypeShared
DefaultGojaInterfaces["excelize"]["Bar"] = github_com_qax_os_excelize.Bar
DefaultGojaInterfaces["excelize"]["BarStacked"] = github_com_qax_os_excelize.BarStacked
DefaultGojaInterfaces["excelize"]["BarPercentStacked"] = github_com_qax_os_excelize.BarPercentStacked
DefaultGojaInterfaces["excelize"]["Bar3DClustered"] = github_com_qax_os_excelize.Bar3DClustered
DefaultGojaInterfaces["excelize"]["Bar3DStacked"] = github_com_qax_os_excelize.Bar3DStacked
DefaultGojaInterfaces["excelize"]["Bar3DPercentStacked"] = github_com_qax_os_excelize.Bar3DPercentStacked
DefaultGojaInterfaces["excelize"]["Col"] = github_com_qax_os_excelize.Col
DefaultGojaInterfaces["excelize"]["ColStacked"] = github_com_qax_os_excelize.ColStacked
DefaultGojaInterfaces["excelize"]["ColPercentStacked"] = github_com_qax_os_excelize.ColPercentStacked
DefaultGojaInterfaces["excelize"]["Col3DClustered"] = github_com_qax_os_excelize.Col3DClustered
DefaultGojaInterfaces["excelize"]["Col3D"] = github_com_qax_os_excelize.Col3D
DefaultGojaInterfaces["excelize"]["Col3DStacked"] = github_com_qax_os_excelize.Col3DStacked
DefaultGojaInterfaces["excelize"]["Col3DPercentStacked"] = github_com_qax_os_excelize.Col3DPercentStacked
DefaultGojaInterfaces["excelize"]["Doughnut"] = github_com_qax_os_excelize.Doughnut
DefaultGojaInterfaces["excelize"]["Line"] = github_com_qax_os_excelize.Line
DefaultGojaInterfaces["excelize"]["Pie"] = github_com_qax_os_excelize.Pie
DefaultGojaInterfaces["excelize"]["Pie3D"] = github_com_qax_os_excelize.Pie3D
DefaultGojaInterfaces["excelize"]["Radar"] = github_com_qax_os_excelize.Radar
DefaultGojaInterfaces["excelize"]["Scatter"] = github_com_qax_os_excelize.Scatter
DefaultGojaInterfaces["excelize"]["DataValidationTypeCustom"] = github_com_qax_os_excelize.DataValidationTypeCustom
DefaultGojaInterfaces["excelize"]["DataValidationTypeDate"] = github_com_qax_os_excelize.DataValidationTypeDate
DefaultGojaInterfaces["excelize"]["DataValidationTypeDecimal"] = github_com_qax_os_excelize.DataValidationTypeDecimal
DefaultGojaInterfaces["excelize"]["DataValidationTypeTextLeng"] = github_com_qax_os_excelize.DataValidationTypeTextLeng
DefaultGojaInterfaces["excelize"]["DataValidationTypeTime"] = github_com_qax_os_excelize.DataValidationTypeTime
DefaultGojaInterfaces["excelize"]["DataValidationTypeWhole"] = github_com_qax_os_excelize.DataValidationTypeWhole
DefaultGojaInterfaces["excelize"]["DataValidationOperatorBetween"] = github_com_qax_os_excelize.DataValidationOperatorBetween
DefaultGojaInterfaces["excelize"]["DataValidationOperatorEqual"] = github_com_qax_os_excelize.DataValidationOperatorEqual
DefaultGojaInterfaces["excelize"]["DataValidationOperatorGreaterThan"] = github_com_qax_os_excelize.DataValidationOperatorGreaterThan
DefaultGojaInterfaces["excelize"]["DataValidationOperatorGreaterThanOrEqual"] = github_com_qax_os_excelize.DataValidationOperatorGreaterThanOrEqual
DefaultGojaInterfaces["excelize"]["DataValidationOperatorLessThan"] = github_com_qax_os_excelize.DataValidationOperatorLessThan
DefaultGojaInterfaces["excelize"]["DataValidationOperatorLessThanOrEqual"] = github_com_qax_os_excelize.DataValidationOperatorLessThanOrEqual
DefaultGojaInterfaces["excelize"]["DataValidationOperatorNotBetween"] = github_com_qax_os_excelize.DataValidationOperatorNotBetween
DefaultGojaInterfaces["excelize"]["DataValidationOperatorNotEqual"] = github_com_qax_os_excelize.DataValidationOperatorNotEqual
DefaultGojaInterfaces["excelize"]["SourceRelationship"] = github_com_qax_os_excelize.SourceRelationship
DefaultGojaInterfaces["excelize"]["SourceRelationshipChart"] = github_com_qax_os_excelize.SourceRelationshipChart
DefaultGojaInterfaces["excelize"]["SourceRelationshipComments"] = github_com_qax_os_excelize.SourceRelationshipComments
DefaultGojaInterfaces["excelize"]["SourceRelationshipImage"] = github_com_qax_os_excelize.SourceRelationshipImage
DefaultGojaInterfaces["excelize"]["SourceRelationshipTable"] = github_com_qax_os_excelize.SourceRelationshipTable
DefaultGojaInterfaces["excelize"]["SourceRelationshipDrawingML"] = github_com_qax_os_excelize.SourceRelationshipDrawingML
DefaultGojaInterfaces["excelize"]["SourceRelationshipDrawingVML"] = github_com_qax_os_excelize.SourceRelationshipDrawingVML
DefaultGojaInterfaces["excelize"]["SourceRelationshipHyperLink"] = github_com_qax_os_excelize.SourceRelationshipHyperLink
DefaultGojaInterfaces["excelize"]["SourceRelationshipWorkSheet"] = github_com_qax_os_excelize.SourceRelationshipWorkSheet
DefaultGojaInterfaces["excelize"]["SourceRelationshipChart201506"] = github_com_qax_os_excelize.SourceRelationshipChart201506
DefaultGojaInterfaces["excelize"]["SourceRelationshipChart20070802"] = github_com_qax_os_excelize.SourceRelationshipChart20070802
DefaultGojaInterfaces["excelize"]["SourceRelationshipChart2014"] = github_com_qax_os_excelize.SourceRelationshipChart2014
DefaultGojaInterfaces["excelize"]["SourceRelationshipCompatibility"] = github_com_qax_os_excelize.SourceRelationshipCompatibility
DefaultGojaInterfaces["excelize"]["NameSpaceDrawingML"] = github_com_qax_os_excelize.NameSpaceDrawingML
DefaultGojaInterfaces["excelize"]["NameSpaceDrawingMLChart"] = github_com_qax_os_excelize.NameSpaceDrawingMLChart
DefaultGojaInterfaces["excelize"]["NameSpaceDrawingMLSpreadSheet"] = github_com_qax_os_excelize.NameSpaceDrawingMLSpreadSheet
DefaultGojaInterfaces["excelize"]["NameSpaceSpreadSheet"] = github_com_qax_os_excelize.NameSpaceSpreadSheet
DefaultGojaInterfaces["excelize"]["NameSpaceXML"] = github_com_qax_os_excelize.NameSpaceXML
DefaultGojaInterfaces["excelize"]["EMU"] = github_com_qax_os_excelize.EMU
DefaultGojaInterfaces["excelize"]["HSLModel"] = github_com_qax_os_excelize.HSLModel
DefaultGojaInterfaces["excelize"]["HSLToRGB"] = github_com_qax_os_excelize.HSLToRGB
DefaultGojaInterfaces["excelize"]["RGBToHSL"] = github_com_qax_os_excelize.RGBToHSL
DefaultGojaInterfaces["excelize"]["ReadZipReader"] = github_com_qax_os_excelize.ReadZipReader
DefaultGojaInterfaces["excelize"]["ThemeColor"] = github_com_qax_os_excelize.ThemeColor
DefaultGojaInterfaces["excelize"]["TitleToNumber"] = github_com_qax_os_excelize.TitleToNumber
DefaultGojaInterfaces["excelize"]["ToAlphaString"] = github_com_qax_os_excelize.ToAlphaString
DefaultGojaInterfaces["excelize"]["Comment"] = github_com_qax_os_excelize.Comment{}
DefaultGojaInterfaces["excelize"]["DataValidation"] = github_com_qax_os_excelize.DataValidation{}
DefaultGojaInterfaces["excelize"]["NewDataValidation"] = github_com_qax_os_excelize.NewDataValidation
DefaultGojaInterfaces["excelize"]["DataValidationErrorStyleStop"] = github_com_qax_os_excelize.DataValidationErrorStyleStop
DefaultGojaInterfaces["excelize"]["DataValidationErrorStyleWarning"] = github_com_qax_os_excelize.DataValidationErrorStyleWarning
DefaultGojaInterfaces["excelize"]["DataValidationErrorStyleInformation"] = github_com_qax_os_excelize.DataValidationErrorStyleInformation
DefaultGojaInterfaces["excelize"]["ErrSheetNotExist"] = github_com_qax_os_excelize.ErrSheetNotExist{}
DefaultGojaInterfaces["excelize"]["File"] = github_com_qax_os_excelize.File{}
DefaultGojaInterfaces["excelize"]["NewFile"] = github_com_qax_os_excelize.NewFile
DefaultGojaInterfaces["excelize"]["OpenFile"] = github_com_qax_os_excelize.OpenFile
DefaultGojaInterfaces["excelize"]["OpenReader"] = github_com_qax_os_excelize.OpenReader
DefaultGojaInterfaces["excelize"]["HSL"] = github_com_qax_os_excelize.HSL{}
DefaultGojaInterfaces["excelize"]["Rows"] = github_com_qax_os_excelize.Rows{}
DefaultGojaInterfaces["excelize"]["HSLToRGB"] = github_com_qax_os_excelize.HSLToRGB
DefaultGojaInterfaces["excelize"]["RGBToHSL"] = github_com_qax_os_excelize.RGBToHSL
DefaultGojaInterfaces["excelize"]["ReadZipReader"] = github_com_qax_os_excelize.ReadZipReader
DefaultGojaInterfaces["excelize"]["ThemeColor"] = github_com_qax_os_excelize.ThemeColor
DefaultGojaInterfaces["excelize"]["TitleToNumber"] = github_com_qax_os_excelize.TitleToNumber
DefaultGojaInterfaces["excelize"]["ToAlphaString"] = github_com_qax_os_excelize.ToAlphaString
DefaultGojaInterfaces["excelize"]["NewDataValidation"] = github_com_qax_os_excelize.NewDataValidation
DefaultGojaInterfaces["excelize"]["NewFile"] = github_com_qax_os_excelize.NewFile
DefaultGojaInterfaces["excelize"]["OpenFile"] = github_com_qax_os_excelize.OpenFile
DefaultGojaInterfaces["excelize"]["OpenReader"] = github_com_qax_os_excelize.OpenReader

			 			}	
			 		}
			 	}
			 }
		}
	}
}