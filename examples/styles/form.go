package styles

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
)

var Form = []gcss.Style{
	{
		Selector: ".input",
		Props: gcss.Props{
			BackgroundColor: input,
			Border: props.Border{
				Width: props.UnitPx(1),
				Style: props.BorderStyleSolid,
				Color: borderColor,
			},
			BorderRadius:  radius,
			Display:       props.DisplayFlex,
			FontSize:      fontSm,
			Height:        spacing10,
			LineHeight:    leadingTight,
			PaddingTop:    spacing2,
			PaddingRight:  spacing3,
			PaddingBottom: spacing2,
			PaddingLeft:   spacing3,
			Width:         props.UnitPercent(100),
		},
	},
	{
		Selector: ".input::file-selector-button",
		Props: gcss.Props{
			BackgroundColor: props.ColorTransparent(),
			BorderWidth:     props.UnitRaw(0),
			FontSize:        fontSm,
			FontWeight:      props.FontWeightMedium,
		},
	},
	{
		Selector: ".input-label",
		Props: gcss.Props{
			FontSize:   fontSm,
			FontWeight: props.FontWeightMedium,
			LineHeight: leadingTight,
		},
	},
	{
		Selector: ".input-help",
		Props: gcss.Props{
			Color:    mutedForeground,
			FontSize: fontSm,
		},
	},
	{
		Selector: ".input-error",
		Props: gcss.Props{
			Color:    destructive,
			FontSize: fontSm,
		},
	},
	{
		Selector: ".select",
		Props: gcss.Props{
			BackgroundColor: input,
			Border: props.Border{
				Width: props.UnitPx(1),
				Style: props.BorderStyleSolid,
				Color: borderColor,
			},
			BorderRadius:  radius,
			Display:       props.DisplayFlex,
			FontSize:      fontSm,
			Height:        spacing10,
			LineHeight:    leadingTight,
			PaddingTop:    spacing2,
			PaddingRight:  spacing3,
			PaddingBottom: spacing2,
			PaddingLeft:   spacing3,
			Width:         props.UnitPercent(100),
		},
	},
	{
		Selector: ".select:not([size])",
		Props: gcss.Props{
			Appearance:       props.AppearanceNone,
			PaddingRight:     spacing10,
			PrintColorAdjust: props.PrintColorAdjustExact,
			BackgroundImage:  iconChevronDown,
			BackgroundPosition: props.BackgroundPositionEdges(
				props.BackgroundPositionEdge{Position: props.BackgroundPositionRight, Unit: spacing3},
				props.BackgroundPositionEdge{Position: props.BackgroundPositionCenter},
			),
			BackgroundRepeat: props.BackgroundRepeatNoRepeat,
			BackgroundSize:   props.BackgroundSizeDimension(props.UnitEm(1), props.UnitEm(1)),
		},
	},
}