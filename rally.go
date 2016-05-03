package rally

import (
	"fmt"
	"net/url"
	"strconv"
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

// QueryHierarchicalRequirement query the hierarchicalrequirement
func (r *Client) QueryHierarchicalRequirement(formattedID string, startIndex int) (*[]HierarchicalRequirement, int, error) {
	if startIndex < 1 {
		startIndex = 1
	}
	params := make(url.Values)
	params.Add("fetch", "true")
	if formattedID != "" {		
		params.Add("query", fmt.Sprintf("(FormattedID = %s)", formattedID))
	} else {
		params.Add("query", "")
	}
	params.Add("start", strconv.Itoa(startIndex))

	var out struct {
		QueryResult struct {
			Results          []HierarchicalRequirement
			TotalResultCount int
		}
	}

	_, err := r.getRequest("hierarchicalrequirement", params, &out)
	if err != nil {
		return nil, 0, err
	}

	/*if total := out.QueryResult.TotalResultCount; total != 1 {
		return nil, fmt.Errorf("not found, or found too many (%v)", total)
	}*/

	//jsonOut, _ := json.MarshalIndent(out, "", "  ")
	//os.Stdout.Write(jsonOut)
	return &out.QueryResult.Results, out.QueryResult.TotalResultCount, nil
}

// QueryPortfolioItemFeature query the PortfolioItemFeature
func (r *Client) QueryPortfolioItemFeature(formattedID string) (*PortfolioItemFeature, int, error) {
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
		return nil, 0, err
	}

	if total := out.QueryResult.TotalResultCount; total != 1 {
		return nil, out.QueryResult.TotalResultCount, fmt.Errorf("not found, or found too many (%v)", total)
	}

	//jsonOut, _ := json.MarshalIndent(out, "", "  ")
	//os.Stdout.Write(jsonOut)
	return &out.QueryResult.Results[0], out.QueryResult.TotalResultCount, nil
}

// QueryTestCase query the test case
func (r *Client) QueryTestCase(formattedID string, startIndex int) (*[]TestCase, int, error) {
	if startIndex < 1 {
		startIndex = 1
	}
    params := make(url.Values)
    params.Add("fetch", "true")
	if formattedID != "" { 
    	params.Add("query", fmt.Sprintf("(FormattedID = \"%s\")", formattedID))
	} else {
		params.Add("query", "")
	}
	params.Add("start", strconv.Itoa(startIndex))
    
    var out struct {
        QueryResult struct {
            Results          []TestCase
            TotalResultCount int
        }
    }
    
    _, err := r.getRequest("testcase", params, &out)
    if err != nil {
        return nil, 0, err
    }
    
 /*   if total := out.QueryResult.TotalResultCount; total != 1 {
        return nil, fmt.Errorf("not found, or found too many (%v)", total)
    }*/
    
    return &out.QueryResult.Results, out.QueryResult.TotalResultCount, nil
}

// QueryTestCaseSteps query the test case step
func (r *Client) QueryTestCaseSteps(testCase string, startIndex int) (*[]TestCaseStep, int, error) {
	if startIndex < 1 {
		startIndex = 1
	}
	
    params := make(url.Values)
    params.Add("fetch", "true")
    params.Add("query", fmt.Sprintf("(TestCase = \"%s\")", testCase))
    params.Add("start", strconv.Itoa(startIndex))
	
    var out struct {
        QueryResult struct {
            Results          []TestCaseStep
            TotalResultCount int
        }
    }
    
    _, err := r.getRequest("testcasestep", params, &out)
    if err != nil {
        return nil, 0, err
    }
    
 /*   if total := out.QueryResult.TotalResultCount; total != 1 {
        return nil, fmt.Errorf("not found, or found too many (%v)", total)
    }*/
    
    return &out.QueryResult.Results, out.QueryResult.TotalResultCount, nil
}

// QueryIteration query the iteration
func (r *Client) QueryIteration(formattedID string, startIndex int) (*[]Iteration, int, error) {
	if startIndex < 1 {
		startIndex = 1
	}
    params := make(url.Values)
    params.Add("fetch", "true")
    if formattedID != "" {
    	params.Add("query", fmt.Sprintf("(Name = \"%s\")", formattedID))
	} else {
		params.Add("query","")
	}
    params.Add("start", strconv.Itoa(startIndex))
    var out struct {
        QueryResult struct {
            Results          []Iteration
            TotalResultCount int
        }
    }
    
    _, err := r.getRequest("iteration", params, &out)
    if err != nil {
        return nil, 0, err
    }
    
 /*   if total := out.QueryResult.TotalResultCount; total != 1 {
        return nil, fmt.Errorf("not found, or found too many (%v)", total)
    }*/
    
    return &out.QueryResult.Results, out.QueryResult.TotalResultCount, nil
}

// QueryProject query for project based on project name
func (r *Client) QueryProject(formattedID string, startIndex int) (*[]Project, int, error) {
	if startIndex < 1 {
		startIndex = 1
	}
    params := make(url.Values)
    params.Add("fetch", "true")
	if formattedID != "" {
    	params.Add("query", fmt.Sprintf("(Name = \"%s\")", formattedID))
	} else {
		params.Add("query","")
	}
    params.Add("start",strconv.Itoa(startIndex))
    var out struct {
        QueryResult struct {
            Results          []Project
            TotalResultCount int
        }
    }
    
    _, err := r.getRequest("project", params, &out)
    if err != nil {
        return nil, 0, err
    }
    /*
    if total := out.QueryResult.TotalResultCount; total != 1 {
        return nil, fmt.Errorf("not found, or found too many (%v)", total)
    }*/
    
    return &out.QueryResult.Results, out.QueryResult.TotalResultCount, nil
}

// QueryDefect query the Defect
func (r *Client) QueryDefect(formattedID string) (*[]Defect, int, error) {
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
		return nil, 0, err
	}

	/*if total := out.QueryResult.TotalResultCount; total != 1 {
		return nil, fmt.Errorf("not found, or found too many (%v)", total)
	}*/

	return &out.QueryResult.Results, out.QueryResult.TotalResultCount, nil
}
