package ownframe

func repaintCases() []repaintCase {
	return []repaintCase{
		loginCase(),
		platformCase(),
		statesCase(),
		formsCase(),
		scrollCase(),
		themeCase(),
	}
}
