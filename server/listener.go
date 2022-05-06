package main

// Listener will listen to on-chain transactions and update the mysql database to reflect that

type Listener struct {
	*Store
}

func NewListener(store *Store) *Listener {
	return &Listener {
		Store: store,
	}
}

func (l *Listener) Start() {
	for {

		time.Sleep(100000)
	}
}