# how to run 

1.) `doc-rob.sig` has the signature for the following commit https://github.com/lemiesz/trusty/commit/9334dcc3a649676078fdc2621c3057f579cca097

2.) `commit-rob.txt` has the commit data 

3.) `rob-pub.key` has the public key this was signed with 

to verify commit first setup go dependencys 

`go mod vendor`  

Then run `go run test.go`
