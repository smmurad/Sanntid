package hardware

import(
	"../setup"
	"time"
)

func ReadButtons(ch_button_polling chan<- setup.ButtonStruct){
	var last_pushed[setup.NUM_FLOORS][setup.NUM_BUTTONS] int
	var button setup.ButtonType
	for{
		for button=setup.ButtonUp; button <= setup.ButtonCommand; button++{
			for floor:=0; floor<setup.NUM_FLOORS; floor++{
				if GetButtonSignal(button, floor) == 1{
					if (last_pushed[floor][button] == 0){
						ch_button_polling <- setup.ButtonStruct{Button_Type : button , Floor : floor}
					}
					last_pushed[floor][button] = 1
				} else {
					last_pushed[floor][button] = 0
				}
			}
		}
		time.Sleep(100*time.Millisecond)
	}
}