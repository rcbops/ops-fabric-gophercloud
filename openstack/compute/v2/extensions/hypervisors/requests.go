package hypervisors

import (
	"strconv"
	"log"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/pagination"
)

// List makes a request against the API to list hypervisors.
// TODO I bet the third argument here is wrong ... supposed to be a:
// func(r PageResult) Page. Look at how createPage is implemented in places
// where it works properly...
func List(client *gophercloud.ServiceClient) pagination.Pager {
	log.Println("Listing hypervisors...")
	return pagination.NewPager(client, hypervisorsListDetailURL(client), func(r pagination.PageResult) pagination.Page {
		log.Println("Constructing a hypervisor page from this pagination.PageResult:")
		log.Println(r)
		return HypervisorPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Statistics makes a request against the API to get hypervisors statistics.
func GetStatistics(client *gophercloud.ServiceClient) (r StatisticsResult) {
	_, r.Err = client.Get(hypervisorsStatisticsURL(client), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Get makes a request against the API to get details for specific hypervisor.
func Get(client *gophercloud.ServiceClient, hypervisorID int) (r HypervisorResult) {
	v := strconv.Itoa(hypervisorID)
	_, r.Err = client.Get(hypervisorsGetURL(client, v), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// GetUptime makes a request against the API to get uptime for specific hypervisor.
func GetUptime(client *gophercloud.ServiceClient, hypervisorID int) (r UptimeResult) {
	v := strconv.Itoa(hypervisorID)
	_, r.Err = client.Get(hypervisorsUptimeURL(client, v), &r.Body, &gophercloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
