package styles

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
)

var Badges = []gcss.Style{
	{
		Selector: ".badge",
		Props: gcss.Props{
			AlignItems:    props.AlignItemsCenter,
			Display:       props.DisplayInlineFlex,
			BorderRadius:  radiusFull,
			FontSize:      fontSm,
			FontWeight:    props.FontWeightMedium,
			LineHeight:    leadingNone,
			PaddingTop:    spacing1,
			PaddingRight:  spacing3,
			PaddingBottom: spacing1,
			PaddingLeft:   spacing3,
		},
	},
	{
		Selector: ".badge-primary",
		Props: gcss.Props{
			BackgroundColor: primary,
			Color:           primaryForeground,
		},
	},
	{
		Selector: ".badge-secondary",
		Props: gcss.Props{
			BackgroundColor: secondary,
			Color:           secondaryForeground,
		},
	},
	{
		Selector: ".badge-destructive",
		Props: gcss.Props{
			BackgroundColor: destructive,
			Color:           destructiveForeground,
		},
	},
	{
		Selector: ".badge-outline",
		Props: gcss.Props{
			BackgroundColor: props.ColorTransparent(),
			Border: props.Border{
				Color: borderColor,
				Style: props.BorderStyleSolid,
				Width: props.UnitPx(1),
			},
		},
	},
}