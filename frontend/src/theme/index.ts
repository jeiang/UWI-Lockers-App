import { extendTheme, withDefaultColorScheme } from "@chakra-ui/react";
import { FontsProvider, chakraFonts as fonts } from "./fonts";

const theme = extendTheme(
  {
    fonts,
  },
  withDefaultColorScheme({ colorScheme: "purple" })
);

export { theme, FontsProvider };
