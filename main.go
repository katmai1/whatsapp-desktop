package main

import (
	// "github.com/zserge/webview"
	"github.com/visualfc/goqt/ui"
)

func main() {
	s := make([]string, 1)
	winn := ui.NewApplication(s)
	winn.SetAttributeWithAttributeOn(ui.Qt_AA_MSWindowsUseDirect3DByDefault, true)
	// winn.
	loader := ui.NewUiLoader()
	w := loader.Load(ui.NewFileWithName("form.ui"))
	// winn.SetAttribute(ui.Qt_AA_MSWindowsUseDirect3DByDefault)
	// winn.SetAttributeWithAttributeOn(ui.Qt_AA_, true)
	// Open wikipedia in a 800x600 resizable window
	// webview.Open("Whatsapp Desktop",
	// "https://web.whatsapp.com/", 800, 600, true)
	w.Show()
	winn.Exec()
}
