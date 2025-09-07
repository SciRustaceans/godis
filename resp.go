const (
	STRING  = '+'
	ERROR   = '-'
	INTEGER = ':'
	BULK    = "$"
	ARRAY   = "*"
)

type Value struct {
	typ 	string // Determines the data type carried by value 
	str 	string // holds the value of the string recieve from the simple strings 
	num 	int  // holds the value of the integer received from the integers 
	bulk 	string // is usedto store the sring recieved from the bulk string 
	array 	[]Value // holds all the values recieved from the arrays
}

// Reader method to contain all methods that will help us read from the buffer and store it in the Value struct

type Resp struct {
	reader *bufio.Reader
}

func NewResp(rd io.Reader) *Resp {
	return &Resp{reader: bufio.NewReader(rd)}
}

//Reads line from the buffer 
// method: Readline

func (r*Resp) readline() (line []byte, n int, err error) {
	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}
		n += 1
		line = append(line, b)
		if len(line) >= 2 && line[len(line) -2] == '\r' {
			break
		}
	}
	return line[:len(line)-2], n, nil
}
