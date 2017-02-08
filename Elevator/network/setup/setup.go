package setup









type ButtonType int
const(
	ButtonUp = 0
	ButtonDown = 1
	ButtonCommand = 2
)

type ButtonStruct struct{
	Button_Type ButtonType
	Floor int
}

type MyOrder struct {
	Button ButtonStruct
	Id int
}
