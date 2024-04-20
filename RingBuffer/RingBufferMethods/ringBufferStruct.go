package ringbuffer

type ringBuffer struct {
    array []int
    head, tail ,bufferSize int
}

//constructor
func RingBuffer(buffersize int) *ringBuffer {
    return &ringBuffer{
      array:make([]int,buffersize), 
      head:0, 
      tail:0, 
      bufferSize:buffersize,
    }
}
