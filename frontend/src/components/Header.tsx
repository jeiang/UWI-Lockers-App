import { useLocation } from "preact-iso";
import {
  ButtonGroup,
  Flex,
  FlexProps,
  Heading,
  Spacer,
  Stack,
} from "@chakra-ui/react";
import { LinkButton } from "@/components/LinkButton";

export interface HeaderProps extends FlexProps {
  pagesToNav?: Array<{ to: string; name: string }>;
}

export const Header = ({ pagesToNav, ...props }: HeaderProps) => {
  const { url, route } = useLocation();

  return (
    <Flex
      direction="column"
      {...props}
    >
      <Flex
        as="header"
        direction="row"
        padding={3}
        background="purple.600"
      >
        <Heading>SAC Locker Room App</Heading>
        <Spacer />
        <Stack
          direction={{ base: "column", md: "row" }}
          display="flex"
          alignItems="center"
        >
          <ButtonGroup flexDirection={{ base: "column", md: "row" }}>
            {pagesToNav?.map(({ to, name }) => (
              <LinkButton
                to={to}
                from={url}
                navigate={route}
              >
                {name}
              </LinkButton>
            ))}
          </ButtonGroup>
        </Stack>
      </Flex>
      <Spacer padding={2} />
    </Flex>
  );
};
