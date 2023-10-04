import { Header } from "@/components/Header";
import { Text } from "@chakra-ui/react";

const paths = [
  { to: "/", name: "Home" },
  { to: "/404", name: "404" },
];

// Empty Home for now
export const Home = () => (
  <>
    <Header pagesToNav={paths} />
    <Text>Hello World</Text>
  </>
);
