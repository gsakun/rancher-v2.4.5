package client

import (
	"github.com/rancher/norman/types"
)

const (
	ApplicationType                      = "application"
	ApplicationFieldAnnotations          = "annotations"
	ApplicationFieldComponents           = "components"
	ApplicationFieldCreated              = "created"
	ApplicationFieldCreatorID            = "creatorId"
	ApplicationFieldLabels               = "labels"
	ApplicationFieldName                 = "name"
	ApplicationFieldNamespaceId          = "namespaceId"
	ApplicationFieldOptTraits            = "optTraits"
	ApplicationFieldOwnerReferences      = "ownerReferences"
	ApplicationFieldProjectID            = "projectId"
	ApplicationFieldRemoved              = "removed"
	ApplicationFieldState                = "state"
	ApplicationFieldStatus               = "status"
	ApplicationFieldTransitioning        = "transitioning"
	ApplicationFieldTransitioningMessage = "transitioningMessage"
	ApplicationFieldUUID                 = "uuid"
)

type Application struct {
	types.Resource
	Annotations          map[string]string      `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Components           []Component            `json:"components,omitempty" yaml:"components,omitempty"`
	Created              string                 `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID            string                 `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Labels               map[string]string      `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name                 string                 `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId          string                 `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`
	OptTraits            *ComponentTraitsForOpt `json:"optTraits,omitempty" yaml:"optTraits,omitempty"`
	OwnerReferences      []OwnerReference       `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	ProjectID            string                 `json:"projectId,omitempty" yaml:"projectId,omitempty"`
	Removed              string                 `json:"removed,omitempty" yaml:"removed,omitempty"`
	State                string                 `json:"state,omitempty" yaml:"state,omitempty"`
	Status               *ApplicationStatus     `json:"status,omitempty" yaml:"status,omitempty"`
	Transitioning        string                 `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage string                 `json:"transitioningMessage,omitempty" yaml:"transitioningMessage,omitempty"`
	UUID                 string                 `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type ApplicationCollection struct {
	types.Collection
	Data   []Application `json:"data,omitempty"`
	client *ApplicationClient
}

type ApplicationClient struct {
	apiClient *Client
}

type ApplicationOperations interface {
	List(opts *types.ListOpts) (*ApplicationCollection, error)
	ListAll(opts *types.ListOpts) (*ApplicationCollection, error)
	Create(opts *Application) (*Application, error)
	Update(existing *Application, updates interface{}) (*Application, error)
	Replace(existing *Application) (*Application, error)
	ByID(id string) (*Application, error)
	Delete(container *Application) error
}

func newApplicationClient(apiClient *Client) *ApplicationClient {
	return &ApplicationClient{
		apiClient: apiClient,
	}
}

func (c *ApplicationClient) Create(container *Application) (*Application, error) {
	resp := &Application{}
	err := c.apiClient.Ops.DoCreate(ApplicationType, container, resp)
	return resp, err
}

func (c *ApplicationClient) Update(existing *Application, updates interface{}) (*Application, error) {
	resp := &Application{}
	err := c.apiClient.Ops.DoUpdate(ApplicationType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ApplicationClient) Replace(obj *Application) (*Application, error) {
	resp := &Application{}
	err := c.apiClient.Ops.DoReplace(ApplicationType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *ApplicationClient) List(opts *types.ListOpts) (*ApplicationCollection, error) {
	resp := &ApplicationCollection{}
	err := c.apiClient.Ops.DoList(ApplicationType, opts, resp)
	resp.client = c
	return resp, err
}

func (c *ApplicationClient) ListAll(opts *types.ListOpts) (*ApplicationCollection, error) {
	resp := &ApplicationCollection{}
	resp, err := c.List(opts)
	if err != nil {
		return resp, err
	}
	data := resp.Data
	for next, err := resp.Next(); next != nil && err == nil; next, err = next.Next() {
		data = append(data, next.Data...)
		resp = next
		resp.Data = data
	}
	if err != nil {
		return resp, err
	}
	return resp, err
}

func (cc *ApplicationCollection) Next() (*ApplicationCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ApplicationCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ApplicationClient) ByID(id string) (*Application, error) {
	resp := &Application{}
	err := c.apiClient.Ops.DoByID(ApplicationType, id, resp)
	return resp, err
}

func (c *ApplicationClient) Delete(container *Application) error {
	return c.apiClient.Ops.DoResourceDelete(ApplicationType, &container.Resource)
}
