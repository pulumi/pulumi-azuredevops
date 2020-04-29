PROJECT_NAME := azuredevops Package
include build/common.mk

PACK             := azuredevops
PACKDIR          := sdk
PROJECT          := github.com/pulumi/pulumi-${PACK}
NODE_MODULE_NAME := @pulumi/${PACK}
TF_NAME          := ${PACK}
TF_MOD			 := "vendor"

TFGEN           := pulumi-tfgen-${PACK}
PROVIDER        := pulumi-resource-${PACK}
VERSION         := $(shell scripts/get-version)
PYPI_VERSION    := $(shell scripts/get-py-version)

DOTNET_PREFIX  := $(firstword $(subst -, ,${VERSION:v%=%})) # e.g. 1.5.0
DOTNET_SUFFIX  := $(word 2,$(subst -, ,${VERSION:v%=%}))    # e.g. alpha.1

ifeq ($(strip ${DOTNET_SUFFIX}),)
	DOTNET_VERSION := $(strip ${DOTNET_PREFIX})
else
	DOTNET_VERSION := $(strip ${DOTNET_PREFIX})-$(strip ${DOTNET_SUFFIX})
endif

TESTPARALLELISM := 4

# NOTE: Since the plugin is published using the nodejs style semver version
# We set the PLUGIN_VERSION to be the same as the version we use when building
# the provider (e.g. x.y.z-dev-... instead of x.y.zdev...)
build:: build_dependencies build_languages

build_dependencies:: patch tfgen provider
build_languages:: build_go build_nodejs build_python build_dotnet

build_clean::
	$(call STEP_MESSAGE)
	cd sdk && for LANGUAGE in "nodejs" "python" "go" "dotnet" ; do \
		rm -fr $$LANGUAGE ; \
	done

build_go:: build_dependencies
	$(call STEP_MESSAGE)
	(cd sdk && rm -fr go)
	(cd provider && $(TFGEN) go --overlays overlays/go/ --out ../${PACKDIR}/go/ || exit 3) ; \

build_nodejs:: build_dependencies
	$(call STEP_MESSAGE)
	(cd sdk && rm -fr nodejs)
	(cd provider && $(TFGEN) nodejs --overlays overlays/nodejs/ --out ../${PACKDIR}/nodejs/ || exit 3) ; \
	(cd ${PACKDIR}/nodejs/ && \
		yarn install && \
		yarn run tsc && \
		cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/ && \
		sed -i.bak "s/\$${VERSION}/$(VERSION)/g" ./bin/package.json)

build_python:: build_dependencies
	$(call STEP_MESSAGE)
	(cd sdk && rm -fr python)
	(cd provider && $(TFGEN) python --overlays overlays/python/ --out ../${PACKDIR}/python/ || exit 3) ; \
	(cd ${PACKDIR}/python/ && \
		cp ../../README.md . && \
		$(PYTHON) setup.py clean --all 2>/dev/null && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && cp -R ../python.bin ./bin && rm -rf ../python.bin && \
		sed -i.bak -e "s/\$${VERSION}/$(PYPI_VERSION)/g" -e "s/\$${PLUGIN_VERSION}/$(VERSION)/g" ./bin/setup.py && \
		rm ./bin/setup.py.bak && \
		cd ./bin && $(PYTHON) setup.py build sdist)

build_dotnet:: build_dependencies
	$(call STEP_MESSAGE)
	(cd sdk && rm -fr dotnet)
	(cd provider && $(TFGEN) dotnet --overlays overlays/dotnet/ --out ../${PACKDIR}/dotnet/ || exit 3) ; \
	(cd ${PACKDIR}/dotnet/ && \
  		echo "${VERSION:v%=%}" >version.txt && \
  		dotnet build /p:Version=${DOTNET_VERSION})

patch::
	$(call STEP_MESSAGE)
	scripts/patch-provider.sh

tfgen::
	$(call STEP_MESSAGE)
	cd provider && go install -mod=${TF_MOD} -ldflags "-X github.com/pulumi/pulumi-${PACK}/provider/pkg/version.Version=${VERSION}" ${PROJECT}/provider/cmd/${TFGEN}

generate_schema:: tfgen
	$(call STEP_MESSAGE)
	$(TFGEN) schema --out ./provider/cmd/${PROVIDER}

provider:: generate_schema
	$(call STEP_MESSAGE)
	cd provider && go generate cmd/${PROVIDER}/main.go
	cd provider && go install -mod=${TF_MOD} -ldflags "-X github.com/pulumi/pulumi-${PACK}/provider/pkg/version.Version=${VERSION}" ${PROJECT}/provider/cmd/${PROVIDER}

lint:: patch
	$(call STEP_MESSAGE)
	cd provider && golangci-lint --modules-download-mode=${TF_MOD} run

install:: install_nodejs install_python install_dotnet

install_nodejs:: build_nodejs
	$(call STEP_MESSAGE)
	[ ! -e "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)" ] || rm -rf "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)"
	mkdir -p "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)"
	cp -r ${PACKDIR}/nodejs/bin/. "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)"
	rm -rf "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)/node_modules"
	cd "$(PULUMI_NODE_MODULES)/$(NODE_MODULE_NAME)" && \
		yarn install --offline --production && \
		(yarn unlink > /dev/null 2>&1 || true) && \
		yarn link

install_python:: build_python
	$(call STEP_MESSAGE)
	cd ${PACKDIR}/python/bin && $(PIP) install --user -e .

install_dotnet:: build_dotnet
	$(call STEP_MESSAGE)
	[ -d "$(PULUMI_NUGET)" ] || mkdir "$(PULUMI_NUGET)"
	[ ! -e "$(PULUMI_NUGET)" ] || find "$(PULUMI_NUGET)" -type f -iname "*${PACK}*" -exec rm {} \;
	find ${PACKDIR}/dotnet -name '*.nupkg' -exec cp -p {} ${PULUMI_NUGET} \;

test_fast::
	cd examples && $(GO_TEST_FAST) .

test_all::
	cd examples && $(GO_TEST) .

.PHONY: publish_tgz
publish_tgz:
	$(call STEP_MESSAGE)
	./scripts/publish_tgz.sh

.PHONY: publish_packages
publish_packages:
	$(call STEP_MESSAGE)
	rm -rf "$$(go env GOPATH)/src/github.com/pulumi/scripts" && git clone https://github.com/pulumi/scripts "$$(go env GOPATH)/src/github.com/pulumi/scripts"
	$$(go env GOPATH)/src/github.com/pulumi/scripts/ci/publish-tfgen-package .
	$$(go env GOPATH)/src/github.com/pulumi/scripts/ci/build-package-docs.sh ${PACK}

.PHONY: check_clean_worktree
check_clean_worktree:
	$$(go env GOPATH)/src/github.com/pulumi/scripts/ci/check-worktree-is-clean.sh

# The travis_* targets are entrypoints for CI.
.PHONY: travis_cron travis_push travis_pull_request travis_api
travis_cron: all
travis_push: only_build check_clean_worktree publish_tgz only_test publish_packages
travis_pull_request: all check_clean_worktree
travis_api: all
