# Matrix Permutation Finder

Prints every permutation of HTTP matrix parameters

# Usage

```
% go build
% ./matrix-permutation-printer -path ";m=exact;h=([0-9]+);w=([0-9]+)"
/v1/asset/(.*)/resize;h=([0-9]+);m=exact;w=([0-9]+)
/v1/asset/(.*)/resize;h=([0-9]+);w=([0-9]+);m=exact
/v1/asset/(.*)/resize;m=exact;h=([0-9]+);w=([0-9]+)
/v1/asset/(.*)/resize;m=exact;w=([0-9]+);h=([0-9]+)
/v1/asset/(.*)/resize;w=([0-9]+);h=([0-9]+);m=exact
/v1/asset/(.*)/resize;w=([0-9]+);m=exact;h=([0-9]+)


For a total of 6 permutations
```
