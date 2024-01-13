import { FlashCard } from "@/component/card";
import { Heading } from "@chakra-ui/react";

export default function Cards() {
  return (
    <div>
      <Heading as="h1" paddingBottom="5px">FlashCards</Heading>
      <FlashCard />
    </div>
  );
}
