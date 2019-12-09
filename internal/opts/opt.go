package opts

type Opt func() error
type Opts []Opt

func (opts Opts) Apply(ptr interface{}) error {
	for _, opt := range opts {
		if err := opt(); err != nil {
			return err
		}
	}
	return nil
}
