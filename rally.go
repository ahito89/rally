package rally

import (
	"fmt"
	"net/url"
)

var Debug = false

type Client struct {
	apikey string
}

func New(apikey string) *Client {
	return &Client{
		apikey: apikey,
	}
}

func (r *Client) GetHierarchicalRequirement(objectID string) (*HierarchicalRequirement, error) {
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

func (r *Client) QueryHierarchicalRequirement(formattedID string) (*HierarchicalRequirement, error) {
	params := make(url.Values)
	params.Add("fetch", "true")
	params.Add("query", fmt.Sprintf("(FormattedID = %s)", formattedID))

	var out struct {
		QueryResult struct {
			Results          []HierarchicalRequirement
			TotalResultCount int
		}
	}

	_, err := r.getRequest("hierarchicalrequirement", params, &out)
	if err != nil {
		return nil, err
	}

	if total := out.QueryResult.TotalResultCount; total != 1 {
		return nil, fmt.Errorf("not found, or found too many (%v)", total)
	}

	//jsonOut, _ := json.MarshalIndent(out, "", "  ")
	//os.Stdout.Write(jsonOut)
	return &out.QueryResult.Results[0], nil
}

func (r *Client) QueryPortfolioItemFeature(formattedID string) (*PortfolioItemFeature, error) {
	params := make(url.Values)
	params.Add("fetch", "true")
	params.Add("query", fmt.Sprintf("(FormattedID = %s)", formattedID))

	var out struct {
		QueryResult struct {
			Results          []PortfolioItemFeature
			TotalResultCount int
		}
	}

	_, err := r.getRequest("portfolioitem/feature", params, &out)
	if err != nil {
		return nil, err
	}

	if total := out.QueryResult.TotalResultCount; total != 1 {
		return nil, fmt.Errorf("not found, or found too many (%v)", total)
	}

	//jsonOut, _ := json.MarshalIndent(out, "", "  ")
	//os.Stdout.Write(jsonOut)
	return &out.QueryResult.Results[0], nil
}

func (r *Client) QueryDefect(formattedID string) (*Defect, error) {
	params := make(url.Values)
	params.Add("fetch", "true")
	params.Add("query", fmt.Sprintf("(FormattedID = %s)", formattedID))

	var out struct {
		QueryResult struct {
			Results          []Defect
			TotalResultCount int
		}
	}

	_, err := r.getRequest("defect", params, &out)
	if err != nil {
		return nil, err
	}

	if total := out.QueryResult.TotalResultCount; total != 1 {
		return nil, fmt.Errorf("not found, or found too many (%v)", total)
	}

	return &out.QueryResult.Results[0], nil
}
