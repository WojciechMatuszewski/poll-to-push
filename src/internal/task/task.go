package task

type Item struct {
	ID string `json:"id"`
	Input
}

type Input struct {
	Name string `json:"name"`
}

func (i Input) IsValid() bool {
	return i.Name != ""
}

type Result struct {
	Greeting string `json:"greeting"`
}

type Connection struct {
	ID     string `json:"id"`
	TaskID string `json:"taskID"`
}
