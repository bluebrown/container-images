CLI=docker
REGISTRY=docker.io
NAMESPACE=bluebrown

prefix=$(REGISTRY)/$(NAMESPACE)

build.%:
	$(CLI) build \
		-t $(prefix)/$*:latest \
		-t $(prefix)/$*:$(shell grep version $*/meta.txt | cut -d'=' -f2)  \
		$*/

push.%:
	$(CLI) push $(prefix)/$*:latest
	$(CLI) push $(prefix)/$*:$(shell grep version $*/meta.txt | cut -d'=' -f2)
