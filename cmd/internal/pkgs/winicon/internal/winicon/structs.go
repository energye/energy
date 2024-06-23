package winicon

// IconFileHeader is the icon file's 48bit header
type IconFileHeader struct {
	_          uint16
	ImageType  uint16
	ImageCount uint16
}

// IconHeader is the header for the icon data
type IconHeader struct {
	Width        uint8
	Height       uint8
	Colours      uint8
	_            uint8
	Planes       uint16
	BitsPerPixel uint16
	Size         uint32
	Offset       uint32
}
