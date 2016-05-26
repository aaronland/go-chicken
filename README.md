# go-chicken

🐔 🐔 🐔

## Caveats

It's early days, still. There is lots left to do. `chicken` is not yet able to distinguish between actual words and things like punctuation or numbers. It will, but not today.

## Install

```
make build
```

This will build a copy of `chicken` and put it in the [bin](bin) directory.

## Usage

```
$> ./bin/chicken -h
Usage of ./bin/chicken:
  -language string
      	    A valid ISO-639-3 language code. (default "zxx")
```

### chicken

```
$> ./bin/chicken ./cmd/chicken.go
🐔 🐔

🐔 🐔
🐔
🐔
🐔
🐔
🐔
🐔

🐔 🐔 🐔

🐔 🐔 🐔 🐔 🐔 🐔

🐔

🐔 🐔 🐔

🐔 🐔 🐔
🐔 🐔 🐔
🐔 🐔 🐔
🐔 🐔 🐔
🐔

🐔 🐔 🐔 🐔 🐔 🐔 🐔

🐔 🐔 🐔

🐔 🐔 🐔 🐔 🐔
🐔 🐔 🐔
🐔 🐔 🐔

🐔 🐔 🐔 🐔

🐔 🐔 🐔 🐔 🐔
🐔
🐔

🐔 🐔 🐔
🐔

🐔 🐔

🐔 🐔 🐔
🐔 🐔

🐔 🐔

🐔 🐔 🐔 🐔

🐔 🐔 🐔
🐔 🐔 🐔

🐔 🐔 🐔 🐔 🐔 🐔 🐔
🐔 🐔 🐔 🐔
🐔

🐔 🐔 🐔
🐔
🐔
🐔
```

### languages

```
$> ./bin/chicken -language eng ./cmd/chicken.go
chicken chicken

chicken chicken
chicken
chicken
chicken
chicken
chicken
chicken

chicken chicken chicken

chicken chicken chicken chicken chicken chicken

chicken

chicken chicken chicken

chicken chicken chicken
chicken chicken chicken
chicken chicken chicken
chicken chicken chicken
chicken

chicken chicken chicken chicken chicken chicken chicken

chicken chicken chicken

chicken chicken chicken chicken chicken
chicken chicken chicken
chicken chicken chicken

chicken chicken chicken chicken

chicken chicken chicken chicken chicken
chicken
chicken

chicken chicken chicken
chicken

chicken chicken

chicken chicken chicken
chicken chicken

chicken chicken

chicken chicken chicken chicken

chicken chicken chicken
chicken chicken chicken

chicken chicken chicken chicken chicken chicken chicken
chicken chicken chicken chicken
chicken

chicken chicken chicken
chicken
chicken
chicken
```

### reading from STDIN

You can tell `chicken` to read from STDIN by passing `-` as its input argument.

```
$> yes | ./bin/chicken -language hbo -
עוף
עוף
עוף
עוף
עוף
עוף
עוף
```

## Translations

* The current [list of translations](strings/strings.go) is very short and [your help is welcome](https://github.com/thisisaaronland/go-chicken/pulls)! The goal is to have a 🐔 for every language listed in [ISO 639-3](https://en.wikipedia.org/wiki/List_of_ISO_639-3_codes).
 
## See also

* http://www.fileformat.info/info/unicode/char/1f414/index.htm
