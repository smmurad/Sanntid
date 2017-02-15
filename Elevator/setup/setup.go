package setup

const NUM_FLOORS 	= 4
const NUM_BUTTONS 	= 3

type MotorDir int
const(
	DIR_DOWN 	= -1
	DIR_UP 		= 1
	DIR_STOP 	= 0	
)

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

