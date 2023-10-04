import { Global } from "@emotion/react";

export const FontsProvider = () => (
  <Global
    styles={`
      @import url('https://fonts.googleapis.com/css2?family=Poppins&family=Red+Hat+Display&family=Source+Code+Pro&display=swap');
    `}
  />
);

export const chakraFonts = {
  heading: `'Poppins', system-ui, sans-serif`,
  body: `'Red Hat Display', system-ui, sans-serif`,
  monospace: `'Source Code Pro', monospace`,
};
