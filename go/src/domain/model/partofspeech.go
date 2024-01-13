package model

type PartOfSpeech uint8

const (
	PartOfSpeechUnspecified PartOfSpeech = iota
	PartOfSpeechNoun
	PartOfSpeechVerb
	PartOfSpeechAdjective
	PartOfSpeechAdverb
	PartOfSpeechPronoun
	PartOfSpeechPreposition
	PartOfSpeechConjunction
	PartOfSpeechInterjection
)

var partOfSpeeches = map[PartOfSpeech]string{
	PartOfSpeechUnspecified: "unspecified",
	PartOfSpeechNoun:        "noun",
	PartOfSpeechVerb:        "verb",
	PartOfSpeechAdjective:   "adjective",
	PartOfSpeechAdverb:      "adverb",
	PartOfSpeechPronoun:     "pronoun",
	PartOfSpeechPreposition: "preposition",
	PartOfSpeechConjunction: "conjunction",
	PartOfSpeechInterjection: "interjection",
}

func (p PartOfSpeech) String() string {
	switch p {
	case PartOfSpeechNoun:
		return partOfSpeeches[PartOfSpeechNoun]
	case PartOfSpeechVerb:
		return partOfSpeeches[PartOfSpeechVerb]
	case PartOfSpeechAdjective:
		return partOfSpeeches[PartOfSpeechAdjective]
	case PartOfSpeechAdverb:
		return partOfSpeeches[PartOfSpeechAdverb]
	case PartOfSpeechPronoun:
		return partOfSpeeches[PartOfSpeechPronoun]
	case PartOfSpeechPreposition:
		return partOfSpeeches[PartOfSpeechPreposition]
	case PartOfSpeechConjunction:
		return partOfSpeeches[PartOfSpeechConjunction]
	case PartOfSpeechInterjection:
		return partOfSpeeches[PartOfSpeechInterjection]
	default:
		return partOfSpeeches[PartOfSpeechUnspecified]
	}
}

func (p PartOfSpeech) Uint8() uint8 {
	switch p {
	case PartOfSpeechNoun:
		return uint8(PartOfSpeechNoun)
	case PartOfSpeechVerb:
		return uint8(PartOfSpeechVerb)
	case PartOfSpeechAdjective:
		return uint8(PartOfSpeechAdjective)
	case PartOfSpeechAdverb:
		return uint8(PartOfSpeechAdverb)
	case PartOfSpeechPronoun:
		return uint8(PartOfSpeechPronoun)
	case PartOfSpeechPreposition:
		return uint8(PartOfSpeechPreposition)
	case PartOfSpeechConjunction:
		return uint8(PartOfSpeechConjunction)
	case PartOfSpeechInterjection:
		return uint8(PartOfSpeechInterjection)
	default:
		return uint8(PartOfSpeechUnspecified)
	}
}


func PartOfSpeechFromString(s string) PartOfSpeech {
	switch s {
	case partOfSpeeches[PartOfSpeechNoun]:
		return PartOfSpeechNoun
	case partOfSpeeches[PartOfSpeechVerb]:
		return PartOfSpeechVerb
	case partOfSpeeches[PartOfSpeechAdjective]:
		return PartOfSpeechAdjective
	case partOfSpeeches[PartOfSpeechAdverb]:
		return PartOfSpeechAdverb
	case partOfSpeeches[PartOfSpeechPronoun]:
		return PartOfSpeechPronoun
	case partOfSpeeches[PartOfSpeechPreposition]:
		return PartOfSpeechPreposition
	case partOfSpeeches[PartOfSpeechConjunction]:
		return PartOfSpeechConjunction
	case partOfSpeeches[PartOfSpeechInterjection]:
		return PartOfSpeechInterjection
	default:
		return PartOfSpeechUnspecified
	}
}

func PartOfSpeechFromUint8(pos uint8) PartOfSpeech {
	switch pos {
	case uint8(PartOfSpeechNoun):
		return PartOfSpeechNoun
	case uint8(PartOfSpeechVerb):
		return PartOfSpeechVerb
	case uint8(PartOfSpeechAdjective):
		return PartOfSpeechAdjective
	case uint8(PartOfSpeechAdverb):
		return PartOfSpeechAdverb
	case uint8(PartOfSpeechPronoun):
		return PartOfSpeechPronoun
	case uint8(PartOfSpeechPreposition):
		return PartOfSpeechPreposition
	case uint8(PartOfSpeechConjunction):
		return PartOfSpeechConjunction
	case uint8(PartOfSpeechInterjection):
		return PartOfSpeechInterjection
	default:
		return PartOfSpeechUnspecified
	}
}