package physical

func LeisNewton(lei int) string {
	if lei == 1 {
		return "1º Lei de Newton - Lei da inércia"
	}
	if lei == 2 {
		return "2º Lei de Newton - Princípio fundamental da dinâmica"
	}
	if lei == 3 {
		return "3º Lei de Newton - Lei da ação e reação"
	}
	return "lei não existe!"
}
