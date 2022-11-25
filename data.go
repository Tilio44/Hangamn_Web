package hangman

type Loop struct {
	Letter  string
	Visible bool
}
type Sub struct {
	Letter       string
	Text         string
	Keyboard1    []Loop
	NumOfBButton []interface{}
}

var data Sub

func button() {
	for i := 0; i < 10; i++ {
		data.NumOfBButton = append(data.NumOfBButton, 0)
	}
}
