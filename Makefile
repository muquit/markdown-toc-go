#====================================================================
# Reqquires https://github.com/muquit/go-xbuild-go for cross compiling
# for other platforms.
# Mar-29-2025 muquit@muquit.com 
#====================================================================
README_ORIG=./docs/README.md
MAIN_MD=./docs/main.md
README=./README.md
BINARY=./markdown-toc-go
TEST_MD=./test/Test.md
# v 1.0.3
GLOSSARY_FILE=./docs/glossary.txt
TEST_GLOSSARY_FILE=./test/TEST_GLOSSARY.md
TEST_GLOSSARY_EXPANDED_FILE=./TEST_GLOSSARY_EXPANDED.md


all: build build_all doc

build:
	@echo "*** Compiling markdon-toc-go ...."
	@/bin/rm -f bin/*
	go build -o $(BINARY)

build_all:
	@echo "*** Cross Compiling markdon-toc-go ...."
	@/bin/rm -f bin/*
	go-xbuild-go

release:
	go-xbuild-go -release

doc:
	echo "*** Generating README.md with TOC ..."
	chmod 600 $(README)
	$(BINARY) -i $(MAIN_MD) -o $(README) --glossary ${GLOSSARY_FILE} -f
	chmod 444 $(README)
	$(BINARY) -i ${TEST_MD} -o ./Test.md -f
	$(BINARY) -i ${TEST_GLOSSARY_FILE} -o ${TEST_GLOSSARY_EXPANDED_FILE} --glossary ${GLOSSARY_FILE} -f

clean:
	/bin/rm -f $(BINARY)
	/bin/rm -rf ./bin
