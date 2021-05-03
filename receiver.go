package main

import "log"

type Lamp interface {
	on()
	off()
	changeColor()
}

type Switch struct {
	lampOn bool
	color string
}

func (s *Switch) on()  {
	log.Println("[Switch] turning on...")
	s.lampOn = true
}

func (s *Switch) off()  {
	log.Println("[Switch] turning off...")
	s.lampOn = false
}

func (s *Switch) changeColor()  {
	log.Println("[Switch] changing color...")
	s.color = "green"
}


type Alexa struct {
	lampOn bool
	lampColor string
}

func (s *Alexa) on()  {
	log.Println("[Alexa] turning on...")
	s.lampOn = true
}

func (s *Alexa) off()  {
	log.Println("[Alexa] turning off...")
	s.lampOn = false
}

func (s *Alexa) changeColor()  {
	log.Println("[Alexa] changing color...")
	s.lampColor = "red"
}