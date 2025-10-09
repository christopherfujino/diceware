WORDLIST = wordlist.txt

.PHONY: run
run: $(WORDLIST)
	go run .

.PHONY: build
build:
	go build -o diceware .

$(WORDLIST):
	curl -L 'https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt' -o $(WORDLIST)
