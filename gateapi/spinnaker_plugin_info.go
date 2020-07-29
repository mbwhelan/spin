/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type SpinnakerPluginInfo struct {

	Description string `json:"description,omitempty"`

	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	ProjectUrl string `json:"projectUrl,omitempty"`

	Provider string `json:"provider,omitempty"`

	Releases []SpinnakerPluginRelease `json:"releases"`

	RepositoryId string `json:"repositoryId,omitempty"`
}
