package main

import dv "devdrift/device" // ← только так Go найдёт ваш пакет

var devices = []dv.Device{

	dv.Pc{dv.Cpu{
		Prod: "Intel",
		Core: 16,
	}},
	dv.Mobile{dv.Cpu{
		Prod: "AMP",
		Core: 4,
	}},
}

func main() {
for _, dvc:= range devices {
	dvc.Off()
}
}