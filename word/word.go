package word

import (
	"math/rand"
)

// funnyWords is a list of words that can be used to describe a release.
var funnyWords = []string{
	"awesome",
	"stunning",
	"bug-free",
	"amazing",
	"fantastic",
	"unbelievable",
	"mind-blowing",
	"epic",
	"legendary",
	"superb",
	"phenomenal",
	"sensational",
	"extraordinary",
	"incredible",
	"spectacular",
	"jaw-dropping",
	"breathtaking",
	"outstanding",
	"marvelous",
	"wondrous",
	"fabulous",
	"terrific",
	"splendid",
	"fantabulous",
	"radical",
	"groovy",
	"stellar",
	"wicked",
	"dazzling",
	"brilliant",
	"smashing",
	"majestic",
	"excellent",
	"awesome",
	"stellar",
	"phenomenal",
	"surreal",
	"supreme",
	"grand",
	"top-notch",
	"mega",
	"zany",
	"outrageous",
	"funky",
	"fantasmic",
	"spectacular",
	"bodacious",
}

// RandomReleaseWord makes no sense, but it's fun
func RandomReleaseWord() string {
	return funnyWords[rand.Intn(len(funnyWords))]
}
