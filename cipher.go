package cipher

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Solver struct {
	Data    []Row
	Message []string
	Keyword []string
	Result  string
}

type Row struct {
	Letter string
	Value  []string
}

func CreateSolver() (Solver, error) {
	s := Solver{}

	file, err := os.Open("data.txt")
	if err != nil {
		return Solver{}, err
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		entries := strings.Split(strings.TrimSpace(scanner.Text()), "")
		fmt.Println("letter")
		//fix this its ugly
		if entries[2] != strings.ToUpper(entries[2]) {
			fmt.Println(entries[2:])
			k := Row{
				Letter: entries[0],
				Value:  entries[2:],
			}
			s.Data = append(s.Data, k)
		}
	}

	err = file.Close()
	if err != nil {
		return Solver{}, err
	}

	//print the thing so its easy to see
	for i := 1; i < len(s.Data); i++ {
		fmt.Println(s.Data[i])
	}

	return s, nil
}

func createKeyword(word string, requiredLength int) []string {
	wordAsArray := strings.Split(word, "")
	newArray := []string{}
	k := 0

	for i := 1; i < requiredLength+1; i++ {
		newArray = append(newArray, wordAsArray[k])
		if k == len(word)-1 {
			k = 0
		} else {
			k++
		}
	}
	return newArray
}

func (s Solver) Encode(message string, keyword string) string {
	s.Message = strings.Split(message, "")
	s.Keyword = createKeyword(keyword, len(s.Message))

	answer := []string{}
	for i, k := range s.Message {
		//	look up column s (i.e. k here) and then row m (i.e. keyword[i])
		letterFromMessage := GetIndexOfLetter(k)
		letter := s.GetElement(letterFromMessage, s.Keyword[i])
		answer = append(answer, letter)
	}
	return strings.Join(answer, "")
}

func GetIndexOfLetter(letter string) int {
	return strings.Index("ABCDEFGHIJKLMNOPQRSTUVWXYZ", strings.ToUpper(letter))
}

func (s Solver) GetElement(letterLocation int, letter string) string {
	for _, k := range s.Data {
		if k.Letter == strings.ToUpper(letter) {
			return k.Value[letterLocation]
		}
	}
	return ""
}
