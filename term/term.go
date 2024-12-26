package term

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sulicat/goboi/colors"
	"github.com/sulicat/goboi/utils"
	"golang.org/x/term"
)

// the term package allows you to fill pixels on a terminal
// get terminal input, including mouse
// get info on the terminal (mostly a passthrough from go/x/term)

type RGB [3]int // rgb

type KeyCommand struct {
	Keycode int
	Buffer  []byte
}

type MouseCommand struct {
	IsMousePress   bool
	IsMouseRelease bool
	IsMouseMove    bool
	IsScrollUp     bool
	IsScrollDown   bool

	MouseX int
	MouseY int
	Button int

	Buffer []byte
}

type Cell struct {
	FGColor RGB
	BGColor RGB
	Char    string
	Depth   int

	cached      string
	has_changed bool
}

func (c *Cell) Str() string {

	if !c.has_changed {
		return c.cached
	}

	// FG color, BG color, char
	out := colors.Color(c.FGColor[0], c.FGColor[1], c.FGColor[2])
	out += colors.ColorBG(c.BGColor[0], c.BGColor[1], c.BGColor[2])
	out += c.Char
	c.cached = out
	return out
}

type FrameBuffer [][]Cell

func (f *FrameBuffer) Make(width int, height int) {
	*f = make([][]Cell, height)
	for r := range height {
		(*f)[r] = make([]Cell, width)
	}
}

func (f *FrameBuffer) Clear() {
	for r := range len(*f) {
		for c := range len((*f)[r]) {
			(*f)[r][c].has_changed = true
			(*f)[r][c].Char = ""
			(*f)[r][c].BGColor = RGB{0, 0, 0}
		}
	}
}

func (f *FrameBuffer) Overlay(other *FrameBuffer, x int, y int) {
	f_r := y
	f_c := x

	for other_r := 0; other_r < len(*other); other_r++ {
		for other_c := 0; other_c < len((*other)[other_r]); other_c++ {

			if f_r >= 0 && f_r < len(*f) {
				if f_c >= 0 && f_c < len((*f)[f_r]) {
					// fmt.Printf("%d:%d   %d:%d\n", f_r, len(*f), f_c, len((*f)[f_r]))
					(*f)[f_r][f_c] = (*other)[other_r][other_c]
				}
			}

			f_c += 1
		}
		f_r += 1
		f_c = x
	}
}

func CheckInside(
	px int, py int,
	cx int, cy int, cw int, ch int,
) bool {
	if px >= cx && px < cx+cw {
		if py >= cy && py < cy+ch {
			return true
		}
	}
	return false
}

type Term struct {
	width       int
	height      int
	start_x     int
	start_y     int
	framerate_s float32
	fullScreen  bool

	front FrameBuffer
	back  FrameBuffer

	sb               strings.Builder
	writer           *bufio.Writer
	old_state        *term.State
	frame_rate_timer utils.WaitTimer

	// input channel
	key_input_buff   chan KeyCommand
	mouse_input_buff chan MouseCommand

	term_state        TermState
	term_state_inital TermState
}

func Create(width int, height int) Term {
	out := Term{
		width:       width,
		height:      height,
		start_x:     0,
		start_y:     0,
		framerate_s: 1 / 30.0,
		fullScreen:  true,
	}

	out.front.Make(out.width, out.height)
	out.back.Make(out.width, out.height)

	out.writer = bufio.NewWriter(os.Stdout)
	out.sb = strings.Builder{}
	out.sb.Grow(out.width * out.height * 20)
	out.SetOffset(0, 0)
	out.frame_rate_timer = utils.CreateWaitTimer(float64(out.framerate_s))

	out.key_input_buff = make(chan KeyCommand, 10)     // buffer 10 keys
	out.mouse_input_buff = make(chan MouseCommand, 10) // buffer 10 moves

	out.SetColor(RGB{255, 255, 255})
	out.SetBackgroundColor(RGB{0, 0, 0})

	out.term_state_inital = out.term_state
	out.term_state_inital.MouseDown = false

	out.Start()
	return out
}

func (t *Term) SetColor(c RGB) {
	t.term_state.fg_color = c
}

func (t *Term) SetBackgroundColor(c RGB) {
	t.term_state.bg_color = c
}

func (t *Term) SameLine() {
	// next item to be added is same line
	t.term_state.SameLine()
}

func (t *Term) SetOffset(x, y int) {
	t.start_x = x
	t.start_y = y + 1
}

func (t *Term) Width() int {
	return t.width
}

func (t *Term) Height() int {
	return t.height
}

func (t *Term) Resize(new_w int, new_h int) {
	if new_w != t.width || new_h != t.height {
		t.width = new_w
		t.height = new_h

		// resize the cell buffers
		t.front.Make(t.width, t.height)
		t.back.Make(t.width, t.height)

		//clear the terminal
		t.sb.Reset()
		t.sb.WriteString(colors.ColorBG(0, 0, 0))
		t.sb.WriteString(Clear())
		t.writer.Write([]byte(t.sb.String()))
		fmt.Fprint(t.writer, t.sb.String())
		t.writer.Flush()

	}
}

