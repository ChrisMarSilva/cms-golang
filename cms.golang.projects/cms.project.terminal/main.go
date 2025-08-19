package main

// go mod init github.com/chrismarsilva/cms.project.terminal
// go get -u github.com/charmbracelet/bubbletea
// go get -u github.com/charmbracelet/bubbles/spinner
// go get -u github.com/charmbracelet/bubbles/help
// go get -u github.com/charmbracelet/bubbles/key
// go get -u github.com/charmbracelet/bubbles/timer
// go get -u github.com/charmbracelet/bubbles/stopwatch
// go get -u github.com/charmbracelet/bubbles/progress
// go get -u github.com/charmbracelet/bubbles/paginator
// go get -u github.com/charmbracelet/lipgloss
// go get -u github.com/fogleman/ease
// go get -u github.com/lucasb-eyer/go-colorful
// go mod tidy

// go run main.go

// go install github.com/air-verse/air@latest
// air init
// air

import (
	"fmt"
	"log/slog"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/stopwatch"
	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fogleman/ease"
	"github.com/lucasb-eyer/go-colorful"
)

const (
	timeout           = 5 * time.Second
	padding           = 2
	maxWidth          = 80
	progressBarWidth  = 71
	progressFullChar  = "█"
	progressEmptyChar = "░"
	dotChar           = " • "
)

var (
	spinnerStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	helpStyleList     = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Margin(1, 0)
	dotStyle          = helpStyleList.UnsetMargins()
	helpStyleProgress = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render
	durationStyle     = dotStyle
	appStyle          = lipgloss.NewStyle().Margin(1, 2, 0, 2)
	keywordStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	subtleStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	ticksStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("79"))
	checkboxStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	progressEmpty     = subtleStyle.Render(progressEmptyChar)
	dotStyleView      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
	mainStyle         = lipgloss.NewStyle().MarginLeft(2)
	ramp              = makeRampStyles("#B14FFF", "#00FFA3", progressBarWidth)
)

func main() {
	m := newModel()
	p := tea.NewProgram(m)

	go func() {
		for {
			pause := time.Duration(rand.Int63n(899)+100) * time.Millisecond // nolint:gosec
			time.Sleep(pause)

			// Send the Bubble Tea program a message from outside the
			// tea.Program. This will block until it is ready to receive
			// messages.
			p.Send(resultMsg{food: randomFood(), duration: pause})
		}
	}()

	if _, err := p.Run(); err != nil {
		slog.Error("Error running program", slog.Any("err", err))
		os.Exit(1)
	}
}

type (
	tickMsg     time.Time
	tickMsgView struct{}
	frameMsg    struct{}
)

func randomFood() string {
	food := []string{
		"an apple", "a pear", "a gherkin", "a party gherkin",
		"a kohlrabi", "some spaghetti", "tacos", "a currywurst", "some curry",
		"a sandwich", "some peanut butter", "some cashews", "some ramen",
	}
	return food[rand.Intn(len(food))] // nolint:gosec
}

type model struct {
	sub       chan struct{} // where we'll receive activity notifications
	responses int           // how many responses we've received
	spinner   spinner.Model
	timer     timer.Model
	stopwatch stopwatch.Model
	keymap    keymap
	help      help.Model
	results   []resultMsg
	progress  progress.Model
	Choice    int
	Chosen    bool
	Ticks     int
	Frames    int
	Progress  float64
	Loaded    bool
	//Quitting bool
	items     []string
	paginator paginator.Model
	quitting  bool
}

func newModel() model {
	const numLastResults = 5

	s := spinner.New()
	s.Style = spinnerStyle

	var items []string
	for i := 1; i < 101; i++ {
		text := fmt.Sprintf("Item %d", i)
		items = append(items, text)
	}

	p := paginator.New()
	p.Type = paginator.Dots
	p.PerPage = 10
	p.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("•")
	p.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("•")
	p.SetTotalPages(len(items))

	m := model{
		//sub
		//responses
		spinner:   s,
		timer:     timer.NewWithInterval(timeout, time.Millisecond),
		stopwatch: stopwatch.NewWithInterval(time.Millisecond),
		keymap: keymap{
			start: key.NewBinding(key.WithKeys("s"), key.WithHelp("s", "start")),
			stop:  key.NewBinding(key.WithKeys("s"), key.WithHelp("s", "stop")),
			reset: key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "reset")),
			quit:  key.NewBinding(key.WithKeys("q", "ctrl+c"), key.WithHelp("q", "quit")),
		},
		help:     help.New(),
		results:  make([]resultMsg, numLastResults),
		progress: progress.New(progress.WithDefaultGradient()),
		Choice:   0,
		Chosen:   false,
		Ticks:    10,
		Frames:   0,
		Progress: 0,
		Loaded:   false,
		//Quitting :false,
		paginator: p,
		items:     items,
		//quitting
	}

	m.keymap.start.SetEnabled(false)

	return m
}

