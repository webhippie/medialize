package photo

import (
	"fmt"

	"github.com/metakeule/fmtdate"
	"github.com/webhippie/medialize/util"
)

// NextName tries to find an unused next file name.
func NextName(file, dest string, counter int) (string, error) {
	taken, err := CreationTime(file)

	if err != nil {
		s, err := util.Checksum(file)

		if err != nil {
			return "", fmt.Errorf("Failed to get checksum")
		}

		d := util.TrimSuffix(dest, "/")
		e := util.CleanExt(file)

		return fmt.Sprintf(
				"%s/0000/%s%s",
				d,
				s,
				e),
			err
	}

	d := util.TrimSuffix(dest, "/")
	f := fmtdate.Format("YYYY/MM", taken)
	t := taken.Format("20060102-150405-0700")
	s := fmt.Sprintf("%05d", counter)
	e := util.CleanExt(file)

	return fmt.Sprintf(
			"%s/%s/%s-%s%s",
			d,
			f,
			t,
			s,
			e),
		nil
}
