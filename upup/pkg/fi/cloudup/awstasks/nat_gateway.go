package awstasks

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"k8s.io/kops/upup/pkg/fi/cloudup/awsup"
)

//go:generate fitask -type=NatGateway
type NatGateway struct {
	Name      *string
	ID        *string
	Subnet    *Subnet
	ElasticIP *ElasticIP
}

func (s *NatGateway) CompareWithID() *string {
	return s.ID
}

func findNatGateway(cloud *awsup.AWSCloud, request *ec2.DescribeNatGatewaysInput) (*ec2.NatGateway, error) {
	return nil, nil
}
