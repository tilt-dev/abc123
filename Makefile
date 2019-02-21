.PHONY: build

build:
	docker build ./fe -t abc123/fe
	docker build ./letters -t abc123/letters
	docker build ./numbers -t abc123/numbers