func (t *Term) TermWidth() int {
	w, _, _ := term.GetSize(0)
	return w
}

func (t *Term) TermHeight() int {
	_, h, _ := term.GetSize(0)
	return h
}

func (t *Term) SetFullscreen(is_fullscreen bool) {
	t.fullScreen = is_fullscreen
}

func (t *Term) SetFramerate(framerate_s float32) {
	t.framerate_s = framerate_s
	t.frame_rate_timer.SetDuration(float64(t.framerate_s))
}

func (t *Term) Step() {

	t.term_state.reset_cursor_pos()

	if t.frame_rate_timer.Check() {
		t.frame_rate_timer.Reset()
		if t.fullScreen {
			t.Resize(t.TermWidth(), t.TermHeight())
		} else {
			t.Resize(t.width, t.height)
		}

		t.Draw()
	}

}

// parse input ansi sequeces
func (t *Term) process_key_command(in KeyCommand) {
}

// parse input ansi sequeces
func (t *Term) process_mouse_command(in MouseCommand) {

	t.term_state.MouseX = in.MouseX
	t.term_state.MouseY = in.MouseY

	// on rising edge, the clicked signal is high
	if in.IsMousePress && !t.term_state.last_mouse_down {

		t.term_state.MouseClicked = true

	} else if t.term_state.MouseClicked {
		// otherwise, we are gonna trigger the click low
		t.term_state.MouseClicked = false
	}

	t.term_state.last_mouse_down = in.IsMousePress

}

func (t *Term) InputLoop() {
	buf := make([]byte, 1024)
	for {
		n, _ := os.Stdin.Read(buf)
		if n == 1 {

			// make sure we can still escape out
			input := KeyCommand{Keycode: int(buf[0]), Buffer: append([]byte(nil), buf[:n]...)}
			if input.Keycode == 3 {
				t.Close()
			}

			if len(t.key_input_buff) < cap(t.key_input_buff) {
				t.key_input_buff <- input
			}

		} else if n > 3 {

			prefix := string(buf[:3])
			switch prefix {
			case MouseInputPrefix:
				// mouse input
				input := MouseCommand{Buffer: append([]byte(nil), buf[:n]...)}
				input.MouseX = int(buf[4]) - 33
				input.MouseY = int(buf[5]) - 33

				switch buf[3] {
				case 97: // scroll up
					input.IsScrollUp = true
				case 96: // scroll down
					input.IsScrollDown = true
				case 67: // MouseMove
					input.IsMouseMove = true
				case 32:
					input.IsMousePress = true
					input.Button = 0
				case 34:
					input.IsMousePress = true
					input.Button = 1
				case 35:
					input.IsMouseRelease = true
				}

				if len(t.mouse_input_buff) < cap(t.mouse_input_buff) {
					t.mouse_input_buff <- input
				}
			}

		}
	}
}

func (t *Term) Start() {

	// set raw mode
	s, err := term.MakeRaw(int(os.Stdin.Fd()))
	t.old_state = s
	if err != nil {
		fmt.Println(colors.BgRed+"Error setting raw mode:", err, colors.Reset)
		return
	}

	// start a go routine that fills an input buffer
	go func(t *Term) {
		t.InputLoop()
	}(t)

	// enable mouse tracking
	fmt.Print(EnableMouseTracking())

	// clear the terminal
	t.sb.Reset()
	t.sb.WriteString(Clear())
	t.writer.Write([]byte(t.sb.String()))
	fmt.Fprint(t.writer, t.sb.String())
	t.writer.Flush()

}
func (t *Term) Close() {
	fmt.Print(DisableMouseTracking())
	term.Restore(int(os.Stdin.Fd()), t.old_state)
}

func (t *Term) Draw() {

	// flush the buffer for input
	for {
		if len(t.key_input_buff) > 0 {
			input := <-t.key_input_buff
			t.process_key_command(input)
		} else {
			break
		}
	}

	// flush mouse events
	for {
		if len(t.mouse_input_buff) > 0 {
			input := <-t.mouse_input_buff
			t.process_mouse_command(input)
		} else {
			break
		}
	}

	t.sb.Reset()

	// move the cursor
	t.sb.WriteString(MoveCursor(t.start_x, t.start_y))

	for r := range t.height {
		for c := range t.width {

			// if we have an empty cell, make it really empty
			if t.front[r][c].Char == "" {
				t.front[r][c].Char = " "
				t.front[r][c].has_changed = true
			}

			t.sb.WriteString(t.front[r][c].Str())
		}
		t.sb.WriteString(MoveCursor(
			t.start_x,
			t.start_y+r+1))
	}

	t.writer.Write([]byte(t.sb.String()))
	fmt.Fprint(t.writer, t.sb.String())
	t.writer.Flush()

	// TODO: suli, can we fix this??? slow and busts the caching
	t.front.Clear()

}
