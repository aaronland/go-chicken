package chicken

var CHICKENS = map[string]string{
	"ara": "فروج",
	"eng": "chicken",
	"fre": "poule",
	"hbo": "עוף",
	"jpn": "鶏",
	"kor": "치킨",
	"pol": "kurczak",
	"spa": "pollo",
	"und": "🐔", // https://en.wikipedia.org/wiki/ISO_639-2#Special_situations
	"zxx": "🐔", // https://en.wikipedia.org/wiki/ISO_639-2#Special_situations
}

// https://en.wikipedia.org/wiki/Cross-linguistic_onomatopoeias#Chicken_clucking

var CLUCKING = map[string][]string{
	"ara": []string{"قرقر", "بق بق بيق"},
	"eng": []string{"cluck cluck", "bok bok bok", "bok bok b'gawk"},
	"fre": []string{"cot cot cot", "cot cot codet"},
	"hbo": []string{"קוּר־קוּר־קוּר", "בָּק־בָּק־בָּק"},
	"jpn": []string{"コッコッ"},
	"kor": []string{"꼬꼬댁"},
	"pol": []string{"ko ko"},
	"spa": []string{"kikiriki"},
}
