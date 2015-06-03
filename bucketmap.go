package deliverytools

// BucketMap represents the mapping of prefixes to buckets
type BucketMap struct {
	// The default or fallthrough bucket
	Default string

	// Bucket mappings
	Mounts []BucketMount
}

// BucketMount is a prefix -> bucket mapping
type BucketMount struct {
	Prefix string
	Bucket string
}

// BucketMapping is the current BucketMap
var ProdBucketMap = BucketMap{
	Default: "archive",
	Mounts: []BucketMount{
		BucketMount{"firefox/try-builds/", "firefox-try"},
		BucketMount{"firefox/", "firefox"},
		BucketMount{"mobile/try-builds/", "firefox-android-try"},
		BucketMount{"mobile/", "firefox-android"},
		BucketMount{"opus/", "opus"},
		BucketMount{"thunderbird/try-builds/", "thunderbird-try"},
		BucketMount{"thunderbird/", "thunderbird"},
		BucketMount{"xulrunner/try-builds/", "xulrunner-try"},
		BucketMount{"xulrunner/", "xulrunner"},
	},
}