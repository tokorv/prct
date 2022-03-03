package main

type Config struct {
	Languages []LanguageConfig `json:"config"`
}

type LanguageConfig struct {
	Language  string        `json:"language"`
	Extension string        `json:"extension"`
	Colours   ColourSetting `json:"colours"`
}

type ColourSetting struct {
	Keywords  []ColourOption `json:"keywords"`
	Strings   string         `json:"strings"`
	Integers  string         `json:"integers"`
	Floats    string         `json:"floats"`
	Operators []ColourOption `json:"operators"`
}

type ColourOption struct {
	Name   string `json:"option"`
	Colour string `json:"colours"`
}
