package main

func commandPokedex(c *Config, args ...string) error {
	err := c.trainer.PrintContainPokemon()
	if err != nil {
		return err
	}
	return nil

}
