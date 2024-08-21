//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package mime type
package mime

import "path/filepath"

var mimes = map[string]string{
	".css":     "text/css; charset=utf-8",
	".htm":     "text/html; charset=utf-8",
	".html":    "text/html; charset=utf-8",
	".mjs":     "text/javascript; charset=utf-8",
	".js":      "text/javascript; charset=utf-8",
	".xml":     "text/xml; charset=utf-8",
	".txt":     "text/plain",
	".csv":     "text/csv",
	".tsv":     "text/tab-separated-values",
	".php":     "text/x-php",
	".rtf":     "text/rtf",
	".vtt":     "text/vtt",
	".lua":     "text/x-lua",
	".pl":      "text/x-perl",
	".tcl":     "text/x-tcl",
	".vcf":     "text/vcard",
	".ics":     "text/calendar",
	".jpeg":    "image/jpeg",
	".jpg":     "image/jpeg",
	".gif":     "image/gif",
	".avif":    "image/avif",
	".png":     "image/png",
	".svg":     "image/svg+xml",
	".webp":    "image/webp",
	".jxl":     "image/jxl",
	".jp2":     "image/jp2",
	".jpf":     "image/jpx",
	".jpm":     "image/jpm",
	".xpm":     "image/x-xpixmap",
	".bpg":     "image/bpg",
	".tiff":    "image/tiff",
	".bmp":     "image/bmp",
	".ico":     "image/x-icon",
	".icns":    "image/x-icns",
	".psd":     "image/vnd.adobe.photoshop",
	".heic":    "image/heic",
	".heif":    "image/heif",
	".hdr":     "image/vnd.radiance",
	".dwg":     "image/vnd.dwg",
	".xcf":     "image/x-xcf",
	".pat":     "image/x-gimp-pat",
	".gbr":     "image/x-gimp-gbr",
	".djvu":    "image/vnd.djvu",
	".wasm":    "application/wasm",
	".pdf":     "application/pdf",
	".json":    "application/json",
	".xz":      "application/x-xz",
	".gz":      "application/gzip",
	".7z":      "application/x-7z-compressed",
	".zip":     "application/zip",
	".tar":     "application/x-tar",
	".xar":     "application/x-xar",
	".bz2":     "application/x-bzip2",
	".fdf":     "application/vnd.fdf",
	".xlsx":    "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	".docx":    "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	".pptx":    "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	".epub":    "application/epub+zip",
	".jar":     "application/jar",
	".msi":     "application/x-ms-installer",
	".aaf":     "application/octet-stream",
	".doc":     "application/msword",
	".ppt":     "application/vnd.ms-powerpoint",
	".pub":     "application/vnd.ms-publisher",
	".xls":     "application/vnd.ms-excel",
	".msg":     "application/vnd.ms-outlook",
	".ps":      "application/postscript",
	".fits":    "application/fits",
	".ogg":     "application/ogg",
	".har":     "application/json",
	".geojson": "application/geo+json",
	".ndjson":  "application/x-ndjson",
	".srt":     "application/x-subrip",
	".py":      "application/x-python",
	".rss":     "application/rss+xml",
	".owl":     "application/owl+xml",
	".atom":    "application/atom+xml",
	".kml":     "application/vnd.google-earth.kml+xml",
	".xlf":     "application/x-xliff+xml",
	".gml":     "application/gml+xml",
	".gpx":     "application/gpx+xml",
	".tcx":     "application/vnd.garmin.tcx+xml",
	".amf":     "application/x-amf",
	".3mf":     "application/vnd.ms-package.3dmanufacturing-3dmodel+xml",
	".rmvb":    "application/vnd.rn-realmedia-vbr",
	".class":   "application/x-java-applet",
	".swf":     "application/x-shockwave-flash",
	".crx":     "application/x-chrome-extension",
	".eot":     "application/vnd.ms-fontobject",
	".shp":     "application/octet-stream",
	".shx":     "application/octet-stream",
	".dbf":     "application/x-dbf",
	".exe":     "application/vnd.microsoft.portable-executable",
	".so":      "application/x-sharedlib",
	".a":       "application/x-archive",
	".deb":     "application/vnd.debian.binary-package",
	".rpm":     "application/x-rpm",
	".dcm":     "application/dicom",
	".odt":     "application/vnd.oasis.opendocument.text",
	".ott":     "application/vnd.oasis.opendocument.text-template",
	".ods":     "application/vnd.oasis.opendocument.spreadsheet",
	".ots":     "application/vnd.oasis.opendocument.spreadsheet-template",
	".odp":     "application/vnd.oasis.opendocument.presentation",
	".otp":     "application/vnd.oasis.opendocument.presentation-template",
	".odg":     "application/vnd.oasis.opendocument.graphics",
	".otg":     "application/vnd.oasis.opendocument.graphics-template",
	".odf":     "application/vnd.oasis.opendocument.formula",
	".odc":     "application/vnd.oasis.opendocument.chart",
	".sxc":     "application/vnd.sun.xml.calc",
	".rar":     "application/x-rar-compressed",
	".mobi":    "application/x-mobipocket-ebook",
	".lit":     "application/x-ms-reader",
	".sqlite":  "application/vnd.sqlite3",
	".warc":    "application/warc",
	".nes":     "application/vnd.nintendo.snes.rom",
	".lnk":     "application/x-ms-shortcut",
	".macho":   "application/x-mach-binary",
	".mrc":     "application/marc",
	".mdb":     "application/x-msaccess",
	".accdb":   "application/x-msaccess",
	".zst":     "application/zstd",
	".cab":     "application/vnd.ms-cab-compressed",
	".lz":      "application/lzip",
	".torrent": "application/x-bittorrent",
	".cpio":    "application/x-cpio",
	".p7s":     "application/pkcs7-signature",
	".xfdf":    "application/vnd.adobe.xfdf",
	".m3u":     "application/vnd.apple.mpegurl",
	".oga":     "audio/ogg",
	".mp3":     "audio/mpeg",
	".flac":    "audio/flac",
	".midi":    "audio/midi",
	".ape":     "audio/ape",
	".mpc":     "audio/musepack",
	".wav":     "audio/wav",
	".aiff":    "audio/aiff",
	".au":      "audio/basic",
	".amr":     "audio/amr",
	".aac":     "audio/aac",
	".voc":     "audio/x-unknown",
	".m4a":     "audio/x-m4a",
	".qcp":     "audio/qcelp",
	".m4v":     "video/x-m4v",
	".mp4":     "video/mp4",
	".webm":    "video/webm",
	".mpeg":    "video/mpeg",
	".mov":     "video/quicktime",
	".mqv":     "video/quicktime",
	".3gp":     "video/3gpp",
	".3g2":     "video/3gpp2",
	".avi":     "video/x-msvideo",
	".flv":     "video/x-flv",
	".mkv":     "video/x-matroska",
	".asf":     "video/x-ms-asf",
	".ogv":     "video/ogg",
	".ttf":     "font/ttf",
	".woff":    "font/woff",
	".woff2":   "font/woff2",
	".otf":     "font/otf",
	".ttc":     "font/collection",
	".glb":     "model/gltf-binary",
	".x3d":     "model/x3d+xml",
	".dae":     "model/vnd.collada+xml",
	//"application/x-ole-storage", ""
	//"application/x-elf", ""
	//"application/x-object", ""
	//"application/x-executable", ""
	//"application/x-coredump", ""
	//"application/tzif", ""
}

func GetMimeType(filename string) string {
	return mimes[filepath.Ext(filename)]
}