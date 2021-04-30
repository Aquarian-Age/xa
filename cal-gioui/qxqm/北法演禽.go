package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/outlay"
	"image"
	"strings"
	"time"
)

//北法演禽主窗口 也是程序主窗口
type BFYQWindow struct {
	win    *Window
	items  []*BFYQ
	list   layout.List
	days   []*cald
	months []*calm
	hours  []*calh
	years  YearBtns
}

//北法演禽二十八宿日禽口诀图
type BFYQ struct {
	Number int
	Text   string //二十八宿日禽名称
	Click  widget.Clickable
}

var (
	t                          = time.Now().Local()
	y                          = t.Year()
	swapY, swapM, swapD, swapH int
)

//主窗口
func NewBFYQWindow() *BFYQWindow {
	view := &BFYQWindow{
		list: layout.List{Axis: layout.Vertical},
	}
	//北法演禽二十八宿日禽口诀图
	for i := 0; i < len(starNames); i++ {
		view.items = append(view.items, &BFYQ{Text: starNames[i]})
	}

	//月份
	for i := 1; i < 13; i++ {
		view.months = append(view.months, &calm{month: i})
	}
	//日
	for i := 1; i <= 31; i++ {
		view.days = append(view.days, &cald{day: i})
	}
	//hour
	for i := 0; i < 24; i++ {
		view.hours = append(view.hours, &calh{hour: i})
	}

	return view
}

//
func (v *BFYQWindow) Run(w *Window) error {
	v.win = w
	return bfyqView(v.LayoutBfyqRiQin).Run(w)
}

type bfyqView func(gtx layout.Context) layout.Dimensions

//方法
func (view bfyqView) Run(w *Window) error {
	var ops op.Ops
	ticker := time.NewTicker(time.Second)
	th := w.App.Theme
	applicationClose := w.App.Context.Done()
	for {
		select {
		case <-applicationClose:
			return nil
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				realTimeLayout(gtx, material.NewTheme(gofont.Collection()))
				Show(gtx, th)
				view(gtx)
				e.Frame(gtx.Ops)
			}
		case <-ticker.C:
			w.Invalidate()
		}

	}
}

