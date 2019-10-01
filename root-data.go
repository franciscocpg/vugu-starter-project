package main

type RootData struct{ Show bool }

func (data *RootData) Toggle() {
	data.Show = !data.Show
}
