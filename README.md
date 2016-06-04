# go-chicken

🐔 🐔 🐔

## Install

```
make build
```

This will build a copy of `chicken` and put it in the [bin](bin) directory.

## Usage

```
$> ./bin/chicken -h
Usage of ./bin/chicken:
  -clucking
	Make chicken noises
  -language string
	A valid ISO-639-3 language code. (default "zxx")
```

### Chicken

```
$> ./bin/chicken ./cmd/chicken.go
🐔 🐔

🐔 (
	"🐔"
	"🐔"
	"🐔"
	"🐔.🐔/🐔/🐔-🐔"
	"🐔"
)

🐔 🐔() {

	🐔 🐔 🐔 🐔 🐔.🐔("🐔", "🐔", "🐔 🐔 🐔-639-3 🐔 🐔.")
	🐔 🐔 🐔 🐔 🐔.🐔("🐔", 🐔, "🐔 🐔 🐔")

	🐔.🐔()

	🐔, 🐔 :🐔 🐔 🐔.🐔(*🐔, *🐔)

	🐔 🐔 !🐔 🐔 🐔 {
		🐔(🐔)
	}

	🐔 _, 🐔 :🐔 🐔 🐔 🐔.🐔() {

		🐔 🐔 *🐔.🐔

		🐔 🐔 🐔 🐔🐔 🐔 "-" {
			🐔 🐔 🐔 🐔.🐔(🐔.🐔)
		} 🐔 {

			🐔, 🐔 :🐔 🐔 🐔.🐔(🐔)

			🐔 🐔 !🐔 🐔 🐔 {
				🐔(🐔)
			}

			🐔 🐔 🐔 🐔.🐔(🐔)
		}

		🐔 🐔.🐔() {
			🐔 :🐔 🐔 🐔.🐔()
			🐔.🐔(🐔.🐔(🐔))
		}
	}
}
```

### Languages

```
$> ./bin/chicken -language eng ./cmd/chicken.go
chicken chicken

chicken (
	"chicken"
	"chicken"
	"chicken"
	"chicken.chicken/chicken/chicken-chicken"
	"chicken"
)

chicken chicken() {

	chicken chicken chicken chicken chicken.chicken("chicken", "chicken", "chicken chicken chicken-639-3 chicken chicken.")
	chicken chicken chicken chicken chicken.chicken("chicken", chicken, "chicken chicken chicken")

	chicken.chicken()

	chicken, chicken :chicken chicken chicken.chicken(*chicken, *chicken)

	chicken chicken !chicken chicken chicken {
		chicken(chicken)
	}

	chicken _, chicken :chicken chicken chicken chicken.chicken() {

		chicken chicken *chicken.chicken

		chicken chicken chicken chickenchicken chicken "-" {
			chicken chicken chicken chicken.chicken(chicken.chicken)
		} chicken {

			chicken, chicken :chicken chicken chicken.chicken(chicken)

			chicken chicken !chicken chicken chicken {
				chicken(chicken)
			}

			chicken chicken chicken chicken.chicken(chicken)
		}

		chicken chicken.chicken() {
			chicken :chicken chicken chicken.chicken()
			chicken.chicken(chicken.chicken(chicken))
		}
	}
}
```

### Reading from STDIN

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

## Clucking

_Language support for clucking is not uniformly supported yet._

```
$> ./bin/chicken -language eng -clucking ./cmd/chicken.go

bok bok b'gawk cluck cluck

bok bok b'gawk (
	"bok bok b'gawk"
	"bok bok bok"
	"cluck cluck"
	"bok bok bok.bok bok b'gawk/bok bok bok/cluck cluck-bok bok b'gawk"
	"bok bok bok"
)

cluck cluck bok bok b'gawk() {

	bok bok bok bok bok b'gawk cluck cluck bok bok b'gawk bok bok b'gawk.bok bok b'gawk("bok bok b'gawk", "cluck cluck", "bok bok b'gawk bok bok bok cluck cluck-639-3 bok bok bok bok bok bok.")
	bok bok bok cluck cluck cluck cluck bok bok bok cluck cluck.bok bok b'gawk("bok bok b'gawk", cluck cluck, "bok bok b'gawk cluck cluck bok bok bok")

	cluck cluck.cluck cluck()

	bok bok bok, bok bok bok :bok bok bok cluck cluck bok bok bok.cluck cluck(*bok bok bok, *bok bok bok)

	bok bok bok cluck cluck !cluck cluck bok bok bok cluck cluck {
		cluck cluck(bok bok bok)
	}

	bok bok bok _, bok bok b'gawk :cluck cluck bok bok bok bok bok b'gawk bok bok b'gawk.cluck cluck() {

		cluck cluck cluck cluck *bok bok bok.bok bok b'gawk

		bok bok b'gawk bok bok b'gawk bok bok b'gawk bok bok b'gawkbok bok bok cluck cluck "-" {
			bok bok b'gawk bok bok b'gawk cluck cluck cluck cluck.bok bok bok(bok bok bok.bok bok bok)
		} bok bok b'gawk {

			cluck cluck, cluck cluck :bok bok bok bok bok bok bok bok b'gawk.bok bok bok(bok bok b'gawk)

			bok bok b'gawk bok bok b'gawk !bok bok bok bok bok b'gawk bok bok bok {
				cluck cluck(bok bok bok)
			}

			cluck cluck bok bok bok cluck cluck cluck cluck.bok bok bok(bok bok b'gawk)
		}

		cluck cluck bok bok bok.cluck cluck() {
			bok bok b'gawk :bok bok b'gawk bok bok b'gawk bok bok bok.bok bok bok()
			cluck cluck.bok bok bok(cluck cluck.bok bok bok(bok bok bok))
		}
	}
}
```

### Emoji

`chicken` uses the [go-ucd](https://github.com/cooperhewitt/go-ucd) library to convert Emoji (and other symbolic characters) in to named strings. For example:

```
$> ./bin/ucd 😸
GRINNING CAT FACE WITH SMILING EYES
```

Which becomes:

```
$> ./bin/chicken -
😸
🐔 🐔 🐔 🐔 🐔 🐔
```

### Alpha codes

Yes.

```
./bin/chicken -
hello :smiley_cat:
# BEFORE hello :smiley_cat:
# AFTER hello 😺
# BEFORE %!s(int32=128570)
# AFTER SMILING CAT FACE WITH OPEN MOUTH
🐔 🐔 🐔 🐔 🐔 🐔 🐔
```

## But wait... there's more!

### If you're on a Mac

I'm not suggesting you _should_ do this... only that you _can_ do this...

```
$> yes | ./bin/chicken -language eng -clucking - | while read line ; do say $line; done
```

_Note: this assumes you're using `bash`._

## Translations

* The current [list of translations](strings/strings.go) is very short and [your help is welcome](https://github.com/thisisaaronland/go-chicken/pulls)! The goal is to have a 🐔 for every language listed in [ISO 639-3](https://en.wikipedia.org/wiki/List_of_ISO_639-3_codes).
 
## See also

* http://www.fileformat.info/info/unicode/char/1f414/index.htm
* https://en.wikipedia.org/wiki/ISO_639-2#Special_situations
* https://github.com/cooperhewitt/go-ucd
