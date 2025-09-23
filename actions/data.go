package action

type Task struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Actions interface {
	Add(string)
	Update(int, string)
	Delete(int)
	MarkProgress(int)
	MarkDone(int)
	List()
}
