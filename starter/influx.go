package starter

type Influx struct {
	Mode      string
	Frequency int
}

func (m *Influx) Builder() error {
	return nil
}
