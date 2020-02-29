package main

//Stack represents LIFO data strucutre
type Stack struct {
	list []int
}

/* func main() {
	st := Stack{
		list: []int{},
	}

	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)

	fmt.Println("stack size :", st.Size())
	fmt.Println("stack elements :", st.list)

	st.Pop()
	st.Pop()
	st.Pop()
	if st.IsEmpty() {
		fmt.Println("Stack is empty")
	}
	st.Pop()
	if st.IsEmpty() {
		fmt.Println("Stack is empty")
	}
	fmt.Println("stack elements after first POP :", st.list)

} */

//Size retruns size of stack
func (s *Stack) Size() int {
	return len(s.list)
}

//Push will add the elements in stack
func (s *Stack) Push(i int) {
	s.list = append(s.list, i)
}

//Pop removes element from stack
func (s *Stack) Pop() {
	s.list = s.list[:s.Size()-1]
}

//IsEmpty returns boolean true if stack is empty otherwise false
func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}
