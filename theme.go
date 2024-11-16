package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type fysionTheme struct {
	fyne.Theme
}

var themex fyne.ThemeColorName = "name"

func newFysionTheme() fyne.Theme {
	return &fysionTheme{Theme: theme.DefaultTheme()}
}

func (t *fysionTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return hexColor("#EFF5F7FD")

		}
		return hexColor("#FCCFAAFD")
	}
	if name == theme.ColorNameHeaderBackground {
		return hexColor("#FDF6CAFD")
	}
	if name == theme.ColorNameInputBackground {
		return hexColor("#FFFFFFFF")
	}
	if name == theme.ColorNameSeparator {
		return hexColor("#585858FD")
	}
	if name == theme.ColorNameInputBorder {
		return hexColor("#585858FD")
	}
	if name == theme.ColorNameForeground {
		return hexColor("#585858FD")
	}
	if name == theme.ColorNameButton {
		return hexColor("#8CD2FAFF")
	}
	if name == themex {
		return hexColor("#FF0000FF")
	}
	return t.Theme.Color(name, variant)
}

func (t *fysionTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 12
	}

	return t.Theme.Size(name)
}

func hexColor(s string) (c color.RGBA) {
	c.A = 0xff

	if s[0] != '#' {
		return c
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		return 0
	}
	switch len(s) {
	case 9:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
		c.A = hexToByte(s[7])<<4 + hexToByte(s[8])
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	}
	return
}

/*
type interiorTheme struct {
	fyne.Theme
}

func newInteriorTheme() fyne.Theme {
	return &fysionTheme{Theme: theme.DefaultTheme()}
}

func (t *interiorTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return hexColor("#D2EAF0FD")

		}
		return hexColor("#FCCFAAFD")
	}
	if name == theme.ColorNameHeaderBackground {
		return hexColor("#FDF6CAFD")
	}
	if name == theme.ColorNameInputBorder {
		return hexColor("#585858FD")
	}
	if name == theme.ColorNameForeground {
		return hexColor("#585858FD")
	}
	if name == theme.ColorNameButton {
		return hexColor("#ffb900")
	}
	return t.Theme.Color(name, variant)
}

func (t *interiorTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 12
	}

	return t.Theme.Size(name)
}
*/
