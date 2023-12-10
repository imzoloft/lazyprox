package inout

import (
	"fmt"
	"os"
	"time"

	"github.com/imzoloft/lazyprox/common"
)

func DisplayBanner() {
	fmt.Printf("%s%s%s", common.TextBlue, common.BANNER, common.TextReset)
}

func DisplayAuthor() {
	fmt.Printf("%s%s%s\n\n", common.TextBlue, common.AUTHOR, common.TextReset)
}

func DisplayUsage() {
	fmt.Printf("%s\n", common.USAGE)
}

func DisplayVersion() {
	fmt.Printf("lazyprox version %s\n", common.VERSION)
	os.Exit(0)
}

func DisplayMessage(color string, message string) {
	fmt.Printf("[%s!%s] %s\n", color, common.TextReset, message)
}

func DisplayOverview() {
	DisplayMessage(
		common.TextGreen,
		fmt.Sprintf("Validated %d proxies in %f seconds | Working %d proxies | Dead %d proxies",
			common.Stats.ValidatedProxy,
			GetTimeElapsed(),
			common.Stats.WorkingProxy,
			common.Stats.DeadProxy))
}

func FatalError(errorProp string) {
	fmt.Fprintf(os.Stderr, "[%s!%s] %sERROR%s | %s\n", common.TextRed, common.TextReset, common.TextRed, common.TextReset, errorProp)
	os.Exit(1)
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func GetTimeElapsed() float64 {
	finishTime := time.Now()
	timeElapsed := finishTime.Sub(common.Opts.StartTime)

	return timeElapsed.Seconds()
}
