terraform-hashcode
===

For a given string, compute and output the value of [hashcode.String](https://github.com/hashicorp/terraform-plugin-sdk/blob/master/internal/helper/hashcode/hashcode.go). This utility can be convenient when building tests against provider schemas that include sets.

Install:

```
$ go get -u github.com/roberth-k/terraform-hashcode
```

Usage:

```
$ terraform-hashcode '2.2.0.0/16-Test-'
2841319978
```
