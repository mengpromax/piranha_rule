all: clean
	chmod +x build.sh
	bash build.sh

.PHONY: clean
clean:
	rm -rf target