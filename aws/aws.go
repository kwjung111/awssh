package aws

import (
	"cfg"
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

var client *ec2.Client

func InitClient() *ec2.Client {
	fmt.Println("client init")

	conf := *cfg.GetConf()

	profile := conf.Profile

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(profile),
	)
	if err != nil {
		fmt.Printf("Unable to Load SDK config : %v \n", err)
	}

	client = ec2.NewFromConfig(cfg)

	return client
}

func GetAllInstances(client *ec2.Client) ([]types.Instance, error) {
	var instances []types.Instance

	paginator := ec2.NewDescribeInstancesPaginator(client, &ec2.DescribeInstancesInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(context.Background())
		if err != nil {
			return nil, err
		}

		for _, reservation := range output.Reservations {
			for _, instance := range reservation.Instances {
				instances = append(instances, instance)
			}
		}
	}

	return instances, nil
}

// TODO condition
func GetEc2List() {

	instances, err := GetAllInstances(client)
	if err != nil {
		log.Fatalf("Unable to get instances, %v", err)
	}

	for _, instance := range instances {

		var name *string

		for _, tag := range instance.Tags {
			if *tag.Key == "Name" {
				name = tag.Value
				break
			}
		}

		fmt.Printf("Name : %-25s, ip : %s \n", *name, *instance.PrivateIpAddress)
		//fmt.Printf("Instance ID: %s, State: %s\n", *instance.InstanceId, instance.State.Name)
	}

}
