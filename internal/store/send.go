package store

var SendImpl func(cmd any)

func Send(cmd any) {
	if SendImpl != nil {
		SendImpl(cmd)
	}
}
