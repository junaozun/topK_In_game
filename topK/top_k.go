package topK

import "container/heap"
/*
  需求：
  1.维护战斗力前k的小根堆
  2.堆中同一武将的数量不能超过3个
*/
type TopCommander struct {
	K int
	CommanderCountMap map[int]int //堆中武将ID：数量
	CommanderPowerHeap
}

//小根堆
type CommanderPowerHeap []*CommanderPowerNode

//根据武将战斗力维护小根堆
type CommanderPowerNode struct {
	AssistCommander *AssistCommander
}

func (h *CommanderPowerHeap) Len() int {
	return len(*h)
}

func (h *CommanderPowerHeap) Less(i, j int) bool {
	return (*h)[i].AssistCommander.Power < (*h)[j].AssistCommander.Power
}

func (h *CommanderPowerHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *CommanderPowerHeap) Push(x interface{}) {
	data := x.(*CommanderPowerNode)
	if data != nil{
		*h = append(*h, data)
	}
}

func (h *CommanderPowerHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = nil //避免内存泄漏
	*h = old[0:n-1]
	return x
}

func NewTopMessage(k int) *TopCommander {
	c := &TopCommander{
		K:                  k,
		CommanderCountMap:make(map[int]int),
		CommanderPowerHeap: make(CommanderPowerHeap,0,k),
	}
	return c
}

//找到变化的武将在堆中的index
func (t *TopCommander) Find(commanderID int) (int, *CommanderPowerNode) {
	for i, v := range t.CommanderPowerHeap {
		if v.AssistCommander.CommanderId == commanderID {
			return i, v
		}
	}
	return -1, nil
}

//返回堆顶的武将的战斗力（即堆中最小战斗力）
func (t *TopCommander)MinPower() int {
	if t.Len() > 0{
		return t.CommanderPowerHeap[0].AssistCommander.Power
	}
	return 0
}

//向堆中添加武将
func (t *TopCommander) AddCommander(commander *CommanderPowerNode) bool {
	if t.Len() >= t.K{ //堆满
		if commander.AssistCommander.Power >= t.MinPower(){//该武将战斗力大于堆顶武将战斗力
			t.Update(t.CommanderPowerHeap[0], commander)
			return true
		}else {//该武将战斗力小于堆顶武将战斗力
			return false
		}
	}
	//堆未满
	heap.Push(&t.CommanderPowerHeap,commander)
	t.CommanderCountMap[commander.AssistCommander.CommanderId]++
	return true
}

//实际需求：B玩家的关羽战斗力大于堆中A玩家的关羽的战斗力，那么更新这个关羽的节点信息
//更新武将战斗力
func(t *TopCommander) Update (daityNode *CommanderPowerNode,newdataNode *CommanderPowerNode)  {
	//更新节点，map动态维护
	t.CommanderCountMap[daityNode.AssistCommander.CommanderId]--
	t.CommanderCountMap[newdataNode.AssistCommander.CommanderId]++
	index,oldCommander := t.Find(daityNode.AssistCommander.CommanderId)
	oldCommander.AssistCommander = newdataNode.AssistCommander
	heap.Fix(t,index)
}

//找到堆中commanderID相同，战斗力最低的那个武将节点
func (t *TopCommander)FindMinxPowerWithSameCommanderID(commanderID int) *CommanderPowerNode {
	var tempData *CommanderPowerNode
	for _, v := range t.CommanderPowerHeap {
		if v.AssistCommander.CommanderId != commanderID {
			continue
		}

		if tempData == nil {
			tempData = v
			continue
		}

		if v.AssistCommander.Power < tempData.AssistCommander.Power {
			tempData = v
		}
	}
	return tempData
}
