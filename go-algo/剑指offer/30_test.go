package main

type MinStack struct {
	stack []int
	mins  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		mins:  []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	oldMin := this.mins[len(this.stack)-1]
	min := min(oldMin, x)
	this.mins = append(this.mins, min)
}

func min(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.mins = this.mins[:len(this.mins)-1]
}

func (this *MinStack) Top() int {
	return this.mins[len(this.mins)-1]
}

func (this *MinStack) Min() int {
	return this.mins[0]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

func Test(t *testing.T) {
	obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
}