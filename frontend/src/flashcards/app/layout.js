import { ChakraProvider, ColorModeScript, Container } from "@chakra-ui/react";
import { Inter } from "next/font/google";
import "./globals.css";
import theme from "./theme";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "FlashCards",
  description: "Awesome flashcards app",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className={inter.className} style={{ height: "100vh" }}>
        <ChakraProvider theme={theme}>
          {/* <ColorModeScript initialColorMode={theme.config.initialColorMode} /> */}
          <Container maxW="container.md" minW="container.xs">
            {children}
          </Container>
        </ChakraProvider>
      </body>
    </html>
  );
}
