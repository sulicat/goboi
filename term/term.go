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

type InputCommand struct {
	Keycode int
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
	key_input_buff   chan InputCommand
	mouse_input_buff chan InputCommand
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

	out.key_input_buff = make(chan InputCommand, 10) // buffer 10 keys

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
func (t *Term) parse(in InputCommand) {
	// 3 is crt-c
	if in.Keycode == 3 {
		t.Close()
	}

}

func (t *Term) InputLoop() {
	buf := make([]byte, 1024)
	for {
		n, _ := os.Stdin.Read(buf)
		if n == 1 {
			t.key_input_buff <- InputCommand{Keycode: int(buf[0])}
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
			t.parse(input)
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
