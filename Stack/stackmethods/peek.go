package stackmethods

func (s *stack) Peek() int {
    if(s.head != nil) {
        return s.head.value 
    }
    return 0
}
