package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

const sideWidth = 300

type appLayout struct {
	title, toolbar, content, actions, cases, divider fyne.CanvasObject
	//divider                  [3]fyne.CanvasObject
}

func newAppLayout(title, toolbar, content, actions, cases, divider fyne.CanvasObject) fyne.Layout {
	return &appLayout{title: title, toolbar: toolbar, content: content, actions: actions, cases: cases, divider: divider}
}

func (l *appLayout) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	//middleHeight := size.Height
	wh := l.title.MinSize().Height
	pd := l.title.MinSize().Height * .1
	l.title.Resize(fyne.NewSize(size.Width-sideWidth, wh))
	l.title.Move(fyne.NewPos(0, 0))
	l.toolbar.Resize(fyne.NewSize(size.Width-sideWidth, wh*1.5))
	l.toolbar.Move(fyne.NewPos(0, wh+pd))
	l.content.Resize(fyne.NewSize(size.Width-sideWidth, size.Height-3.5*wh-3*pd))
	l.content.Move(fyne.NewPos(0, wh*2.5+pd*2))
	l.divider.Move(fyne.NewPos(size.Width-sideWidth+theme.SeparatorThicknessSize()*2, 0))
	l.divider.Resize(fyne.NewSize(theme.SeparatorThicknessSize()*2, size.Height))
	l.actions.Resize(fyne.NewSize(size.Width-sideWidth, wh))
	l.actions.Move(fyne.NewPos(0, size.Height-wh))
	l.cases.Resize(fyne.NewSize(sideWidth-theme.SeparatorThicknessSize()*4, size.Height))
	l.cases.Move(fyne.NewPos(size.Width-sideWidth+theme.SeparatorThicknessSize()*4, 0))
}

func (l *appLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	borders := fyne.NewSize(
		sideWidth*3.5,
		l.title.MinSize().Height*3+l.content.MinSize().Height,
	)
	return borders.AddWidthHeight(100, 100)
}

type asLayout struct {
	label, entry, errmsg fyne.CanvasObject
	//dividers                  [3]fyne.CanvasObject
}

func newAsLayout(label, entry, errmsg fyne.CanvasObject) fyne.Layout {
	return &asLayout{label: label, entry: entry, errmsg: errmsg}
}

func (l *asLayout) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	wh := size.Height
	l.label.Resize(fyne.NewSize(size.Width/2, wh))
	l.label.Move(fyne.NewPos(0, 0))
	l.entry.Resize(fyne.NewSize(size.Width*1/8, wh))
	l.entry.Move(fyne.NewPos(size.Width/2, 0))
	l.errmsg.Resize(fyne.NewSize(size.Width*3/8, wh))
	l.errmsg.Move(fyne.NewPos(size.Width*(1.0-3.0/8.0), 0))
}

func (l *asLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	borders := fyne.NewSize(
		300,
		l.label.MinSize().Height,
	)
	return borders
}
