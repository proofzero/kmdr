		// Type of the label that notes a kind of each lines.
labelType: "long"

// Style of "Build" lines.
buildStyle: {
	// Hide lines
	hide: false
	// Bold or increased intensity.
	bold:       true
	faint:      true
	italic:     true
	underline:  true
	blinkSlow:  true
	blinkRapid: false
	// Swap the foreground color and background color.
	inverse:  false
	conceal:  false
	crossOut: false
	frame:    false
	encircle: false
	overline: false
}
// Fore-color of text
// foreground: (#xxxxxx | rgb(0-256,0-256,0-256) | rgb(0x00-0xFF,0x00-0xFF,0x00-0xFF) | (name of colors))
// Back-color of text
// background: # Same format as `foreground`
// Style of the "Start" lines.
// startStyle:
// Same format as `buildStyle`
// Style of the "Pass" lines.
// passStyle:
// Same format as `buildStyle`
// Style of the "Fail" lines.
// failStyle:
// Same format as `buildStyle`
// Style of the "Skip" lines.
// skipStyle:
// Same format as `buildStyle`
// Style of the "File" lines.
// fileStyle:
// Same format as `buildStyle`
// Style of the "Line" lines.
// lineStyle:
// Same format as `buildStyle`
// Style of the "Pass" package lines.
// passPackageStyle:
// Same format as `buildStyle`
// Style of the "Fail" package lines.
// failPackageStyle:
// Same format as `buildStyle`
// A threashold of the coverage
coverThreshold: 80

// Style of the "Cover" lines with the coverage that is higher than coverThreshold.
// coveredStyle:
// Same format as `buildStyle`
// Style of the "Cover" lines with the coverage that is lower than coverThreshold.
// uncoveredStyle:
// Same format as `buildStyle`
// If you want to delete lines, write the regular expressions.
// removals:
//   - (regexp)
// If you want to leave `Test` prefixes, set it "true".
leaveTestPrefix: true
