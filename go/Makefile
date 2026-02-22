WORDLIST = wordlist.txt
OUT = diceware

.PHONY: run
run: $(WORDLIST) $(OUT)
	go run .

$(OUT): $(WORDLIST)
	go build -o $@ .

$(WORDLIST):
	curl -L 'https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt' -o $(WORDLIST)
