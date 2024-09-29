package pointer

// Dumb af
func PushInt(value int) (int, *int){
	// pointer
	var ptr *int;
	var integer int = value;

	ptr = &value;
	return integer, ptr

}