package main

type Command interface {
	execute()
}

// OnCommand - turns the lamp on
type OnCommand struct {
	lamp Lamp
}

func (o *OnCommand) execute() {
	o.lamp.on()
}


// OffCommand - turns the lamp off
type OffCommand struct {
	lamp Lamp
}

func (o *OffCommand) execute() {
	o.lamp.off()
}

// ChangeColorCommand - turns the lamp off
type ChangeColorCommand struct {
	lamp Lamp
}

func (o *ChangeColorCommand) execute() {
	o.lamp.changeColor()
}