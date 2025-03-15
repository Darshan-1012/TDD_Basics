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
	if language == Spain {
		return SpainPrefix + name
	}
	if language == French {
		return FrenchPrefix + name
	}
	return EnglishPrefix + name
}

/*The main is first step*/
// func main() {
// 	fmt.Println(hello("Darshan"))
// }
