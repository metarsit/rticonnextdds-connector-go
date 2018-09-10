package main

import (
	"github.com/kyoungho/rticonnextdds-connector"
	"log"
	"path"
	"runtime"
	"time"
)

type Shape struct {
	Color     string `json:"color"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Shapesize int    `json:"shapesize"`
}

func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Panic("runtime.Caller error")
	}
	filepath := path.Join(path.Dir(filename), "../ShapeExample.xml")
	connector, err := rti.NewConnector("MyParticipantLibrary::Zero", filepath)
	if err != nil {
		log.Panic(err)
	}
	defer connector.Delete()
	input, err := connector.GetInput("MySubscriber::MySquareReader")
	if err != nil {
		log.Panic(err)
	}
	output, err := connector.GetOutput("MyPublisher::MySquareWriter")
	if err != nil {
		log.Panic(err)
	}

	for i := 0; i < 500; i++ {
		input.Take()
		numOfSamples := input.Samples.GetLength()
		for j := 0; j < numOfSamples; j++ {
			if input.Infos.IsValid(j) {
				var shape Shape
				err := input.Samples.Get(j, &shape)
				if err != nil {
					log.Println(err)
				}
				log.Println("---Received Sample---")
				log.Printf("color: %s\n", shape.Color)
				log.Printf("x: %d\n", shape.X)
				log.Printf("y: %d\n", shape.Y)
				log.Printf("shapesize: %d\n", shape.Shapesize)

				if shape.Color == "BLUE" {
					tmp := shape.X
					shape.X = shape.Y
					shape.Y = tmp
					shape.Color = "RED"
					err = output.Instance.Set(shape)
					if err != nil {
						log.Println(err)
					}
					output.Write()
				}
			}
		}
		time.Sleep(time.Second * 1)
	}
}
