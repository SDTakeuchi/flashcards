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
	switch status {
	case cardStatuses[CardStatusRemembered]:
		return CardStatusRemembered
	case cardStatuses[CardStatusLearnAgain]:
		return CardStatusLearnAgain
	case cardStatuses[CardStatusNotRemembered]:
		return CardStatusNotRemembered
	default:
		return CardStatusUnspecified
	}
}

func CardStatusFromUint8(status uint8) CardStatus {
	switch status {
	case uint8(CardStatusRemembered):
		return CardStatusRemembered
	case uint8(CardStatusLearnAgain):
		return CardStatusLearnAgain
	case uint8(CardStatusNotRemembered):
		return CardStatusNotRemembered
	default:
		return CardStatusUnspecified
	}
}
