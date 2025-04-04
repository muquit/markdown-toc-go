#====================================================================
# Reqquires https://github.com/muquit/go-xbuild-go for cross compiling
# for other platforms.
# Mar-29-2025 muquit@muquit.com 
#====================================================================
README_ORIG=./docs/README.md
README=./README.md
BINARY=./markdown-toc-go

all: build build_all doc

build:
	echo "*** Compiling markdon-to-co ...."
	go build -o $(BINARY)

build_all:
	echo "*** Cross Compiling markdon-to-co ...."
	go-xbuild-go

doc:
	echo "*** Generating README.md with TOC ..."
	chmod 600 $(README)
	$(BINARY) -i $(README_ORIG) -o $(README) -f
	chmod 444 $(README)
	$(BINARY) -i ./test/Test.md -o ./Test.md -f

clean:
	/bin/rm -f $(BINARY)
	/bin/rm -rf ./bin
