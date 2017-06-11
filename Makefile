# Makefile contains targets for making all landgrab executables.
#
# All executables will have the form run_landgrab_<executable> and will be
# placed in $GOPATH/bin.
#
# all target makes all targets in the Makefile.

all: landgrab_run_cli landgrab_run_web landgrab_run_arena doc

# landgrab_run_cli target makes the cli app.
landgrab_run_cli:
	$(call log,$@)
	$(call make,$@)

# landgrab_run_web target makes the client web app and API which is served from
# it.
landgrab_run_web:
	$(call log,$@)
	$(call make,$@)

# landgrab_run_arena target makes the run-arena app.
landgrab_run_arena:
	$(call log,$@)
	$(call make,$@)

# clean built files.
clean:
	$(call log, $@)
	rm -rf pubspeck.lock
	rm -rf build
	rm -rf .pub
	rm -rf .packages

# make a go target in the cmd directory with the passed name.
#
# A go main package must exist with the passed name as a subpackage of cmd.
#
# An example call is:
#   $(call make,<name>)
define make
	cd cmd/$(1) && go install
	@echo
endef

# log the target with the passed name.
#
# An example call is:
#   $(call log,<name>)
define  log
	@echo Making $(1):
	@echo -----------------------------------------
endef