//
func (v *BFYQWindow) LayoutBfyqRiQin(gtx layout.Context) layout.Dimensions {
	hGrid := outlay.Grid{
		Num:  7, //每行的最大数
		Axis: layout.Horizontal,
	}
	th := v.win.App.Theme

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		//年
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return v.years.LayoutYear(gtx, th)
		}),
		//月
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return hGrid.Layout(gtx, 12, func(gtx layout.Context, indexM int) layout.Dimensions {
				m := v.months[indexM]
				if m.Click.Clicked() {
					//fmt.Printf("月:%d\n", m.month)
					swapM = m.month
				}
				var mWidth int
				ms := fmt.Sprintf("%d月", m.month)
				return layout.Stack{Alignment: layout.NE}.Layout(gtx,
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						dims := material.Clickable(gtx, &m.Click, func(gtx layout.Context) layout.Dimensions {
							return layout.UniformInset(unit.Dp(6)).Layout(gtx,
								material.Body1(th, ms).Layout,
							)
						})
						mWidth = dims.Size.X
						return dims
					}),
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						nameHeight := gtx.Px(unit.Dp(3))
						return layout.Dimensions{
							Size: image.Point{X: mWidth, Y: nameHeight},
						}
					}),
				)
			})
		}),
		//日
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return hGrid.Layout(gtx, len(v.days), func(gtx layout.Context, index int) layout.Dimensions {
				d := v.days[index]
				if d.Click.Clicked() {
					//fmt.Printf("日: %d\n", d.day)
					swapD = d.day
					txt := " 的子窗口"
					SubWD(d.day, th, v, txt)
				}
				var dWidth int
				ds := fmt.Sprintf("%d日", d.day)
				return layout.Stack{Alignment: layout.NE}.Layout(gtx,
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						dims := material.Clickable(gtx, &d.Click, func(gtx layout.Context) layout.Dimensions {
							return layout.UniformInset(unit.Dp(6)).Layout(gtx,
								material.Body1(th, ds).Layout,
							)
						})
						dWidth = dims.Size.X
						return dims
					}),
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						nameHeight := gtx.Px(unit.Dp(3))
						//tabRect := image.Rect(0, 0, calWidth, nameHeight)
						//paint.FillShape(gtx.Ops, th.Palette.ContrastBg, clip.Rect(tabRect).Op())
						return layout.Dimensions{
							Size: image.Point{X: dWidth, Y: nameHeight},
						}
					}),
				)
			})
		}),
		//时辰
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return hGrid.Layout(gtx, 24, func(gtx layout.Context, indexH int) layout.Dimensions {
				h := v.hours[indexH]
				if h.Click.Clicked() {
					//	fmt.Printf("时辰: %d\n", h.hour)
					swapH = h.hour
					///新窗口内容 备用
					switch h.hour {
					case 23, 0:
						title := "子时"
						txt := "23 0"
						subWindowH(th, v, title, txt)
					case 1, 2:
						title := "丑时"
						txt := "1 2"
						subWindowH(th, v, title, txt)
					case 3, 4:
						title := "寅时"
						txt := "3 5"
						subWindowH(th, v, title, txt)
					case 5, 6:
						title := "卯时"
						txt := "5 7"
						subWindowH(th, v, title, txt)
					case 7, 8:
						title := "辰时"
						txt := "7 9"
						subWindowH(th, v, title, txt)
					case 9, 10:
						title := "巳时"
						txt := "9 10"
						subWindowH(th, v, title, txt)
					case 11, 12: //示例
						info := material.H6(th, "11,12")
						size := info.TextSize
						size.V *= 36 //窗口初始大小
						v.win.App.NewWindow("午时",
							bfyqView(func(gtx layout.Context) layout.Dimensions {
								gtx.Reset() //清除主窗口内容
								return layout.Center.Layout(gtx, info.Layout)
							}),
							app.Size(size, size),
						)
					case 13, 14:
						title := "未时"
						txt := "13 14"
						subWindowH(th, v, title, txt)
					case 15, 16:
						title := "申时"
						txt := "15,16"
						subWindowH(th, v, title, txt)
					case 17, 18:
						title := "酉时"
						txt := "17,18"
						subWindowH(th, v, title, txt)
					case 19, 20:
						title := "戌时"
						txt := "19 20"
						subWindowH(th, v, title, txt)
					case 21, 22:
						title := "亥时"
						txt := "21 22"
						subWindowH(th, v, title, txt)
					}
				}
				var hWidth int
				hours := fmt.Sprintf("%d点", h.hour)
				return layout.Stack{Alignment: layout.NE}.Layout(gtx,
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						dims := material.Clickable(gtx, &h.Click, func(gtx layout.Context) layout.Dimensions {
							return layout.UniformInset(unit.Dp(6)).Layout(gtx,
								material.Body1(th, hours).Layout,
							)
						})
						hWidth = dims.Size.X
						return dims
					}),
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						nameHeight := gtx.Px(unit.Dp(3))
						return layout.Dimensions{
							Size: image.Point{X: hWidth, Y: nameHeight},
						}
					}),
				)
			})
		}),
		//新窗口显示二十八宿图片
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return hGrid.Layout(gtx, len(v.items), func(gtx layout.Context, index int) layout.Dimensions {
				item := v.items[index]
				if item.Click.Clicked() {
					for i := 0; i < len(starNames); i++ {
						if strings.EqualFold(starNames[i], item.Text) {
							showBfyqRiQin(item.Text)
							break
						}
					}
				}
				///
				var nameWidth int
				return layout.Stack{Alignment: layout.S}.Layout(gtx,
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						dims := material.Clickable(gtx, &item.Click, func(gtx layout.Context) layout.Dimensions {
							return layout.UniformInset(unit.Dp(6)).Layout(gtx,
								material.Body1(th, item.Text).Layout,
							)
						})
						nameWidth = dims.Size.X
						return dims
					}),
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						nameHeight := gtx.Px(unit.Dp(2))
						//tabRect := image.Rect(0, 0, nameWidth, nameHeight)
						//paint.FillShape(gtx.Ops, th.Palette.ContrastBg, clip.Rect(tabRect).Op())
						return layout.Dimensions{
							Size: image.Point{X: nameWidth, Y: nameHeight},
						}
					}),
				)
			})
		}),
	)
}

//
var (
	starNames = []string{
		"角木蛟", "亢金龙", "氐土貉", "房日兔", "心月狐", "尾火虎", "箕水豹", //东方青龙
		"斗木獬", "牛金牛", "女土蝠", "虚日鼠", "危月燕", "室火猪", "壁水貐", //北方玄武
		"奎木狼", "娄金狗", "胃土雉", "昴日鸡", "畢月乌", "觜火猴", "参水猿", //西方白虎
		"井木犴", "鬼金羊", "柳土獐", "星日马", "张月鹿", "翼火蛇", "轸水蚓", //南方朱雀
	}
)

//新窗口显示图片
//北法演禽 日禽
func showBfyqRiQin(name string) {
	go func() {
		w := app.NewWindow(app.Title(name), app.Size(unit.Dp(1600), unit.Dp(850)))
		events := w.Events()
		var ops op.Ops
		for {
			e := <-events
			switch e := e.(type) {
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				ns := name + ".png"
				showPNG(ns, gtx)
				e.Frame(gtx.Ops)
			}
		}
	}()
}
