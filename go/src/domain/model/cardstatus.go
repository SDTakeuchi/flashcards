package model

type CardStatus uint8

const (
	CardStatusUnspecified CardStatus = iota
	CardStatusRemembered
	CardStatusLearnAgain
	CardStatusNotRemembered
)

var cardStatuses = map[CardStatus]string{
	CardStatusUnspecified:   "unspecified",
	CardStatusRemembered:    "remembered",
	CardStatusLearnAgain:    "learn again",
	CardStatusNotRemembered: "not remembered",
}

func (c CardStatus) String() string {
	switch c {
	case CardStatusRemembered:
		return cardStatuses[CardStatusRemembered]
	case CardStatusLearnAgain:
		return cardStatuses[CardStatusLearnAgain]
	case CardStatusNotRemembered:
		return cardStatuses[CardStatusNotRemembered]
	default:
		return cardStatuses[CardStatusUnspecified]
	}
}

func (c CardStatus) Uint8() uint8 {
	switch c {
	case CardStatusRemembered:
		return uint8(CardStatusRemembered)
	case CardStatusLearnAgain:
		return uint8(CardStatusLearnAgain)
	case CardStatusNotRemembered:
		return uint8(CardStatusNotRemembered)
	default:
		return uint8(CardStatusUnspecified)
	}
}

func CardStatusFromString(status string) CardStatus {
	for k, v := range cardStatuses {
		if v == status {
			return k
		}
	}
	return CardStatusUnspecified
}

func CardStatusFromUint8(status uint8) CardStatus {
	for k := range cardStatuses {
		if uint8(k) == status {
			return k
		}
	}
	return CardStatusUnspecified
}
