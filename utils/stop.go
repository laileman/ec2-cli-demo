package utils

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/pterm/pterm"
)

func StopInstance(id string) {
	// This function will start all stopped EC2 instances
	ec2client := EC2Client()

	// Implement the logic to stop instances using the client
	input := &ec2.StopInstancesInput{
		InstanceIds: []string{id},
	}
	_, err := ec2client.StopInstances(context.TODO(), input)
	if err != nil {
		// Handle error
		pterm.Error.Printf("instance stop failed. Instance ID: %s\n", id)
		pterm.Error.Println("Error stopping instance:", err)
		return
	}
	// Optionally, you can print a message indicating that the instances have been stopped
	pterm.Success.Printf("instance stop success. Instance ID: %s\n", id)
}
