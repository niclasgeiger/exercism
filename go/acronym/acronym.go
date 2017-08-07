package acronym
import(
	"strings"
	"regexp"
)
const testVersion = 2

func Abbreviate(s string) string {
	//replace "-" with " "
	re := regexp.MustCompile("[-]")
	word := re.ReplaceAllString(s," ")
	//create space for uppercase following lower case (HyperText = Hyper Text)
	re = regexp.MustCompile("([a-z]+)([A-Z]+)")
	word = re.ReplaceAllString(word," $2")
	//split string for spaces and use first letters
	words := strings.Fields(word)
	acro := ""
	for i := 0; i<len(words);i++ {
		acro += words[i][:1]
	}
	//convert acronym to upper case
	acro = strings.ToUpper(acro)
	return acro

}
