source_folder := ./../../nmd
output_folder := ./../../json
config_folder := ./config
convert: cleanup
	mkdir -p $(output_folder)
  # NOT WORKING
	hugo convert toJSON --output $(output_folder) --source $(source_folder) --configDir $(config_folder)

cleanup:
	rm -rf $(output_folder)