func (m model) View() string {
	var result string

	result += "timer: " + m.timer.View() + "\n"
	result += "stopwatch: " + m.stopwatch.View() + "\n"

	if m.timer.Timedout() {
		result += "timer: All done!\n"
	}

	if m.quitting {
		result += "That’s all for today!"
	} else {
		// result += "Exiting in " + m.timer.View() + "\n"
		// result += "Elapsed " + m.stopwatch.View() + "\n"
		result += m.spinner.View() + " Eating food..."
		result += m.helpView()
	}

	result += "\n\n"
	for _, res := range m.results {
		result += res.String() + "\n"
	}

	if !m.quitting {
		result += helpStyleList.Render("Press any key to exit")
	}

	result = appStyle.Render(result)

	pad := strings.Repeat(" ", padding)
	result += "\n" + pad + m.progress.View() + "\n\n" + pad + helpStyleProgress("Press any key to quit")

	if !m.Chosen {
		result += mainStyle.Render("\n" + choicesView(m) + "\n\n")
	} else {
		result += mainStyle.Render("\n" + chosenView(m) + "\n\n")
	}

	var b strings.Builder
	b.WriteString("\n  Paginator Example\n\n")
	start, end := m.paginator.GetSliceBounds(len(m.items))
	for _, item := range m.items[start:end] {
		b.WriteString("  • " + item + "\n\n")
	}
	b.WriteString("  " + m.paginator.View())
	b.WriteString("\n\n  h/l ←/→ page • q: quit\n")
	result += b.String()

	return result
}

func (m model) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{m.keymap.start, m.keymap.stop, m.keymap.reset, m.keymap.quit})
}

func (m model) Init() tea.Cmd {
	var result tea.Cmd
	result = tea.Batch(result, m.timer.Init())
	result = tea.Batch(result, m.stopwatch.Init())
	result = tea.Batch(result, m.spinner.Tick)
	result = tea.Batch(result, tickCmd())
	result = tea.Batch(result, tick())

	return result
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			//m.Quitting = true
			return m, tea.Quit
		}
	}

	if !m.Chosen {
		return updateChoices(msg, m)
	}

	switch msg := msg.(type) {
	case timer.TickMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, cmd

	case timer.StartStopMsg:
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		m.keymap.stop.SetEnabled(m.timer.Running())
		m.keymap.start.SetEnabled(!m.timer.Running())
		return m, cmd

	case timer.TimeoutMsg:
		m.quitting = true
		return m, tea.Quit

	case resultMsg:
		m.results = append(m.results[1:], msg)
		return m, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case tickMsg:
		if m.progress.Percent() == 1.0 {
			return m, nil // return m, tea.Quit
		}
		cmd := m.progress.IncrPercent(0.25)
		return m, tea.Batch(tickCmd(), cmd)

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.quit):
			m.quitting = true
			return m, tea.Quit

		case key.Matches(msg, m.keymap.reset):
			m.timer.Timeout = timeout
			return m, m.stopwatch.Reset()

		case key.Matches(msg, m.keymap.start, m.keymap.stop):
			m.keymap.stop.SetEnabled(!m.stopwatch.Running())
			m.keymap.start.SetEnabled(m.stopwatch.Running())
			m.timer.Toggle()
			return m, m.stopwatch.Toggle()
		}
	}

	updateChosen(msg, m)

	var cmd tea.Cmd
	m.stopwatch, cmd = m.stopwatch.Update(msg)
	m.paginator, cmd = m.paginator.Update(msg)

	return m, cmd
}

type keymap struct {
	start key.Binding
	stop  key.Binding
	reset key.Binding
	quit  key.Binding
}

type resultMsg struct {
	duration time.Duration
	food     string
}

