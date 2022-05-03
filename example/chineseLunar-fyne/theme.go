package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type Biu struct{}

func (b Biu) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.DefaultTheme().Font(s)
		}
		return resourceNotoSansSCRegularTtf
	}
	if s.Italic {
		return theme.DefaultTheme().Font(s)
	}
	return resourceNotoSansSCRegularTtf
}

var (
	white = &color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	black = &color.RGBA{A: 50}
	snow4 = &color.RGBA{R: 139, G: 137, B: 137, A: 150}
)

func (b Biu) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return white
	case theme.ColorNameButton, theme.ColorNameDisabled:
		return black
	case theme.ColorNamePlaceHolder, theme.ColorNameScrollBar:
		return black
	case theme.ColorNameInputBackground:
		return snow4
	case theme.ColorNamePrimary, theme.ColorNameHover, theme.ColorNameFocus:
		return snow4
	case theme.ColorNameShadow:
		return snow4
	case theme.ColorNameForeground:
		return &color.RGBA{R: 0, G: 0, B: 0, A: 255}
	default:
		return snow4
	}
}

func (b Biu) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (b Biu) Size(name fyne.ThemeSizeName) float32 {
	if name == "text" {
		return 13.0
	}
	return theme.DefaultTheme().Size(name)
}
