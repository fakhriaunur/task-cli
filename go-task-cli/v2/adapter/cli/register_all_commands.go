package cli

func (c *commands) RegisterAllCommands() {
	c.register("add", handlerAdd)
	c.register("update", handlerUpdate)
	c.register("delete", handlerDelete)
	c.register("mark-in-progress", handlerMarkInProgress)
	c.register("mark-done", handlerMarkDone)
	c.register("list", handlerList)
	c.register("help", handlerHelp)
}

// func (c *standaloneCommands) RegisterAllStandaloneCommands() {
// 	c.register("help", handlerHelp)
// }
