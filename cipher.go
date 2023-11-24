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
}

type Row struct {
	Letter string
	Value  []string
}

func CreateSolver(message string, keyword string) (Solver, error) {
	s := Solver{}
	s.Message = strings.Split(message, "")
	s.Keyword = strings.Split(keyword, "")

	file, err := os.Open("data.txt")
	if err != nil {
		return Solver{}, err
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		entries := strings.Split(strings.TrimSpace(scanner.Text()), "")
		k := Row{
			Letter: entries[0],
			Value:  entries[2:],
		}
		s.Data = append(s.Data, k)
	}

	err = file.Close()
	if err != nil {
		return Solver{}, err
	}

	for i := 1; i < len(s.Data); i++ {
		fmt.Println(s.Data[i])
	}

	return s, nil
}

func (s Solver) Encode() []string {
	answer := []string{}
	for i, k := range s.Message {
		fmt.Println("letter in message")
		fmt.Println(k)
		fmt.Println("letter in keyword")
		fmt.Println(s.Keyword[i])
		//	look up column s (i.e. k here) and then row m (i.e. keyword[i])
		letterFromMessage := GetIndexOfLetter(k)

		letter := s.GetElement(letterFromMessage, s.Keyword[i])
		answer = append(answer, letter)
		return answer
	}
	return s.Keyword
}

func GetIndexOfLetter(letter string) int {
	return strings.Index("ABCDEFGHIJKLMNOPQRSTUVWXYZ", strings.ToUpper(letter))
}

func (s Solver) GetElement(letterLocation int, letter string) string {
	fmt.Println("letter")
	fmt.Println(letter)
	for _, k := range s.Data {
		if k.Letter == strings.ToUpper(letter) {
			return k.Value[letterLocation]
		}
	}
	return ""
}
