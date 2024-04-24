package style

import (
	"fmt"
	"io"
	"reflect"

	"github.com/AccentDesign/gostyle/props"
)

type (
	Props struct {
		AlignItems      props.AlignItems     `css:"align-items"`
		BackgroundColor props.Color          `css:"background-color"`
		BorderRadius    props.Unit           `css:"border-radius"`
		BorderColor     props.Color          `css:"border-color"`
		BorderStyle     props.BorderStyle    `css:"border-style"`
		BorderWidth     props.Unit           `css:"border-width"`
		Color           props.Color          `css:"color"`
		Display         props.Display        `css:"display"`
		Height          props.Unit           `css:"height"`
		JustifyContent  props.JustifyContent `css:"justify-content"`
		Margin          props.Margin         `css:"margin"`
		MarginTop       props.Margin         `css:"margin-top"`
		MarginRight     props.Margin         `css:"margin-right"`
		MarginBottom    props.Margin         `css:"margin-bottom"`
		MarginLeft      props.Margin         `css:"margin-left"`
		MaxWidth        props.Unit           `css:"max-width"`
		MinWidth        props.Unit           `css:"min-width"`
		Padding         props.Unit           `css:"padding"`
		PaddingTop      props.Unit           `css:"padding-top"`
		PaddingRight    props.Unit           `css:"padding-right"`
		PaddingBottom   props.Unit           `css:"padding-bottom"`
		PaddingLeft     props.Unit           `css:"padding-left"`
		Width           props.Width          `css:"width"`
	}
	Style struct {
		Selector string
		Props    Props
	}
)

// CSS returns the CSS string representation of the style.
func (s *Style) CSS(w io.Writer) error {
	propsValue := reflect.ValueOf(s.Props)
	propsType := reflect.TypeOf(s.Props)

	if _, err := fmt.Fprintf(w, "%s{", s.Selector); err != nil {
		return err
	}

	for i := 0; i < propsValue.NumField(); i++ {
		fieldValue := propsValue.Field(i)
		fieldType := propsType.Field(i)

		if fieldValue.IsZero() {
			continue
		}

		fieldName := fieldType.Tag.Get("css")

		if v, ok := fieldValue.Interface().(fmt.Stringer); ok {
			if _, err := fmt.Fprintf(w, "%s:%s;", fieldName, v.String()); err != nil {
				return err
			}
		}
	}

	if _, err := fmt.Fprint(w, "}"); err != nil {
		return err
	}

	return nil
}
