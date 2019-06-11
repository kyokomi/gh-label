gh-label
============================
A github cli tool only to label for golang

## Install

```bash
go get github.com/kyokomi/gh-label
```

## Usage

```bash
$ gh-label -h
Usage of ./gh-lebel:
  -b string
    	github branch name
  -c string
    	comment
  -o string
    	github owner name
  -r string
    	github repository name
  -t string
    	github access token
```

## Example

```
$ gh-label \
	-t "<github token>" \
	-o "kyokomi" \
	-r "github.com/kyokomi/gh-label" \
	-b "feature/test" \
	-l "test"
```


## License

[MIT](/LICENSE)
