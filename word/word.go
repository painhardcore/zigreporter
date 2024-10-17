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
	"blazingly fast",
	"ziggy",
	"cutting-edge",
	"innovative",
	"next-gen",
	"revolutionary",
	"state-of-the-art",
	"dynamic",
	"powerful",
	"robust",
	"sleek",
	"efficient",
	"user-friendly",
	"intuitive",
	"advanced",
	"progressive",
	"trailblazing",
	"groundbreaking",
	"premium",
	"refined",
	"optimized",
	"enhanced",
	"swift",
	"agile",
	"versatile",
	"flexible",
	"adaptive",
	"seamless",
	"reliable",
	"resilient",
	"impactful",
	"game-changing",
	"unique",
	"distinctive",
	"exceptional",
	"elite",
	"high-performance",
	"smart",
	"bright",
	"brisk",
	"charming",
	"crisp",
	"elegant",
	"flawless",
	"glorious",
	"impeccable",
	"lustrous",
	"magnificent",
	"nifty",
	"polished",
	"pristine",
	"radiant",
	"sharp",
	"slick",
	"smooth",
	"sparkling",
	"lit af",
	"no cap",
	"yeet",
	"gucci",
	"savage",
	"fomo",
	"flex",
	"ghost",
	"bruh",
	"litty",
	"boujee",
	"dank",
	"hype",
	"slay",
	"viral",
	"meme-worthy",
	"epic win",
	"og",
	"thicc",
	"savage af",
	"clout",
	"stan",
	"lit squad",
	"bop",
	"vibe",
	"szn",
	"flexin’",
	"bet",
	"lowkey",
	"highkey",
	"wavy",
	"savage mode",
	"finesse",
	"drip",
	"hundo",
	"squad goals",
	"glow up",
	"fire",
	"banger",
	"turbo",
	"hyper",
	"quantum",
	"galactic",
	"cosmic",
	"ninja",
	"turbocharged",
	"rocket",
	"blazing",
	"thunder",
	"lightning",
	"inferno",
	"phoenix",
	"dragon",
	"titan",
	"vortex",
	"nebula",
	"meteor",
	"aurora",
	"pulse",
	"zenith",
	"eclipse",
	"fusion",
	"horizon",
	"nova",
	"orbit",
	"radiance",
	"surge",
	"tempest",
	"zen",
	"aura",
	"blitz",
	"cascade",
	"dynamo",
	"echo",
	"flux",
	"glide",
	"halo",
	"ignite",
	"jolt",
	"kinetic",
	"luminary",
	"momentum",
	"nebulous",
	"opus",
	"quasar",
	"rift",
	"synergy",
	"torrent",
	"uplift",
	"vertex",
	"wave",
	"xenon",
	"yield",
	"zephyr",
}

// RandomReleaseWord makes no sense, but it's fun
func RandomReleaseWord() string {
	return funnyWords[rand.Intn(len(funnyWords))]
}
