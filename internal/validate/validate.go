package validate

import "regexp"

var (
	rxLikeAWSRegion = regexp.MustCompile(`^[a-z]{2}(-gov)?-(central|(north|south)?(east|west)?)-\d{1,3}$`)
	rxEmail         = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// Returns a boolean value indicating whether the provided region
// name is valid according to AWS naming conventions.
func LikeAWSRegion(region string) bool {
	return rxLikeAWSRegion.MatchString(region)
}

// Checks whether a given string is a valid email address.
func IsEmail(email string) bool {
	return rxEmail.MatchString(email)
}
