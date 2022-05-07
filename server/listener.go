package main

// Listener will listen to on-chain transactions and update the mysql database to reflect that

type Updater struct {
	*Store
}

func NewUpdater(store *Store) *Updater {
	return &Listener {
		Store: store,
	}
}

func (u *Updater) Start() {
	// Setup ethereum event log/listener for an update transaction

	for {

		time.Sleep(100000)
	}
}

