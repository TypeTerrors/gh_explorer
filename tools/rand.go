package tools

import (
	"math/rand"
)

// RandomlySelectString randomly selects a string from the provided slice.
func RandomlySelectString() string {
	words := []string{"lol",
		"hehe",
		"omg",
		"haha",
		"huh?",
		"wut?",
		"boo",
		"yup",
		"yolo",
		"rofl",
		"hmm",
		"tasty",
		"cool",
		"bruh",
		"tmi",
		"rip",
		"afk",
		"meh",
		"eww"}
	return words[rand.Intn(len(words))] // Select a random index in the slice.
}

// RandomlySelectString randomly selects a string from the provided slice.
func RadnomnLoadingMessage() string {
	loadingMessage := []string{"Reticulating splines...",
		"Summoning internet fairies...",
		"Herding cats...",
		"Counting backwards from Infinity",
		"Splitting the atom...",
		"Ordering 1s and 0s...",
		"Building a fort out of boxes...",
		"Consulting the oracle...",
		"Warming up the hamsters...",
		"Aligning covfefe particles...",
		"Compressing fish files...",
		"Deciding what's for lunch...",
		"Trying to find the 'Any' key...",
		"Updating the flux capacitor...",
		"Downloading more RAM...",
		"Bribing the guards...",
		"Waiting for the stars to align...",
		"Loading the loading message...",
		"Stretching our pixels...",
		"Bending space-time continuum...",
		"Distracting you with a funny message...",
		"Looking for the plot device...",
		"Spinning the hamster wheel...",
		"Charging the laser beams...",
		"Catching some Z's...",
		"Engaging hyperdrive in 3... 2... 1...",
		"Preparing snarky remarks...",
		"Locating the required gigapixels to render...",
		"Spawning internet trolls...",
		"Wrangling unicorns...",
		"Buffering the buffer...",
		"Digging for digital gold...",
		"Poking clouds with a stick...",
		"Twiddling thumbs in synchronization...",
		"Calibrating the doodads...",
		"Running around in circles...",
		"Staring contest with the server...",
		"Polishing the pixels...",
		"Programming the coffee machine...",
		"Negotiating with time travelers..."}
	// Initialize the random number generator.
	return loadingMessage[rand.Intn(len(loadingMessage))] // Select a random index in the slice.
}
