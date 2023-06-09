package main

import (
	_ "embed"
	"fmt"
	"os/exec"

	"github.com/getlantern/systray"
)

//go:embed image/icon.png
var icon []byte

func main() {
	onExit := func() {}
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTemplateIcon(icon, icon)
	go func() {
		wechat := systray.AddMenuItem("微信", "微信多开")
		qq := systray.AddMenuItem("QQ", "QQ多开")
		dd := systray.AddMenuItem("钉钉", "钉钉多开")
		systray.AddSeparator()
		quit := systray.AddMenuItem("Quit", "Quit the whole app")
		for {
			select {
			case <-wechat.ClickedCh:
				go func() {
					cmd := exec.Command("nohup", "/Applications/WeChat.app/Contents/MacOS/WeChat", ">", "/dev/null", "2>&1")
					err := cmd.Run()
					if err != nil {
						fmt.Println(err)
					}
				}()
			case <-qq.ClickedCh:
				go func() {
					cmd := exec.Command("nohup", "/Applications/QQ.app/Contents/MacOS/QQ", ">", "/dev/null", "2>&1")
					err := cmd.Run()
					if err != nil {
						fmt.Println(err)
					}
				}()
			case <-dd.ClickedCh:
				go func() {
					cmd := exec.Command("nohup", "/Applications/DingTalk.app/Contents/MacOS/DingTalk", ">", "/dev/null", "2>&1")
					err := cmd.Run()
					if err != nil {
						fmt.Println(err)
					}
				}()
			case <-quit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}
