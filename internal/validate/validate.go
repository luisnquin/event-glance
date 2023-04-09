package validate

import "regexp"

var rxLikeAWSRegion = regexp.MustCompile(`^[a-z]{2}(-gov)?-(central|(north|south)?(east|west)?)-\d{1,3}$`)

// Returns a boolean value indicating whether the provided region
// name is valid according to AWS naming conventions.
func LikeAWSRegion(region string) bool {
	return rxLikeAWSRegion.MatchString(region)
}
