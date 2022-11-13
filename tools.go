package casdmwsgo

// NewInt is a wrapper to create inline pointer to int
func NewInt(n int) *int {
	return &n
}

// NewString is a wrapper to create inline pointer to string
func NewString(s string) *string {
	return &s
}

// ArrToAOS is a wrapper to create inline pointer to Array of String
func ArrToAOS(arr []string) *ArrayOfString {
	arrToPTR := make([]*string, 0, len(arr))
	for k := range arr {
		arrToPTR = append(arrToPTR, &arr[k])
	}
	return &ArrayOfString{
		String: arrToPTR,
	}
}

// ArrToAOI is a wrapper to create inline pointer to Array of Int
func ArrToAOI(arr []int) *ArrayOfInt {
	arrToPTR := make([]*int, 0, len(arr))
	for k := range arr {
		arrToPTR = append(arrToPTR, &arr[k])
	}
	return &ArrayOfInt{
		Integer: arrToPTR,
	}
}
