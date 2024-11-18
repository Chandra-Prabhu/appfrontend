package main

import (
	"fyne.io/fyne/v2"
)

const sideWidth = 300

type appLayout struct {
	title, toolbar, content, actions, cases fyne.CanvasObject
	//dividers                  [3]fyne.CanvasObject
}

func newAppLayout(title, toolbar, content, actions, cases fyne.CanvasObject) fyne.Layout {
	return &appLayout{title: title, toolbar: toolbar, content: content, actions: actions, cases: cases}
}

func (l *appLayout) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	//middleHeight := size.Height
	wh := l.title.MinSize().Height
	l.title.Resize(fyne.NewSize(size.Width-sideWidth, wh))
	l.title.Move(fyne.NewPos(0, 0))
	l.toolbar.Resize(fyne.NewSize(size.Width-sideWidth, wh))
	l.toolbar.Move(fyne.NewPos(0, wh))
	l.content.Resize(fyne.NewSize(size.Width-sideWidth, size.Height-3*wh))
	l.content.Move(fyne.NewPos(0, wh*2))
	l.actions.Resize(fyne.NewSize(size.Width-sideWidth, wh))
	l.actions.Move(fyne.NewPos(0, size.Height-wh))
	l.cases.Resize(fyne.NewSize(sideWidth, size.Height))
	l.cases.Move(fyne.NewPos(size.Width-sideWidth, 0))
}

func (l *appLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	borders := fyne.NewSize(
		sideWidth*3,
		l.title.MinSize().Height*3+l.content.MinSize().Height,
	)
	return borders.AddWidthHeight(100, 100)
}
