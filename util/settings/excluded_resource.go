package settings

import (
	"github.com/gobwas/glob"
	log "github.com/sirupsen/logrus"
)

type ExcludedResource struct {
	ApiGroups []string `json:"apiGroups"`
	Kinds     []string `json:"kinds"`
	Clusters  []string `json:"clusters"`
}

func (r ExcludedResource) matchGroup(apiGroup string) bool {
	for _, excludedApiGroup := range r.ApiGroups {
		if match(excludedApiGroup, apiGroup) {
			return true
		}
	}
	return false
}

func match(pattern, text string) bool {
	compiledGlob, err := glob.Compile(pattern)
	if err != nil {
		log.Warnf("failed to compile pattern %s due to error %v", pattern, err)
		return false
	}
	return compiledGlob.Match(text)
}

func (r ExcludedResource) matchKind(kind string) bool {
	for _, excludedKind := range r.Kinds {
		if excludedKind == "*" || excludedKind == kind {
			return true
		}
	}
	return false
}

func (r ExcludedResource) matchCluster(cluster string) bool {
	for _, excludedCluster := range r.Clusters {
		if match(excludedCluster, cluster) {
			return true
		}
	}
	return false
}

func (r ExcludedResource) Match(apiGroup, kind, cluster string) bool {
	return r.matchGroup(apiGroup) && r.matchKind(kind) && r.matchCluster(cluster)
}
