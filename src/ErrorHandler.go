package main

func handleError(errorCode [2]byte, err error) int {
	switch errorCode[0] {
	case 0x00: // non fatal
		switch errorCode[1] {
		case 0x01:
			warningLog.Printf("Initializing warning, %v. Kara continue to work\n", err)
		case 0x02:
			warningLog.Printf("Reading response warning, %v. Kara continue to work\n", err)
		default:
			warningLog.Printf("Unexpected minor error code %v, Kara continue to work\n", int(errorCode[0]))
		}
		return 1
	case 0x01: // fatal
		switch errorCode[1] {
		case 0x01:
			errorLog.Printf("Initializing error, %v, Kara is shutting down\n", err)
		case 0x02:
			errorLog.Printf("Initializing error, %v, Kara is shutting down\n", err)
		case 0x03:
			errorLog.Printf("Creating request error, %v, Kara is shutting down\n", err)
		default:
			errorLog.Printf("Unexpected minor code %v, Kara is shutting down\n", int(errorCode[0]))
		}
		return 2
	default:
		errorLog.Printf("Unexpected senior error code %v, Kara is shutting down\n", int(errorCode[0]))
		return 2
	}
}
