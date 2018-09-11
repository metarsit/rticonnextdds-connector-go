/*****************************************************************************
*    (c) 2005-2015 Copyright, Real-Time Innovations, All rights reserved.    *
*                                                                            *
*  RTI grants Licensee a license to use, modify, compile, and create         *
*  derivative works of the Software.  Licensee has the right to distribute   *
*  object form only for use with RTI products. The Software is provided      *
*  "as is", with no warranty of any type, including any warranty for fitness *
*  for any purpose. RTI is under no obligation to maintain or support the    *
*  Software.  RTI shall not be liable for any incidental or consequential    *
*  damages arising out of the use or inability to use the software.          *
*                                                                            *
*****************************************************************************/

package main

import (
	"github.com/rticommunity/rticonnextdds-connector-go"
	"path"
	"runtime"
	"time"
	"log"
)

type Shape struct {
        Color     string `json:"color"`
        X         int    `json:"x"`
        Y         int    `json:"y"`
        Shapesize int    `json:"shapesize"`
}

func main() {
	// Find the file path to the XML configuration
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Panic("runtime.Caller error")
	}
	filepath := path.Join(path.Dir(filename), "../ShapeExample.xml")

	// Create a connector defined in the XML configuration
	connector, err := rti.NewConnector("MyParticipantLibrary::Zero", filepath)
	if err != nil {
		log.Panic(err)
	}
	// Delete the connector when this main function returns
	defer connector.Delete()

	// Get an output from the connector
	output, err := connector.GetOutput("MyPublisher::MySquareWriter")
	if err != nil {
		log.Panic(err)
	}

	// Set values to the instance and write the instance
	for i := 0; i < 500; i++ {
		var shape Shape
		shape.X = i
		shape.Y = i*2
		shape.Shapesize = 30
		shape.Color = "BLUE"
		output.Instance.Set(&shape)
		output.Write()
		log.Println("Writing...")
		time.Sleep(time.Second * 1)
	}
}
