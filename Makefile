GOCMD=go
GOBUILD=$(GOCMD) build -v
GOINSTALL=$(GOCMD) install -v
GOTEST=$(GOCMD) test -v
GOFMT=$(GOCMD) fmt
GOCLEAN=$(GOCMD) clean

TOP_LEVEL_PKG := algo_ds
PKG_LIST := util sort array collections

LIST := $(foreach pkg, $(PKG_LIST), $(TOP_LEVEL_PKG)/$(pkg)) 

all: clean format test

clean: 
	$(GOCLEAN) $(LIST)
	
format:
	$(GOFMT) $(LIST)
	
test:
	$(GOTEST) $(LIST)