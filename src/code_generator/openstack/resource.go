package openstack

// resource in openstack
// server, img ...
type ResourceInfo struct {
	ResourceName string
	Actions      []ActionInfo
	ImportPaths  []string
}

func NewResourceInfo(resourceName string) *ResourceInfo {
	ri := &ResourceInfo{
		ResourceName: resourceName,
	}
	ri.Actions = make([]ActionInfo, 0)
	ri.ImportPaths = make([]string, 0)
	return ri
}

// describe a action to the resouce
// list, get, create ...
type ActionInfo struct {
	ActionName       string
	ActionParameters map[string]string
	ActionReturns    map[string]string
}

func NewActionInfo(actioonName string) *ActionInfo {
	ai := &ActionInfo{
		ActionName: actioonName,
	}
	ai.ActionParameters = make(map[string]string, 0)
	ai.ActionReturns = make(map[string]string, 0)
	return ai
}
