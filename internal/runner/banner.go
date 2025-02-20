package runner

import (
	"fmt"

	"github.com/zan8in/afrog/pkg/config"
	"github.com/zan8in/afrog/pkg/log"
	"github.com/zan8in/afrog/pkg/upgrade"
	"github.com/zan8in/afrog/pkg/utils"
	"github.com/zan8in/gologger"
)

var banner = fmt.Sprintf(`
      _/_/    _/_/_/_/  _/_/_/      _/_/      _/_/_/   
   _/    _/  _/        _/    _/  _/    _/  _/          
  _/_/_/_/  _/_/_/    _/_/_/    _/    _/  _/  _/_/     
 _/    _/  _/        _/    _/  _/    _/  _/    _/      
_/    _/  _/        _/    _/    _/_/      _/_/_/  %s
`, config.Version)

func ShowBanner() {
	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msgf("\t\t\tLost in your uniqueness even though we are all ordinary.\n\n")
}

func ShowBanner2(upgrade *upgrade.Upgrade) {
	messageStr := ""
	if utils.Compare(upgrade.LastestAfrogVersion, ">", config.Version) {
		messageStr = " (" + log.LogColor.Red(upgrade.LastestAfrogVersion) + ")"
	} else {
		messageStr = " (" + log.LogColor.Green("latest") + ")"
	}
	gologger.Info().Msgf("Using afrog Engine %s%s", config.Version, messageStr)

	messageStr2 := ""
	if utils.Compare(upgrade.LastestVersion, ">", upgrade.CurrVersion) {
		messageStr2 = " (" + log.LogColor.Red(upgrade.LastestVersion) + ")"
	} else {
		messageStr2 = " (" + log.LogColor.Green("latest") + ")"
	}
	gologger.Info().Msgf("Using afrog-pocs %s%s", upgrade.CurrVersion, messageStr2)
}

// func ShowUsage() string {
// 	return "\nUSAGE:\n   afrog -t example.com -o result.html\n   afrog -T urls.txt -o result.html\n   afrog -T urls.txt -s -o result.html\n   afrog -t example.com -P ./pocs/poc-test.yaml -o result.html\n   afrog -t example.com -P ./pocs/ -o result.html\n"
// }

// func ShowTips() string {
// 	return "\nTIPS:\n   " + utils.GetRandomTips() + "\n"
// }
