#!/bin/bash

for filename in ./*; do
	gpx_filename="${filename/fit/gpx}"
	if [[ ${filename} == *fit ]]
	then
		echo "${filename} ${gpx_filename}"
		gpsbabel -i garmin_fit -f ${filename} -o gpx -F ${gpx_filename}
	fi
done
