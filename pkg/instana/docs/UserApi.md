# \UserApi

All URIs are relative to *https://unit-tenant.instana.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInvitations**](UserApi.md#GetInvitations) | **Get** /api/settings/users/invitations | All pending invitations
[**GetUserById**](UserApi.md#GetUserById) | **Get** /api/settings/users/{userId} | Get single user
[**GetUsers**](UserApi.md#GetUsers) | **Get** /api/settings/users | All users (without invitations)
[**GetUsersIncludingInvitations**](UserApi.md#GetUsersIncludingInvitations) | **Get** /api/settings/users/overview | All users (incl. invitations)
[**RemoveUserFromTenant**](UserApi.md#RemoveUserFromTenant) | **Delete** /api/settings/users/{userId} | Remove user from tenant
[**RevokePendingInvitations**](UserApi.md#RevokePendingInvitations) | **Delete** /api/settings/users/invitations | Revoke pending invitation
[**SendInvitation**](UserApi.md#SendInvitation) | **Post** /api/settings/users/invitations | Send user invitation
[**UpdateUser**](UserApi.md#UpdateUser) | **Put** /api/settings/users/{email} | Change user name of single user



## GetInvitations

> []UserResult GetInvitations(ctx, )

All pending invitations

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]UserResult**](UserResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserById

> []UserResult GetUserById(ctx, userId)

Get single user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**|  | 

### Return type

[**[]UserResult**](UserResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> []UserResult GetUsers(ctx, )

All users (without invitations)

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]UserResult**](UserResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersIncludingInvitations

> UsersResult GetUsersIncludingInvitations(ctx, )

All users (incl. invitations)

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**UsersResult**](UsersResult.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUserFromTenant

> RemoveUserFromTenant(ctx, userId)

Remove user from tenant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokePendingInvitations

> RevokePendingInvitations(ctx, optional)

Revoke pending invitation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RevokePendingInvitationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RevokePendingInvitationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **optional.String**|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendInvitation

> SendInvitation(ctx, optional)

Send user invitation

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SendInvitationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SendInvitationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | [**optional.Interface of []string**](string.md)|  | 
 **roleId** | [**optional.Interface of []string**](string.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> UpdateUser(ctx, email, optional)

Change user name of single user

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string**|  | 
 **optional** | ***UpdateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **editUser** | [**optional.Interface of EditUser**](EditUser.md)|  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

