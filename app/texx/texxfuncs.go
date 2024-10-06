package texx

import "github.com/kudrykv/latex-yearly-planner/app/tex"

func EmphCell(text string) string {
	return tex.CellColor("white", tex.TextColor("black", text))
}
