"use client";

import { useState, useEffect } from "react";
import Cookies from "js-cookie";
import {
  Button,
  Card,
  CardHeader,
  CardBody,
  Text,
  CardFooter,
  Spacer,
  Flex,
  Heading,
} from "@chakra-ui/react";

const backendHost = process.env.BACKEND_HOST || "http://localhost/backend_api";

export function FlashCard() {
  const [isFront, setIsFront] = useState(true);

  const flip = () => setIsFront(!isFront);
  const [word, setWord] = useState("");
  const [pronunciation, setPronunciation] = useState("");
  const [definition, setDefinition] = useState("");
  const [partOfSpeech, setPartOfSpeech] = useState("");
  const [example, setExample] = useState("");

  const getFlashCard = () => {
    setExample("");
    setPronunciation("");
    setIsFront(true);
    setWord("Loading...");

    const urlParams = new URLSearchParams(window.location.search);
    const url =
      urlParams.get("cards") === "remembered"
        ? backendHost + "/flashcard/remembered"
        : backendHost + "/flashcard";
    const mockRes = {
      word: "error",
      description: "a mistake",
      part_of_speech: "noun",
      example: "He admitted that he'd made an error.",
      pronunciation: "/ˈer.ɚ/",
    };

    try {
      fetch(url, {
        headers: {
          Authorization: `Bearer ${Cookies.get("flashcard_access_token")}`,
        },
      })
        .then((response) => response.json())
        .then((data) => {
          console.log(data);
          const res = data.data || mockRes;
          setWord(res.word);
          setDefinition(res.description);
          setPartOfSpeech(res.part_of_speech);
          setExample(res.example);
          setPronunciation(res.pronunciation);
        });
    } catch (error) {
      console.log(error);
      setWord(mockRes.word);
      setDefinition(mockRes.description);
      setPartOfSpeech(mockRes.part_of_speech);
      setExample(mockRes.example);
      setPronunciation(mockRes.pronunciation);
    }
  };

  useEffect(() => {
    getFlashCard();
  }, []);

  return (
    <Card onClick={flip} maxH="80vh" minH="300px">
      {isFront ? (
        <Front
          text={word}
          pronunciation={pronunciation}
          example={example}
          footer={partOfSpeech}
        />
      ) : (
        <Back
          definition={definition}
          example={example}
          word={word}
          getFlashCard={getFlashCard}
        />
      )}
    </Card>
  );
}

function Front({ text, pronunciation, example, footer }) {
  if (
    pronunciation &&
    !pronunciation.startsWith("/") &&
    !pronunciation.endsWith("/")
  ) {
    pronunciation = `/${pronunciation}/`;
  }

  return (
    <>
      <CardHeader>
        <Heading as="h2" size="md">
          WORD
        </Heading>
      </CardHeader>
      <CardBody>
        <Text>
          {text}
          {pronunciation && ` ${pronunciation}`}
        </Text>
        {example !== "" && (
          <>
            <Heading as="h4" size="sm" marginTop="10px" marginBottom="5px">
              Example
            </Heading>
            <Text>{example}</Text>
          </>
        )}
      </CardBody>
      {footer !== "unspecified" && <CardFooter>as {footer}</CardFooter>}
    </>
  );
}

function Back({ definition, example, word, getFlashCard }) {
  return (
    <>
      <CardHeader>
        <Heading as="h2" size="md">
          DEFINITION
        </Heading>
      </CardHeader>
      <CardBody>
        <Text>{definition}</Text>
        {example !== "" && (
          <>
            <Heading as="h4" size="sm" marginTop="10px" marginBottom="5px">
              Example
            </Heading>
            <Text>{example}</Text>
          </>
        )}
      </CardBody>
      <CardFooter>
        <FlashCardButtons
          word={word}
          definition={definition}
          getFlashCard={getFlashCard}
        />
      </CardFooter>
    </>
  );
}

function FlashCardButtons({ word, definition, getFlashCard }) {
  return (
    <Flex width="100%">
      <UpdateStatusButton
        status={1}
        word={word}
        definition={definition}
        getFlashCard={getFlashCard}
      >
        Remembered
      </UpdateStatusButton>

      <Spacer />

      <UpdateStatusButton
        status={2}
        word={word}
        definition={definition}
        getFlashCard={getFlashCard}
      >
        Learn Again
      </UpdateStatusButton>

      <Spacer />

      <UpdateStatusButton
        status={3}
        word={word}
        definition={definition}
        getFlashCard={getFlashCard}
      >
        Not Remembered
      </UpdateStatusButton>
    </Flex>
  );
}

function UpdateStatusButton({
  children,
  status,
  word,
  definition,
  getFlashCard,
}) {
  const url = backendHost + "/flashcard";
  const updateStatus = () => {
    fetch(url, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${Cookies.get("flashcard_access_token")}`,
      },
      body: JSON.stringify({ word, status, description: definition }),
    }).then(() => {
      getFlashCard();
    });
  };

  return (
    <Button as="a" onClick={updateStatus}>
      {children}
    </Button>
  );
}
