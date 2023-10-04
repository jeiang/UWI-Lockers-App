import { Header } from "@/components/Header";
import { Text, Heading, Center, VStack } from "@chakra-ui/react";

export const NotFound = () => (
  <>
    <Header />
    <Center>
      <VStack>
        <Heading>404: Not Found</Heading>
        <Text>It's not here :(</Text>
      </VStack>
    </Center>
  </>
);
