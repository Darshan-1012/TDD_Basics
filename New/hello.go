package new

const EnglishPrefix = "Hello "
const SpainPrefix = "Hola "
const FrenchPrefix = "Bonjour "
const Spain = "Spain"
const French = "French"

func hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetprefix(language) + name
}

func greetprefix(language string) (prefix string) {
	switch language {
	case Spain:
		prefix = SpainPrefix
	case French:
		prefix = FrenchPrefix
	default:
		prefix = EnglishPrefix
	}
	return
}

/*The main is first step*/
// func main() {
// 	fmt.Println(hello("Darshan"))
// }
