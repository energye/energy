//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package application

type AppearanceNamed string

const (
	// NSAppearanceNameAqua - The standard light system appearance.
	NSAppearanceNameAqua AppearanceNamed = "NSAppearanceNameAqua"
	// NSAppearanceNameDarkAqua - The standard dark system appearance.
	NSAppearanceNameDarkAqua AppearanceNamed = "NSAppearanceNameDarkAqua"
	// NSAppearanceNameVibrantLight - The light vibrant appearance
	NSAppearanceNameVibrantLight AppearanceNamed = "NSAppearanceNameVibrantLight"
	// NSAppearanceNameAccessibilityHighContrastAqua - A high-contrast version of the standard light system appearance.
	NSAppearanceNameAccessibilityHighContrastAqua AppearanceNamed = "NSAppearanceNameAccessibilityHighContrastAqua"
	// NSAppearanceNameAccessibilityHighContrastDarkAqua - A high-contrast version of the standard dark system appearance.
	NSAppearanceNameAccessibilityHighContrastDarkAqua AppearanceNamed = "NSAppearanceNameAccessibilityHighContrastDarkAqua"
	// NSAppearanceNameAccessibilityHighContrastVibrantLight - A high-contrast version of the light vibrant appearance.
	NSAppearanceNameAccessibilityHighContrastVibrantLight AppearanceNamed = "NSAppearanceNameAccessibilityHighContrastVibrantLight"
	// NSAppearanceNameAccessibilityHighContrastVibrantDark - A high-contrast version of the dark vibrant appearance.
	NSAppearanceNameAccessibilityHighContrastVibrantDark AppearanceNamed = "NSAppearanceNameAccessibilityHighContrastVibrantDark"
)

type MacOS struct {
	AppearanceNamed AppearanceNamed
	// TODO add titlebar
}
