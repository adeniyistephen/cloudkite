package business

import "log"

// NewCloudKiteService manages the set of API's for Passenger access.
type NewCloudKiteService struct {
	log *log.Logger
}

// New constructs a CloudKiteservice for api access.
func New(log *log.Logger) NewCloudKiteService {
	return NewCloudKiteService{log}
}

type Message struct {
	Message string `json:"message"`
}

// GetVowels returns the number of vowels in a string.
func (nc NewCloudKiteService) ReverseVowels(m Message) (Message, error) {
	resverstring := m.Message
	b := []byte(resverstring)
	for i, j := 0, len(b)-1; i < j; {
		if IsVowels(b[i]) && IsVowels(b[j]) {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		} else if IsVowels(b[i]) && !IsVowels(b[j]) {
			j--
		} else if !IsVowels(b[i]) && IsVowels(b[j]) {
			i++
		} else {
			i++
			j--
		}
	}
	nc.log.Printf("%s", "vowels reversed:" + string(b))
	return Message{string(b)}, nil
}

// Hello returns a string of world when hit.
func (nc NewCloudKiteService) Hello() (string, error) {
	nc.log.Printf("%s", "hello world")
	return "world", nil
}

// IsVowels returns true if the rune is a vowel.
func IsVowels(s byte) bool {
	if s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' || s == 'A' || s == 'E' || s == 'I' || s == 'O' || s == 'U' {
		return true
	}
	return false
}
