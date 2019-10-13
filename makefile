test:
	cd .. && cookiecutter ./cookiecutter-cloud-run-go --no-input
	cd ../sample && make setup && go build
	cd .. && rm -r sample

	cd .. && cookiecutter ./cookiecutter-cloud-run-go --no-input use_secret=yes
	cd ../sample && make setup && go build
	cd .. && rm -r sample

	cd .. && cookiecutter ./cookiecutter-cloud-run-go --no-input use_gcs=yes
	cd ../sample && make setup && go build
	cd .. && rm -r sample

	cd .. && cookiecutter ./cookiecutter-cloud-run-go --no-input golang_mod_name=github.com/lol/lol type=msg
	cd ../lol && make setup && go build
	cd .. && rm -r lol

	cd .. && cookiecutter ./cookiecutter-cloud-run-go --no-input golang_mod_name=github.com/lol/lol type=msg
	cd ../lol && make setup && go build
	cd .. && rm -r lol

	cd .. && cookiecutter ./cookiecutter-cloud-run-go --no-input golang_mod_name=github.com/lol/lol type=msg use_gcs=yes
	cd ../lol && make setup && go build
	cd .. && rm -r lol