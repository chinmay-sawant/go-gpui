package remote

import (
	"fmt"
	"math"
)

func (a *App) setSensitivity(value float64) {
	a.view.Sensitivity = clampInt(int(math.Round(value/5))*5, 25, 300)
	a.view.SensitivityLabel = fmt.Sprintf("%.2fx", float64(a.view.Sensitivity)/100)
	a.view.SensitivityFill = (a.view.Sensitivity - 25) * 100 / 275
}

func (a *App) sensitivityAt(x float64) {
	for _, b := range a.page.Boxes() {
		if b.ID == "sensitivity" && b.W > 0 {
			a.sliderPosition = math.Max(25, math.Min(300, 25+(x-b.X)*275/b.W))
			a.setSensitivity(a.sliderPosition)
			return
		}
	}
}

func (a *App) moveSensitivity(dx float64) {
	for _, b := range a.page.Boxes() {
		if b.ID == "sensitivity" && b.W > 0 {
			a.sliderPosition = math.Max(25, math.Min(300, a.sliderPosition+dx*275/b.W))
			a.setSensitivity(a.sliderPosition)
			a.setData()
			return
		}
	}
}
