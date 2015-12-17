package photo

import (
	"fmt"
	"github.com/metakeule/fmtdate"
	"github.com/webhippie/medialize/util"
)

func NextName(file, dest string, counter int) (string, error) {
	taken, err := CreationTime(file)

	if err != nil {
		d := util.TrimSuffix(dest, "/")
		s := fmt.Sprintf("%05d", counter)
		e := util.CleanExt(file)

		return fmt.Sprintf(
				"%s/0000/%s%s",
				d,
				s,
				e),
			err
	} else {
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
}
