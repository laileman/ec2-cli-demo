package utils

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/pterm/pterm"
)

func StartInstance(id string) {
	// This function will start all stopped EC2 instances
	ec2client := EC2Client()
	// Implement the logic to start instances using the client
	// For example, you can use the StartInstances API call
	// ...
	input := &ec2.StartInstancesInput{
		InstanceIds: []string{id},
	}
	_, err := ec2client.StartInstances(context.TODO(), input)
	if err != nil {
		// Handle error
		pterm.Error.Printf("instance start failed. Instance ID: %s\n", id)
		pterm.Error.Println("Error starting instance:", err)
		return
	}
	// Print the output or handle it as needed
	//fmt.Println(output)
	// Optionally, you can print a message indicating that the instances have been started
	pterm.Success.Printf("instance start success. Instance ID: %s\n", id)
	//fmt.Printf("instance %v been started.", id)
}
