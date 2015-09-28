package rally

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)

var Debug = false

type Rally struct {
	apikey string
}

func New(apikey string) *Rally {
	return &Rally{
		apikey: apikey,
	}
}

func (r *Rally) GetHierarchicalRequirement(objectID string) (*HierarchicalRequirement, error) {
	var hr struct {
		HierarchicalRequirement *HierarchicalRequirement
	}
	params := make(url.Values)
	params.Add("fetch", "Name,FormattedID,Tasks")
	_, err := r.getRequest(fmt.Sprintf("hierarchicalrequirement/%s", objectID), params, &hr)
	if err != nil {
		return nil, err
	}

	return hr.HierarchicalRequirement, nil
}

func (r *Rally) QueryHierarchicalRequirement(formattedID string) (*HierarchicalRequirement, error) {
	var out interface{}
	params := make(url.Values)
	params.Add("query", fmt.Sprintf("(FormattedID = %s)", formattedID))
	_, err := r.getRequest("hierarchicalrequirement", params, &out)
	if err != nil {
		return nil, err
	}
	jsonOut, _ := json.MarshalIndent(out, "", "  ")
	os.Stdout.Write(jsonOut)
	return nil, nil
}

func (r *Rally) QueryPortfolioItemFeature(formattedID string) (*PortfolioItemFeature, error) {
	var out interface{}
	params := make(url.Values)
	params.Add("fetch", "true")
	params.Add("query", fmt.Sprintf("(FormattedID = %s)", formattedID))
	_, err := r.getRequest("portfolioitem/feature", params, &out)
	if err != nil {
		return nil, err
	}
	jsonOut, _ := json.MarshalIndent(out, "", "  ")
	os.Stdout.Write(jsonOut)
	return nil, nil
}
