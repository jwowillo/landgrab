# Makefile contains targets for making all landgrab executables.
#
# All executables will have the form run_landgrab_<executable> and will be
# placed in $GOPATH/bin.

# all target makes all targets in the Makefile.
all: run_cli run_web run_arena

# run_cli target makes the cli app.
run_cli:
	$(call log,$@)
	$(call make,$@)

# run_web target makes the client web app and API which is served from it.
run_web: pub
	$(call log,$@)
	$(call make,$@)

#run_arena target makes the run-arena app.
run_arena:
	$(call log,$@)
	$(call make,$@)

# pub installs pub dependencies if necessary.
pub:
	$(call log,$@)
	pub get
	pub build --mode=debug
	@echo

# clean built files.
clean:
	$(call log, $@)
	rm -rf pubspeck.lock
	rm -rf build
	rm -rf .pub
	rm -rf .packages

# make a go target in the app directory with the passed name.
#
# A go main package must exist with the passed name as a subpackage of app.
#
# An example call is:
#   $(call make,<name>)
define make
	cd app/$(1) && go build; \
	mv $(1) $$GOPATH/bin/landgrab_$(1)
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
