package client

import (
	"strconv"
	"strings"
)

const (
	// Authentication
	EndpointAdminToken = "/api/admin/token"

	// Admin Management
	EndpointAdmin              = "/api/admin"
	EndpointAdminByUsername    = "/api/admin/{username}"
	EndpointAdmins             = "/api/admins"
	EndpointAdminUsersDisable  = "/api/admin/{username}/users/disable"
	EndpointAdminUsersActivate = "/api/admin/{username}/users/activate"
	EndpointAdminUsageReset    = "/api/admin/usage/reset/{username}"
	EndpointAdminUsage         = "/api/admin/usage/{username}"

	// Core Management
	EndpointCore        = "/api/core"
	EndpointCoreRestart = "/api/core/restart"
	EndpointCoreConfig  = "/api/core/config"

	// Node Management
	EndpointNodeSettings  = "/api/node/settings"
	EndpointNode          = "/api/node"
	EndpointNodeByID      = "/api/node/{node_id}"
	EndpointNodes         = "/api/nodes"
	EndpointNodeReconnect = "/api/node/{node_id}/reconnect"
	EndpointNodesUsage    = "/api/nodes/usage"

	// Subscription
	EndpointSubscription           = "/sub/{token}/"
	EndpointSubscriptionInfo       = "/sub/{token}/info"
	EndpointSubscriptionUsage      = "/sub/{token}/usage"
	EndpointSubscriptionClientType = "/sub/{token}/{client_type}"

	// System
	EndpointSystem   = "/api/system"
	EndpointInbounds = "/api/inbounds"
	EndpointHosts    = "/api/hosts"

	// User Template
	EndpointUserTemplate     = "/api/user_template"
	EndpointUserTemplateByID = "/api/user_template/{template_id}"

	// User Management
	EndpointUser                  = "/api/user"
	EndpointUserByUsername        = "/api/user/{username}"
	EndpointUserReset             = "/api/user/{username}/reset"
	EndpointUserRevokeSubsription = "/api/user/{username}/revoke_sub"
	EndpointUsers                 = "/api/users"
	EndpointUsersReset            = "/api/users/reset"
	EndpointUserUsage             = "/api/user/{username}/usage"
	EndpointUserActiveNext        = "/api/user/{username}/active-next"
	EndpointUsersUsage            = "/api/users/usage"
	EndpointUserSetOwner          = "/api/user/{username}/set-owner"
	EndpointUsersExpired          = "/api/users/expired"

	// Base
	EndpointBase = "/"
)

// Helper functions for parameterized endpoints

// GetAdminByUsernameEndpoint returns the endpoint for a specific admin
func GetAdminByUsernameEndpoint(username string) string {
	return strings.Replace(EndpointAdminByUsername, "{username}", username, 1)
}

// GetAdminUsersDisableEndpoint returns the endpoint to disable all users for an admin
func GetAdminUsersDisableEndpoint(username string) string {
	return strings.Replace(EndpointAdminUsersDisable, "{username}", username, 1)
}

// GetAdminUsersActivateEndpoint returns the endpoint to activate all users for an admin
func GetAdminUsersActivateEndpoint(username string) string {
	return strings.Replace(EndpointAdminUsersActivate, "{username}", username, 1)
}

// GetAdminUsageResetEndpoint returns the endpoint to reset admin usage
func GetAdminUsageResetEndpoint(username string) string {
	return strings.Replace(EndpointAdminUsageReset, "{username}", username, 1)
}

// GetAdminUsageEndpoint returns the endpoint to get admin usage
func GetAdminUsageEndpoint(username string) string {
	return strings.Replace(EndpointAdminUsage, "{username}", username, 1)
}

// GetNodeByIDEndpoint returns the endpoint for a specific node
func GetNodeByIDEndpoint(nodeID int) string {
	return strings.Replace(EndpointNodeByID, "{node_id}", strconv.Itoa(nodeID), 1)
}

// GetNodeReconnectEndpoint returns the endpoint to reconnect a node
func GetNodeReconnectEndpoint(nodeID int) string {
	return strings.Replace(EndpointNodeReconnect, "{node_id}", strconv.Itoa(nodeID), 1)
}

// GetSubscriptionEndpoint returns the subscription endpoint for a token
func GetSubscriptionEndpoint(token string) string {
	return strings.Replace(EndpointSubscription, "{token}", token, 1)
}

// GetSubscriptionInfoEndpoint returns the subscription info endpoint for a token
func GetSubscriptionInfoEndpoint(token string) string {
	return strings.Replace(EndpointSubscriptionInfo, "{token}", token, 1)
}

// GetSubscriptionUsageEndpoint returns the subscription usage endpoint for a token
func GetSubscriptionUsageEndpoint(token string) string {
	return strings.Replace(EndpointSubscriptionUsage, "{token}", token, 1)
}

// GetSubscriptionClientTypeEndpoint returns the subscription endpoint for a specific client type
func GetSubscriptionClientTypeEndpoint(token, clientType string) string {
	endpoint := strings.Replace(EndpointSubscriptionClientType, "{token}", token, 1)
	return strings.Replace(endpoint, "{client_type}", clientType, 1)
}

// GetUserTemplateByIDEndpoint returns the endpoint for a specific user template
func GetUserTemplateByIDEndpoint(templateID int) string {
	return strings.Replace(EndpointUserTemplateByID, "{template_id}", strconv.Itoa(templateID), 1)
}

// GetUserByUsernameEndpoint returns the endpoint for a specific user
func GetUserByUsernameEndpoint(username string) string {
	return strings.Replace(EndpointUserByUsername, "{username}", username, 1)
}

// GetUserResetEndpoint returns the endpoint to reset user data usage
func GetUserResetEndpoint(username string) string {
	return strings.Replace(EndpointUserReset, "{username}", username, 1)
}

// GetUserRevokeSubscriptionEndpoint returns the endpoint to revoke user subscription
func GetUserRevokeSubscriptionEndpoint(username string) string {
	return strings.Replace(EndpointUserRevokeSubsription, "{username}", username, 1)
}

// GetUserUsageEndpoint returns the endpoint to get user usage
func GetUserUsageEndpoint(username string) string {
	return strings.Replace(EndpointUserUsage, "{username}", username, 1)
}

// GetUserActiveNextEndpoint returns the endpoint to activate next plan
func GetUserActiveNextEndpoint(username string) string {
	return strings.Replace(EndpointUserActiveNext, "{username}", username, 1)
}

// GetUserSetOwnerEndpoint returns the endpoint to set user owner
func GetUserSetOwnerEndpoint(username string) string {
	return strings.Replace(EndpointUserSetOwner, "{username}", username, 1)
}
