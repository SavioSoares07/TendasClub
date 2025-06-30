package enum

type Status string

const (
	StatusMarked   Status = "marcado"
	StatusUnMarked Status = "nao_marcado"
	StatusPending  Status = "Pedente"
)