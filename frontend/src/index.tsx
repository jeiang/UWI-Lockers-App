import { render } from "preact";
import { LocationProvider, Router, Route } from "preact-iso";

import { Home } from "@/pages/Home/index.jsx";
import { NotFound } from "@/pages/_404.jsx";
import { ChakraProvider } from "@chakra-ui/react";
import { theme, FontsProvider } from "@/theme";

export const App = () => (
  <ChakraProvider theme={theme}>
    <FontsProvider />
    <LocationProvider>
      <Router>
        <Route
          path="/"
          component={Home}
        />
        <Route
          default
          component={NotFound}
        />
      </Router>
    </LocationProvider>
  </ChakraProvider>
);

render(<App />, document.getElementById("app"));
