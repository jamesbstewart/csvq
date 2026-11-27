package csvq

import (
	"database/sql"
)

// DriverName is the SQL registration name for this fork.
const DriverName = "csvq-jamesbstewart"

func init() {
	sql.Register(DriverName, &Driver{})
}
