# Makefile contains targets for making all landgrab executables.
#
# All executables will have the form run_landgrab_<executable> and will be
# placed in $GOPATH/bin.

# all target makes all targets in the Makefile.
all: run_cli run_web

# run_cli target makes the cli app.
run_cli:
	$(call make,$@)

# run_web target makes the client web app and API which is served from it.
run_web:
	$(call make,$@)

# make a go target in the app directory with the passed name.
#
# A go main package must exist with the passed name as a subpackage of app.
#
# An example call is:
#   $(call make,<name>)
define make
	@echo Making $(1):
	@echo -----------------------------------------
	cd app/$(1) && go build -o landgrab_$(1); \
	mv landgrab_$(1) $$GOPATH/bin
	@echo
endef
