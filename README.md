# Simple checks for PagerDuty

Simple show case on connection to PagetDuty, quering on current incidents and performing simple operations

Requires your `userID`, `serviceID` and `authtoken`, with you can add in the separate file `credentials.go`, which is not checked in and ignored by git for obvious reasons.
To get those values go to your profile page.
* `userID` would be a string in browser's address bar, after `/users/`
* `authToken` can be created in `API Access` section of profile
* `serviceID` can be found in browse's address bar when you view details page of the selected service, after `/services/`
