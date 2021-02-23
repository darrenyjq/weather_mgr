package xconfig

type (
	CollectCardConfig struct {
		Cards map[string]struct {
			Triggers []struct {
				Grade  uint32 `json:"grade"`
				Amount uint32 `json:"amount"`
			} `json:"triggers"`
		}
	}
)
