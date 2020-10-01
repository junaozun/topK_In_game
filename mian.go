package main

import (
	"fmt"
	"top_k_commander/topK"
)

func main()  {
/*------------------------------------------------------------test----------------------------------------------------------------------*/
	//PlayerId、CommanderId、Power
	initCommander := [][]int{{12,1,99333},{13,1,453464},{22,1,453459},{10,13,5445},{43,565,744},{34,1,676567}}

	topPowerHeap := topK.NewTopMessage(4)

	for _,arr := range initCommander {
		commander := &topK.CommanderPowerNode{
			AssistCommander:&topK.AssistCommander{
				PlayerId:    arr[0],
				CommanderId: arr[1],
				Power:       arr[2],
			},
		}
		//入堆
		//堆中相同武将的数量不能超过3个，若堆中已有3个相同的武将（关羽），则新入堆的武将也是关羽，且战力高于3个中最小的，则替换那个最小的武将
		if count,ok := topPowerHeap.CommanderCountMap[commander.AssistCommander.CommanderId];ok && count >= 3 {
			//找到commanderID相同，战斗力最小的节点
			minPowerNode := topPowerHeap.FindMinxPowerWithSameCommanderID(commander.AssistCommander.CommanderId)
			//更新堆
			topPowerHeap.Update(minPowerNode,commander)
		}else {
			topPowerHeap.AddCommander(commander)
		}

	}

	for _,commander := range topPowerHeap.CommanderPowerHeap {
		fmt.Printf("%v %v %v\t",commander.AssistCommander.PlayerId,commander.AssistCommander.CommanderId,commander.AssistCommander.Power)
	}

}
