package structs

type Score struct {
	SubjectName string
	FullScore   float32
	Score       float32

	Student *Student
}