func (r resultMsg) String() string {
	if r.duration == 0 {
		return dotStyle.Render(strings.Repeat(".", 30))
	}

	return fmt.Sprintf("🍔 Ate %s %s", r.food, durationStyle.Render(r.duration.String()))
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

func frame() tea.Cmd {
	return tea.Tick(time.Second/60, func(time.Time) tea.Msg {
		return frameMsg{}
	})
}

func updateChosen(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case frameMsg:
		if !m.Loaded {
			m.Frames++
			m.Progress = ease.OutBounce(float64(m.Frames) / float64(100))
			if m.Progress >= 1 {
				m.Progress = 1
				m.Loaded = true
				m.Ticks = 3
				return m, tick()
			}
			return m, frame()
		}

	case tickMsg:
		if m.Loaded {
			if m.Ticks == 0 {
				//m.Quitting = true
				return m, tea.Quit
			}
			m.Ticks--
			return m, tick()
		}
	}

	return m, nil
}

// Sub-views

// The first view, where you're choosing a task
func choicesView(m model) string {
	c := m.Choice

	tpl := "What to do today?\n\n"
	tpl += "%s\n\n"
	tpl += "Program quits in %s seconds\n\n"
	tpl += subtleStyle.Render("j/k, up/down: select") + dotStyleView +
		subtleStyle.Render("enter: choose") + dotStyleView +
		subtleStyle.Render("q, esc: quit")

	choices := fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		checkbox("Plant carrots", c == 0),
		checkbox("Go to the market", c == 1),
		checkbox("Read something", c == 2),
		checkbox("See friends", c == 3),
	)

	return fmt.Sprintf(tpl, choices, ticksStyle.Render(strconv.Itoa(m.Ticks)))
}

func chosenView(m model) string {
	var msg string

	switch m.Choice {
	case 0:
		msg = fmt.Sprintf("Carrot planting?\n\nCool, we'll need %s and %s...", keywordStyle.Render("libgarden"), keywordStyle.Render("vegeutils"))
	case 1:
		msg = fmt.Sprintf("A trip to the market?\n\nOkay, then we should install %s and %s...", keywordStyle.Render("marketkit"), keywordStyle.Render("libshopping"))
	case 2:
		msg = fmt.Sprintf("Reading time?\n\nOkay, cool, then we’ll need a library. Yes, an %s.", keywordStyle.Render("actual library"))
	default:
		msg = fmt.Sprintf("It’s always good to see friends.\n\nFetching %s and %s...", keywordStyle.Render("social-skills"), keywordStyle.Render("conversationutils"))
	}

	label := "Downloading..."
	if m.Loaded {
		label = fmt.Sprintf("Downloaded. Exiting in %s seconds...", ticksStyle.Render(strconv.Itoa(m.Ticks)))
	}

	return msg + "\n\n" + label + "\n" + progressbar(m.Progress) + "%"
}

func checkbox(label string, checked bool) string {
	if checked {
		return checkboxStyle.Render("[x] " + label)
	}
	return fmt.Sprintf("[ ] %s", label)
}

func progressbar(percent float64) string {
	w := float64(progressBarWidth)

	fullSize := int(math.Round(w * percent))
	var fullCells string
	for i := 0; i < fullSize; i++ {
		fullCells += ramp[i].Render(progressFullChar)
	}

	emptySize := int(w) - fullSize
	emptyCells := strings.Repeat(progressEmpty, emptySize)

	return fmt.Sprintf("%s%s %3.0f", fullCells, emptyCells, math.Round(percent*100))
}

func makeRampStyles(colorA, colorB string, steps float64) (s []lipgloss.Style) {
	cA, _ := colorful.Hex(colorA)
	cB, _ := colorful.Hex(colorB)

	for i := 0.0; i < steps; i++ {
		c := cA.BlendLuv(cB, i/steps)
		s = append(s, lipgloss.NewStyle().Foreground(lipgloss.Color(colorToHex(c))))
	}
	return
}

func colorToHex(c colorful.Color) string {
	return fmt.Sprintf("#%s%s%s", colorFloatToHex(c.R), colorFloatToHex(c.G), colorFloatToHex(c.B))
}

func colorFloatToHex(f float64) (s string) {
	s = strconv.FormatInt(int64(f*255), 16)
	if len(s) == 1 {
		s = "0" + s
	}
	return
}

func updateChoices(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.Choice++
			if m.Choice > 3 {
				m.Choice = 3
			}
		case "k", "up":
			m.Choice--
			if m.Choice < 0 {
				m.Choice = 0
			}
		case "enter":
			m.Chosen = true
			return m, frame()
		}

	case tickMsg:
		if m.Ticks == 0 {
			//m.Quitting = true
			return m, tea.Quit
		}
		m.Ticks--
		return m, tick()
	}

	return m, nil
}
