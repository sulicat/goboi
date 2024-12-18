package term

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sulicat/goboi/colors"
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

type Term struct {
	width   int
	height  int
	start_x int
	start_y int
	pixels  []RGB

	sb        strings.Builder
	writer    *bufio.Writer
	old_state *term.State

	// input channel
	key_input_buff   chan KeyCommand
	mouse_input_buff chan MouseCommand
}

func Create(width int, height int) Term {
	out := Term{
		width:   width,
		height:  height,
		start_x: 0,
		start_y: 0,
	}

	out.pixels = make([]RGB, width*height)
	out.writer = bufio.NewWriter(os.Stdout)
	out.sb = strings.Builder{}
	out.sb.Grow(out.width * out.height * 20)
	out.SetOffset(0, 0)

	out.key_input_buff = make(chan KeyCommand, 10)     // buffer 10 keys
	out.mouse_input_buff = make(chan MouseCommand, 10) // buffer 10 moves

	out.Start()
	return out
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
	t.width = new_w
	t.height = new_h
}

func (t *Term) TermWidth() int {
	w, _, _ := term.GetSize(0)
	return w
}

func (t *Term) TermHeight() int {
	_, h, _ := term.GetSize(0)
	return h
}

// parse input ansi sequeces
func (t *Term) process_key_command(in KeyCommand) {
}

// parse input ansi sequeces
func (t *Term) process_mouse_command(in MouseCommand) {
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
				input.MouseX = int(buf[4])
				input.MouseY = int(buf[5])

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
	go t.InputLoop()

	// enable mouse tracking
	fmt.Print(EnableMouseTracking())

	// clear the terminal
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

	// TODO suli: clearing bad, keeps history
	//t.sb.WriteString(Clear())

	// move the cursor
	t.sb.WriteString(MoveCursor(t.start_x, t.start_y))

	for y := range t.height {
		for range t.width {
			t.sb.WriteString(DrawBlock(0, 0, 255))
		}
		t.sb.WriteString(MoveCursor(t.start_x, t.start_y+y+1))
		// t.sb.WriteString("\n")
	}

	t.writer.Write([]byte(t.sb.String()))
	fmt.Fprint(t.writer, t.sb.String())
	t.writer.Flush()

}
