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
	for k, v := range partOfSpeeches {
		if v == s {
			return k
		}
	}
	return PartOfSpeechUnspecified
}

func PartOfSpeechFromUint8(pos uint8) PartOfSpeech {
	for k := range partOfSpeeches {
		if uint8(k) == pos {
			return k
		}
	}
	return PartOfSpeechUnspecified
}