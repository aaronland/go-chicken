package chicken

var CHICKENS = map[string]string{
	"ara": "دجاج",
	"eng": "chicken",
	"fre": "poule",
	"hbo": "עוף",
	"jpn": "鶏",
	"pol": "kurczak",
	"spa": "pollo",
	"und": "🐔", // https://en.wikipedia.org/wiki/ISO_639-2#Special_situations
	"zxx": "🐔", // https://en.wikipedia.org/wiki/ISO_639-2#Special_situations
}

// https://en.wikipedia.org/wiki/Cross-linguistic_onomatopoeias#Chicken_clucking

var CLUCKING = map[string][]string{
	"eng": []string{"cluck cluck", "bok bok bok", "bok bok b'gawk"},
	"fre": []string{"cot cot cot", "cot cot codet"},
	"hbo": []string{"קוּר־קוּר־קוּר", "בָּק־בָּק־בָּק"},
	"jpn": []string{"コッコッ"},
	"pol": []string{"ko ko"},
	"spa": []string{"kikiriki"},
}
