/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PluginDependency struct {
	Optional bool `json:"optional,omitempty"`
	PluginId string `json:"pluginId,omitempty"`
	PluginVersionSupport string `json:"pluginVersionSupport,omitempty"`
}
