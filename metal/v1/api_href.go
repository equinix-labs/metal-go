package v1

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type HrefApiService service

type href interface {
	GetHref() string
}

type HrefRequest struct {
	ctx        context.Context
	ApiService *HrefApiService
	model      href
	include    *[]string
	exclude    *[]string
}

// Nested attributes to include. Included objects will return their full attributes. Attribute names can be dotted (up to 3 levels) to included deeply nested objects.
func (r HrefRequest) Include(include []string) HrefRequest {
	r.include = &include
	return r
}

// Nested attributes to exclude. Excluded objects will return only the href attribute. Attribute names can be dotted (up to 3 levels) to exclude deeply nested objects.
func (r HrefRequest) Exclude(exclude []string) HrefRequest {
	r.exclude = &exclude
	return r
}

func (r HrefRequest) Execute() (*http.Response, error) {
	return r.ApiService.FindByHrefExecute(r)
}

func (a *HrefApiService) FindByHref(ctx context.Context, model href) HrefRequest {
	return HrefRequest{
		ApiService: a,
		ctx:        ctx,
		model:      model,
	}
}

func (a *HrefApiService) FindByHrefExecute(r HrefRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}
	localVarPath := localBasePath + strings.TrimPrefix(r.model.GetHref(), "/metal/v1")

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.include != nil {
		parameterAddToQuery(localVarQueryParams, "include", r.include, "csv")
	}

	if r.exclude != nil {
		parameterAddToQuery(localVarQueryParams, "exclude", r.exclude, "csv")
	}

	request, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	response, err := a.client.callAPI(request)
	if err != nil {
		return response, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return response, err
	}

	a.client.decode(r.model, body, response.Header.Get("Content-Type"))

	return response, nil
}
