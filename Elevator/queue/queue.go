package queue

import("../setup")

var Queue [setup.NUM_FLOORS][setup.NUM_BUTTONS][setup.MAX_NUM_ELEVATORS] int

func GetOptimalElevator() int{
	return 0
}


func IsNewOrder(order setup.MyOrder) int{
	if order.Button.Button_Type == setup.ButtonCommand{
		for i := 0; i<setup.MAX_NUM_ELEVATORS; i++{
			if Queue[order.Button.Floor][order.Button.Button_Type][i] == order.Id{
				return 0
			}
		}
	}else{
		if Queue[order.Button.Floor][order.Button.Button_Type][0] != -1{
				return 0
			}
	}
	return 1
}

func AddOrder(newOrder setup.MyOrder){
	if newOrder.Button.Button_Type == setup.ButtonCommand{
		for i:=0; i<setup.MAX_NUM_ELEVATORS;i++{
			if Queue[newOrder.Button.Floor][newOrder.Button.Button_Type][i] == -1{
				Queue[newOrder.Button.Floor][newOrder.Button.Button_Type][i] = newOrder.Id
			}
		}
	}else{
		Queue[newOrder.Button.Floor][newOrder.Button.Button_Type][0] = newOrder.Id
	}
}