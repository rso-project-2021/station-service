package server

func Start(address, ginMode string) {
	router := NewRouter(ginMode)
	router.Run(address)
}
