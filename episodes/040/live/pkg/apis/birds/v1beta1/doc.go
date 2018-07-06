// Api versions allow the api contract for a resource to be changed while keeping
// backward compatibility by support multiple concurrent versions
// of the same resource

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=github.com/heptio/tgik/episodes/040/live/pkg/apis/birds
// +k8s:defaulter-gen=TypeMeta
// +groupName=birds.fabulous.af
package v1beta1 // import "github.com/heptio/tgik/episodes/040/live/pkg/apis/birds/v1beta1"
