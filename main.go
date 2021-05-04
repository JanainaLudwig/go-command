package main

func main()  {
	alexa := &Alexa{}
	swit := &Switch{}

	onAlexa := OnCommand{lamp: alexa}
	onSwitch := OnCommand{lamp: swit}

	onAlexa.execute()
	onSwitch.execute()


	changeColorAlexa := ChangeColorCommand{lamp: alexa}
	changeColorSwitch := ChangeColorCommand{lamp: swit}

	changeColorAlexa.execute()
	changeColorSwitch.execute()


	offAlexa := OffCommand{lamp: alexa}
	offSwitch := OffCommand{lamp: swit}

	offAlexa.execute()
	offSwitch.execute()
}
