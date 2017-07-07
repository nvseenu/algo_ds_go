GOCMD=go
GOBUILD=$(GOCMD) build -v
GOINSTALL=$(GOCMD) install -v
GOTEST=$(GOCMD) test -v
GOFMT=$(GOCMD) fmt
GOCLEAN=$(GOCMD) clean

TOP_LEVEL_PKG := algo_ds
PKG_LIST := util sort array collections
ALL_LIST := $(PKG_LIST)

CLEAN_LIST := $(foreach pkg, $(ALL_LIST), $(pkg)_clean) 
TEST_LIST := $(foreach pkg, $(PKG_LIST), $(pkg)_test) 
FORMAT_LIST := $(foreach pkg, $(ALL_LIST), $(pkg)_format) 

.PHONY: $(CLEAN_LIST) $(FORMAT_LIST) $(TEST_LIST)

all: clean format test

clean: $(CLEAN_LIST)
format: $(FORMAT_LIST)
test: $(TEST_LIST)	

$(CLEAN_LIST): %_clean:
	$(GOCLEAN) $(TOP_LEVEL_PKG)/$*

$(TEST_LIST): %_test: 
	$(GOTEST) $(TOP_LEVEL_PKG)/$*

$(FORMAT_LIST): %_format:
	$(GOFMT) $(TOP_LEVEL_PKG)/$*