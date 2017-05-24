package htmq

type Asseter interface {
	Asset(string) ([]byte, error)
}

func AScript(a Asseter, ss ...string) (*Tag, error) {
	var rErr error
	var inners []string
	for _, v := range ss {
		as, err := a.Asset(v)
		if err != nil {
			rErr = err
			continue
		}
		inners = append(inners, string(as))
	}
	return QScript(inners...), rErr
}
