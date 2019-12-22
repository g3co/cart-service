# variables
GOCMD			=	go
GOPATH			:=	${shell pwd}
BINPATH			=	$(GOPATH)/bin

# parameters
GODEP			=	$(GOCMD) get
GOTEST			=	$(GOCMD) test -v 
GOBUILD			=	$(GOCMD) build
GOCONVEY        =   $(BINPATH)/goconvey
GOMOCKGEN       =   $(BINPATH)/mockgen

GOCONVEY_PORT   =   8180
GOCONVEY_TARGET =   "src/main/cart-service"

export GOPATH

# buildable packages
MAIN_PKGS 		:=	main/cart-service

# dependencies packages
DEPS_PKGS 		:=	gopkg.in/yaml.v2 \
					github.com/satori/go.uuid \
					github.com/smartystreets/goconvey \
					github.com/golang/mock/gomock \
					github.com/golang/mock/mockgen

# packages for testing
TEST_PKGS		:=	$(MAIN_PKGS)

# buildable lists
DEPS_LIST		=	$(foreach int, $(DEPS_PKGS), $(int)_deps)
TEST_LIST		=	$(foreach int, $(TEST_PKGS), $(int)_test)
BUILD_LIST		=	$(foreach int, $(MAIN_PKGS), $(int)_build)

all:			init deps build

deps:			$(DEPS_LIST)
test:			$(TEST_LIST)
build:			$(BUILD_LIST)

init:

$(DEPS_LIST): %_deps:
	$(GODEP) $*;

$(TEST_LIST): %_test:
	$(GOTEST) $*/...

$(BUILD_LIST): %_build:
	$(GOBUILD) -o $(BINPATH)/$(shell basename $*) $*
