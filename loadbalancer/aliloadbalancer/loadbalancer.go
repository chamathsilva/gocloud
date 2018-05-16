package aliloadbalancer

import (
	"strconv"
	"github.com/cloudlibz/gocloud/aliauth"
)

//TODO
func (aliloadbalancer *Aliloadbalancer) Creatloadbalancer(request interface{}) (resp interface{}, err error) {
	var options CreateLoadBalancer

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "RegionId":
			regionID, _ := value.(string)
			options.RegionID = regionID
		case "MasterZoneId":
			masterZoneID, _ := value.(string)
			options.MasterZoneID = masterZoneID
		case "SlaveZoneId":
			slaveZoneID, _ := value.(string)
			options.SlaveZoneID = slaveZoneID
		case "LoadBalancerName":
			loadBalancerName, _ := value.(string)
			options.LoadBalancerName = loadBalancerName
		case "AddressType":
			addressType, _ := value.(string)
			options.AddressType = addressType
		case "VSwitchId":
			VSwitchID, _ := value.(string)
			options.VSwitchID = VSwitchID
		case "PayType":
			payType, _ := value.(string)
			options.PayType = payType
		case "PricingCycle":
			PricingCycle, _ := value.(string)
			options.PricingCycle = PricingCycle
		case "Duration":
			Duration, _ := value.(string)
			options.Duration = Duration
		case "Autopay":
			switch value.(type) {
			case bool:
				options.Autopay = value.(bool)
			case string:
				options.Autopay = value.(string) == "true" || value.(string) == "True"
			}
		case "InternetChargeType":
			InternetChargeType, _ := value.(string)
			options.InternetChargeType = InternetChargeType
		case "Bandwidth":
			switch value.(type) {
			case int:
				options.Bandwidth = value.(int)
			case string:
				options.Bandwidth, _ = strconv.Atoi(value.(string))
			}
		case "ClientToken":
			ClientToken, _ := value.(string)
			options.ClientToken = ClientToken
		case "ResourceGroupID":
			ResourceGroupID, _ := value.(string)
			options.ResourceGroupID = ResourceGroupID
		}
	}

	// Put all of options into params
	params := aliauth.PutStructIntoMap(&options)

	response := make(map[string]interface{})
	err = aliauth.LoadBalancerSignAndDoRequest("CreateLoadBalancer", params, response)
	resp = response
	return resp, err
}

//TODO
func (aliloadbalancer *Aliloadbalancer) Deleteloadbalancer(request interface{}) (resp interface{}, err error) {
	var options DeleteLoadBalancer

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "RegionId":
			regionID, _ := value.(string)
			options.RegionID = regionID
		case "LoadBalancerId":
			LoadBalancerID, _ := value.(string)
			options.LoadBalancerID = LoadBalancerID
		}
	}

	// Put all of options into params
	params := aliauth.PutStructIntoMap(&options)

	response := make(map[string]interface{})
	err = aliauth.LoadBalancerSignAndDoRequest("DeleteLoadBalancer", params, response)
	resp = response
	return resp, err
}

//TODO
func (aliloadbalancer *Aliloadbalancer) Listloadbalancer(request interface{}) (resp interface{}, err error) {
	var options ListLoadBalancer

	param := make(map[string]interface{})

	param = request.(map[string]interface{})
	params := make(map[string]interface{})

	for key, value := range param {
		switch key {
		case "RegionId":
			regionID, _ := value.(string)
			options.RegionID = regionID
		case "LoadBalancerId":
			LoadBalancerID, _ := value.(string)
			options.LoadBalancerID = LoadBalancerID
		default:
			switch value.(type) {
			case string:
				params[key] = value.(string)
			case int:
				params[key] = strconv.Itoa(value.(int))
			case bool:
				params[key] = strconv.FormatBool(value.(bool))
			}
		}
	}

	// Put all of options into params
	params = aliauth.PutStructIntoMap(&options)

	response := make(map[string]interface{})
	err = aliauth.LoadBalancerSignAndDoRequest("DescribeLoadBalancers", params, response)
	resp = response
	return resp, err
}

//TODO
func (aliloadbalancer *Aliloadbalancer) Detachnodewithloadbalancer(request interface{}) (resp interface{}, err error) {
	return resp, err
}

//TODO
func (aliloadbalancer *Aliloadbalancer) Attachnodewithloadbalancer(request interface{}) (resp interface{}, err error) {
	return resp, err
}
