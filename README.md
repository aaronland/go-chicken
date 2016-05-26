# go-chicken

## Caveats

It's early days, still. There is lots left to do...

## Install

```
make build
```

## Usage

### chicken

```
$> chicken ./cmd/chicken.go
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

### language tags

```
$> chicken -language eng ./cmd/chicken.go
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

```
$> yes | ./chicken -language hbo -
עוף
עוף
עוף
עוף
עוף
עוף
עוף
```

## Translations

* The current [list of translations](strings/strings.go) is very short and [your help is welcome](https://github.com/thisisaaronland/go-chicken/pulls)! The goal is to have a 🐔 for every language listed in [ISO 639-3](https://en.wikipedia.org/wiki/List_of_ISO_639-3_codes)
 
