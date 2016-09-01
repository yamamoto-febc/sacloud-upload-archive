package api

/************************************************
  generated by IDE. for [InternetAPI]
************************************************/

import (
	"github.com/yamamoto-febc/libsacloud/sacloud"
)

/************************************************
   To support influent interface for Find()
************************************************/

func (api *InternetAPI) Reset() *InternetAPI {
	api.reset()
	return api
}

func (api *InternetAPI) Offset(offset int) *InternetAPI {
	api.offset(offset)
	return api
}

func (api *InternetAPI) Limit(limit int) *InternetAPI {
	api.limit(limit)
	return api
}

func (api *InternetAPI) Include(key string) *InternetAPI {
	api.include(key)
	return api
}

func (api *InternetAPI) Exclude(key string) *InternetAPI {
	api.exclude(key)
	return api
}

func (api *InternetAPI) FilterBy(key string, value interface{}) *InternetAPI {
	api.filterBy(key, value, false)
	return api
}

// func (api *InternetAPI) FilterMultiBy(key string, value interface{}) *InternetAPI {
// 	api.filterBy(key, value, true)
// 	return api
// }

func (api *InternetAPI) WithNameLike(name string) *InternetAPI {
	return api.FilterBy("Name", name)
}

func (api *InternetAPI) WithTag(tag string) *InternetAPI {
	return api.FilterBy("Tags.Name", tag)
}
func (api *InternetAPI) WithTags(tags []string) *InternetAPI {
	return api.FilterBy("Tags.Name", []interface{}{tags})
}

// func (api *InternetAPI) WithSizeGib(size int) *InternetAPI {
// 	api.FilterBy("SizeMB", size*1024)
// 	return api
// }

// func (api *InternetAPI) WithSharedScope() *InternetAPI {
// 	api.FilterBy("Scope", "shared")
// 	return api
// }

// func (api *InternetAPI) WithUserScope() *InternetAPI {
// 	api.FilterBy("Scope", "user")
// 	return api
// }

func (api *InternetAPI) SortBy(key string, reverse bool) *InternetAPI {
	api.sortBy(key, reverse)
	return api
}

func (api *InternetAPI) SortByName(reverse bool) *InternetAPI {
	api.sortByName(reverse)
	return api
}

// func (api *InternetAPI) SortBySize(reverse bool) *InternetAPI {
// 	api.sortBy("SizeMB", reverse)
// 	return api
// }

/************************************************
  To support CRUD(Create/Read/Update/Delete)
************************************************/

func (api *InternetAPI) New() *sacloud.Internet {
	return &sacloud.Internet{}
}

func (api *InternetAPI) Create(value *sacloud.Internet) (*sacloud.Internet, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.create(api.createRequest(value), res)
	})
}

func (api *InternetAPI) Read(id string) (*sacloud.Internet, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.read(id, nil, res)
	})
}

func (api *InternetAPI) Update(id string, value *sacloud.Internet) (*sacloud.Internet, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.update(id, api.createRequest(value), res)
	})
}

func (api *InternetAPI) Delete(id string) (*sacloud.Internet, error) {
	return api.request(func(res *sacloud.Response) error {
		return api.delete(id, nil, res)
	})
}

/************************************************
  Inner functions
************************************************/

func (api *InternetAPI) setStateValue(setFunc func(*sacloud.Request)) *InternetAPI {
	api.baseAPI.setStateValue(setFunc)
	return api
}

func (api *InternetAPI) request(f func(*sacloud.Response) error) (*sacloud.Internet, error) {
	res := &sacloud.Response{}
	err := f(res)
	if err != nil {
		return nil, err
	}
	return res.Internet, nil
}

func (api *InternetAPI) createRequest(value *sacloud.Internet) *sacloud.Request {
	req := &sacloud.Request{}
	req.Internet = value
	return req
}
