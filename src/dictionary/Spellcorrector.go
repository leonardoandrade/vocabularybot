package dictionary

type  SpellCorrector struct {
    termFrequencies map[string]string
}

func (c *SpellCorrector) Correct(word string) ([]string) {
        ret := []string{"österreich"}
        return ret
}
