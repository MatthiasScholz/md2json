source_folder := ./../../md
output_folder := ./../../json
convert: cleanup shell
	earthly +deploy

shell:
	poetry shell

dependencies:
	poetry add markdown-to-json
