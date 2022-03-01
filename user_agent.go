package faker

import "strconv"

// UserAgent will generate a random broswer user agent
func UserAgent() string {
	randNum := randIntRange(0, 3)
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
	var ver = []string{"81.0.4044.138","81.0.4044.129","81.0.4044.122", "81.0.4044.92", "80.0.3987.163", "80.0.3987.149", "80.0.3987.132", "80.0.3987.122", "80.0.3987.116", "80.0.3987.87", "79.0.3945.130", "79.0.3945.117", "79.0.3945.88", "79.0.3945.79", "78.0.3904.108", "78.0.3904.97", "78.0.3904.87", "78.0.3904.70", "77.0.3865.120", "77.0.3865.90", "77.0.3865.75", "76.0.3809.132", "76.0.3809.100", "75.0.3770.142", "75.0.3770.100", "75.0.3770.90", "75.0.3770.80"}
	return "Mozilla/5.0 " + "(" + randomPlatform() + ") AppleWebKit/537.36 (KHTML, like Gecko) Chrome/" + RandString(ver)+ " Safari/537.36"
}

// FirefoxUserAgent will generate a random firefox broswer user agent string
func FirefoxUserAgent() string {
	rv := strconv.Itoa(randIntRange(55, 76))
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
	randNum := strconv.Itoa(randIntRange(531, 536)) + "." + strconv.Itoa(randIntRange(1, 51)) + "." + strconv.Itoa(randIntRange(1, 8))
	ver := strconv.Itoa(randIntRange(4, 6)) + "." + strconv.Itoa(randIntRange(0, 2))

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
	return "Macintosh; " + getRandValue([]string{"computer", "mac_processor"}) + " Mac OS X 10_" + strconv.Itoa(randIntRange(5, 15)) + "_" + strconv.Itoa(randIntRange(1, 3))
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
