package rally

import (
	"fmt"
	"net/url"
)

// Debug mode prints messages
var Debug = false

// Client Definition
type Client struct {
	apikey string
    username string
    password string
}

// New rally object
func New(apikey string, username string, password string) *Client {
	return &Client{
		apikey: apikey,
        username: username,
        password: password,
	}
}

// GetHierarchicalRequirement gets the HierarchicalRequirement by object ID
func (r *Client) GetHierarchicalRequirement(objectID int64) (*HierarchicalRequirement, error) {
	var hr struct {
		HierarchicalRequirement *HierarchicalRequirement
	}
	params := make(url.Values)
	params.Add("fetch", "Name,FormattedID,Tasks")
	_, err := r.getRequest(fmt.Sprintf("hierarchicalrequirement/%d", objectID), params, &hr)
	if err != nil {
		return nil, err
	}

	return hr.HierarchicalRequirement, nil
}

// QueryHierarchicalRequirement get the hierarchicalrequirement
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

// QueryPortfolioItemFeature get the PortfolioItemFeature
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

// QueryDefect get the Defect
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
