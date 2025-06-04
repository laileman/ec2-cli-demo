package utils

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/pterm/pterm"
	"io/ioutil"
)

func Ssh(file string, user string) {
	// This function will list all EC2 instances in the specified region.
	// It will use the AWS SDK to connect to the EC2 service and retrieve
	// the list of instances.
	ec2Client := EC2Client()
	// Create a request to describe instances
	input := &ec2.DescribeInstancesInput{}
	// Call the DescribeInstances API
	resp, err := ec2Client.DescribeInstances(context.TODO(), input)
	if err != nil {
		// Handle error
		panic(err)
	}
	//publicIPs := make([]string, 0)
	// Iterate through the reservations and instances to extract details
	//publicIPs[0] = *resp.Reservations[0].Instances[0].PrivateIpAddress

	if resp.Reservations[0].Instances[0].PublicIpAddress == nil {
		fmt.Println("No public IP address available for the instance. skipping SSH config generation.")
		return
	}
	//fmt.Println(*resp.Reservations[0].Instances[0].PublicIpAddress)

	sshCfg := `
Host aws-instance
  HostName %s
  User %s`
	sshCfg = fmt.Sprintf(sshCfg, *resp.Reservations[0].Instances[0].PublicIpAddress, user)
	pterm.Info.Println("ssh config generate ...")
	err = ioutil.WriteFile(file, []byte(sshCfg), 777)
	if err != nil {
		pterm.Error.Println("Error writing SSH config file:", err)
		return
	}
	pterm.Success.Println("SSH config file generated successfully. ", file)
}
