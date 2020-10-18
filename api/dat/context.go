package dat

type Context struct {
	Txn      *Transaction
	Users    *UserMapper
	Sessions *SessionMapper
}

func NewDatabaseContext() (*Context, error) {
	m := &DbMap{}

	t, err := m.BeginTxn()
	if err != nil {
		return nil, err
	}

	c := Context{
		Txn:      t,
		Users:    &UserMapper{},
		Sessions: &SessionMapper{},
	}
	return &c, nil
}
