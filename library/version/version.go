package version

import "fmt"

func ShowLogo(buildVersion, buildTime, commitID string) {
	//版本号
	fmt.Println("     _______.     ___       _______   ______     ______   \n    /       |    /   \\     /  _____| /  __  \\   /  __  \\  \n   |   (----`   /  ^  \\   |  |  __  |  |  |  | |  |  |  | \n    \\   \\      /  /_\\  \\  |  | |_ | |  |  |  | |  |  |  | \n.----)   |    /  _____  \\ |  |__| | |  `--'  | |  `--'  | \n|_______/    /__/     \\__\\ \\______|  \\______/   \\______/  \n                                                          ")
	fmt.Println("Version   ：", buildVersion)
	fmt.Println("BuildTime ：", buildTime)
	fmt.Println("CommitID  ：", commitID)
	fmt.Println("")

}
