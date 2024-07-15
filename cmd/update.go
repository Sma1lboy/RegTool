package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type updateModel struct {
	confirming bool
}

func (m updateModel) Init() tea.Cmd {
	return nil
}

func (m updateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			fmt.Println("Checking all apps and querying local software")
			//source.GetAllRegisteredApp();
			//TODO update here
			//获取所有本地存在的app->更新本地库（这个有时候会用到 直接从本地获取而不是每次都check）-> 更新结束
			//同时需要加上动画
			return GetCommand("mainMenu")
		case "q", "esc":
			return GetCommand("mainMenu")
		}

	}
	return m, nil
}

func (m updateModel) View() string {
	return "Init/Update All Apps Recording \n\nAre you sure you want to check all apps and query local software? Press 'Enter' to confirm, 'q' or 'esc' to go back.\n"
}

func init() {
	RegisterCommand("update", "Init/Update All Apps Recording", updateModel{})
}
