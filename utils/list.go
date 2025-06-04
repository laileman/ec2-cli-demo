package utils

import (
	"context"
	//"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/pterm/pterm"
	//"gopkg.in/yaml.v3"
)

type EC2Instance struct {
	InstanceID string `yaml:"instanceID"`
	State      string `yaml:"state"`
	PublicIP   string `yaml:"publicIP,omitempty"`
	PrivateIP  string `yaml:"privateIP,omitempty"`
}

func ListInstances() {
	// This function will list all EC2 instances in the specified region.
	// It will use the AWS SDK to connect to the EC2 service and retrieve
	// the list of instances.
	//var cfg := cfg.NewConfig("config")

	ec2Client := EC2Client()
	// Create a request to describe instances
	input := &ec2.DescribeInstancesInput{}
	// Call the DescribeInstances API
	resp, err := ec2Client.DescribeInstances(context.TODO(), input)
	if err != nil {
		// Handle error
		panic(err)
	}
	// Iterate through the reservations and instances to extract details
	//var ec2Instances []EC2Instance
	dataTable := pterm.TableData{
		{"Instance ID", "State", "Public IP", "Private IP"},
	}
	for _, reservation := range resp.Reservations {
		for _, instance := range reservation.Instances {
			// Create an EC2Instance struct for each instance

			//spinner, _ := pterm.DefaultSpinner.Start("Loading instances...")
			//time.Sleep(2 * time.Second)
			//bar, _ := pterm.DefaultProgressbar.WithTotal(100).WithTitle("Loading instances...").Start()
			var publicIP string
			if instance.PublicIpAddress == nil {
				publicIP = "n/a"
			} else {
				publicIP = *instance.PublicIpAddress
			}
			dataTable = append(dataTable, []string{*instance.InstanceId, string(instance.State.Name), publicIP, *instance.PrivateIpAddress})
			//	dataTable = append(dataTable, []string{*instance.InstanceId, string(instance.State.Name), publicIP, *instance.PrivateIpAddress})
			_ = pterm.DefaultTable.WithHasHeader().WithBoxed(true).WithRowSeparator("-").WithData(dataTable).Render()

		}
	}
}
