package main

import (
	"fmt"
	"log"

	"github.com/MadhavKrishanGoswami/Cash-CLI/currency"
	"github.com/MadhavKrishanGoswami/Cash-CLI/models"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

var in models.Input

type model struct {
	BaseCurrency    string
	TargetCurrency  string
	ConversionRate  float64
	ConvertedAmount float64
}

func (m model) Init() tea.Cmd { return nil }
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Automatically quit after displaying the message
	return m, tea.Quit
}

func (m model) View() string {
	// Define a style for the output
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("208")).
		Bold(true).
		Padding(1, 2)

	// Construct a fun output message
	output := fmt.Sprintf(
		"💰 You're converting from %s to %s!\n"+
			"📈 Current magic rate: %f.\n"+
			"🎉Ta-da! Your converted amount is: %.2f!\n"+
			"Keep on converting those currencies! 🌟",
		m.BaseCurrency, m.TargetCurrency, m.ConversionRate, m.ConvertedAmount,
	)

	return "\n" + style.Render(output) + "\n"
}

func main() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("What is your base currency").
				Options(
					huh.NewOption("$ USD", "USD"),
					huh.NewOption("₹ INR", "INR"),
					huh.NewOption("¥ JPY", "JPY"),
					huh.NewOption("€ EURO", "EUR"),
				).
				Value(&in.Base),
			huh.NewSelect[string]().
				Title("What do you want to convert into").
				Options(
					huh.NewOption("$ USD", "USD"),
					huh.NewOption("₹ INR", "INR"),
					huh.NewOption("¥ JPY", "JPY"),
					huh.NewOption("€ EURO", "EUR"),
				).
				Value(&in.Convert),
			huh.NewInput().
				Title("How much to convert?").
				Value(&in.Input),
		))
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	out := currency.Convertor(in)
	conversionModel := model{
		BaseCurrency:    in.Base,
		TargetCurrency:  in.Convert,
		ConversionRate:  out.Converstion,
		ConvertedAmount: out.Value,
	}
	if err := tea.NewProgram(conversionModel).Start(); err != nil {
		panic(err)
	}
}
