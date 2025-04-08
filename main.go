package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	choices []string
	cursor  int
}

func ExecuteScript(fileName string) (string, error) {
	cmd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	shell := filepath.Join(cmd, "./", fileName)
	// fmt.Println(cmd, shell)
	out, err := exec.Command("/bin/bash", shell).Output()
	if err != nil {
		return "", err
	} else {
		return string(out), nil
	}
}

type Styles struct {
	BorderColor lipgloss.Color
	InputrField lipgloss.Style
}

var styleTrue = lipgloss.NewStyle().Foreground(lipgloss.Color("36"))
var styleName = lipgloss.NewStyle().Foreground(lipgloss.Color("3"))
var styleFalse = lipgloss.NewStyle()
var ss = new(Styles)
var styleHeader = lipgloss.NewStyle().BorderForeground(ss.BorderColor).BorderStyle(lipgloss.NormalBorder()).Padding(1).Width(80)

func initialModel() model {
	return model{
		choices: []string{"1. START SERVICES", "2. QUIT SERVICES", "3. EXIT"},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			if m.cursor == 0 {
				out, err := ExecuteScript("start_main.sh")
				if err != nil {
					log.Fatal(err)
				}
				fmt.Print(out)
			} else if m.cursor == 1 {
				out, err := ExecuteScript("quit_main.sh")
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(out)
			} else if m.cursor == 2 {
				return m, tea.Quit
			}
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	// s := "\n\n"
	// s += "Welcome to ManSyc"
	// s += "\n\n"
	// s += "Linux Process management tool.\n\n"
	// s += "Bash Scripting by - Dishan S Samuel, CLI UI built by - Vincent Samuel Paul\n"
	s := styleHeader.Render(styleTrue.Render("Welcome to ManSyc\n") + "\nA Linux Process management tool.\n\nBash Scripting by -" + styleName.Render(" Dishan S Samuel") + "\nCLI UI built by -" + styleName.Render(" Vincent Samuel Paul") + "\nTech Stack: Bash, GOlang, BubbleTea\n\n© ManSyc v1 2025 All Rights Reserved\n")
	s += "\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
			rendered := styleTrue.Render(fmt.Sprintf("%s %s", cursor, choice))
			s += fmt.Sprintf("%s\n", rendered)
		} else {
			rendered := styleFalse.Render(fmt.Sprintf("%s %s", cursor, choice))
			s += fmt.Sprintf("%s\n", rendered)
		}
		// s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
	s += "\nPress enter to run.\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
