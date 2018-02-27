// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
//
// Example code for Load Balancing Service API
//

package example

func ExampleCreateLoadbalancer() {

	/*c, clerr := loadbalancer.NewLoadBalancerClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.LogIfError(clerr)

	request := loadbalancer.CreateLoadBalancerRequest{}
	request.CompartmentId = helpers.CompartmentID()
	request.DisplayName = common.String("OCI-GO-Sample-LB")

	vcn := CreateVcn(common.String("10.0.1.0/16"))
	subnet1 := CreateSubnet(vcn.Id, common.String("10.0.1.0/16"))

	availableDomain := helpers.AvailabilityDomain()

	subnet2 := createOrGetSubnetWithDetails(t, common.String(subnetDisplayName2), common.String("10.0.1.0/24"), common.String("subnetdns2"), availableDomain)
	request.SubnetIds = []string{*subnet1.Id, *subnet2.Id}

	shapes := listLoadBalancerShapes(t)
	request.ShapeName = shapes[0].Name

	resp, err := c.CreateLoadBalancer(context.Background(), request)
	helpers.LogIfError(err)*/

}
