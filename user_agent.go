package faker

import "strconv"

// UserAgent will generate a random broswer user agent
func UserAgent() string {
	randNum := randIntRange(0, 1)
	switch randNum {
	case 0:
		return ChromeUserAgent()
	case 1:
		return FirefoxUserAgent()
	case 2:
		return SafariUserAgent()
	case 3:
		return OperaUserAgent()
	default:
		return ChromeUserAgent()
	}
}

// ChromeUserAgent will generate a random chrome browser user agent string
func ChromeUserAgent() string {
	var ver = []string{"99.0.4844.58", "99.0.4844.48", "98.0.4758.101", "98.0.4758.87", "97.0.4692.99", "97.0.4692.87", "97.0.4692.71", "96.0.4664.110", "96.0.4664.93", "95.0.4638.74", "95.0.4638.54", "94.0.4606.81", "94.0.4606.71", "94.0.4606.61", "94.0.4606.50", "94.0.4606.61", "93.0.4577.82", "93.0.4577.63", "92.0.4515.159", "92.0.4515.131", "92.0.4515.107", "91.0.4472.164", "91.0.4472.124", "91.0.4472.114", "91.0.4472.101", "91.0.4472.77", "90.0.4430.212", "90.0.4430.85", "90.0.4430.72", "89.0.4389.114", "89.0.4389.90", "89.0.4389.82", "87.0.4280.141", "87.0.4280.88", "87.0.4280.66", "86.0.4240.198", "86.0.4240.111", "86.0.4240.75", "85.0.4183.121", "85.0.4183.102", "85.0.4183.83", "84.0.4147.135", "84.0.4147.125", "84.0.4147.105", "84.0.4147.89", "83.0.4103.106", "83.0.4103.97", "83.0.4103.61", "81.0.4044.138","81.0.4044.129","81.0.4044.122", "81.0.4044.92", "80.0.3987.163", "80.0.3987.149", "80.0.3987.132", "80.0.3987.122", "80.0.3987.116", "80.0.3987.87"}
	return "Mozilla/5.0 " + "(" + randomPlatform() + ") AppleWebKit/537.36 (KHTML, like Gecko) Chrome/" + RandString(ver)+ " Safari/537.36"
}

// FirefoxUserAgent will generate a random firefox broswer user agent string
func FirefoxUserAgent() string {
	rv := strconv.Itoa(randIntRange(82, 96))
	ver := "Gecko/20100101 Firefox/" + rv + ".0"
	platforms := []string{
		"(" + windowsPlatformToken() + "; rv:" + rv + ".0) " + ver,
		"(" + linuxPlatformToken() + "; rv:" + rv + ".0) " + ver,
		"(" + macPlatformToken() + "; rv:" + rv + ".0) " + ver,
	}

	return "Mozilla/5.0 " + RandString(platforms)
}

// SafariUserAgent will generate a random safari browser user agent string
func SafariUserAgent() string {
	randNum := strconv.Itoa(randIntRange(531, 604)) + "." + strconv.Itoa(randIntRange(1, 2)) + "." + strconv.Itoa(randIntRange(1, 10))
	ver := strconv.Itoa(randIntRange(9, 12)) + "." + strconv.Itoa(randIntRange(0, 2))

	mobileDevices := []string{
		"iPhone; CPU iPhone OS",
		"iPad; CPU OS",
	}

	platforms := []string{
		"(Windows; U; " + windowsPlatformToken() + ") AppleWebKit/" + randNum + " (KHTML, like Gecko) Version/" + ver + " Safari/" + randNum,
		"(" + macPlatformToken() + " rv:" + strconv.Itoa(randIntRange(4, 7)) + ".0; en-US) AppleWebKit/" + randNum + " (KHTML, like Gecko) Version/" + ver + " Safari/" + randNum,
		"(" + RandString(mobileDevices) + " " + strconv.Itoa(randIntRange(7, 9)) + "_" + strconv.Itoa(randIntRange(0, 3)) + "_" + strconv.Itoa(randIntRange(1, 3)) + " like Mac OS X; " + "en-US" + ") AppleWebKit/" + randNum + " (KHTML, like Gecko) Version/" + strconv.Itoa(randIntRange(3, 5)) + ".0.5 Mobile/8B" + strconv.Itoa(randIntRange(111, 120)) + " Safari/6" + randNum,
	}

	return "Mozilla/5.0 " + RandString(platforms)
}

// OperaUserAgent will generate a random opera browser user agent string
func OperaUserAgent() string {
	platform := "(" + randomPlatform() + "; en-US) Presto/2." + strconv.Itoa(randIntRange(8, 13)) + "." + strconv.Itoa(randIntRange(160, 355)) + " Version/" + strconv.Itoa(randIntRange(10, 13)) + ".00"

	return "Opera/" + strconv.Itoa(randIntRange(8, 10)) + "." + strconv.Itoa(randIntRange(10, 99)) + " " + platform
}

// linuxPlatformToken will generate a random linux platform
func linuxPlatformToken() string {
	return "X11; Linux " + getRandValue([]string{"computer", "linux_processor"})
}

// macPlatformToken will generate a random mac platform
func macPlatformToken() string {
	return "Macintosh; " + getRandValue([]string{"computer", "mac_processor"}) + " Mac OS X 10_" + strconv.Itoa(randIntRange(8, 15)) + "_" + strconv.Itoa(randIntRange(1, 5))
}

// windowsPlatformToken will generate a random windows platform
func windowsPlatformToken() string {
	return getRandValue([]string{"computer", "windows_platform"}) + getRandValue([]string{"computer", "windows_version"})
}

// randomPlatform will generate a random platform
func randomPlatform() string {
	platforms := []string{
		linuxPlatformToken(),
		macPlatformToken(),
		windowsPlatformToken(),
	}

	return RandString(platforms)
}
