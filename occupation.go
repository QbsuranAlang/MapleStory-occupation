package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	math_rand "math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"unicode"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

const (
	title         = "楓之谷職業隨機選擇器"
	maxOcuupation = 15
	maxDpX        = 800
	maxDpY        = 800
	version       = "v0.0.1"
	reachMax      = "數量太多啦"
)

var (
	o = []string{
		"江湖 - 墨玄",
		"超新星 - 天使破壞者",
		"江湖 - 琳恩",
		"雷普族 - 卡莉",
		"皇家騎士團 - 聖魂劍士",
		"皇家騎士團 - 烈焰巫師",
		"皇家騎士團 - 破風使者",
		"皇家騎士團 - 暗夜行者",
		"皇家騎士團 - 閃雷悍將",
		"皇家騎士團長 - 米哈逸",
		"冒險家 - 劍士 英雄",
		"冒險家 - 劍士 聖騎士",
		"冒險家 - 劍士 黑騎士",
		"冒險家 - 法師 冰雷大魔導士",
		"冒險家 - 法師 火毒大魔導士",
		"冒險家 - 法師 主教",
		"冒險家 - 弓箭手 箭神",
		"冒險家 - 弓箭手 神射手",
		"冒險家 - 弓箭手 開拓者",
		"冒險家 - 盜賊 夜使者",
		"冒險家 - 盜賊 暗影神偷",
		"冒險家 - 盜賊 影武者",
		"冒險家 - 海盜 拳霸",
		"冒險家 - 海盜 槍神",
		"冒險家 - 海盜 重砲指揮官",
		"阿尼瑪 - 菈菈",
		"超新星 - 凱殷",
		"雷普族 - 阿戴爾",
		"阿尼瑪 - 虎影",
		"雷普族 - 亞克",
		"雷普族 - 伊利恩",
		"超新星 - 卡蒂娜",
		"英雄團 - 狂狼勇士",
		"英雄團 - 龍魔導士",
		"英雄團 - 精靈遊俠",
		"英雄團 - 幻影俠盜",
		"英雄團 - 夜光",
		"英雄團 - 隱月",
		"末日反抗軍 - 爆拳槍神",
		"末日反抗軍 - 煉獄巫師",
		"末日反抗軍 - 狂豹獵人",
		"末日反抗軍 - 機甲戰神",
		"末日反抗軍支援者 - 惡魔 - 惡魔復仇者",
		"末日反抗軍支援者 - 惡魔 - 惡魔殺手",
		"末日反抗軍支援者 - 傑諾",
		"超新星 - 凱撒",
		"朋友世界 - 凱內西斯",
		"超越者 - 神之子",
		"曉之陣 - 劍豪",
		"曉之陣 - 陰陽師",
		"初心者 - 初心者(冒險家)",
		"初心者 - 貴族(皇家騎士團)",
		"初心者 - 市民(末日反抗軍)",
		"初心者 - 傳說(狂狼勇士)",
		"初心者 - ？？？(隱月)",
	} //end slice
	occupationColor = map[string]color.NRGBA{
		"冒險家":      color.NRGBA{R: 244, G: 169, B: 0, A: 255},
		"曉之陣":      color.NRGBA{R: 108, G: 70, B: 117, A: 255},
		"朋友世界":     color.NRGBA{R: 47, G: 53, B: 59, A: 255},
		"末日反抗軍":    color.NRGBA{R: 127, G: 181, B: 181, A: 255},
		"末日反抗軍支援者": color.NRGBA{R: 127, G: 181, B: 181, A: 255},
		"江湖":       color.NRGBA{R: 37, G: 109, B: 123, A: 255},
		"皇家騎士團":    color.NRGBA{R: 199, G: 180, B: 70, A: 255},
		"皇家騎士團長":   color.NRGBA{R: 199, G: 180, B: 70, A: 255},
		"英雄團":      color.NRGBA{R: 65, G: 34, B: 39, A: 255},
		"超新星":      color.NRGBA{R: 164, G: 125, B: 144, A: 255},
		"超越者":      color.NRGBA{R: 204, G: 6, B: 5, A: 255},
		"阿尼瑪":      color.NRGBA{R: 2, G: 86, B: 105, A: 255},
		"雷普族":      color.NRGBA{R: 114, G: 20, B: 34, A: 255},
		"初心者":      color.NRGBA{R: 255, G: 93, B: 115, A: 255},
	}
	loop       = 1
	once       widget.Bool
	blackColor = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	issueColor = color.NRGBA{R: 0, G: 145, B: 147, A: 255}
)

func init() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	} //end if
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
	once.Value = false
} //end init

