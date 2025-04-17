package winicon

import (
	"encoding/binary"
	"io"

	"github.com/cyber-xxm/energy/v2/pkgs/winicon/internal/winicon"
)

// Icon stores the data for a single icon
type Icon struct {
	Width        uint16
	Height       uint16
	Colours      uint8
	Planes       uint16
	BitsPerPixel uint16
	Data         []byte `json:"-"`
	Format       string
	Offset       uint32
	size         uint32
}

// GetFileData reads in the given .ico filename and returns information
// about the icon/icons
func GetFileData(r io.Reader) ([]*Icon, error) {

	var result []*Icon

	// Parse the .ico file
	var header winicon.IconFileHeader
	err := binary.Read(r, binary.LittleEndian, &header)
	if err != nil {
		return nil, err
	}

	// Loop over icons
	for index := 0; index < (int)(header.ImageCount); index++ {
		// Read in icon headers
		var iconHeader winicon.IconHeader
		err = binary.Read(r, binary.LittleEndian, &iconHeader)
		if err != nil {
			return nil, err
		}
		icon := Icon{
			Width:        (uint16)(iconHeader.Width),
			Height:       (uint16)(iconHeader.Height),
			BitsPerPixel: iconHeader.BitsPerPixel,
			Planes:       iconHeader.Planes,
			Offset:       iconHeader.Offset,
			size:         iconHeader.Size,
		}

		// Width/Height of 256 is encoded as 0 in the icon header
		if icon.Width == 0 {
			icon.Width = 256
		}
		if icon.Height == 0 {
			icon.Height = 256
		}

		result = append(result, &icon)
	}

	// Loop over Icons to read in image data
	for _, icon := range result {
		icon.Data = make([]byte, icon.size)
		_, err := r.Read(icon.Data)
		if err != nil {
			return nil, err
		}
		if string(icon.Data[1:4]) == "PNG" {
			icon.Format = "PNG"
		} else {
			icon.Format = "BMP"
		}
	}

	return result, nil
}
