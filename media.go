package main

import "regexp"

func IsImage(ext string) bool {
	found, err := regexp.MatchString("(bmp|jpg|png|tif|gif|pcx|tga|exif|fpx|svg|psd|cdr|pcd|dxf|ufo|eps|ai|raw|WMF|webp|jpeg)", ext)
	if err != nil {
		return false
	}
	return found
}