func getOccupation() []string {

	/* shuffle the occupation slice */
	rand.Shuffle(len(o), func(i, j int) { o[i], o[j] = o[j], o[i] })

	copyO := make([]string, len(o))
	copy(copyO, o)

	output := []string{}
	for i := 0; i < loop; i++ {
		min := 0
		max := len(o) - 1
		randomIndex := rand.Intn(max-min+1) + min
		choose := o[randomIndex]

		output = append(output, choose)

		if !once.Value {
			o = append(o[:randomIndex], o[randomIndex+1:]...)
		} //end if
	} //end for

	o = make([]string, len(copyO))
	copy(o, copyO)

	return output
} //end getOccupation

func openURL(url string) {

	err := (error)(nil)

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = exec.Command("open", url).Start()
	} //end switch

	if err != nil {
		log.Fatal(err)
	} //end if
} //end openURL

func main() {
	/* create new window */
	w := new(app.Window)

	/* handle window events */
	if err := eventLoop(w); err != nil {
		log.Fatal(err)
	} //end if

	app.Main()
} //end main

func eventLoop(w *app.Window) error {
	go func() {
		/* create new window */
		w := new(app.Window)
		w.Option(app.Title(title))
		w.Option(app.Size(unit.Dp(maxDpX), unit.Dp(maxDpY)))

		/* ops are the operations from the UI */
		ops := op.Ops{}

		/* button is a clickable widget */
		btn := widget.Clickable{}
		linkBtn := widget.Clickable{}

		/* th defines the material design style */
		th := material.NewTheme()

		/* create an Editor widget for input */
		numberEditor := widget.Editor{}
		numberEditor.SingleLine = true /* make sure it's a single line input */

		text := getOccupation()
		for {
			/* first grab the event */
			e := w.Event()

			/* then detect the type */
			switch e := e.(type) {
			/* this is sent when the application should re-render. */
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)

				/* get the current text from the editor */
				input := numberEditor.Text()
				if input == "" {
					input = "1"
				} // end if

				if btn.Clicked(gtx) {
					if i, err := strconv.Atoi(input); err != nil {
						loop = 1
						text = getOccupation()
					} else if i > maxOcuupation || i > len(o) {
						text = []string{reachMax}
						loop = 1
					} else {
						loop = i
						text = getOccupation()
					} // end else
				} // end if

				if linkBtn.Clicked(gtx) {
					/* open browser */
					openURL("https://github.com/QbsuranAlang/occupation")
				} //end if

				/* filter the input to only allow numeric characters */
				filteredInput := filterNumbers(input)
				if input != filteredInput {
					numberEditor.SetText(filteredInput)
				} // end if

				/* center layout for the entire content */
				layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{
						Axis:      layout.Vertical,
						Alignment: layout.Middle,
					}.Layout(gtx,
						/* use rigid to add label */
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							flexChildren := []layout.FlexChild{} /* create a slice of flexChild */
							for _, occupation := range text {
								s := strings.Split(occupation, " ")
								color, ok := occupationColor[s[0]]
								if !ok {
									occupation = reachMax
									color = blackColor
								} // end if

								label := material.H3(th, occupation)
								label.Color = color
								label.TextSize = unit.Sp(26)

								/* add Inset to add spacing */
								inset := layout.Inset{
									Top: unit.Dp(2), /* top spacing */
								} // end if

								/* use anonymity function to wrap every label to layout.FlexChild */
								flexChildren = append(flexChildren, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return inset.Layout(gtx, label.Layout)
								}))
							} // end for
							return layout.Flex{
								Axis: layout.Vertical,
							}.Layout(gtx, flexChildren...) // 使用 flexChildren
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							cb := material.CheckBox(th, &once, "職業是否可重複")
							return cb.Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								editor := material.Editor(th, &numberEditor, fmt.Sprintf("請輸入要抽幾個(最多%v位)", maxOcuupation))
								editor.TextSize = unit.Sp(14)
								gtx.Constraints.Max.X = gtx.Dp(300)
								return editor.Layout(gtx)
							})
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return material.Button(th, &btn, "繼續抽").Layout(gtx)
							})
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								inset := layout.Inset{
									Top:    unit.Dp(10),
									Bottom: unit.Dp(10),
								}
								btn := material.Button(th, &linkBtn, "如有問題或建議，點我開issue")
								btn.TextSize = unit.Sp(8)
								btn.Background = issueColor
								return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									return btn.Layout(gtx)
								})
							})
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							versionLabel := material.Body1(th, fmt.Sprintf("版本: %v", version))
							versionLabel.Color = blackColor
							versionLabel.TextSize = unit.Sp(8)
							return layout.Center.Layout(gtx, versionLabel.Layout)
						}),
					)
				})
				e.Frame(gtx.Ops)
			/* and this is sent when the application should exit */
			case app.DestroyEvent:
				os.Exit(0)
			} // end switch
		} // end for
	}() // end go

	app.Main()
	return nil
} // end eventLoop

/* helper function to filter non-numeric characters from a string */
func filterNumbers(input string) string {
	// Remove any non-numeric characters
	return strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		} //end if
		return -1
	}, input)
} //end filterNumbers
