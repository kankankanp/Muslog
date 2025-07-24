# BlogsApi

All URIs are relative to *http://localhost:8080*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**blogsGet**](#blogsget) | **GET** /blogs | Get all blogs|
|[**blogsIdDelete**](#blogsiddelete) | **DELETE** /blogs/{id} | Delete a blog|
|[**blogsIdGet**](#blogsidget) | **GET** /blogs/{id} | Get a blog by ID|
|[**blogsIdPut**](#blogsidput) | **PUT** /blogs/{id} | Update a blog|
|[**blogsPagePageGet**](#blogspagepageget) | **GET** /blogs/page/{page} | Get blogs by page|
|[**blogsPost**](#blogspost) | **POST** /blogs | Create a new blog|

# **blogsGet**
> { [key: string]: any; } blogsGet()

Get all blog posts

### Example

```typescript
import {
    BlogsApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new BlogsApi(configuration);

const { status, data } = await apiInstance.blogsGet();
```

### Parameters
This endpoint does not have any parameters.


### Return type

**{ [key: string]: any; }**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **blogsIdDelete**
> { [key: string]: any; } blogsIdDelete()

Delete a blog post by its ID

### Example

```typescript
import {
    BlogsApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new BlogsApi(configuration);

let id: number; //Blog ID (default to undefined)

const { status, data } = await apiInstance.blogsIdDelete(
    id
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **id** | [**number**] | Blog ID | defaults to undefined|


### Return type

**{ [key: string]: any; }**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **blogsIdGet**
> { [key: string]: any; } blogsIdGet()

Get a single blog post by its ID

### Example

```typescript
import {
    BlogsApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new BlogsApi(configuration);

let id: number; //Blog ID (default to undefined)

const { status, data } = await apiInstance.blogsIdGet(
    id
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **id** | [**number**] | Blog ID | defaults to undefined|


### Return type

**{ [key: string]: any; }**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **blogsIdPut**
> { [key: string]: any; } blogsIdPut(blog)

Update an existing blog post

### Example

```typescript
import {
    BlogsApi,
    Configuration,
    ModelPost
} from './api';

const configuration = new Configuration();
const apiInstance = new BlogsApi(configuration);

let id: number; //Blog ID (default to undefined)
let blog: ModelPost; //Blog post to update

const { status, data } = await apiInstance.blogsIdPut(
    id,
    blog
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **blog** | **ModelPost**| Blog post to update | |
| **id** | [**number**] | Blog ID | defaults to undefined|


### Return type

**{ [key: string]: any; }**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **blogsPagePageGet**
> { [key: string]: any; } blogsPagePageGet()

Get blog posts paginated

### Example

```typescript
import {
    BlogsApi,
    Configuration
} from './api';

const configuration = new Configuration();
const apiInstance = new BlogsApi(configuration);

let page: number; //Page number (default to undefined)

const { status, data } = await apiInstance.blogsPagePageGet(
    page
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **page** | [**number**] | Page number | defaults to undefined|


### Return type

**{ [key: string]: any; }**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **blogsPost**
> { [key: string]: any; } blogsPost(blog)

Create a new blog post

### Example

```typescript
import {
    BlogsApi,
    Configuration,
    ModelPost
} from './api';

const configuration = new Configuration();
const apiInstance = new BlogsApi(configuration);

let blog: ModelPost; //Blog post to create

const { status, data } = await apiInstance.blogsPost(
    blog
);
```

### Parameters

|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **blog** | **ModelPost**| Blog post to create | |


### Return type

**{ [key: string]: any; }**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
|**201** | Created |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

