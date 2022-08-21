package company

// Worker — интерфейс работника компании
type Worker interface {
	// всё, что он должен уметь делать, — это работать
	Work(tasks []string) string
}

// Company — структура компании
type Company struct {
	// personal — сотрудники компании
	// обратите внимание, мы создали слайс сотрудников компании, то есть слайс переменных интерфейсного типа Worker
	personal []Worker
}

var c Company

// Hire — наём нового сотрудника
// Сотрудник может быть любого типа: человек, робот или сторожевая собака. Главное, чтобы он умел работать, то есть удовлетворял интерфейсу Worker
// Go ещё на этапе компиляции проверяет, соответствует ли интерфейсу переданная переменная
func (с *Company) Hire(newbie Worker) {
	c.personal = append(c.personal, newbie)
}

// Process — работа конкретного сотрудника
//func (с Company) Process(id int, tasks []string) (res string) {
//return c.children[id].Work(tasks)
//}
