package fsm

import (
	"../queue"
	"../setup"
	"fmt"
	)

func FsmOrderHandler(ch_button_polling <-chan setup.ButtonStruct){
	var newOrder setup.MyOrder
	var button setup.ButtonStruct
	Id := 0
	for{
		select {
			case button = <-ch_button_polling:
				fmt.Println(button)
		}
		if button.Button_Type == setup.ButtonCommand{
			//Using ip-adress as id
			Id = 0
		}else{
			Id = queue.GetOptimalElevator() // Taking order as argument
		}
		newOrder = setup.MyOrder{Button:button, Id:Id}

		if queue.IsNewOrder(newOrder) == 1{
			fmt.Println("isnew")
			queue.AddOrder(newOrder)

			fmt.Println("Added new order, new queue:")
			fmt.Println(queue.Queue)
		}
	}
}