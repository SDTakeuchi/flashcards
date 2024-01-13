"use client";

import { useState, useEffect } from "react";
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

const backendHost = process.env.BACKEND_HOST || "http://localhost/api";

export function FlashCard() {
  const [isFront, setIsFront] = useState(true);

  const flip = () => setIsFront(!isFront);
  const [word, setWord] = useState("");
  const [definition, setDefinition] = useState("");
  const [partOfSpeech, setPartOfSpeech] = useState("");
  const [example, setExample] = useState("");

  const getFlashCard = () => {
    setIsFront(true);
    setWord("Loading...");
    const url = backendHost + "/flashcard";
    fetch(url)
      .then((response) => response.json())
      .then((data) => {
        console.log(data);
        const res = data.data;
        setWord(res.word);
        setDefinition(res.description);
        setPartOfSpeech(res.part_of_speech);
        setExample(res.example);
      });
  };

  useEffect(() => {
    getFlashCard();
  }, []);

  return (
    <Card onClick={flip}>
      {isFront ? (
        <Front text={word} footer={partOfSpeech} />
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

function Front({ text, footer }) {
  return (
    <>
      <CardHeader>
        <Heading as="h2" size="md">
          WORD
        </Heading>
      </CardHeader>
      <CardBody>
        <Text>{text}</Text>
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
    <Flex>
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
      headers: { "Content-Type": "application/json" },
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