WORDLIST = wordlist.txt

.PHONY: run
run: $(WORDLIST)
	go run .

$(WORDLIST):
	curl -L 'https://www.eff.org/files/2016/07/18/eff_large_wordlist.txt' -o $(WORDLIST)
