CDS_VERSION=$(shell curl -s https://consumerdatastandardsaustralia.github.io/standards/includes/swagger/cds_energy.json | jq -r '.info.version')
ARCHIVE_VERSION=$(ENERGY_VERSION)

.PHONY: requirements clean cdr update-latest update

requirements:
	@if ! command -v curl &>/dev/null; then \
		echo "curl is not installed"; \
		exit 1; \
	fi

clean:
	rm -f ./cdr

cdr: clean
	go build -o cdr ./cmd

# Updates the root industry package to the latest CDS version
update-latest: requirements
	mkdir -p energy/$(CDS_VERSION)
	mkdir -p banking/$(CDS_VERSION)
	mkdir -p common/$(CDS_VERSION)

	curl -o energy/$(CDS_VERSION)/cdr_energy.swagger.$(CDS_VERSION).json https://consumerdatastandardsaustralia.github.io/standards/includes/swagger/cds_energy.json
	curl -o common/$(CDS_VERSION)/cdr_common.swagger.$(CDS_VERSION).json https://consumerdatastandardsaustralia.github.io/standards/includes/swagger/cds_common.json
	curl -o banking/$(CDS_VERSION)/cdr_banking.swagger.$(CDS_VERSION).json https://consumerdatastandardsaustralia.github.io/standards/includes/swagger/cds_banking.json

	cd energy/$(CDS_VERSION) && go tool oapi-codegen -config ../config.yaml cdr_energy.swagger.$(CDS_VERSION).json
	cd banking/$(CDS_VERSION) && go tool oapi-codegen -config ../config.yaml cdr_banking.swagger.$(CDS_VERSION).json
	cd common/$(CDS_VERSION) && go tool oapi-codegen cdr_common.swagger.$(CDS_VERSION).json > common.gen.go

# Generates models for a specific version of the CDS for each industry in a sub-package
update: requirements
	mkdir -p energy/$(ARCHIVE_VERSION)
	mkdir -p banking/$(ARCHIVE_VERSION)
	mkdir -p common/$(ARCHIVE_VERSION)

	curl -o energy/$(ARCHIVE_VERSION)/cdr_energy.swagger.$(ARCHIVE_VERSION).json https://consumerdatastandardsaustralia.github.io/standards-archives/standards-$(ARCHIVE_VERSION)/includes/swagger/cds_energy.json
	curl -o banking/$(ARCHIVE_VERSION)/cdr_banking.swagger.$(ARCHIVE_VERSION).json https://consumerdatastandardsaustralia.github.io/standards-archives/standards-$(ARCHIVE_VERSION)/includes/swagger/cds_banking.json
	curl -o common/$(ARCHIVE_VERSION)/cdr_common.swagger.$(ARCHIVE_VERSION).json https://consumerdatastandardsaustralia.github.io/standards-archives/standards-$(ARCHIVE_VERSION)/includes/swagger/cds_common.json

	cd energy/$(ARCHIVE_VERSION) && go tool oapi-codegen -config ../config.yaml cdr_energy.swagger.$(ARCHIVE_VERSION).json
	cd banking/$(ARCHIVE_VERSION) && go tool oapi-codegen -config ../config.yaml cdr_banking.swagger.$(ARCHIVE_VERSION).json
	cd common/$(ARCHIVE_VERSION) && go tool oapi-codegen cdr_common.swagger.$(ARCHIVE_VERSION).json > common.gen.go
