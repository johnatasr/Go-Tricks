// BAD
var delayMillis int64 = 15000

// GOOD
var delay time.Duration = 15 * time.Second