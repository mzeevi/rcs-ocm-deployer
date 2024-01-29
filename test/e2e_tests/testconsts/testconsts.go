package testconsts

import "time"

const (
	TimeoutCapp              = 60 * time.Second
	CappCreationInterval     = 2 * time.Second
	DefaultEventuallySeconds = 2
)

const (
	CappPlacement = "test-placement"
	CappCluster   = "cluster1"
)
