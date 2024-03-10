import { getServerSession } from "next-auth";
import Credentials from "next-auth/providers/credentials";
import { userService } from "./services/userService";

export const authOptions = {
  session: {
    strategy: "jwt",
  },
  callbacks: {
    async jwt({ token, account, profile }) {
      if (account && account.type === "credentials") {
        token.userId = account.providerAccountId; // this is Id that coming from authorize() callback
      }
      return token;
    },
    async session({ session, token, user }) {
      session.user.id = token.userId;
      return session;
    },
  },
  pages: {
    signIn: "/login",
  },
  providers: [
    Credentials({
      name: "Credentials",
      credentials: {
        username: { label: "Username", type: "text", placeholder: "username" },
        password: { label: "Password", type: "password" },
      },
      async authorize(credentials, req) {
        const { username, password } = credentials;
        console.log("name and password: ", username, password);
        return userService.authenticate(username, password);
      },
    }),
  ],
};

export const getServerAuthSession = () => getServerSession(authOptions);
