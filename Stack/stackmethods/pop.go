package stackmethods

import (
  "math"
)

func (s *stack) Pop() int {
    //if you try to pop an empty stack, this throws a runtime error
    s.Length = int( math.Max(0, float64(s.Length-1)) )
    if( s.Length == 0 ) {
        top := s.head
        s.head = nil
        return top.value
    }

    //use this instead. uncomment the below, and comment the above
    /*
    if(s.head == nil) {
        println("nothing to pop")
        return 0
    }
    */
    
    top := s.head
    s.head = s.head.prev
    //s.Length--

    return top.value
}
