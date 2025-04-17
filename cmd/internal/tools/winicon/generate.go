package winicon

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"image/png"
	"io"

	"github.com/cyber-xxm/energy/v2/cmd/internal/tools/winicon/internal/winicon"
	"golang.org/x/image/draw"
)

// GenerateIcon reads image data from the given reader and generates
// a .ico file that is written to the given writer. The .ico file will include
// a number of icons at the sizes given.
func GenerateIcon(r io.Reader, w io.Writer, sizes []int) error {
	header := &winicon.IconFileHeader{
		ImageType:  1,
		ImageCount: uint16(len(sizes)),
	}

	iconheaders := make([]winicon.IconHeader, len(sizes))

	var imageData bytes.Buffer

	// Decode to internal image
	imagedata, _, err := image.Decode(r)
	if err != nil {
		return err
	}

	// Loop over sizes desired
	for index, size := range sizes {

		// Check target size
		if size == 0 {
			return fmt.Errorf("a size of 0 is not valid")
		}

		// Scale image
		rect := image.Rect(0, 0, size, size)
		rawdata := image.NewRGBA(rect)
		scale := draw.CatmullRom
		scale.Scale(rawdata, rect, imagedata, imagedata.Bounds(), draw.Over, nil)

		// Convert back to PNG
		icondata := new(bytes.Buffer)
		writer := bufio.NewWriter(icondata)
		err = png.Encode(writer, rawdata)
		if err != nil {
			return err
		}
		err = writer.Flush()
		if err != nil {
			return err
		}

		// Save image data
		imageData.Write(icondata.Bytes())

		// Save header information
		if size >= 256 {
			size = 0
		}
		iconheaders[index].Width = (uint8)(size)
		iconheaders[index].Height = (uint8)(size)
		iconheaders[index].BitsPerPixel = 32
		iconheaders[index].Size = uint32(len(icondata.Bytes()))
	}

	// Update the offsets. Start by skipping header+icon headers
	var currentOffset uint32 = (uint32)(6 + (16 * len(iconheaders)))

	for index := range iconheaders {
		iconheaders[index].Offset = currentOffset
		currentOffset += iconheaders[index].Size
	}

	// Write out the header
	err = binary.Write(w, binary.LittleEndian, header)
	if err != nil {
		return err
	}

	// Write out the icon headers
	for _, iconheader := range iconheaders {
		err = binary.Write(w, binary.LittleEndian, iconheader)
		if err != nil {
			return err
		}
	}

	// Write out the image data
	_, err = w.Write(imageData.Bytes())
	if err != nil {
		return err
	}
	return nil
}
